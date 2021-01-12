# Contents hotel_management_api
- [Contents hotel_management_api](#contents-hotel_management_api)
- [Cmd](#cmd)
- [Docker-compose](#docker-compose)
- [API Documentation](#api-documentation)
  - [`POST` hotel_room/create](#post-hotel_roomcreate)
  - [`DELETE` hotel_room/delete/<room_id>](#delete-hotel_roomdeleteroom_id)
  - [`GET` hotel_room/list](#get-hotel_roomlist)
  - [`POST` bookings/create](#post-bookingscreate)
  - [`DELETE` bookings/delete/<booking_id>](#delete-bookingsdeletebooking_id)
  - [`GET` bookings/list](#get-bookingslist)
- [Database scheme](#database-scheme)
# Cmd
Prepare
```sh
go mod download
```
and setting up config files in ./config

Test
```sh
...
```
Run
```sh
go run app/main.go
```
**[⬆ Back to Top](#Contents-hotel_management_api)**
# Docker-compose

Build and Run
```sh
docker-compose up --build -d
```
Down
```sh
docker-compose down
```
**[⬆ Back to Top](#Contents-hotel_management_api)**
# API Documentation
## `POST` hotel_room/create
**Метод для добавления номера отеля.**
Принимает на вход текстовое описание и цену за ночь. Возвращает ID номера отеля.


### Parameters

| Name                    | Description        |
| ----------------------- | ------------------ |
| discription `*required` | Текстовое описание |
| price `*required`       | Цена за ночь       |


### Example

Valid
```sh
curl -X POST -d "discription=test_discription" -d "price=99.99" http://localhost:8080/hotel_room/create
```

Invalid
```sh
curl -X POST -d "discription=test_discription" -d "price=99.b99" http://localhost:8080/hotel_room/create
```

### Response

<table>
    <tr>
        <th> Status </th> 
        <th> Response </th>
    </tr>
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

## `DELETE` hotel_room/delete/<room_id>
**Метод для удаления номера отеля и всех его броней.**
Принимает на вход ID номера отеля.

### Parameters

<table>
    <tr>
        <th> Name </th>
        <th> Description </th>
    </tr>
    <tr>
        <td> room_id <code>*required</code> </td>
        <td> ID номера отеля. Передается в адресной строке </td>
    <tr>
</table>

### Example

Valid
```sh
curl -X DELETE http://localhost:8080/hotel_room/delete/99
```

### Response

<table>
<tr>
    <th> Status </th>
    <th> Response </th>
</tr>
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
**Метод для получения списка номеров отеля.** Есть возможность отсортировать по цене или по дате добавления (по возрастанию и убыванию)

### Parameters

| Name     | Description                                                                                                      |
| -------- | ---------------------------------------------------------------------------------------------------------------- |
| order_by | Параметр по которому сортировать. Возможные варианты: `price`, `price DESC`, `data_created`, `data_created DESC` |

### Example

Valid
```sh
curl http://localhost:8080/hotel_room/list?order_by=price
```

Invalid
```sh
curl http://localhost:8080/hotel_room/list?order_by=udmurtia
```

### Response

<table>
<tr>
    <th> Status </th> 
    <th> Response </th>
</tr>
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
    "details": "Error 1054: Unknown column 'udmurtia' in 'order clause'"
}
```
</td>
</tr>
</table>

**[⬆ Back to Top](#Contents-hotel_management_api)**
<hr>


## `POST` bookings/create
**Метод для добавления брони.** Принимает на вход существующий ID номера отеля, дату начала, дату окончания брони. Возвращает ID брони.

### Parameters

| Name                   | Description                  |
| ---------------------- | ---------------------------- |
| room_id `*required`    | существующий ID номера отеля |
| date_start `*required` | дата начала брони            |
| date_end `*required`   | дата окончания брони         |

### Example

```sh
curl -X POST -d "room_id=1" -d "date_start=2021-12-30" -d "date_end=2022-01-02" http://localhost:8080/bookings/create
```

### Response

<table>
<tr>
    <th> Status </th>
    <th> Response </th>
</tr>
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
            "booking_id": "1"
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

## `DELETE` bookings/delete/<booking_id>
**Метод для удаления брони.** Принимает на вход ID брони.

### Parameters

| Name                   | Description                            |
| ---------------------- | -------------------------------------- |
| booking_id `*required` | ID брони. Передается в адресной строке |

### Example

```sh
curl -X DELETE http://localhost:8080/bookings/delete/99
```

### Response

<table>
<tr>
    <th> Status </th>
    <th> Response </th>
</tr>
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

## `GET` bookings/list
**Метод для получения списока броней номера отеля.** Принимает на вход ID номера отеля. Возвращает список бронирований, каждое бронирование содержит ID, дату начала, дату окончания. Бронирования отсортированы по дате начала.

### Parameters

| Name                | Description     |
| ------------------- | --------------- |
| room_id `*required` | ID номера отеля |

### Example

Valid
```sh
curl http://localhost:8080/bookings/list?room_id=1
```

Invalid
```sh
curl http://localhost:8080/bookings/list?room_id=word
```

### Response

<table>
<tr>
    <th> Status </th>
    <th> Response </th>
</tr>
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
            "booking_id": 2,
            "date_start": "2020-05-10",
            "date_end": "2020-10-10"
        },
        {
            "booking_id": 1,
            "date_start": "2020-06-10",
            "date_end": "2020-07-10"
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
    "details": "strconv.ParseUint: parsing \"word\": invalid syntax"
}
```
</td>
</tr>
</table>


**[⬆ Back to Top](#Contents-hotel_management_api)**

# Database scheme

[<img src="https://live.staticflickr.com/65535/50827301056_0122698ce8_h.jpg" width=900>](https://live.staticflickr.com/65535/50827301056_0122698ce8_h.jpg)

**[⬆ Back to Top](#Contents-hotel_management_api)**