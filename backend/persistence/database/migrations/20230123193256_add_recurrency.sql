-- +goose Up
ALTER TABLE daps_todos ADD recurrency VARCHAR(25) DEFAULT '' NOT NULL;