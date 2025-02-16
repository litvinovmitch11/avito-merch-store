#!/bin/bash

PGPASSWORD=$PGPASSWORD psql -h $PGHOST -U $PGUSER -p $PGPORT -c 'DROP DATABASE IF EXISTS merch_store;'
PGPASSWORD=$PGPASSWORD psql -h $PGHOST -U $PGUSER -p $PGPORT -d merch_store -c 'CREATE DATABASE merch_store;'

# TODO: add roles: robot - reader, writer; user - reader; owner - all
