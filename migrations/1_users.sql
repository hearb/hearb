-- +migrate Up
CREATE TABLE `users` (
	`id`			INT UNSIGNED NOT NULL AUTO_INCREMENT,
	`facebook_id`	BIGINT UNSIGNED NOT NULL,
	`name`			VARCHAR(50) NOT NULL,
	`icon`			VARCHAR(255) NOT NULL,

	`created`		TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`updated`		TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	`deleted`		TIMESTAMP,
	`last_logined`	TIMESTAMP NOT NULL,

	PRIMARY KEY(`id`)
) ENGINE=InnoDB;

CREATE UNIQUE INDEX facebook_id ON `users`(`facebook_id`);
CREATE INDEX user_deleted ON `users`(`deleted`);