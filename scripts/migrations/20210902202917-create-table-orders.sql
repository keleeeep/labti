
-- +migrate Up
CREATE TABLE IF NOT EXISTS `workshop-labti`.`orders` (
                                               `id` INT NOT NULL,
                                               `reference_number` VARCHAR(45) NOT NULL,
    `created_by` INT NOT NULL,
    `updated_by` INT NOT NULL,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE INDEX `reference_number_UNIQUE` (`reference_number` ASC) VISIBLE)
    ENGINE = InnoDB;

-- +migrate Down
