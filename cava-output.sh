#!/usr/bin/env bash

FILE="/tmp/cava_file"
TMP="$FILE.tmp"

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

cava -p $SCRIPT_DIR/cava_config | while IFS= read -r line; do
    printf '%s\n' "$line" > "$TMP"
    mv -f "$TMP" "$FILE"
done
