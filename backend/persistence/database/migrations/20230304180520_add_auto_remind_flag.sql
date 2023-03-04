-- +goose Up
ALTER TABLE daps_user_configs ADD auto_remind INT NOT NULL DEFAULT 0;