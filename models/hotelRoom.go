package models

// HotelRoom table
type HotelRoom struct {
    HotelRoomID uint    `json:"hotel_id"`
    Discription string  `json:"discription"`
    Price       float32 `json:"price"`
    DateCreated string  `json:"date_created"`
}

type hotelRoomDaoInterface interface {
    CreateHotelRoom(discription string, price string) (string, error)
    DeleteHotelRoom(roomID string) (string, error)
    GetOrderedByParamHotelRooms(orderByParam string) ([]HotelRoom, error)
}
