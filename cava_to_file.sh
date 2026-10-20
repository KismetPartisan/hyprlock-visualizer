FILE="/tmp/cava_file"
TMP="$FILE.tmp"

cava -p cava_config | while IFS= read -r line; do
    printf '%s\n' "$line" > "$TMP"
    mv -f "$TMP" "$FILE"
done
echo "TEST"
