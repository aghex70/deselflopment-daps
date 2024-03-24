-- Create the database
CREATE DATABASE IF NOT EXISTS deselflopment;

-- Drop the user from the database
DROP USER 'deselflopment'@'%';

-- Flush privileges
FLUSH PRIVILEGES;

-- Create the user for the database
CREATE USER 'deselflopment'@'%' IDENTIFIED BY 'deselflopment';

-- Grant permissions to the user
GRANT CREATE, ALTER, INDEX, LOCK TABLES, REFERENCES, UPDATE, DELETE, DROP, SELECT, INSERT ON `deselflopment`.* TO 'deselflopment'@'%';

-- Flush privileges
FLUSH PRIVILEGES;
