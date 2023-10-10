CREATE TABLE `companies` (
    `company_id` VARCHAR(255) NOT NULL,
    `name` VARCHAR(255) NOT NULL,
    `representative_name` VARCHAR(255) NOT NULL,
    `phone_number` VARCHAR(20) NOT NULL,
    `postal_code` VARCHAR(10) NOT NULL,
    `address` VARCHAR(255) NOT NULL,
    `created_at` DATETIME DEFAULT NULL,
    `updated_at` DATETIME DEFAULT NULL,
    `deleted_at` DATETIME DEFAULT NULL,
    PRIMARY KEY (`company_id`)
);

CREATE TABLE `users` (
    `user_id` VARCHAR(255) NOT NULL,
    `company_id` VARCHAR(255) NOT NULL,
    `name` VARCHAR(255) NOT NULL,
    `email` VARCHAR(255) NOT NULL,
    `password` VARCHAR(255) NOT NULL,
    `created_at` DATETIME DEFAULT NULL,
    `updated_at` DATETIME DEFAULT NULL,
    `deleted_at` DATETIME DEFAULT NULL,
    PRIMARY KEY (`user_id`),
    FOREIGN KEY (`company_id`) REFERENCES `companies`(`company_id`),
    UNIQUE (`email`)
);

CREATE TABLE `clients` (
    `company_id` VARCHAR(255) NOT NULL,
    `client_id` VARCHAR(255) NOT NULL,
    PRIMARY KEY (`company_id`, `client_id`),
    FOREIGN KEY (`company_id`) REFERENCES `companies`(`company_id`),
    FOREIGN KEY (`client_id`) REFERENCES `companies`(`company_id`)
);

CREATE TABLE `company_banks` (
    `company_bank_id` VARCHAR(255) NOT NULL,
    `company_id` VARCHAR(255) NOT NULL,
    `bank_code` VARCHAR(255) NOT NULL,
    `branch_code` VARCHAR(255) NOT NULL,
    `bank_name` VARCHAR(255) NOT NULL,
    `branch_name` VARCHAR(255) NOT NULL,
    `account_type` VARCHAR(255) NOT NULL,
    `account_number` VARCHAR(20) NOT NULL,
    `account_holder` VARCHAR(255) NOT NULL,
    `created_at` DATETIME DEFAULT NULL,
    `updated_at` DATETIME DEFAULT NULL,
    `deleted_at` DATETIME DEFAULT NULL,
    PRIMARY KEY (`company_bank_id`),
    FOREIGN KEY (`company_id`) REFERENCES `companies`(`company_id`)
);

CREATE TABLE `invoices` (
    `invoice_id` VARCHAR(255) NOT NULL,
    `company_id` VARCHAR(255) NOT NULL,
    `client_id` VARCHAR(255) NOT NULL,
    `payment_amount` decimal(10, 2) NOT NULL,
    `fee` decimal(10, 2) NOT NULL,
    `fee_rate` decimal(5, 2) NOT NULL,
    `tax` decimal(10, 2) NOT NULL,
    `tax_rate` decimal(5, 2) NOT NULL,
    `invoice_amount` decimal(10, 2) NOT NULL,
    -- `status` VARCHAR(20) NOT NULL,
    `due_at` DATETIME NOT NULL,
    `created_at` DATETIME DEFAULT NULL,
    `updated_at` DATETIME DEFAULT NULL,
    `deleted_at` DATETIME DEFAULT NULL,
    PRIMARY KEY (`invoice_id`),
    FOREIGN KEY (`company_id`) REFERENCES `companies`(`company_id`),
    FOREIGN KEY (`client_id`) REFERENCES `companies`(`company_id`)
);

CREATE TABLE `invoice_status_logs` (
    `invoice_id` VARCHAR(255) NOT NULL,
    `user_id` VARCHAR(255) NOT NULL,
    `status` VARCHAR(20) NOT NULL,
    `created_at` DATETIME DEFAULT NULL,
    `updated_at` DATETIME DEFAULT NULL,
    `deleted_at` DATETIME DEFAULT NULL,
    PRIMARY KEY (`invoice_id`, `status`),
    FOREIGN KEY (`invoice_id`) REFERENCES `invoices`(`invoice_id`),
    FOREIGN KEY (`user_id`) REFERENCES `users`(`user_id`)
);