package controllers

import (
    "encoding/json"
    "net/http"

    "github.com/gorilla/mux"

    models "github.com/zoglam/hotel_management_api/models"
)

// CreateHotelRoom creates hotel_room and returns room_id of new row
func CreateHotelRoom(w http.ResponseWriter, r *http.Request) {
    var err error
    defer func() {
        if err != nil {
            w.WriteHeader(http.StatusBadRequest)
            w.Write([]byte(GetErrorJSON(err.Error())))
        }
    }()

    r.ParseForm()
    discription := r.PostForm.Get("discription")
    price := r.PostForm.Get("price")

    res, err := models.HotelRoomDao.CreateHotelRoom(discription, price)
    if err != nil {
        return
    }
    w.Write([]byte(GetSuccessJSON(res)))
}

// DeleteHotelRoom deletes hotel_room by room_id and returns room_id for details
func DeleteHotelRoom(w http.ResponseWriter, r *http.Request) {
    var err error
    defer func() {
        if err != nil {
            w.WriteHeader(http.StatusBadRequest)
            w.Write([]byte(GetErrorJSON(err.Error())))
        }
    }()

    vars := mux.Vars(r)
    hotelRoomID := vars["hotelRoomID"]

    res, err := models.HotelRoomDao.DeleteHotelRoom(hotelRoomID)
    if err != nil {
        return
    }
    w.Write([]byte(GetSuccessJSON(res)))
}

// GetListHotelRooms returns list of hotel_room-s which are sorted by price/date_created
func GetListHotelRooms(w http.ResponseWriter, r *http.Request) {
    var err error
    defer func() {
        if err != nil {
            w.WriteHeader(http.StatusBadRequest)
            w.Write([]byte(GetErrorJSON(err.Error())))
        }
    }()

    vars := r.URL.Query()
    orderByParam := vars.Get("order_by")
    if orderByParam == "" {
        orderByParam = "price"
    }

    listOfHotelRooms, err := models.HotelRoomDao.GetOrderedByParamHotelRooms(orderByParam)
    if err != nil {
        return
    }

    answer := HTTPHotelRoomResponse{
        Status:  "True",
        Details: listOfHotelRooms,
    }

    js, err := json.Marshal(answer)
    if err != nil {
        return
    }

    w.Write(js)
}
