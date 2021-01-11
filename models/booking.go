package models

// Booking table
type Booking struct {
    BookingID uint   `json:"booking_id"`
    DateStart string `json:"date_start"`
    DateEnd   string `json:"date_end"`
    HotelID   int    `json:"hotel_id"`
}
