package models

import (
	"time"

	"gorm.io/gorm"
)

type Reservation struct {
	gorm.Model
	OrderID      string    `gorm:"order_id;unique_index" json:"order_id"`
	RoomID       uint      `gorm:"room_id" json:"room_id"`
	HotelID      uint      `gorm:"hotel_id" json:"hotel_id"`
	CustomerName string    `gorm:"customer_name" json:"customer_name"`
	CheckinDate  time.Time `gorm:"checkin_date" json:"checkin_date"`
	CheckoutDate time.Time `gorm:"checkout_date" json:"checkout_date"`
	Stays        Stays     `json:"stays,omitempty"`
}

type Reservations []Reservation
