#!/bin/bash

set -e

HOME="$(git rev-parse --show-toplevel)"
cd "$HOME"

echo ""
cat <<'EOF'
        _ _                 _                                            _ _
   __ _(_) |_    __ _ _   _| |_ ___         ___ ___  _ __ ___  _ __ ___ (_) |_
  / _` | | __|  / _` | | | |   __/ _ \ _____ / __/ _ \| '_ ` _ \| '_ ` _ \| | __|
 | (_| | | |_  | (_| | |_| | || (_) |_____| (_| (_) | | | | | | | | | | | | |_
  \__, |_|\__|  \__,_|\__,_|\__\___/       \___\___/|_| |_| |_|_| |_| |_|_|\__|
  |___/
EOF
echo ""

echo -e "\e[33mGit Auto-Commit is an extension for the Git version control system designed to automatically generate meaningful and context-sensitive commit messages based on changes made to the codebase. The tool simplifies developers' workflows by allowing them to focus on the content of edits rather than on the formulation of descriptions for commits.\e[0m"

BINARY_NAME="auto-commit"
HOOKS_DIR=".git/hooks"
HOOK_PATH="$HOOKS_DIR/$BINARY_NAME"

URL="https://github.com/thefuture-industries/git-auto-commit/blob/main/bin/auto-commit?raw=true"

if [ ! -d .git ]; then
  echo "[!] There is no .git. Run it in the root of the Git repository."
  exit 1
fi

read -p "Should I install git auto-commit in the project? (Y/N) " answer

if [[ "$answer" == "Y" || "$answer" == "y" ]]; then
  echo -e "\e[32mInstall $URL...\e[0m"
  curl -L "$URL" -o "$HOOK_PATH"
  chmod +x "$HOOK_PATH"
  echo -e "\e[33mFile saved as $HOOK_PATH\e[0m"

  git config --local alias.auto "!$HOOK_PATH"

  echo -e "\e[32mSuccessful installation and settings alias for auto-commit.\e[0m"
  echo ""
  echo -e "\e[33mMore detailed: https://github.com/thefuture-industries/git-auto-commit\e[0m"
  echo -e "\e[33mNow you can run: git auto\e[0m"
elif [[ "$answer" == "N" || "$answer" == "n" ]]; then
  echo -e "\e[33mSkipping installation.\e[0m"
  exit 0
else
  echo -e "\e[31mInvalid input. Please enter Y or N.\e[0m"
  exit 1
fi
