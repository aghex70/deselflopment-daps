#!/bin/bash

BACKUP_HOST_PATH="/home/ubuntu/bk/daps"
BACKUP_CONTAINER_PATH="/bk/daps"

docker-compose exec db mkdir -p $BACKUP_CONTAINER_PATH

docker cp /home/ubuntu/daps/database-backup.sh daps_db_1:/database-backup.sh

# Run the backup script in the container
docker-compose exec db /database-backup.sh

# Copy generated file from the container to the host
docker cp daps_db_1:$BACKUP_CONTAINER_PATH/daps-backup-$(date +%Y-%m-%d).sql $BACKUP_HOST_PATH/daps-backup-$(date +%Y-%m-%d).sql

# Remove backups older than 7 days from the host
find $BACKUP_HOST_PATH -type f -mtime +7 -delete
