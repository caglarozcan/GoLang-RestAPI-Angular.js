CREATE DATABASE `golang`;


CREATE TABLE `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `firstname` varchar(45) NOT NULL,
  `lastname` varchar(45) NOT NULL,
  `email` varchar(45) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8_general_ci;


DELIMITER $$
CREATE PROCEDURE `userlist`()
BEGIN

	SELECT * FROM users;

END$$
DELIMITER ;
