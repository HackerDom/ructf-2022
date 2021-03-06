#!/bin/bash
while true
do
 flock -n -w 10 /tmp/auto_kill.lock psql -U postgres -c "SELECT now(), query, pg_terminate_backend(pid) FROM pg_stat_activity WHERE coalesce(xact_start, query_start) < current_timestamp - '1 minute'::interval AND usename NOT IN (SELECT rolname FROM pg_roles WHERE rolsuper) AND usename NOT IN ('') AND backend_type = 'client backend';"  >> /var/log/postgresql/autokill.log 2>&1
 sleep 5
done