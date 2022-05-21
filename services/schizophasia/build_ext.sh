#!/bin/bash
set -e -x

docker build -t pg_ext -f nice_ext/Dockerfile .
docker run --rm -iv $PWD:/app/out pg_ext  sh -s <<EOF
set -e -x
rm /app/out/deploy/registry/nice_ext/nice_ext.so || true
rm /app/out/deploy/registry/nice_ext/nice_ext.control || true
rm /app/out/deploy/registry/nice_ext/nice_ext--1.0.sql || true
cp /usr/local/pgsql/lib/nice_ext.so /app/out/deploy/registry/nice_ext/nice_ext.so
cp /app/out/nice_ext/nice_ext.control /app/out/deploy/registry/nice_ext/nice_ext.control
cp /app/out/nice_ext/nice_ext--1.0.sql /app/out/deploy/registry/nice_ext/nice_ext--1.0.sql
EOF
