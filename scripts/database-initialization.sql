-- Create the database
CREATE DATABASE IF NOT EXISTS database;

-- Create the user for the database
CREATE USER 'user'@'%' IDENTIFIED BY 'password';

-- Grant permissions to the user
GRANT CREATE, ALTER, INDEX, LOCK TABLES, REFERENCES, UPDATE, DELETE, DROP, SELECT, INSERT ON database.* TO 'user'@'%';

-- Flush privileges
FLUSH PRIVILEGES;
