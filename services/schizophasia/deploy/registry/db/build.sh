#!/bin/bash

sudo apt update && DEBIAN_FRONTEND=noninteractive sudo apt install --yes --no-install-recommends --no-install-suggests build-essential libreadline-dev zlib1g-dev libssl-dev bison flex && \
    make maintainer-clean; \
    ./configure --enable-cassert --enable-debug CFLAGS="-ggdb -Og -g3 -fno-omit-frame-pointer" && \
    make -j4 -s && \
    make install -s
