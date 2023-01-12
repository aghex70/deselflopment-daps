-- +goose Up
ALTER TABLE daps_users ADD reset_password_code VARCHAR(50) UNIQUE;