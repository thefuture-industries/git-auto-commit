#!/bin/bash

set -e

GIT_ROOT=$(git rev-parse --show-toplevel)
cd "$GIT_ROOT"

HOOKS_DIR=".git/hooks"
HOOK_NAME="auto-commit"
HOOK_PATH="$HOOKS_DIR/$HOOK_NAME"
VERSION_FILE="$HOOKS_DIR/auto-commit.version.txt"

if pgrep -f "$HOOK_PATH" > /dev/null; then
  pkill -f "$HOOK_PATH" || true
fi

VERSION_URL="https://api.github.com/repos/thefuture-industries/git-auto-commit/releases/latest"
TAG=$(curl -s "$VERSION_URL" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')

ARCH=$(uname -m)

case "$ARCH" in
  x86_64)
    ARCH="amd64"
    ;;
  aarch64|arm64)
    ARCH="arm64"
    ;;
  *)
    echo "Unsupported architecture: $ARCH"
    exit 1
    ;;
esac

URL="https://github.com/thefuture-industries/git-auto-commit/releases/download/$TAG/${HOOK_NAME}-linux-${ARCH}"

if [ -f "$VERSION_FILE" ]; then
  CURRENT_TAG=$(cat "$VERSION_FILE" | tr -d ' \n\r')
  if [ "$CURRENT_TAG" = "$TAG" ]; then
    echo -e "\033[33m[!] you have the latest version installed $TAG\033[0m"
    exit 0
  fi
fi

download_with_progress() {
  local url="$1"
  local output="$2"
  local bar_width=60

  content_length=$(curl -sI "$url" | grep -i Content-Length | awk '{print $2}' | tr -d '\r')
  [ -z "$content_length" ] && content_length=0

  echo -n "auto-commit update ["
  for ((i=0; i<bar_width; i++)); do echo -n "."; done
  echo -n "] 0%"

  curl -L "$url" --output "$output" --progress-bar | \
  awk -v bar_width=$bar_width -v total=$content_length '
    BEGIN { done=0 }
    {
      if (match($0, /([0-9]+)%/)) {
        percent = substr($0, RSTART, RLENGTH-1)
        filled = int(bar_width * percent / 100)
        empty = bar_width - filled
        bar = sprintf("%s%s", sprintf("%*s", filled, "*"), sprintf("%*s", empty, "."))
        printf "\rauto-commit update [%s] %d%%", bar, percent
        fflush()
      }
    }
    END { print "" }
  '
}

download_with_progress "$URL" "$HOOK_PATH"
chmod +x "$HOOK_PATH"

echo "$TAG" > "$VERSION_FILE"

git config --local alias.auto "!bash -c './.git/hooks/auto-commit \"\$@\"' --"
echo "successful upgrade to version $TAG"
exit 0
