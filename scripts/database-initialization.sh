#!/bin/bash

# Retrieve user, password, and database name from environment variables
user="$MYSQL_USER"
root_user="$MYSQL_ROOT_USER"
password="$MYSQL_PASSWORD"
database="$MYSQL_DATABASE"

# Create the database
mysql -u "${MYSQL_ROOT_USER}" -p"${MYSQL_ROOT_PASSWORD}" -e "CREATE DATABASE IF NOT EXISTS ${database};"

# Create the user for the database
mysql -u "${MYSQL_ROOT_USER}" -p"${MYSQL_ROOT_PASSWORD}" -e "CREATE USER ''${user}''@''%'' IDENTIFIED BY ''${password}'';"

# Grant permissions to the user
mysql -u "${MYSQL_ROOT_USER}" -p"${MYSQL_ROOT_PASSWORD}" -e "GRANT CREATE, ALTER, INDEX, LOCK TABLES, REFERENCES, UPDATE, DELETE, DROP, SELECT, INSERT ON \`${database}\`.* TO ''${user}''@''%'';"

# Flush privileges
mysql -u "${MYSQL_ROOT_USER}" -p"${MYSQL_ROOT_PASSWORD}" -e "FLUSH PRIVILEGES;"
