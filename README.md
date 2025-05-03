# Git auto-commit

`git auto-commit` is a simple and powerful tool for automating commits in Git. With it, you will no longer have to think and write committees. `git auto-commit` will do it for you!

This is an open source project, which is covered by the Mit License version 1 (some parts of it are released under other licenses compatible with GPLv2).

## Install

You will need two things for installation

1. Our ready-made binary
2. Our installer script

On the [releases](https://github.com/thefuture-industries/git-auto-commit/releases) find the release message and install it in Assets

-   git-auto-commit/git-auto-commit
-   If you're on windows -> `x64_x86-setup.ps1`
-   If you're on linux -> `bash-linux.sh `

## Setting up

1. To configure, transfer the twisted files to your project code (after rebuilding, you can delete the files)
2. Run the downloaded script

-   If you're on windows -> `x64_x86-setup.ps1`

```powershell
powershell -ExecutionPolicy Bypass -File .\x64_x86-setup.ps1
```

-   If you're on linux -> `bash-linux.sh`

```bash
bash bash-linux.sh
```

3. Launch

Everything is ready now, after making changes to the code, just run:

```bash
git add .
git ac
git push
```
