CREATE TABLE IF NOT EXISTS folders (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    chat_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT NULL,
    CONSTRAINT fk_folders_chat FOREIGN KEY (chat_id) REFERENCES chats (id) ON DELETE CASCADE ON UPDATE CASCADE
);