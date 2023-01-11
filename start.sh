#!/bin/sh

# script exits imidiately if s command returns not 0 status
set -e

#echo "run db migrations"
/app/migrate -path /app/migrations -database "$DATABASE_URL" -verbose up

echo "start the app"
# take all the params passed to the script and run it
# in our case it's what in CMD of Dockerfile
exec "$@"
