-- +goose Up
INSERT INTO `daps_categories` (`name`,`international_name`, `description`) VALUES ('recordatorios', 'reminders', 'test_description');
INSERT INTO `daps_categories` (`name`,`international_name`, `description`) VALUES ('tareas', 'tasks', 'test_description');
INSERT INTO `daps_categories` (`name`,`international_name`, `description`) VALUES ('enseñanza', 'teaching', 'test_description');
INSERT INTO `daps_categories` (`name`,`international_name`, `description`) VALUES ('hobbies', 'hobbies', 'test_description');
INSERT INTO `daps_categories` (`name`,`international_name`, `description`) VALUES ('literatura', 'literature', 'test_description');
INSERT INTO `daps_categories` (`name`,`international_name`, `description`) VALUES ('tecnología', 'technology', 'test_description');
INSERT INTO `daps_categories` (`name`,`international_name`, `description`) VALUES ('personal', 'personal', 'test_description');


-- +goose Down
TRUNCATE TABLE `daps_categories`;
-- +goose StatementBegin
-- +goose StatementEnd
