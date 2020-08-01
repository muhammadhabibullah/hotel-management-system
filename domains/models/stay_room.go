package models

import (
	"time"

	"gorm.io/gorm"
)

type StayRoom struct {
	gorm.Model
	StayID uint      `gorm:"stay_id" json:"stay_id"`
	RoomID uint      `gorm:"room_id" json:"room_id"`
	Date   time.Time `gorm:"date" json:"date"`
}

type StayRooms []StayRoom
