package requests

import "hotel-management-system/domains/models"

type AddHotelRequest struct {
	Name    string       `json:"name" validate:"required" example:"Hotel A"`
	Address string       `json:"address" validate:"required" example:"Jl. Bandung"`
	Rooms   models.Rooms `json:"rooms"`
}
