#!/bin/bash

USERS_COUNT=200

for i in $(seq 1 $USERS_COUNT)
do
  USER_NAME=user$i
  USER_PASS=$(echo $RANDOM | md5sum | head -c 25)
  echo "generating user: $USER_NAME"
  /usr/local/pgsql/bin/psql -d postgres -U postgres <<< "CREATE USER $USER_NAME LOGIN PASSWORD '$(USER_PASS)';" &

  #/usr/local/pgsql/bin/psql -d postgres -U loh <<<"select pg_sleep(99999);" &
done;