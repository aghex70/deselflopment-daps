-- +goose Up
DROP TABLE IF EXISTS daps_emails;

CREATE TABLE daps_emails (
    id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
    subject VARCHAR(128) NOT NULL,
    body TEXT NOT NULL,
    creation_date TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
    sent INT NOT NULL DEFAULT 0,
    error TEXT,
    user_id INT NOT NULL REFERENCES deselflopment_users(id) ON DELETE CASCADE
);

-- +goose Down
-- DROP TABLE IF EXISTS daps_emails;
