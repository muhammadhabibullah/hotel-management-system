package exceptions

import "errors"

var ErrHotelNotFound = errors.New("hotel not found")

var ErrNoRoomAvailable = errors.New("no room available")

var ErrCheckinAndCheckOutDateInvalid = errors.New("checkin & checkout date are invalid")

var ErrOrderIDNotFound = errors.New("orderID not found")

var ErrCheckInTooEarly = errors.New("check-in too early")

var ErrLateCheckIn = errors.New("late check-in")
