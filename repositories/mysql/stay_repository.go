package mysql

import (
	"context"

	"gorm.io/gorm"

	"hotel-management-system/domains/models"
)

type StayRepository interface {
	Create(context.Context, models.Stay) (models.Stay, error)
	Update(context.Context, models.Stay) error
	Get(context.Context) (models.Stays, error)
	Find(context.Context, models.Stay) (models.Stay, error)
	Delete(context.Context, models.Stay) error
}

type stayRepository struct {
	db *gorm.DB
}

func NewStayRepository(db *gorm.DB) StayRepository {
	return &stayRepository{
		db: db,
	}
}

func (repo *stayRepository) Create(_ context.Context, stay models.Stay) (models.Stay, error) {
	query := repo.db.Create(&stay)
	return stay, query.Error
}

func (repo *stayRepository) Update(_ context.Context, stay models.Stay) error {
	query := repo.db.Save(&stay)
	return query.Error
}

func (repo *stayRepository) Get(_ context.Context) (models.Stays, error) {
	var stays models.Stays
	query := repo.db.Limit(10).
		Find(&stays)
	return stays, query.Error
}

func (repo *stayRepository) Find(_ context.Context, stay models.Stay) (models.Stay, error) {
	query := repo.db.First(&stay, stay.ID)
	return stay, query.Error
}

func (repo *stayRepository) Delete(_ context.Context, stay models.Stay) error {
	query := repo.db.Delete(&stay)
	return query.Error
}
