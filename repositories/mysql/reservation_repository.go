package mysql

import (
	"context"
	"time"

	"gorm.io/gorm"

	"hotel-management-system/domains/models"
)

type ReservationRepository interface {
	Create(context.Context, models.Reservation) (models.Reservation, error)
	GetByHotelAndDates(context.Context, models.Hotel, time.Time, time.Time) (models.Reservations, error)
	Find(context.Context, models.Reservation) (models.Reservation, error)
}

type reservationRepository struct {
	db               *gorm.DB
	tableAssociation string
}

func NewReservationRepository(db *gorm.DB) ReservationRepository {
	return &reservationRepository{
		db:               db,
		tableAssociation: "Reservations",
	}
}

func (repo *reservationRepository) Create(_ context.Context, reservation models.Reservation) (
	models.Reservation, error) {

	query := repo.db.Create(&reservation)
	return reservation, query.Error
}

func (repo *reservationRepository) GetByHotelAndDates(
	_ context.Context,
	hotel models.Hotel,
	checkinDate, checkoutDate time.Time,
) (reservations models.Reservations, err error) {

	err = repo.db.Model(&hotel).
		Where("checkin_date <= ?", checkoutDate).
		Where("checkout_date >= ?", checkinDate).
		Association(repo.tableAssociation).
		Find(&reservations)
	return
}

func (repo *reservationRepository) Find(_ context.Context, reservation models.Reservation) (models.Reservation, error) {
	query := repo.db.First(&reservation, reservation)
	return reservation, query.Error
}
