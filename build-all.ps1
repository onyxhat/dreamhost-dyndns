$MyPath = Split-Path $MyInvocation.MyCommand.Definition
$MyProject = if (!$env:CI_REPOSITORY_NAME_SLUG) { $env:CI_REPOSITORY_NAME_SLUG } else { $(Get-Item $MyPath).BaseName }

if (!(Test-Path $MyPath/bin)) { New-Item -Path $MyPath/bin -ItemType Directory | Out-Null }

Push-Location $MyPath

& go get .

ForEach ($g in $(& go tool dist list)) {
    if ($g -match '^(darwin|linux|windows)/(arm|arm64|386|amd64)$') {
        $env:GOOS = $Matches[1]
        $env:GOARCH = $Matches[2]

        if ($env:GOOS -eq "windows") { $Ext = ".exe" } else { $Ext = $null }

        Try {
            Write-Host "Building: ${env:GOOS} (${env:GOARCH})"
            & go build -ldflags="-s -w" -o "${MyPath}/bin/${MyProject}-${env:GOOS}-${env:GOARCH}${Ext}"
        }

        Catch {
            Write-Warning ("Build Failed:`r`n" + $_.Exception.Message)
        }
    }
}

Pop-Location
