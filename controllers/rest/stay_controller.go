package rest

import (
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"

	"hotel-management-system/controllers/rest/exceptions"
	"hotel-management-system/controllers/rest/requests"
	"hotel-management-system/controllers/rest/responses"
	"hotel-management-system/services"
	useCases "hotel-management-system/services/use_cases"
)

type StayController struct {
	validate         *validator.Validate
	stayManagementUC useCases.StayManagementUseCase
}

func NewStayController(
	e *echo.Echo,
	v *validator.Validate,
	svc *services.Services,
) {
	ctrl := &StayController{
		validate:         v,
		stayManagementUC: svc.StayManagement,
	}

	v1 := e.Group("/v1")

	v1Reservation := v1.Group("/reservation")
	{
		v1Reservation.POST("", ctrl.AddReservation)
	}

	v1Stay := v1.Group("/stay")
	{
		v1Stay.POST("/check_in", ctrl.CheckIn)
	}
}

// AddReservation godoc
// @Summary Add reservation
// @Description Create reservation to hotel room
// @Tags Stay
// @Accept json
// @Produce json
// @Param request body requests.AddReservationRequest true "Request Body"
// @Success 201 {object} models.Reservation "Ok"
// @Failure 400 {object} responses.ErrorResponse "Status Bad Request"
// @Failure 422 {object} responses.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} responses.ErrorResponse "Internal Server Error"
// @Router /v1/reservation [post]
func (ctrl *StayController) AddReservation(c echo.Context) error {
	var request requests.AddReservationRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, responses.NewErrorResponse(err))
	}
	if err := ctrl.validate.Struct(request); err != nil {
		return c.JSON(http.StatusBadRequest, responses.NewErrorResponse(err))
	}

	ctx := c.Request().Context()
	response, err := ctrl.stayManagementUC.AddReservation(ctx, request)
	if err != nil {
		if errors.Is(err, exceptions.ErrCheckinAndCheckOutDateInvalid) ||
			errors.Is(err, exceptions.ErrHotelNotFound) ||
			errors.Is(err, exceptions.ErrNoRoomAvailable) {
			return c.JSON(http.StatusBadRequest, responses.NewErrorResponse(err))
		}
		return c.JSON(http.StatusInternalServerError, responses.NewErrorResponse(err))
	}

	return c.JSON(http.StatusCreated, response)
}

// CheckIn godoc
// @Summary Check-in
// @Description Check-in using orderID after reservation
// @Tags Stay
// @Accept json
// @Produce json
// @Param request body requests.CheckinRequest true "Request Body"
// @Success 201 {object} models.Stay "Ok"
// @Failure 400 {object} responses.ErrorResponse "Status Bad Request"
// @Failure 422 {object} responses.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} responses.ErrorResponse "Internal Server Error"
// @Router /v1/stay/check_in [post]
func (ctrl *StayController) CheckIn(c echo.Context) error {
	var request requests.CheckinRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, responses.NewErrorResponse(err))
	}
	if err := ctrl.validate.Struct(request); err != nil {
		return c.JSON(http.StatusBadRequest, responses.NewErrorResponse(err))
	}

	ctx := c.Request().Context()
	response, err := ctrl.stayManagementUC.CheckIn(ctx, request)
	if err != nil {
		if errors.Is(err, exceptions.ErrOrderIDNotFound) ||
			errors.Is(err, exceptions.ErrCheckInTooEarly) ||
			errors.Is(err, exceptions.ErrLateCheckIn) {
			return c.JSON(http.StatusBadRequest, responses.NewErrorResponse(err))
		}
		return c.JSON(http.StatusInternalServerError, responses.NewErrorResponse(err))
	}

	return c.JSON(http.StatusCreated, response)
}
