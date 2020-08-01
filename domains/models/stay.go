package models

import "gorm.io/gorm"

type Stay struct {
	gorm.Model
	ReservationID uint     `gorm:"reservation_id" json:"reservation_id"`
	RoomID        uint     `gorm:"room_id" json:"room_id"`
	GuestName     string   `gorm:"guest_name" json:"guest_name"`
	StayRoom      StayRoom `json:"stay_room,omitempty"`
}

type Stays []Stay
