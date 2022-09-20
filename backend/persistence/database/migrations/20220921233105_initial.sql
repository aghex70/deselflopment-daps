-- +goose Up
DROP TABLE IF EXISTS daps_todos;
DROP TABLE IF EXISTS daps_users;

CREATE TABLE daps_users (
    id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
    email VARCHAR(128) NOT NULL,
    name VARCHAR(128) NOT NULL
);

CREATE TABLE daps_todos (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
--     user_id INT NOT NULL UNIQUE,
--     prerequisite_id INT NULL DEFAULT NULL,
    active INT NOT NULL DEFAULT 0,
    end_date TIMESTAMP NULL DEFAULT NULL,
    category varchar(128),
    completed INT NOT NULL DEFAULT 0,
    creation_date TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
    description VARCHAR(128),
    duration VARCHAR(128),
    link VARCHAR(128),
    name VARCHAR(128),
    priority INT DEFAULT 2,
    start_date TIMESTAMP NULL DEFAULT NULL
--     FOREIGN KEY(user_id) REFERENCES daps_users(id),
--     FOREIGN KEY(prerequisite_id) REFERENCES daps_todos(id)
);

-- +goose Down
DROP TABLE IF EXISTS daps_todos;
DROP TABLE IF EXISTS daps_users;
-- +goose StatementBegin
-- +goose StatementEnd
