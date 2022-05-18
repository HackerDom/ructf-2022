#!/bin/bash

set -e

# Cd into script directory
BASE_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null && pwd )"
cd "$BASE_DIR"

service="$1"

if [ -z "$service" ]; then
    ansible-playbook --key-file keys/id_rsa -i teams1-10.hosts image.yml
else
    ansible-playbook --key-file keys/id_rsa -i teams1-10.hosts -t "$service" image.yml
fi
