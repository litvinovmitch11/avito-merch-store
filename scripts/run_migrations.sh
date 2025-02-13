search_dir=migrations
for migration in "$search_dir"/*
do
  echo "$migration"
  PGPASSWORD=postgres psql -h localhost -d merch_store -U postgres -p 5432 -a -w -f $migration
done

