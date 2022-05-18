#!/bin/bash

set -e

# Cd into script directory
BASE_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null && pwd )"
cd "$BASE_DIR"

service="$1"

if [ -z "$service" ]; then
    echo "Sorry, can't cleanup everything"
    exit 1
fi

REMOTE_CMD="systemctl stop $service; sleep 1; (cd /home/$service && docker-compose down --volumes --rmi local); rm -rf /home/$service/; mkdir -p /home/$service/"
ansible --key-file keys/id_rsa -i teams1-10.hosts all -m shell -a "$REMOTE_CMD"
