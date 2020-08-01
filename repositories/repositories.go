package repositories

import (
	mysqlRepo "hotel-management-system/repositories/mysql"
)

type Repository struct {
	Hotel       mysqlRepo.HotelRepository
	Room        mysqlRepo.RoomRepository
	Reservation mysqlRepo.ReservationRepository
	Stay        mysqlRepo.StayRepository
	StayRoom    mysqlRepo.StayRoomRepository
}

func Init() *Repository {
	mysqlDB := mysqlRepo.InitDB()
	return &Repository{
		Hotel:       mysqlRepo.NewHotelRepository(mysqlDB),
		Room:        mysqlRepo.NewRoomRepository(mysqlDB),
		Reservation: mysqlRepo.NewReservationRepository(mysqlDB),
		Stay:        mysqlRepo.NewStayRepository(mysqlDB),
		StayRoom:    mysqlRepo.NewStayRoomRepository(mysqlDB),
	}
}
