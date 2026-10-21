#!/usr/bin/env bash

setsid ~/Development/hyprlock-visualizer/cava-output.sh &
pid=$!

hyprlock

kill -- "-$pid" 2>/dev/null || true
wait "$pid" 2>/dev/null || true
