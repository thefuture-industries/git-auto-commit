#!/bin/bash

set -e

GIT_ROOT=$(git rev-parse --show-toplevel)
cd "$GIT_ROOT"

HOOKS_DIR=".git/hooks"
HOOK_NAME="auto-commit"
HOOK_PATH="$HOOKS_DIR/$HOOK_NAME"
VERSION_FILE="$HOOKS_DIR/auto-commit.version.txt"

if pgrep -f "$HOOK_PATH" > /dev/null; then
  pkill -f "$HOOK_PATH"
  sleep 2
fi

VERSION_URL="https://api.github.com/repos/thefuture-industries/git-auto-commit/releases/latest"
TAG=$(curl -s "$VERSION_URL" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')
URL="https://github.com/thefuture-industries/git-auto-commit/releases/download/$TAG/auto-commit"
