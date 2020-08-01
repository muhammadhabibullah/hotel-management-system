package usecases

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"

	"gorm.io/gorm"

	"hotel-management-system/controllers/rest/exceptions"
	"hotel-management-system/controllers/rest/requests"
	"hotel-management-system/domains/models"
	"hotel-management-system/repositories"
	mysqlRepo "hotel-management-system/repositories/mysql"
)

type StayManagementUseCase interface {
	AddReservation(c context.Context, request requests.AddReservationRequest) (models.Reservation, error)
	CheckIn(c context.Context, request requests.CheckinRequest) (models.Stay, error)
}

type stayManagementUseCase struct {
	hotelRepo       mysqlRepo.HotelRepository
	roomRepo        mysqlRepo.RoomRepository
	reservationRepo mysqlRepo.ReservationRepository
	stayRepo        mysqlRepo.StayRepository
	stayRoomRepo    mysqlRepo.StayRoomRepository
}

func NewStayManagementUseCase(repo *repositories.Repository) StayManagementUseCase {
	return &stayManagementUseCase{
		hotelRepo:       repo.Hotel,
		roomRepo:        repo.Room,
		reservationRepo: repo.Reservation,
		stayRepo:        repo.Stay,
		stayRoomRepo:    repo.StayRoom,
	}
}

func (uc *stayManagementUseCase) AddReservation(c context.Context, request requests.AddReservationRequest) (
	reservation models.Reservation, err error) {

	ctx, cancel := context.WithTimeout(c, time.Duration(2)*time.Second)
	defer cancel()

	checkinDate, _ := parseDate(request.CheckinDate)
	checkoutDate, _ := parseDate(request.CheckoutDate)
	if !checkinDate.Before(checkoutDate) {
		return reservation, fmt.Errorf("error add reservation with check-in date %s and check-out date %s: %w",
			checkinDate.Format(time.RFC822), checkoutDate.Format(time.RFC822), exceptions.ErrCheckinAndCheckOutDateInvalid)
	}

	var hotel models.Hotel
	hotel.ID = request.HotelID
	hotel, err = uc.hotelRepo.Find(ctx, hotel)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return reservation, fmt.Errorf("error add reservation to hotel with ID %d: %w",
				request.HotelID, exceptions.ErrHotelNotFound)
		}
		return
	}

	reservations, err := uc.reservationRepo.GetByHotelAndDates(ctx, hotel, checkinDate, checkoutDate)
	if err != nil {
		return
	}

	var occupiedRooms models.Rooms
	for _, reservation := range reservations {
		var room models.Room
		room.ID = reservation.RoomID
		occupiedRooms = append(occupiedRooms, room)
	}

	var availableRooms models.Rooms
	availableRooms, err = uc.roomRepo.GetAvailableByHotelAndOccupiedRooms(ctx, hotel, occupiedRooms)
	if err != nil {
		return
	}
	if len(availableRooms) == 0 {
		return reservation, fmt.Errorf("error add reservation to hotel %s: %w",
			hotel.Name, exceptions.ErrNoRoomAvailable)
	}
	availableRoom := availableRooms[0]

	reservation.RoomID = availableRoom.ID
	reservation.HotelID = request.HotelID
	reservation.OrderID = uc.generateUniqueOrderID(ctx)
	reservation.CustomerName = request.CustomerName
	reservation.CheckinDate = checkinDate
	reservation.CheckoutDate = checkoutDate
	reservation, err = uc.reservationRepo.Create(ctx, reservation)
	if err != nil {
		return
	}

	return
}

func (uc *stayManagementUseCase) CheckIn(c context.Context, request requests.CheckinRequest) (models.Stay, error) {

	ctx, cancel := context.WithTimeout(c, time.Duration(2)*time.Second)
	defer cancel()

	reservation, err := uc.reservationRepo.Find(ctx, models.Reservation{
		OrderID: request.OrderID,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Stay{}, fmt.Errorf("error check-in with orderID %s: %w",
				request.OrderID, exceptions.ErrOrderIDNotFound)
		}
		return models.Stay{}, err
	}

	if roundDate(time.Now()).Before(roundDate(reservation.CheckinDate)) {
		return models.Stay{}, fmt.Errorf("error check-in %s: %w",
			reservation.CheckinDate.Format(time.RFC822), exceptions.ErrCheckInTooEarly)
	}
	if roundDate(time.Now()).After(roundDate(reservation.CheckoutDate)) {
		return models.Stay{}, fmt.Errorf("error latest check-in %s: %w",
			reservation.CheckoutDate.Format(time.RFC822), exceptions.ErrLateCheckIn)
	}

	return uc.stayRepo.Create(ctx, models.Stay{
		Model:         gorm.Model{},
		ReservationID: reservation.ID,
		RoomID:        reservation.RoomID,
		GuestName:     request.GuestName,
		StayRoom: models.StayRoom{
			RoomID: reservation.RoomID,
			Date:   time.Now(),
		},
	})
}

const orderIDLength = 8

func (uc *stayManagementUseCase) generateUniqueOrderID(ctx context.Context) string {
	rand.Seed(time.Now().UnixNano())

	var letters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	b := make([]rune, orderIDLength)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	orderID := string(b)

	_, err := uc.reservationRepo.Find(ctx, models.Reservation{
		OrderID: orderID,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return orderID
		}
		return ""
	}

	return uc.generateUniqueOrderID(ctx)
}
