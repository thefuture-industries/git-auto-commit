#!/bin/bash

set -e

GIT_ROOT=$(git rev-parse --show-toplevel)
cd "$GIT_ROOT"

HOOKS_DIR=".git/hooks"
HOOK_NAME="auto-commit"
HOOK_PATH="$HOOKS_DIR/$HOOK_NAME"
VERSION_FILE="$HOOKS_DIR/auto-commit.version.txt"

if pgrep -f "$HOOK_PATH" > /dev/null; then
  echo "Stopping running auto-commit processes..."
  pkill -f "$HOOK_PATH"
  sleep 2
fi
