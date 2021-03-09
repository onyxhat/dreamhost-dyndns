$MyPath = Split-Path $MyInvocation.MyCommand.Definition
$BinDir = if (!$env:BINARY_OUTDIR) { "bin" } else { $env:BINARY_OUTDIR }
$MyProject = if (!$env:CI_REPOSITORY_NAME_SLUG) { $(Get-Item $MyPath).BaseName } else { $env:CI_REPOSITORY_NAME_SLUG }

if (!(Test-Path $MyPath/$BinDir)) { New-Item -Path $MyPath/$BinDir -ItemType Directory | Out-Null }

Push-Location $MyPath

& go get .

ForEach ($g in $(& go tool dist list)) {
    if ($g -match '^(darwin|linux|windows)/(arm|arm64|386|amd64)$') {
        $env:GOOS = $Matches[1]
        $env:GOARCH = $Matches[2]

        if ($env:GOOS -eq "windows") { $Ext = ".exe" } else { $Ext = $null }

        Try {
            Write-Host "Building: ${env:GOOS} (${env:GOARCH})"
            & go build -ldflags="-s -w" -o "${MyPath}/${BinDir}/${MyProject}-${env:GOOS}-${env:GOARCH}${Ext}"
        }

        Catch {
            Write-Warning ("Build Failed:`r`n" + $_.Exception.Message)
        }
    }
}

Pop-Location
