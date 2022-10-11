-- +goose Up
INSERT INTO `daps_categories` (`name`,`international_name`, `description`, `custom`) VALUES ('recordatorios', 'reminders', 'test_description', false);
INSERT INTO `daps_categories` (`name`,`international_name`, `description`, `custom`) VALUES ('tareas', 'tasks', 'test_description', false);
INSERT INTO `daps_categories` (`name`,`international_name`, `description`, `custom`) VALUES ('enseñanza', 'teaching', 'test_description', false);
INSERT INTO `daps_categories` (`name`,`international_name`, `description`, `custom`) VALUES ('hobbies', 'hobbies', 'test_description', false);
INSERT INTO `daps_categories` (`name`,`international_name`, `description`, `custom`) VALUES ('literatura', 'literature', 'test_description', false);
INSERT INTO `daps_categories` (`name`,`international_name`, `description`, `custom`) VALUES ('tecnología', 'technology', 'test_description', false);
INSERT INTO `daps_categories` (`name`,`international_name`, `description`, `custom`) VALUES ('personal', 'personal', 'test_description', false);


-- +goose Down
TRUNCATE TABLE `daps_categories`;
-- +goose StatementBegin
-- +goose StatementEnd
