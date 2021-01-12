package main

import (
    "fmt"
    "log"
    "net/http"
    "net/url"

    "github.com/gorilla/mux"
    "github.com/spf13/viper"

    controllers "github.com/zoglam/hotel_management_api/controllers"
    models "github.com/zoglam/hotel_management_api/models"
)

func init() {
    viper.SetConfigFile(`config/config.json`)
    err := viper.ReadInConfig()
    if err != nil {
        panic(err)
    }
}

func main() {
    portNumber := viper.GetString(`server.address`)
    dbHost := viper.GetString(`database.host`)
    dbPort := viper.GetString(`database.port`)
    dbUser := viper.GetString(`database.user`)
    dbPass := viper.GetString(`database.pass`)
    dbName := viper.GetString(`database.name`)
    connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

    val := url.Values{}
    val.Add("parseTime", "1")
    val.Add("loc", "Europe/Moscow")

    dsn := fmt.Sprintf("%s?%s", connection, val.Encode())

    log.Println("Server started on port:", portNumber)
    defer func() {
        log.Println("Server stopped")
    }()

    models.ConnectDataBase(dsn)
    defer models.CloseConnectionDataBase()

    r := mux.NewRouter()
    r.Headers("Content-Type", "application/json")
    r.NotFoundHandler = http.HandlerFunc(controllers.GetEmptyPage)

    r.HandleFunc("/hotel_room/create", controllers.CreateHotelRoom).Methods("POST")
    r.HandleFunc("/hotel_room/delete/{hotelRoomID:[0-9]+}", controllers.DeleteHotelRoom).Methods("DELETE")
    r.HandleFunc("/hotel_room/list", controllers.GetListHotelRooms).Methods("GET")

    r.HandleFunc("/bookings/create", controllers.CreateBooking).Methods("POST")
    r.HandleFunc("/bookings/delete/{bookingID:[0-9]+}", controllers.DeleteBooking).Methods("DELETE")
    r.HandleFunc("/bookings/list", controllers.GetListOfBookingsForRoomID).Methods("GET")

    log.Fatal(http.ListenAndServe(portNumber, r))
}
