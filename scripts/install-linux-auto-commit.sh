#!/bin/bash

set -e

HOME="$(git rev-parse --show-toplevel)"
cd "$HOME"

BINARY_NAME="git-auto-commit"
HOOKS_DIR=".git/hooks"
TARGET="$HOOKS_DIR/auto-commit"

if [ ! -d .git ]; then
  echo "[!] There is no .git. Run it in the root of the Git repository."
  exit 1
fi

cp "$BINARY_NAME" "$TARGET"
chmod +x "$TARGET"
echo "[+] The binary is copied to $TARGET"

git config alias.ac '!f() { .git/hooks/auto-commit "$@"; }; f'

echo "[+] Git alias 'git ac' is configured. Now you can run: git ac"
