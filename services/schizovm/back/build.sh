#!/usr/bin/env bash

set -uex

make --jobs=9
cp schizovm ../docker/back/schizovm
strip ../docker/back/schizovm
