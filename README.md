# Git auto-commit - automatic commit generation tool

Git Auto-Commit is an extension for the Git version control system designed to automatically generate meaningful and context—sensitive commit messages based on changes made to the codebase. The tool simplifies developers' workflows by allowing them to focus on the content of edits rather than on the formulation of descriptions for commits.

The development is conducted as an open source project and is distributed under the MIT license (or other compatible licensing, depending on the implementation). Git Auto-Commit can be integrated into CI/CD pipelines, hook scripts, or used manually via the command line.

Main features:

-   Analysis of changes in the work tree and automatic generation of commit messages in natural language.
-   Integration with Git via the `git auto` sub-command or configuration of user aliases.
-   Support for templates and configurations for configuring the message structure in accordance with project standards (Conventional Commits, Semantic Commit Messages, etc.).
-   Scalability: works both in small projects and in large monorepositories.

## Install

### If you're on windows

Go to the root of the project and run the command.

```bash
iex ((New-Object Net.WebClient).DownloadString('https://github.com/thefuture-industries/git-auto-commit/blob/main/scripts/install-windows-auto-commit.ps1?raw=true'))
```

### If you're on linux

Go to the root of the project and run the command.

```bash
echo Y | bash <(curl -fsSL https://github.com/thefuture-industries/git-auto-commit/blob/main/scripts/install-linux-auto-commit.sh?raw=true)
```

## Setting up

### Launch

Everything is ready now, after making changes to the code, just run:

-   1 Option

```bash
git add .
git auto
git push
```

-   2 Commands

```bash
git auto -w - Comit observer, you dont have to think and write more `git auto` -w (--watch) will figure it out when to make a comit and commit it yourself!
git auto -v - Viewing the current version of auto-commit
git auto -u - Upgrade to the new auto-commit version
```
