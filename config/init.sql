CREATE DATABASE IF NOT EXISTS `hotel_management`;
USE `hotel_management`;

DROP TABLE IF EXISTS `hotel_room`;
CREATE TABLE `hotel_room`(
    `room_id` INT(11) NOT NULL AUTO_INCREMENT,
    `discription` varchar(255),
    `price` FLOAT(10,2),
    `date_created` datetime DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`room_id`)
);

DROP TABLE IF EXISTS `booking`;
CREATE TABLE `booking`(
    `booking_id` INT(11) NOT NULL AUTO_INCREMENT,
    `date_start` datetime DEFAULT NULL,
    `date_end` datetime DEFAULT NULL,
    `room_id` INT(11),
    PRIMARY KEY (`booking_id`),
    FOREIGN KEY (`room_id`) 
        REFERENCES `hotel_room`(`room_id`) 
        ON DELETE CASCADE
);

insert into `hotel_room`(discription, price) values("lolkek", 12.12);
insert into `booking`(date_start,date_end,room_id) values("2020.06.10", "2020.07.10", 1);
insert into `booking`(date_start,date_end,room_id) values("2020.05.10", "2020.10.10", 1);