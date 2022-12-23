#!/bin/bash

# Read the variables for the MySQL connection from the environment
MYSQL_USERNAME=$MYSQL_ROOT_PASSWORD
MYSQL_PASSWORD=$MYSQL_PASSWORD
MYSQL_HOSTNAME=mysql
MYSQL_DBNAME=$MYSQL_DBNAME

# Read the destination for the backup file from the environment
BACKUP_DESTINATION="/bk/daps"

# Set the filename for the backup file
BACKUP_FILENAME=daps-backup-$(date +%Y-%m-%d).sql

# Dump the database using mysqldump
mysqldump -u $MYSQL_USERNAME -p$MYSQL_PASSWORD -h $MYSQL_HOSTNAME $MYSQL_DBNAME > $BACKUP_FILENAME

# Remove backups older than 7 days from the host
find $BACKUP_DESTINATION -type f -mtime +7 -delete