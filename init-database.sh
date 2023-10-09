#!/bin/bash

# List of databases to create
DATABASES=("fastapi" "gin")

# PostgreSQL connection parameters
PG_USER="postgres"

# Function to check if a database exists
database_exists() {
    psql -U $PG_USER -d $1 -c "SELECT 1" > /dev/null 2>&1
}

# Loop through the databases and create if not exists
for DB_NAME in "${DATABASES[@]}"; do
    if database_exists $DB_NAME; then
        echo "Database $DB_NAME already exists."
    else
        echo "Creating database $DB_NAME..."
        createdb -U $PG_USER $DB_NAME
        echo "Database $DB_NAME created."
    fi
done

echo "Database creation script completed."
