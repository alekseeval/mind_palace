#!/bin/bash
# Script for creating DB, installing scheme and DB functions
# Just run it
# Change parameters below if necessary

# Parameters
HOST='192.168.0.108'
PORT='5432'
POSTGRES='postgres'
POSTGRES_PASSWORD='postgres'
DB_ADMIN='mp_admin'
DB_ADMIN_PASSWORD='mp_admin'
DB_NAME="mind_palace"

# Going to script directory
SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
cd "$SCRIPT_DIR"

# Set log file of script
mkdir -p "log"
SCRIPT_LOG="$(pwd)/log/log_$(date -Iseconds).log"
exec &>"$SCRIPT_LOG"

# Creating DB
echo "---> Start DB creation..."
export PGPASSWORD="$POSTGRES_PASSWORD"
psql -h "$HOST" -p "$PORT" --user "$POSTGRES" --dbname "postgres" -c "CREATE DATABASE $DB_NAME;"
psql -h "$HOST" -p "$PORT" --user "$POSTGRES" --dbname "postgres" -c "create user $DB_ADMIN with encrypted password '$DB_ADMIN_DB_ADMIN_PASSWORD';"
psql -h "$HOST" -p "$PORT" --user "$POSTGRES" --dbname "postgres" -c "grant all privileges on database $DB_NAME to $DB_ADMIN;"
echo -e "---> DB $DB_NAME was created\n"

# Installation DB scheme from create_scheme.sql
echo "---> Start DB scheme creation..."
PGPASSWORD="$DB_ADMIN_PASSWORD"
psql -h "$HOST" -p "$PORT" --user "$DB_ADMIN" --dbname "$DB_NAME" < create_scheme.sql
echo -e "---> DB scheme was created\n"

# Create init DB data by init_data.sql
echo "---> Fill DB with init data..."
PGPASSWORD="$DB_ADMIN_PASSWORD"
psql -h "$HOST" -p "$PORT" --user "$DB_ADMIN" --dbname "$DB_NAME" < init_data.sql
echo -e "---> Init data was inserted\n"

# Installation DB api functions from /db_api directory
echo "---> Installing DB API..."
cd "db_api/"
for FILE in *; do
  psql -h "$HOST" -p "$PORT" --user "$DB_ADMIN" --dbname "$DB_NAME" < "$FILE"
done
echo "---> DB API was successfully installed"

exec &>/dev/tty
cat "$SCRIPT_LOG"