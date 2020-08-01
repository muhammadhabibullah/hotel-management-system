package mysql

import (
	"context"

	"gorm.io/gorm"

	"hotel-management-system/domains/constants"
	"hotel-management-system/domains/models"
)

type RoomRepository interface {
	GetByHotel(context.Context, models.Hotel) (models.Rooms, error)
	GetAvailableByHotelAndOccupiedRooms(context.Context, models.Hotel, models.Rooms) (models.Rooms, error)
}

type roomRepository struct {
	db               *gorm.DB
	tableAssociation string
}

func NewRoomRepository(db *gorm.DB) RoomRepository {
	return &roomRepository{
		db:               db,
		tableAssociation: "Rooms",
	}
}

func (repo *roomRepository) GetByHotel(_ context.Context, hotel models.Hotel) (models.Rooms, error) {
	var rooms models.Rooms
	err := repo.db.Model(&hotel).
		Association(repo.tableAssociation).
		Find(&rooms)
	return rooms, err
}

func (repo *roomRepository) GetAvailableByHotelAndOccupiedRooms(
	_ context.Context,
	hotel models.Hotel,
	occupiedRooms models.Rooms,
) (availableRoom models.Rooms, err error) {

	occupiedRoomIDs := make([]uint, 0)
	for _, room := range occupiedRooms {
		occupiedRoomIDs = append(occupiedRoomIDs, room.ID)
	}

	query := repo.db.Model(&hotel).
		Where("status = ?", constants.AvailableRoomStatus)
	if len(occupiedRoomIDs) != 0 {
		query = query.Where("id NOT IN ?", occupiedRoomIDs)
	}
	err = query.Association(repo.tableAssociation).
		Find(&availableRoom)

	return
}
