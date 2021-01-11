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
    viper.SetConfigFile(`config.json`)
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

    models.ConnectDataBase(dsn)
    defer models.CloseConnectionDataBase()

    r := mux.NewRouter()
    r.HandleFunc("/hotels/create/", controllers.CreateHotel).Methods("POST")
    r.HandleFunc("/hotels/delete/", controllers.DeleteHotel).Methods("POST")
    r.HandleFunc("/hotels/list/", controllers.GetListHotel).Methods("GET")

    r.HandleFunc("/bookings/create/", controllers.CreateBooking).Methods("POST")
    r.HandleFunc("/bookings/delete/", controllers.DeleteBooking).Methods("POST")
    r.HandleFunc("/bookings/list/", controllers.GetListBooking).Methods("GET")

    log.Fatal(http.ListenAndServe(portNumber, r))
}
