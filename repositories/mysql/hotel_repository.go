package mysql

import (
	"context"

	"gorm.io/gorm"

	"hotel-management-system/domains/models"
)

type HotelRepository interface {
	Create(context.Context, models.Hotel) (models.Hotel, error)
	Get(context.Context) (models.Hotels, error)
	Find(context.Context, models.Hotel) (models.Hotel, error)
}

type hotelRepository struct {
	db *gorm.DB
}

func NewHotelRepository(db *gorm.DB) HotelRepository {
	return &hotelRepository{
		db: db,
	}
}

func (repo *hotelRepository) Create(_ context.Context, hotel models.Hotel) (models.Hotel, error) {
	query := repo.db.Create(&hotel)
	return hotel, query.Error
}

func (repo *hotelRepository) Get(_ context.Context) (models.Hotels, error) {
	var hotels models.Hotels
	query := repo.db.Find(&hotels)
	return hotels, query.Error
}

func (repo *hotelRepository) Find(_ context.Context, hotel models.Hotel) (models.Hotel, error) {
	query := repo.db.First(&hotel, hotel.ID)
	return hotel, query.Error
}
