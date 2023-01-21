-- +goose Up
ALTER TABLE daps_todos MODIFY COLUMN description TEXT;