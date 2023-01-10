-- +goose Up
ALTER TABLE daps_users ADD activation_code VARCHAR(50) UNIQUE;
ALTER TABLE daps_users ADD active INT NOT NULL DEFAULT 0;