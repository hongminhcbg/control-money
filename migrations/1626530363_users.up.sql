DROP TABLE IF EXISTS `users`;

CREATE TABLE `users` (
                         `id` BIGINT(20) PRIMARY KEY AUTO_INCREMENT,
                         `name` VARCHAR(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
                         `username` VARCHAR(64) UNIQUE,
                         `password` VARCHAR(64),
                         `money` INT(11),
                         `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                         `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                         INDEX user_name (`username`)
);