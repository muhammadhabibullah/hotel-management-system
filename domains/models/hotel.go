package models

import "gorm.io/gorm"

type Hotel struct {
	gorm.Model
	Name         string       `gorm:"name" json:"name"`
	Address      string       `gorm:"address" json:"address"`
	Rooms        Rooms        `json:"rooms,omitempty"`
	Reservations Reservations `json:"reservations,omitempty"`
}

type Hotels []Hotel
