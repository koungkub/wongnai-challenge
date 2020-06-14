CREATE TABLE IF NOT EXISTS `review` (
    `review_id` INT(11),
    `comment` TEXT NOT NULL,
    `created_at` DATETIME DEFAULT current_timestamp(),
    `updated_at` DATETIME DEFAULT current_timestamp() ON UPDATE current_timestamp(),
    PRIMARY KEY (`review_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1000 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS `food_dictionary` (
    `food_id` INT(11) AUTO_INCREMENT,
    `name` VARCHAR(255) NOT NULL,
    `created_at` DATETIME DEFAULT current_timestamp(),
    `updated_at` DATETIME DEFAULT current_timestamp() ON UPDATE current_timestamp(),
    PRIMARY KEY (`food_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1000 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
