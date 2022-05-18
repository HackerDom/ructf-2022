#!/bin/bash

BASE_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null && pwd )"
cd "$BASE_DIR"

INPUT_IMAGE="images/ructfe2020-deploy.ova"
ENCRYPTED_DIR="images/encrypted"

rm -rf "$ENCRYPTED_DIR"
mkdir -p "$ENCRYPTED_DIR"
ln "$INPUT_IMAGE" "$ENCRYPTED_DIR/ructfe2020.ova"
cd $ENCRYPTED_DIR
time sha256sum ructfe2020.ova > sha256sum.txt
time 7z a -paN06vKefJFul2KUCv9bf9cx2AXM63umk ructfe2020.ova.7z ructfe2020.ova

