package requests

import "hotel-management-system/domains/models"

type AvailableHotelRequest struct {
	CheckinDate  string `validate:"datetime=2006-01-02T15:04:05Z07:00"`
	CheckoutDate string `validate:"datetime=2006-01-02T15:04:05Z07:00"`
}

type AddHotelRequest struct {
	Name    string       `json:"name" validate:"required" example:"Hotel A"`
	Address string       `json:"address" validate:"required" example:"Jl. Bandung"`
	Rooms   models.Rooms `json:"rooms"`
}
