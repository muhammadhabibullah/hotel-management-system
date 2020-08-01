package main

import (
	"github.com/joho/godotenv"

	"hotel-management-system/controllers"
	"hotel-management-system/repositories"
	"hotel-management-system/services"
)

func main() {
	godotenv.Load()
	controllers.Init(services.Init(repositories.Init()))
}
