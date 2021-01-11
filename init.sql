CREATE DATABASE IF NOT EXISTS `hotel_management`
USE `hotel_management`;

DROP TABLE IF EXISTS `HOTEL`;
CREATE TABLE `HOTEL`(
    `hotel_id` INT(11) NOT NULL AUTO_INCREMENT,
    `discription` varchat(45),
    `price` FLOAT(10,2),
    PRIMARY KEY (`hotel_id`)
);

DROP TABLE IF EXISTS `BOOKING`;
CREATE TABLE `BOOKING`(
    `booking_id` INT(11) NOT NULL AUTO_INCREMENT,
    `date_start` datetime DEFAULT NULL,
    `date_end` datetime DEFAULT NULL,
    `hotel_id` INT(11),
    PRIMARY KEY (`booking_id`),
    FOREIGN KEY (`hotel_id`) REFERENCES `HOTEL`(`hotel_id`)
);