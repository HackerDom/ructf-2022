#!/usr/bin/env bash

set -uex

make --jobs=9
cp prosopagnosia ../docker/back/prosopagnosia
strip ../docker/back/prosopagnosia

