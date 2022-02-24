
-- +migrate Up
CREATE TABLE IF NOT EXISTS `workshop-labti`.`invoice_details` (
                                                        `id` INT NOT NULL,
                                                        `invoice_id` INT NOT NULL,
                                                        `order_detail_id` INT NOT NULL,
                                                        `product_price` FLOAT NOT NULL,
                                                        `total_amount` FLOAT NOT NULL,
                                                        `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                                        `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                                        PRIMARY KEY (`id`),
    INDEX `fk_invoice_details_invoices1_idx` (`invoice_id` ASC) VISIBLE,
    INDEX `fk_invoice_details_order_details1_idx` (`order_detail_id` ASC) VISIBLE,
    CONSTRAINT `fk_invoice_details_invoices1`
    FOREIGN KEY (`invoice_id`)
    REFERENCES `workshop-labti`.`invoices` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
    CONSTRAINT `fk_invoice_details_order_details1`
    FOREIGN KEY (`order_detail_id`)
    REFERENCES `workshop-labti`.`order_details` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
    ENGINE = InnoDB;

-- +migrate Down
