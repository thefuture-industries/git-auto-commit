# auto-commit.psm1
$ErrorActionPreference = "Stop"

$gitRoot = & git rev-parse --show-toplevel
Set-Location $gitRoot

function Install-GitHook {
    [CmdletBinding()]
    param (
        [string]$HookName = "auto-commit"
    )

    $Url = "https://github.com/thefuture-industries/git-auto-commit/blob/main/bin/auto-commit?raw=true"

    if (-not (Test-Path ".git/hooks")) {
        Write-Error "The current directory is not a Git repository."
        return
    }

    $hookPath = Join-Path -Path ".git/hooks" -ChildPath $HookName

    try {
        Write-Host "Install $Url..."
        Invoke-WebRequest -Uri $Url -OutFile $hookPath -UseBasicParsing
        Write-Host "File saved as $hookPath"
    } catch {
        Write-Error "Error installing: $_"
        return
    }

    git config --local alias.ac "!./.git/hooks/auto-commit"
    Write-Host "[+] Git alias 'git ac' configured. Now you can run: git ac"
}
