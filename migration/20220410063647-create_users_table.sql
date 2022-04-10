
-- +migrate Up
CREATE TABLE IF NOT EXISTS `users` (
  `id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,  
  `email` VARCHAR(100) NOT NULL,
  `name` VARCHAR(20) NOT NULL,
  `age` int(10) UNSIGNED,
  `birthday` DATETIME,
  `member_number` VARCHAR(100),
  `created_at` DATETIME NOT NULL,
  `updated_at` DATETIME NOT NULL,
  PRIMARY KEY (`id`),
  CONSTRAINT email_unique UNIQUE(email)
)ENGINE = InnoDB DEFAULT CHARSET=utf8mb4;

-- +migrate Down
DROP TABLE IF EXISTS `users`;