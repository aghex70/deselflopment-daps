-- +goose Up
ALTER TABLE daps_notes DROP FOREIGN KEY daps_notes_ibfk_1;
ALTER TABLE daps_notes DROP COLUMN subtopic_id;

-- +goose Down