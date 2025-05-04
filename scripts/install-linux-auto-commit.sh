#!/bin/bash

set -e

HOME="$(git rev-parse --show-toplevel)"
cd "$HOME"

BINARY_NAME="auto-commit"
HOOKS_DIR=".git/hooks"
HOOK_PATH="$HOOKS_DIR/$BINARY_NAME"

URL="https://github.com/thefuture-industries/git-auto-commit/blob/main/bin/auto-commit?raw=true"

if [ ! -d .git ]; then
  echo "[!] There is no .git. Run it in the root of the Git repository."
  exit 1
fi

echo "Installing $URL..."
if curl -fsSL "$URL" -o "$HOOK_PATH"; then
  chmod +x "$HOOK_PATH"
  echo "File saved as $HOOK_PATH"
else
  echo "Error: Failed to download the hook script."
  exit 1
fi

git config --local alias.auto "!$HOOK_PATH"
echo "[+] Git alias 'git auto' is configured. Now you can run: git auto"
