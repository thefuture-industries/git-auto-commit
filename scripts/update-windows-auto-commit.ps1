# auto-commit.psm1
$ErrorActionPreference = "Stop"

$gitRoot = & git rev-parse --show-toplevel
Set-Location $gitRoot

$proc = Get-Process "auto-commit" -ErrorAction SilentlyContinue
if ($proc) {
    $proc | Stop-Process -Force
    Start-Sleep -Seconds 2
}

$versionUrl = "https://api.github.com/repos/thefuture-industries/git-auto-commit/releases/latest"
$tag = (Invoke-RestMethod -Uri $versionUrl -UseBasicParsing).tag_name

$Url = "https://github.com/thefuture-industries/git-auto-commit/releases/download/$tag/auto-commit"

$HookName = "auto-commit"
$hookPath = Join-Path -Path ".git/hooks" -ChildPath $HookName
$versionFile = Join-Path -Path ".git/hooks" -ChildPath "auto-commit.version.txt"

if (Test-Path $versionFile) {
    $currentTag = Get-Content $versionFile | ForEach-Object { $_.Trim() }
    if ($currentTag -eq $tag) {
        Write-Host "[!] you have the latest version installed $tag" -ForegroundColor Yellow
        exit 0
    }
}

function Download-WithProgress {
    param (
        [string]$url,
        [string]$output
    )

    $req = [System.Net.HttpWebRequest]::Create($url)
    $req.UserAgent = "git-auto-commit"
    $resp = $req.GetResponse()
    $total = $resp.ContentLength
    $stream = $resp.GetResponseStream()
    $outStream = [System.IO.File]::Open($output, [System.IO.FileMode]::Create)

    $buffer = New-Object byte[] 8192
    $read = 0
    $downloaded = 0
    $barWidth = 60

    while (($read = $stream.Read($buffer, 0, $buffer.Length)) -gt 0) {
        $outStream.Write($buffer, 0, $read)
        $downloaded += $read
        $percent = [math]::Round(($downloaded / $total) * 100)
        $filled = [math]::Floor($barWidth * $percent / 100)
        $empty = $barWidth - $filled
        $bar = ('*' * $filled) + ('.' * $empty)
        Write-Host -NoNewline "`rauto-commit update [$bar] $percent% "
    }

    $outStream.Close()
    $stream.Close()
    Write-Host ""
}

Download-WithProgress -url $Url -output $hookPath

Set-Content -Path $versionFile -Value $tag

git config --local alias.auto '!./.git/hooks/auto-commit'
Write-Host "successful upgrade to version $tag"
