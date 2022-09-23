-- +goose Up
INSERT INTO `daps_categories` (`name`,`international_name`) VALUES ('recordatorios', 'reminders');
INSERT INTO `daps_categories` (`name`,`international_name`) VALUES ('tareas', 'tasks');
INSERT INTO `daps_categories` (`name`,`international_name`) VALUES ('enseñanza', 'teaching');
INSERT INTO `daps_categories` (`name`,`international_name`) VALUES ('hobbies', 'hobbies');
INSERT INTO `daps_categories` (`name`,`international_name`) VALUES ('literatura', 'literature');
INSERT INTO `daps_categories` (`name`,`international_name`) VALUES ('tecnología', 'technology');
INSERT INTO `daps_categories` (`name`,`international_name`) VALUES ('personal', 'personal');


-- +goose Down
TRUNCATE TABLE `daps_categories`;
-- +goose StatementBegin
-- +goose StatementEnd
