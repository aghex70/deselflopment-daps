-- +goose Up
ALTER TABLE daps_categories ADD notifiable INT NOT NULL DEFAULT 0;