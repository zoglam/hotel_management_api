package models

import (
    "database/sql"

    //mysql driver
    _ "github.com/go-sql-driver/mysql"
)

// DB contains connection to mysql database
var DB *sql.DB

// ConnectDataBase makes connect to database
func ConnectDataBase(dsn string) {
    database, err := sql.Open(`mysql`, dsn)

    if err != nil {
        panic("Failed to connect to database!")
    }

    DB = database
}

// CloseConnectionDataBase closes connection to database
func CloseConnectionDataBase() {
    DB.Close()
}
