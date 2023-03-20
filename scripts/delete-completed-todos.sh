#!/bin/bash

# Read the variables for the MySQL connection from the environment
MYSQL_USERNAME=$MYSQL_ROOT_PASSWORD
MYSQL_PASSWORD=$MYSQL_PASSWORD
MYSQL_HOSTNAME=db
MYSQL_DBNAME=$MYSQL_DATABASE

# Run the DELETE query
mysql -u $MYSQL_USERNAME -p$MYSQL_PASSWORD $MYSQL_DBNAME -e 'DELETE FROM daps_todos WHERE completed = true AND recurring = false AND end_date < DATE_SUB(CURDATE(), INTERVAL 5 DAY);'
