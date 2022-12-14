-- +goose Up
ALTER TABLE daps_users ADD UNIQUE (email);