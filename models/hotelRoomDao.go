package models

import (
    "fmt"
)

type hotelRoomDao struct{}

// HotelRoomDao - data access object
var HotelRoomDao hotelRoomDaoInterface

func init() {
    HotelRoomDao = &hotelRoomDao{}
}

// CreateHotelRoom makes query to DB to create new hotel_room
func (h *hotelRoomDao) CreateHotelRoom(discription string, price string) (string, error) {
    statement, err := DB.Prepare(`
        INSERT INTO hotel_room(discription, price)
        VALUES(?,?)
    `)
    if err != nil {
        return "", err
    }

    res, err := statement.Exec(discription, price)
    if err != nil {
        return "", err
    }

    lid, err := res.LastInsertId()
    if err != nil {
        return "", err
    }

    successMsgForm := "[{\"room_id\": \"%d\"}]"
    successMsg := fmt.Sprintf(successMsgForm, lid)

    return successMsg, nil
}

// DeleteHotelRoom makes query to DB to delete hotel_room by room_id
func (h *hotelRoomDao) DeleteHotelRoom(roomID string) (string, error) {
    statement, err := DB.Prepare(`
        DELETE FROM hotel_room WHERE room_id=?
    `)
    if err != nil {
        return "", err
    }

    _, err = statement.Exec(roomID)
    if err != nil {
        return "", err
    }

    successMsgForm := "[{\"room_id\": \"%s\"}]"
    successMsg := fmt.Sprintf(successMsgForm, roomID)

    return successMsg, nil
}

// GetOrderedByParamHotelRooms makes query to DB to get sorted hotel_room list
func (h *hotelRoomDao) GetOrderedByParamHotelRooms(orderByParam string) ([]HotelRoom, error) {
    query := `
        SELECT 
            room_id, 
            discription, 
            price, 
            DATE_FORMAT(date_created, "%Y-%m-%d") 
        FROM hotel_room
        ORDER BY ` + fmt.Sprintf(`%s`, orderByParam)
    rows, err := DB.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var hrs []HotelRoom
    for rows.Next() {
        var hr HotelRoom
        err := rows.Scan(&hr.HotelRoomID, &hr.Discription, &hr.Price, &hr.DateCreated)
        if err != nil {
            return nil, err
        }
        hrs = append(hrs, hr)
    }
    if err = rows.Err(); err != nil {
        return nil, err
    }
    return hrs, nil
}
