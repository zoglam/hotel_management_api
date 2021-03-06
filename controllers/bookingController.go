package controllers

import (
    "encoding/json"
    "errors"
    "net/http"
    "strconv"

    "github.com/gorilla/mux"

    models "github.com/zoglam/hotel_management_api/models"
)

// CreateBooking ...
func CreateBooking(w http.ResponseWriter, r *http.Request) {

    var err error
    defer func() {
        if err != nil {
            w.WriteHeader(http.StatusBadRequest)
            w.Write([]byte(GetErrorJSON(err.Error())))
        }
    }()

    r.ParseForm()
    dateStart := r.PostForm.Get("date_start")
    dateEnd := r.PostForm.Get("date_end")
    roomID := r.PostForm.Get("room_id")

    if !IsValidDate(dateStart) || !IsValidDate(dateEnd) {
        err = errors.New("Invalid format date")
        return
    }

    res, err := models.BookingDao.CreateBooking(dateStart, dateEnd, roomID)
    if err != nil {
        return
    }
    w.Write([]byte(GetSuccessJSON(res)))
}

// DeleteBooking ...
func DeleteBooking(w http.ResponseWriter, r *http.Request) {
    var err error
    defer func() {
        if err != nil {
            w.WriteHeader(http.StatusBadRequest)
            w.Write([]byte(GetErrorJSON(err.Error())))
        }
    }()

    vars := mux.Vars(r)
    bookingID := vars["bookingID"]

    res, err := models.BookingDao.DeleteBooking(bookingID)
    if err != nil {
        return
    }
    w.Write([]byte(GetSuccessJSON(res)))
}

// GetListOfBookingsForRoomID ...
func GetListOfBookingsForRoomID(w http.ResponseWriter, r *http.Request) {
    var err error
    defer func() {
        if err != nil {
            w.WriteHeader(http.StatusBadRequest)
            w.Write([]byte(GetErrorJSON(err.Error())))
        }
    }()

    vars := r.URL.Query()
    roomID := vars.Get("room_id")
    i, err := strconv.ParseUint(roomID, 10, 32)
    if err != nil {
        return
    }

    listOfBookings, err := models.BookingDao.GetBookingsByRoomID(uint(i))
    if err != nil {
        return
    }

    answer := HTTPBookingResponse{
        Status:  "True",
        Details: listOfBookings,
    }

    js, err := json.Marshal(answer)
    if err != nil {
        return
    }

    w.Write(js)
}
