
-- +migrate Up
CREATE TABLE IF NOT EXISTS `workshop-labti`.`users` (
                                              `id` INT NOT NULL AUTO_INCREMENT,
                                              `username` VARCHAR(45) NOT NULL,
    `password` VARCHAR(45) NOT NULL,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`))
    ENGINE = InnoDB;

-- +migrate Down
