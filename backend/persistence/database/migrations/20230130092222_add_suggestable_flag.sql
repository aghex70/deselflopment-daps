-- +goose Up
ALTER TABLE daps_todos ADD suggestable INT NOT NULL DEFAULT 1;