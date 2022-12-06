-- +goose Up
DELETE FROM `daps_categories` WHERE name IN ('recordatorios', 'tareas', 'enseñanza', 'hobbies', 'literatura', 'tecnología', 'personal');
INSERT INTO `daps_categories` (`id`, `name`,`international_name`, `description`, `custom`) VALUES (1, 'Bugs', 'Bugs', 'User reported bugs', false);
-- +goose Down
TRUNCATE TABLE `daps_categories`;
-- +goose StatementBegin
-- +goose StatementEnd
