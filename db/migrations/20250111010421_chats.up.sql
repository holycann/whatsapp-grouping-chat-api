CREATE TABLE IF NOT EXITS chats(
    `id`  INT UNSIGNED NOT NULL AUTO_INCREMENT
    `user_id` INT UNSIGNED NOT NULL,
    `message` TEXT NOT NULL,
    `created_at` TIMESTAMP,
    `updated_at` TIMESTAMP
)