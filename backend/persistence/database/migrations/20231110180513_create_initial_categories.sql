-- +goose Up
INSERT INTO `daps_categories` (`name`) VALUES ('recordatorios');
INSERT INTO `daps_categories` (`name`) VALUES ('tareas');
INSERT INTO `daps_categories` (`name`) VALUES ('enseñanza');
INSERT INTO `daps_categories` (`name`) VALUES ('hobbies');
INSERT INTO `daps_categories` (`name`) VALUES ('literatura');
INSERT INTO `daps_categories` (`name`) VALUES ('tecnología');
INSERT INTO `daps_categories` (`name`) VALUES ('personal');

-- +goose Down