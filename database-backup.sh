#!/bin/bash

# Read the variables for the MySQL connection from the environment
MYSQL_USERNAME=$MYSQL_USERNAME
MYSQL_PASSWORD=$MYSQL_PASSWORD
MYSQL_HOSTNAME=$MYSQL_HOSTNAME
MYSQL_DBNAME=$MYSQL_DBNAME

# Read the destination for the backup file from the environment
BACKUP_DESTINATION="/home/ubuntu/bk/daps"

# Set the filename for the backup file
BACKUP_FILENAME=mysql-backup-$(date +%Y-%m-%d).sql

# Dump the database using mysqldump
mysqldump -u $MYSQL_USERNAME -p$MYSQL_PASSWORD -h $MYSQL_HOSTNAME $MYSQL_DBNAME > $BACKUP_DESTINATION/$BACKUP_FILENAME

# Remove backups older than 7 days
find $BACKUP_DESTINATION -type f -mtime +7 -delete