package usecases

import (
	"context"
	"sync"
	"time"

	"hotel-management-system/controllers/rest/requests"
	"hotel-management-system/controllers/rest/responses"
	"hotel-management-system/domains/constants"
	"hotel-management-system/domains/models"
	"hotel-management-system/repositories"
	mysqlRepo "hotel-management-system/repositories/mysql"
)

type HotelManagementUseCase interface {
	AddHotel(c context.Context, request requests.AddHotelRequest) (models.Hotel, error)
	GetAvailableHotel(c context.Context, checkin, checkout string) (responses.AvailableHotels, error)
}

type hotelManagementUseCase struct {
	hotelRepo       mysqlRepo.HotelRepository
	roomRepo        mysqlRepo.RoomRepository
	reservationRepo mysqlRepo.ReservationRepository
}

func NewHotelManagementUseCase(repo *repositories.Repository) HotelManagementUseCase {
	return &hotelManagementUseCase{
		hotelRepo:       repo.Hotel,
		roomRepo:        repo.Room,
		reservationRepo: repo.Reservation,
	}
}

func (uc *hotelManagementUseCase) AddHotel(c context.Context, request requests.AddHotelRequest) (
	newHotel models.Hotel, err error) {

	ctx, cancel := context.WithTimeout(c, time.Duration(2)*time.Second)
	defer cancel()

	for i, room := range request.Rooms {
		if room.Status != constants.AvailableRoomStatus &&
			room.Status != constants.OutOfServiceRoomStatus {
			request.Rooms[i].Status = constants.AvailableRoomStatus
		}
	}

	newHotel, err = uc.hotelRepo.Create(ctx, models.Hotel{
		Name:    request.Name,
		Address: request.Address,
		Rooms:   request.Rooms,
	})
	if err != nil {
		return
	}

	newHotel.Rooms, err = uc.roomRepo.GetByHotel(ctx, newHotel)
	if err != nil {
		return
	}

	return
}

func (uc *hotelManagementUseCase) GetAvailableHotel(c context.Context, checkin, checkout string) (
	response responses.AvailableHotels, err error) {

	ctx, cancel := context.WithTimeout(c, time.Duration(5)*time.Second)
	defer cancel()

	checkinDate, _ := parseDate(checkin)
	checkoutDate, _ := parseDate(checkout)

	var allHotels models.Hotels
	allHotels, err = uc.hotelRepo.Get(ctx)
	if err != nil {
		return
	}

	response.Hotels = make(responses.Hotels, 0)
	availableHotelChannel := make(chan models.Hotel)
	errChannel := make(chan error)
	wg := sync.WaitGroup{}
	for _, hotel := range allHotels {
		wg.Add(1)

		go func(wg *sync.WaitGroup, hotel models.Hotel) {
			defer wg.Done()

			reservations, err := uc.reservationRepo.GetByHotelAndDates(ctx, hotel, checkinDate, checkoutDate)
			if err != nil {
				errChannel <- err
			}

			var occupiedRooms models.Rooms
			for _, reservation := range reservations {
				var room models.Room
				room.ID = reservation.RoomID
				occupiedRooms = append(occupiedRooms, room)
			}

			availableRooms, err := uc.roomRepo.GetAvailableByHotelAndOccupiedRooms(ctx, hotel, occupiedRooms)
			if err != nil {
				errChannel <- err
			}

			if len(availableRooms) != 0 {
				hotel.Rooms = availableRooms
				availableHotelChannel <- hotel
			}
		}(&wg, hotel)
	}

	go func(wg *sync.WaitGroup) {
		wg.Wait()
		cancel()
	}(&wg)

	for {
		select {
		case availableHotel := <-availableHotelChannel:
			response.TotalHotels++
			response.Hotels = append(response.Hotels, responses.Hotel{
				Hotel:      availableHotel,
				TotalRooms: len(availableHotel.Rooms),
			})
		case err := <-errChannel:
			return response, err
		case <-ctx.Done():
			return response, nil
		}
	}
}
