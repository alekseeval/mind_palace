#!/bin/bash
# Script for creating DB, installing scheme and DB functions
# Just run it
# Change parameters below if necessary

# TODO: прокинуть параметры ниже в скрипт установки схемы и создания БД
# TODO: Можно в скрипте установки схемы просто брать текущего пользователя и не прокидывать ничего

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

# Creating DB from create_db.sql
export PGPASSWORD="$POSTGRES_PASSWORD"
psql -h "$HOST" -p "$PORT" --user "$POSTGRES" --dbname "postgres" < create_db.sql

# Installation DB scheme from create_scheme.sql
PGPASSWORD="$DB_ADMIN_PASSWORD"
psql -h "$HOST" -p "$PORT" --user "$DB_ADMIN" --dbname "$DB_NAME" < create_scheme.sql

# Installation DB api functions from /db_api directory
cd "db_api/"
for FILE in *; do
  psql -h "$HOST" -p "$PORT" --user "$DB_ADMIN" --dbname "$DB_NAME" < "$FILE"
done