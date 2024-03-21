-- +goose Up
UPDATE daps_todos SET recurrency = 0 WHERE recurrency = '';
ALTER TABLE daps_todos MODIFY COLUMN recurrency INT UNSIGNED NULL DEFAULT NULL;

-- +goose Down