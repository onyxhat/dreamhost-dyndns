$MyPath = Split-Path $MyInvocation.MyCommand.Definition
$SrcDir = if (!$env:SRC_DIR) { "src" } else { $env:SRC_DIR }
$BinDir = if (!$env:BINARY_OUTDIR) { "bin" } else { $env:BINARY_OUTDIR }
$MyProject = if (!$env:CI_REPOSITORY_NAME_SLUG) { $(Get-Item $MyPath).BaseName } else { $env:CI_REPOSITORY_NAME_SLUG }
$BuildArch = if (!$env:BUILD_ARCH) { ((& go tool dist list) -match '^(darwin|linux|windows)/(arm|arm64|386|amd64)$') -join ',' } else { $env:BUILD_ARCH }

if (!(Test-Path $MyPath/$BinDir)) { New-Item -Path $MyPath/$BinDir -ItemType Directory | Out-Null }

Push-Location $MyPath/$SrcDir

& go get .

foreach ($g in $BuildArch.Split(',')) {
    $env:GOOS = $g.split('/')[0]
    $env:GOARCH = $g.split('/')[1]

    if ($env:GOOS -eq "windows") { $Ext = ".exe" } else { $Ext = $null }

    try {
        Write-Host "Building: ${env:GOOS} (${env:GOARCH})"
        & go build -ldflags="-s -w" -o "${MyPath}/${BinDir}/${MyProject}-${env:GOOS}-${env:GOARCH}${Ext}"
    }

    catch {
        Write-Warning ("Build Failed:`r`n" + $_.Exception.Message)
    }
}

Pop-Location
