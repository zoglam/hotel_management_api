package models

// Hotel table
type Hotel struct {
    HotelID     uint    `json:"hotel_id"`
    Discription string  `json:"discription"`
    Price       float32 `json:"price"`
}
