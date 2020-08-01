package requests

type AddReservationRequest struct {
	HotelID      uint   `json:"hotel_id" validate:"gt=0" example:"2"`
	CustomerName string `json:"customer_name" validate:"required" example:"Customer A"`
	CheckinDate  string `json:"checkin_date" validate:"datetime=2006-01-02T15:04:05Z07:00" example:"2020-08-01T05:49:49.053+07:00"`  //nolint:lll
	CheckoutDate string `json:"checkout_date" validate:"datetime=2006-01-02T15:04:05Z07:00" example:"2020-08-02T05:49:49.053+07:00"` //nolint:lll
}

type CheckinRequest struct {
	OrderID   string `json:"order_id" validate:"required" example:"1234ABCD"`
	GuestName string `json:"guest_name" validate:"required" example:"Guest A"`
}
