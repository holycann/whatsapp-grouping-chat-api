CREATE TABLE IF NOT EXISTS folders(
    `id`  INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(50) NOT NULL,
    `chat_id` INT UNSIGNED NOT NULL,
    `created_at` TIMESTAMP,
    `updated_at` TIMESTAMP
)