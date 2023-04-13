-- +goose Up
ALTER TABLE deselflopment_emails ADD source VARCHAR(128) NULL DEFAULT 'deselflopment';
