package services

import (
	"hotel-management-system/repositories"
	useCases "hotel-management-system/services/use_cases"
)

type Services struct {
	HotelManagement useCases.HotelManagementUseCase
	StayManagement  useCases.StayManagementUseCase
}

func Init(repo *repositories.Repository) *Services {
	return &Services{
		HotelManagement: useCases.NewHotelManagementUseCase(repo),
		StayManagement:  useCases.NewStayManagementUseCase(repo),
	}
}
