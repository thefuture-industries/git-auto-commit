# auto-commit.psm1
$ErrorActionPreference = "Stop"

Write-Host @"

        _ _                 _                                            _ _
   __ _(_) |_    __ _ _   _| |_ ___         ___ ___  _ __ ___  _ __ ___ (_) |_
  / _` | | __|  / _` | | | |   __/ _ \ _____ / __/ _ \| '_ ` _ \| '_ ` _ \| | __|
 | (_| | | |_  | (_| | |_| | || (_) |_____| (_| (_) | | | | | | | | | | | | |_
  \__, |_|\__|  \__,_|\__,_|\__\___/       \___\___/|_| |_| |_|_| |_| |_|_|\__|
  |___/

"@

Write-Host "Git Auto-Commit is an extension for the Git version control system designed to automatically generate meaningful and context-sensitive commit messages based on changes made to the codebase. The tool simplifies developers workflows by allowing them to focus on the content of edits rather than on the formulation of descriptions for commits" -ForegroundColor Yellow

$gitRoot = & git rev-parse --show-toplevel
Set-Location $gitRoot

$HookName = "auto-commit-win"

$versionUrl = "https://api.github.com/repos/thefuture-industries/git-auto-commit/releases/latest"
$tag = (Invoke-RestMethod -Uri $versionUrl -UseBasicParsing).tag_name
$Url = "https://github.com/thefuture-industries/git-auto-commit/releases/download/$tag/auto-commit"

if (-not (Test-Path ".git/hooks")) {
    Write-Error "The current directory is not a Git repository."
    return
}

$hookPath = Join-Path -Path ".git/hooks" -ChildPath $HookName

try {
    $answer = Read-Host "Should I install git auto-commit in the project? (Y/N)"

    if ($answer -eq "Y" -or $answer -eq "y") {
        # Install auto-commit
        Write-Host "Install $Url..." -ForegroundColor Green
        Invoke-WebRequest -Uri $Url -OutFile $hookPath -UseBasicParsing
        Write-Host "File saved as $hookPath" -ForegroundColor Yellow

        git config --local alias.auto '!powershell -c ./.git/hooks/auto-commit'

        $versionFile = Join-Path -Path ".git/hooks" -ChildPath "auto-commit.version.txt"
        Set-Content -Path $versionFile -Value $tag

        Write-Host "Successful installation version $tag and settings alias for auto-commit." -ForegroundColor Green

        Write-Host ""
        Write-Host "More detailed: https://github.com/thefuture-industries/git-auto-commit"
        Write-Host "Now you can run: git auto"
    } elseif ($answer -eq "N" -or $answer -eq "n") {
        Write-Host "Skipping installation." -ForegroundColor Yellow
        exit
    } else {
        Write-Error "Invalid input. Please enter Y or N." -ForegroundColor Red
        exit
    }
} catch {
    Write-Error "Error installing: $_"
    return
}
