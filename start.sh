#!/bin/sh

set -e

echo "run db migration"
/app/migrate -path /app/migration -database "$PG_CONNSTRING" -verbose up

echo "start the app"
exec "$@"