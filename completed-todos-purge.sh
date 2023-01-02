#!/bin/bash

docker cp /home/ubuntu/daps/delete-completed-todos.sh daps_db_1:/delete-completed-todos.sh

# Run the backup script in the container
docker-compose exec db /delete-completed-todos.sh
