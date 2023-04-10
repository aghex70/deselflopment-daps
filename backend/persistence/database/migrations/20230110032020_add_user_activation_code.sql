-- +goose Up
ALTER TABLE deselflopment_users ADD activation_code VARCHAR(50) UNIQUE;
ALTER TABLE deselflopment_users ADD active INT NOT NULL DEFAULT 0;
