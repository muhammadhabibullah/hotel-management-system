package rest

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"

	"hotel-management-system/controllers/rest/requests"
	"hotel-management-system/controllers/rest/responses"
	"hotel-management-system/services"
	useCases "hotel-management-system/services/use_cases"
)

type HotelController struct {
	validate          *validator.Validate
	hotelManagementUC useCases.HotelManagementUseCase
}

func NewHotelController(
	e *echo.Echo,
	v *validator.Validate,
	svc *services.Services,
) {
	ctrl := &HotelController{
		validate:          v,
		hotelManagementUC: svc.HotelManagement,
	}

	v1 := e.Group("/v1")

	v1Hotel := v1.Group("/hotel")
	{
		v1Hotel.GET("", ctrl.GetAvailable)
		v1Hotel.POST("", ctrl.AddHotel)
	}
}

// GetAvailable godoc
// @Summary Get available
// @Description Get all available hotel rooms
// @Tags Hotel
// @Accept json
// @Produce json
// @Param checkin_date query string true "Check in date" default("2020-08-01T06:00:00.000+07:00")
// @Param checkout_date query string true "Check out date" default("2020-08-02T06:00:00.000+07:00")
// @Success 200 {object} responses.AvailableHotels "Ok"
// @Failure 400 {object} responses.ErrorResponse "Status Bad Request"
// @Failure 500 {object} responses.ErrorResponse "Internal Server Error"
// @Router /v1/hotel [get]
func (ctrl *HotelController) GetAvailable(c echo.Context) error {
	request := requests.AvailableHotelRequest{
		CheckinDate:  c.QueryParam("checkin_date"),
		CheckoutDate: c.QueryParam("checkout_date"),
	}
	if err := ctrl.validate.Struct(request); err != nil {
		return c.JSON(http.StatusBadRequest, responses.NewErrorResponse(err))
	}

	ctx := c.Request().Context()
	response, err := ctrl.hotelManagementUC.GetAvailableHotel(ctx, request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.NewErrorResponse(err))
	}

	return c.JSON(http.StatusOK, response)
}

// AddHotel godoc
// @Summary Add hotel
// @Description Add new hotel registry
// @Tags Hotel
// @Accept json
// @Produce json
// @Param request body requests.AddHotelRequest true "Request Body"
// @Success 201 {object} models.Hotel "Ok"
// @Failure 400 {object} responses.ErrorResponse "Status Bad Request"
// @Failure 422 {object} responses.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} responses.ErrorResponse "Internal Server Error"
// @Router /v1/hotel [post]
func (ctrl *HotelController) AddHotel(c echo.Context) error {
	var request requests.AddHotelRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, responses.NewErrorResponse(err))
	}
	if err := ctrl.validate.Struct(request); err != nil {
		return c.JSON(http.StatusBadRequest, responses.NewErrorResponse(err))
	}

	ctx := c.Request().Context()
	response, err := ctrl.hotelManagementUC.AddHotel(ctx, request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.NewErrorResponse(err))
	}

	return c.JSON(http.StatusCreated, response)
}
