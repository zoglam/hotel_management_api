package controllers

import (
    "fmt"

    models "github.com/zoglam/hotel_management_api/models"
)

// HTTPBookingResponse is form to prepare response for bookings func-s
type HTTPBookingResponse struct {
    Status  string           `json:"status"`
    Details []models.Booking `json:"details"`
}

// HTTPHotelRoomResponse is form to prepare response for hotel_room func-s
type HTTPHotelRoomResponse struct {
    Status  string             `json:"status"`
    Details []models.HotelRoom `json:"details"`
}

// GetErrorJSON prepares form for error message
func GetErrorJSON(details string) string {
    errorMessage := "{\"status\":\"False\",\"details\":\"%s\"}"
    return fmt.Sprintf(errorMessage, details)
}

// GetSuccessJSON prepares form for success message
func GetSuccessJSON(details string) string {
    successMessage := "{\"status\":\"True\",\"details\": %s}"
    return fmt.Sprintf(successMessage, details)
}
