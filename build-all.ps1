$MyPath = Split-Path $MyInvocation.MyCommand.Definition
$MyProject = $(Get-Item $MyPath).BaseName

if (!(Test-Path $MyPath\bin)) { New-Item -Path $MyPath/bin -ItemType Directory | Out-Null }
#if (!$env:GOPATH) { $env:GOPATH = New-Item -Path $MyPath/go-work/ -ItemType Directory -Force | Select-Object -ExpandProperty FullName }

$BuildOpts = @{
    "linux" = @("386", "amd64", "arm", "arm64");
    "darwin" = @("amd64");
    "windows" = @("386", "amd64")
}

Push-Location $MyPath

& go get .

ForEach ($OS in $BuildOpts.GetEnumerator()) {
    ForEach ($Arch in $OS.Value) {
        $env:GOOS = $OS.Key
        $env:GOARCH = $Arch

        if ($env:GOOS -eq "windows") { $Ext = ".exe" } else { $Ext = $null }

        Try {
            Write-Host "Building: $env:GOOS ($env:GOARCH)"
            & go build -ldflags="-s -w" -o "$MyPath/bin/$MyProject-$env:GOOS-$env:GOARCH$Ext"
        }

        Catch {
            Write-Warning ("Build Failure: $env:GOOS ($env:GOARCH)`r`n`t" + $_.Exception.Message)
        }
    }
}

Pop-Location
