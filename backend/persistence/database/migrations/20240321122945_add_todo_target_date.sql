-- +goose Up
ALTER TABLE daps_todos ADD COLUMN target_date DATE NULL DEFAULT NULL;

-- +goose Down