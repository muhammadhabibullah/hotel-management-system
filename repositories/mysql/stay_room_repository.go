package mysql

import (
	"context"

	"gorm.io/gorm"

	"hotel-management-system/domains/models"
)

type StayRoomRepository interface {
	Create(context.Context, models.StayRoom) (uint, error)
	Update(context.Context, models.StayRoom) error
	Get(context.Context) (models.StayRooms, error)
	Find(context.Context, models.StayRoom) (models.StayRoom, error)
	Delete(context.Context, models.StayRoom) error
}

type stayRoomRepository struct {
	db *gorm.DB
}

func NewStayRoomRepository(db *gorm.DB) StayRoomRepository {
	return &stayRoomRepository{
		db: db,
	}
}

func (repo *stayRoomRepository) Create(_ context.Context, stayRoom models.StayRoom) (uint, error) {
	query := repo.db.Create(&stayRoom)
	return stayRoom.ID, query.Error
}

func (repo *stayRoomRepository) Update(_ context.Context, stayRoom models.StayRoom) error {
	query := repo.db.Save(&stayRoom)
	return query.Error
}

func (repo *stayRoomRepository) Get(_ context.Context) (models.StayRooms, error) {
	var stayRooms models.StayRooms
	query := repo.db.Limit(10).
		Find(&stayRooms)
	return stayRooms, query.Error
}

func (repo *stayRoomRepository) Find(_ context.Context, stayRoom models.StayRoom) (models.StayRoom, error) {
	query := repo.db.First(&stayRoom, stayRoom.ID)
	return stayRoom, query.Error
}

func (repo *stayRoomRepository) Delete(_ context.Context, stayRoom models.StayRoom) error {
	query := repo.db.Delete(&stayRoom)
	return query.Error
}
