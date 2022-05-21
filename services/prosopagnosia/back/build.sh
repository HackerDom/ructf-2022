#!/usr/bin/env bash

set -uex

make --jobs=9
cp prosopagnosia ../deploy/docker/back/prosopagnosia
strip ../deploy/docker/back/prosopagnosia

