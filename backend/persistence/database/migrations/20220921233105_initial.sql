-- +goose Up
-- DROP TABLE IF EXISTS daps_todos;
-- DROP TABLE IF EXISTS daps_users;
-- DROP TABLE IF EXISTS daps_categories;

CREATE TABLE daps_users (
    id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
    email VARCHAR(128) NOT NULL,
    is_admin INT NOT NULL DEFAULT 0,
    registration_date TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
    password VARCHAR(128) NOT NULL
);


CREATE TABLE daps_categories (
     id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
     custom INT NOT NULL DEFAULT 0,
     description VARCHAR(128),
     name VARCHAR(128),
     international_name VARCHAR(128)
);

CREATE TABLE daps_todos (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    category_id INT NOT NULL DEFAULT 2,
--     prerequisite_id INT NULL DEFAULT NULL,
    active INT NOT NULL DEFAULT 0,
    end_date TIMESTAMP NULL DEFAULT NULL,
    completed INT NOT NULL DEFAULT 0,
    recurring INT NOT NULL DEFAULT 0,
    creation_date TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
    description VARCHAR(128),
    duration VARCHAR(128),
    link VARCHAR(128),
    name VARCHAR(128),
    priority INT DEFAULT 2,
    start_date TIMESTAMP NULL DEFAULT NULL,
    suggested INT NOT NULL DEFAULT 0,
    suggestion_date TIMESTAMP NULL DEFAULT NULL,
    FOREIGN KEY(user_id) REFERENCES daps_users(id),
    FOREIGN KEY(category_id) REFERENCES daps_categories(id)
--     FOREIGN KEY(prerequisite_id) REFERENCES daps_todos(id)
);

-- +goose Down
-- DROP TABLE IF EXISTS daps_todos;
-- DROP TABLE IF EXISTS daps_users;
-- DROP TABLE IF EXISTS daps_categories;
-- +goose StatementBegin
-- +goose StatementEnd
