
-- +migrate Up
CREATE TABLE IF NOT EXISTS `workshop-labti`.`products` (
                                                 `id` INT NOT NULL AUTO_INCREMENT,
                                                 `name` VARCHAR(45) NOT NULL,
    `sku` VARCHAR(45) NOT NULL,
    `stock` INT NOT NULL,
    `price` FLOAT NOT NULL,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`))
    ENGINE = InnoDB;

-- +migrate Down
