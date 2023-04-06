-- Create the database
CREATE DATABASE IF NOT EXISTS example_database;

-- Drop the user from the database (to avoid errors when creating the user)
DROP USER 'example_user'@'%';

-- Flush privileges
FLUSH PRIVILEGES;

-- Create the user for the database
CREATE USER 'example_user'@'%' IDENTIFIED BY 'example_password';

-- Grant permissions to the user
GRANT CREATE, ALTER, INDEX, LOCK TABLES, REFERENCES, UPDATE, DELETE, DROP, SELECT, INSERT ON `example_database`.* TO 'example_user'@'%';

-- Flush privileges
FLUSH PRIVILEGES;
