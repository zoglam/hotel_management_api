package models

// Booking table
type Booking struct {
    BookingID uint   `json:"booking_id"`
    DateStart string `json:"date_start"`
    DateEnd   string `json:"date_end"`
    RoomID    uint   `json:"-"`
}

type bookingDaoInterface interface {
    CreateBooking(dateStart string, dateEnd string, roomID string) (string, error)
    DeleteBooking(bookingID string) (string, error)
    GetBookingsByRoomID(RoomID uint) ([]Booking, error)
}
