#!/usr/bin/env bash

set -uex

make --jobs=9
cp backend ../docker/back/backend
strip ../docker/back/backend

