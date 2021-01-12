package models

import (
    "fmt"
)

type bookingDao struct{}

// BookingDao - data access object
var BookingDao bookingDaoInterface

func init() {
    BookingDao = &bookingDao{}
}

// CreateBooking makes query to DB to create new booking
func (b *bookingDao) CreateBooking(dateStart string, dateEnd string, roomID string) (string, error) {
    statement, err := DB.Prepare(`
        INSERT INTO booking(date_start, date_end, room_id)
        VALUES(?,?,?)
    `)
    if err != nil {
        return "", err
    }

    res, err := statement.Exec(dateStart, dateEnd, roomID)
    if err != nil {
        return "", err
    }

    lid, err := res.LastInsertId()
    if err != nil {
        return "", err
    }

    successMsgForm := "[{\"booking_id\": \"%d\"}]"
    successMsg := fmt.Sprintf(successMsgForm, lid)

    return successMsg, nil
}

// DeleteBooking makes query to DB to delete booking
func (b *bookingDao) DeleteBooking(bookingID string) (string, error) {
    statement, err := DB.Prepare(`
        DELETE FROM booking WHERE booking_id=?
    `)
    if err != nil {
        return "", err
    }

    _, err = statement.Exec(bookingID)
    if err != nil {
        return "", err
    }

    successMsgForm := "[{\"booking_id\": \"%s\"}]"
    successMsg := fmt.Sprintf(successMsgForm, bookingID)

    return successMsg, nil
}

// GetBookingsByRoomID allows to get bookings list by room_id
func (b *bookingDao) GetBookingsByRoomID(RoomID uint) ([]Booking, error) {

    rows, err := DB.Query(`
        SELECT 
            booking_id, 
            DATE_FORMAT(date_start, "%Y-%m-%d"), 
            DATE_FORMAT(date_end, "%Y-%m-%d"), 
            room_id FROM booking 
        WHERE room_id=? 
        ORDER BY date_start
    `, RoomID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var bks []Booking
    for rows.Next() {
        var bk Booking
        err := rows.Scan(&bk.BookingID, &bk.DateStart, &bk.DateEnd, &bk.RoomID)
        if err != nil {
            return nil, err
        }
        bks = append(bks, bk)
    }
    if err = rows.Err(); err != nil {
        return nil, err
    }
    return bks, nil
}
