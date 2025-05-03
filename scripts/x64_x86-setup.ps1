# x64_x86-setup.ps1

$ErrorActionPreference = "Stop"

$gitRoot = & git rev-parse --show-toplevel
Set-Location $gitRoot

$binarySource = "git-auto-commit/git-auto-commit"
$hookBinary = ".git/hooks/auto-commit"

if (-Not (Test-Path ".git")) {
    Write-Error "[!] There is no .git. Run it in the root of the Git repository."
    exit 1
}

Copy-Item $binarySource $hookBinary -Force
Write-Output "[+] The binary is copied to .git/hooks/auto-commit"

git config --local alias.ac '!./.git/hooks/auto-commit'

Write-Output "[+] Git alias 'git ac' is configured. Now you can run: git ac"
Read-Host -Prompt "Press Enter to exit"
