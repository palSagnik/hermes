#!/bin/sh
echo "CREATE DATABASE ${POSTGRES_DATABASE};" | psql -U ${POSTGRES_USER}
