CREATE TABLE `notes` (
	`index` INT PRIMARY KEY NOT NULL,
	`note` VARCHAR(256) NOT NULL,
	`created` BIGINT(11) NOT NULL,
	`modified` BIGINT(11) NOT NULL
);
