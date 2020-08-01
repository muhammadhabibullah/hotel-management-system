package controllers

import (
	"os"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

	"hotel-management-system/controllers/rest"
	"hotel-management-system/controllers/rest/validators"
	"hotel-management-system/docs" // swagger
	"hotel-management-system/services"
)

func Init(svc *services.Services) {
	e := echo.New()
	v := validators.InitValidators()

	initDocs(e)
	rest.NewHotelController(e, v, svc)
	rest.NewStayController(e, v, svc)

	e.Logger.Fatal(e.Start(os.Getenv("PORT")))
}

func initDocs(e *echo.Echo) {
	docs.SwaggerInfo.Title = "Hotel Management Service API API"
	docs.SwaggerInfo.Description = "Swagger API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = ""
	docs.SwaggerInfo.BasePath = ""

	e.GET("swagger/*", echoSwagger.WrapHandler)
}
