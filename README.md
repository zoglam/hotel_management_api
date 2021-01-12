# Contents hotel_management_api
- [Cmd](#Native)
- [Docker](#Docker-compose)
- [API Documentation](#API-Documentation)
    - [`POST` hotel_room/create](#POST-hotel_roomcreate)
    - [`DELETE` hotel_room/delete/<hotelRoomID>](#DELETE-hotel_roomdeletehotelRoomID)
    - [`GET` hotel_room/list](#GET-hotel_roomlist)
    - [`POST` bookings/create](#POST-bookingscreate)
    - [`DELETE` bookings/delete/<bookingID>](#DELETE-bookingsdeletebookingID)
    - [`GET` bookings/list](#GET-bookingslist)
- [Database scheme](#Database-scheme)
# Cmd
Prepare
```
go mod download
```
and setting up config files in ./config

Test
```
...
```
Run
```
go run app/main.go
```
**[⬆ Back to Top](#Contents-hotel_management_api)**
# Docker-compose
Build and Run
```
docker-compose up --build -d
```
Down
```
docker-compose down
```
**[⬆ Back to Top](#Contents-hotel_management_api)**
# API Documentation
## `POST` hotel_room/create
**Метод для добавления номера отеля.**
Принимает на вход текстовое описание и цену за ночь. Возвращает ID номера отеля.


### Parameters

|Name|Description|
|----|-----------|
|discription `*required`|Текстовое описание|
|price `*required`|Цена за ночь|


### Example

Valid
```bash
curl -X POST -d "discription=test_discription" -d "price=99.99" http://localhost:8080/hotel_room/create
```

Invalid
```bash
curl -X POST -d "discription=test_discription" -d "price=99.b99" http://localhost:8080/hotel_room/create
```

### Response

<table>
<tr><td> Status </td> <td> Response </td></tr>
<tr>
<td style="text-align:center"> 200 </td>
<td>
Запрос успешно обработан<br/>
<b>Example Value:</b>

```json
{    
    "status": "True",
    "details": [
        {
            "room_id": "3"
        }
    ]
}
```
</td>
</tr>
<tr>
<td style="text-align:center"> 400 </td>
<td>
Ошибка запроса, если отправлена невалидная цена (e.g. 99.b99)<br/>
<b>Example Value:</b>

```json
{    
    "status": "False",
    "details": "Error 1265: Data truncated for column 'price' at row 1"
}
```
</td>
</tr>
</table>



**[⬆ Back to Top](#Contents-hotel_management_api)**
<hr>

## `DELETE` hotel_room/delete/\<hotelRoomID>
**Метод для удаления номера отеля и всех его броней.**
Принимает на вход ID номера отеля.

### Parameters

<table>
<tr><td> <b>Name</b> </td> <td> <b>Description</b> </td></tr>
<tr>
<td>
hotelRoomID <b>*required</b></td>
<td>
ID номера отеля. Передается в адресной строке
</td>
<tr>
</table>

### Example

Valid
```
curl -X DELETE http://localhost:8080/bookings/delete/99
```

### Response

<table>
<tr><td> Status </td> <td> Response </td></tr>
<tr>    
<td style="text-align:center"> 200 </td>
<td>
Запрос успешно обработан<br/>
<b>Example Value:</b>

```json
{
    "status":"True",
    "details": [
        {
            "booking_id": "99"
        }
    ]
}
```
</td>
</tr>
<tr>
<td style="text-align:center"> 400 </td>
<td>
Ошибка запроса<br/>
<b>Example Value:</b>

```json
{
    "status": "False",
    "details": "Error discription"
}
```
</td>
</tr>
</table>

**[⬆ Back to Top](#Contents-hotel_management_api)**
<hr>

## `GET` hotel_room/list
**Метод для получения всех опросов, которые прошел пользователь, с подробным описанием выбранных ответов**

### Parameters

|Name|Description|
|-|-|
|id_user|ID пользователя|
```
Response
```
<table>
<tr><td> Status </td> <td> Response </td></tr>
<tr>
<td style="text-align:center"> 200 </td>
<td>
Запрос успешно обработан<br/>
<b>Example Value:</b>

```json
{
    "status": "True",
    "details": [
        {
            "hotel_id": 1,
            "discription": "test_discription",
            "price": 10.11,
            "date_created": "2021-01-12"
        },
        {
            "hotel_id": 2,
            "discription": "test_discription",
            "price": 69.24,
            "date_created": "2021-01-12"
        },
        {
            "hotel_id": 3,
            "discription": "test_discription",
            "price": 99.99,
            "date_created": "2021-01-12"
        }
    ]
}
```
</td>
</tr>
<tr>
<td style="text-align:center"> 400 </td>
<td>
Ошибка запроса<br/>
<b>Example Value:</b>

```json
{
    "status": "False",
    "details": "Error discription"
}
```
</td>
</tr>
</table>

**[⬆ Back to Top](#Contents-hotel_management_api)**
<hr>


## `POST` bookings/create
**Метод для добавления брони.** Принимает на вход существующий ID номера отеля, дату начала, дату окончания брони. Возвращает ID брони.

**[⬆ Back to Top](#Contents-hotel_management_api)**
<hr>

## `DELETE` bookings/delete/<bookingID>
**Метод для удаления брони.** Принимает на вход ID брони.


**[⬆ Back to Top](#Contents-hotel_management_api)**
<hr>

## `GET` bookings/list
**Метод для получения списока броней номера отеля.** Принимает на вход ID номера отеля. Возвращает список бронирований, каждое бронирование содержит ID, дату начала, дату окончания. Бронирования отсортированы по дате начала.


**[⬆ Back to Top](#Contents-hotel_management_api)**

# Database scheme

**[⬆ Back to Top](#Contents-hotel_management_api)**