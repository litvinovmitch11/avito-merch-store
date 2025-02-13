#!/bin/bash

search_dir=migrations

from=$1
to=$2

for migration in "$search_dir"/*
do
  (( cnt++ ))
  if (( from <= cnt )) && (( (( cnt <= to )) || (( to == 0 )) )); then 
    echo "$migration"
    PGPASSWORD=postgres psql -h localhost -d merch_store -U postgres -p 5432 -a -w -f $migration
  fi
done
