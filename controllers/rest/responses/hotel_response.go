package responses

import "hotel-management-system/domains/models"

type AvailableHotels struct {
	Hotels      Hotels `json:"hotels"`
	TotalHotels int    `json:"total_hotels"`
}

type Hotel struct {
	models.Hotel
	TotalRooms int `json:"total_rooms"`
}

type Hotels []Hotel
