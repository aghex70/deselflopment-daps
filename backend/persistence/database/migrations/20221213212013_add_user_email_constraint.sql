-- +goose Up
ALTER TABLE deselflopment_users ADD UNIQUE (email);
