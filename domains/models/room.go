package models

import (
	"gorm.io/gorm"

	"hotel-management-system/domains/constants"
)

type Room struct {
	gorm.Model
	HotelID      uint                 `gorm:"hotel_id" json:"hotel_id"`
	Number       uint                 `gorm:"number" json:"number"`
	Status       constants.RoomStatus `gorm:"status" json:"status"`
	Reservations Reservations         `json:"reservations,omitempty"`
	Stays        Stays                `json:"stays,omitempty"`
	StayRooms    StayRooms            `json:"stay_rooms,omitempty"`
}

type Rooms []Room
