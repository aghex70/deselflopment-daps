-- +goose Up
ALTER TABLE daps_todos MODIFY COLUMN link TEXT;