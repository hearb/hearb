-- +migrate Up
CREATE TABLE `sessions` (
	`sid`			CHAR(44) NOT NULL,
	`user_id`		INT UNSIGNED,

	`created`		TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`updated`		TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

	PRIMARY KEY(`sid`)
) ENGINE=InnoDB;
