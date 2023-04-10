-- +goose Up
ALTER TABLE deselflopment_users ADD reset_password_code VARCHAR(50) UNIQUE;
