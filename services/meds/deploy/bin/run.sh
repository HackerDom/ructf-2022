#!/bin/bash
set -e

mkdir -p /app/data/
chown -R meds:meds /app/data/

su meds -s /app/meds