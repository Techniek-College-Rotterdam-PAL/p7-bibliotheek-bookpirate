#!/bin/bash

pg_isready -q || {
    echo "PostgreSQL server is not running"
    exit 1
}

PG_USER="root"
PG_PASSWORD=""
PG_HOST="127.0.0.1"
PG_PORT="5432"
DB_NAME="librarydb"

psql -h "$PG_HOST" -p "$PG_PORT" -U "$PG_USER" -c "CREATE DATABASE $DB_NAME;" || {
    echo "Error creating database $DB_NAME"
    exit 1
}

echo "Database $DB_NAME created successfully"
