
-- +migrate Up
CREATE TABLE IF NOT EXISTS `workshop-labti`.`invoices` (
                                                 `id` INT NOT NULL,
                                                 `order_id` INT NOT NULL,
                                                 `payment_method_id` INT NOT NULL,
                                                 `reference_number` VARCHAR(45) NOT NULL,
    `amount` FLOAT NOT NULL,
    `created_by` INT NOT NULL,
    `updated_by` INT NOT NULL,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE INDEX `reference_number_UNIQUE` (`reference_number` ASC) VISIBLE,
    INDEX `fk_invoices_orders1_idx` (`order_id` ASC) VISIBLE,
    INDEX `fk_invoices_payment_methods1_idx` (`payment_method_id` ASC) VISIBLE,
    CONSTRAINT `fk_invoices_orders1`
    FOREIGN KEY (`order_id`)
    REFERENCES `workshop-labti`.`orders` (`id`)
                                                              ON DELETE NO ACTION
                                                              ON UPDATE NO ACTION,
    CONSTRAINT `fk_invoices_payment_methods1`
    FOREIGN KEY (`payment_method_id`)
    REFERENCES `workshop-labti`.`payment_methods` (`id`)
                                                              ON DELETE NO ACTION
                                                              ON UPDATE NO ACTION)
    ENGINE = InnoDB;

-- +migrate Down
