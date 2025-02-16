#!/bin/bash

search_dir=migrations

from=$1
to=$2

for migration in "$search_dir"/*
do
  (( cnt++ ))
  if (( from <= cnt )) && (( (( cnt <= to )) || (( to == 0 )) )); then 
    echo "$migration"
    PGPASSWORD=$PGPASSWORD psql -h $PGHOST -U $PGUSER -p $PGPORT -a -w -f $migration
  fi
done
