
-- +migrate Up
CREATE TABLE IF NOT EXISTS `workshop-labti`.`order_details` (
                                                      `id` INT NOT NULL,
                                                      `order_id` INT NOT NULL,
                                                      `product_id` INT NOT NULL,
                                                      `quantity` INT NOT NULL,
                                                      `created_by` INT NOT NULL,
                                                      `updated_by` INT NOT NULL,
                                                      `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                                      `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                                      PRIMARY KEY (`id`),
    INDEX `fk_order_details_orders_idx` (`order_id` ASC) VISIBLE,
    INDEX `fk_order_details_products1_idx` (`product_id` ASC) VISIBLE,
    CONSTRAINT `fk_order_details_orders`
    FOREIGN KEY (`order_id`)
    REFERENCES `workshop-labti`.`orders` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
    CONSTRAINT `fk_order_details_products1`
    FOREIGN KEY (`product_id`)
    REFERENCES `workshop-labti`.`products` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
    ENGINE = InnoDB;

-- +migrate Down
