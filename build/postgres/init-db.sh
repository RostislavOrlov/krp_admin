#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
	CREATE USER krp_auth ENCRYPTED PASSWORD 'krp_auth' LOGIN;
	CREATE DATABASE krp_auth OWNER krp_auth;
EOSQL

psql -v ON_ERROR_STOP=1 --username "krp_auth" --dbname "krp_auth" -f /app/sql/init-db.sql
