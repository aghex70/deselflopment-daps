-- +goose Up
CREATE TABLE daps_user_configs (
    id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
    language VARCHAR(2) NOT NULL,
    auto_suggest INT NOT NULL DEFAULT 0,
    user_id INT NOT NULL REFERENCES daps_users(id) ON DELETE CASCADE
);