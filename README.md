# Git auto-commit

`git auto-commit` is a simple and powerful tool for automating commits in Git. With it, you will no longer have to think and write committees. `git auto-commit` will do it for you!

This is an open source project, which is covered by the Mit License version 1 (some parts of it are released under other licenses compatible with GPLv2).

## Install

You will need two things for installation

-   If you're on windows

```bash
iex ((New-Object Net.WebClient).DownloadString('https://github.com/thefuture-industries/git-auto-commit/blob/main/scripts/auto-commit.ps1?raw=true'))
```

-   If you're on linux

```bash
bash <(curl -s https://github.com/thefuture-industries/git-auto-commit/blob/main/scripts/auto-commit.sh?raw=true)
```

## Setting up

### Launch

Everything is ready now, after making changes to the code, just run:

```bash
git add .
git ac
git push
```
