-- +goose Up
DROP TABLE IF EXISTS daps_category_users;
DROP TABLE IF EXISTS daps_todos;
DROP TABLE IF EXISTS daps_categories;
DROP TABLE IF EXISTS deselflopment_users;

CREATE TABLE deselflopment_users (
    id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(128),
    email VARCHAR(128) NOT NULL,
    admin INT NOT NULL DEFAULT 0,
    registration_date TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
    password VARCHAR(128) NOT NULL
);


CREATE TABLE daps_categories (
     id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
     owner_id INT NULL DEFAULT NULL,
     custom INT NOT NULL DEFAULT 0,
     description VARCHAR(128),
     name VARCHAR(128),
     international_name VARCHAR(128),
     shared INT NOT NULL DEFAULT 0,
     FOREIGN KEY(owner_id) REFERENCES deselflopment_users(id)
);

CREATE TABLE daps_category_users (
     category_id INT NOT NULL REFERENCES daps_categories(id) ON DELETE CASCADE,
     user_id INT NOT NULL REFERENCES deselflopment_users(id) ON DELETE CASCADE,
     PRIMARY KEY (category_id, user_id)
);

CREATE TABLE daps_todos (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    category_id INT NOT NULL,
    active INT NOT NULL DEFAULT 0,
    end_date TIMESTAMP NULL DEFAULT NULL,
    completed INT NOT NULL DEFAULT 0,
    recurring INT NOT NULL DEFAULT 0,
    creation_date TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
    description VARCHAR(128),
    link VARCHAR(128),
    name VARCHAR(128),
    priority INT DEFAULT 1,
    start_date TIMESTAMP NULL DEFAULT NULL,
    suggested INT NOT NULL DEFAULT 0,
    suggestion_date TIMESTAMP NULL DEFAULT NULL,
    FOREIGN KEY(category_id) REFERENCES daps_categories(id)
);

-- +goose Down
-- DROP TABLE IF EXISTS daps_categories_users_relationships;
-- DROP TABLE IF EXISTS daps_categories;
-- DROP TABLE IF EXISTS deselflopment_users;
-- DROP TABLE IF EXISTS daps_todos;
-- +goose StatementBegin
-- +goose StatementEnd
