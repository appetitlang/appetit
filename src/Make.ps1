# Set essential information
$version = "1"
$date = Get-Date -Format "dd/MM/yyyy HH:mm:ss"
$platform = $env:PROCESSOR_ARCHITECTURE

# Start the build process
Write-Output "Appetit Make Script, building for $platform"
# Clean up any old caches or dist/ directories to enable a "clean" build
Write-Output ":: Cleaning up caches and dist/ directories"
[void](Remove-Item -Path "dist" -Recurse -Force -ErrorAction SilentlyContinue)

# Create the dist folder
# https://lazyadmin.nl/powershell/create-folder/#powershell-create-folder 
# and https://collectingwisdom.com/powershell-new-item-silent/
# and https://stackoverflow.com/questions/16906170/create-directory-if-it-does
# -not-exist
[void](New-Item -Path "." -Name "dist" -ItemType Directory -Force)

# If the architecture is ARM64, build an ARM64 interpreter
if ("arm64" -eq $platform) {
    Write-Output ":: Building Windows ARM64 binary..."
    $Env:GOOS="windows"; $Env:GOARCH="arm64"; go build -buildvcs=false -ldflags="-s -w -X 'main.BuildDate=$date'" -o dist/appetit-$version.exe
# Otherwise, we'll assume that we're on a x86_64 version of Windows
} else {
    Write-Output ":: Building Windows x86_64 binary..."
    # https://stackoverflow.com/questions/50911153/how-to-crosscompile-go-programs-on-windows-10
    $Env:GOOS="windows"; $Env:GOARCH="amd64"; go build -buildvcs=false -ldflags="-s -w -X 'main.BuildDate=$date'" -o dist/appetit-$version.exe
}

Write-Output ":: Adding interpreter to C:\appetit and adding C:\appetit to PATH (user)"
# Create the C:\appetit folder. For now, this assumes that C:\ is available which should be safe given Microsoft's aggressive backwards compatibility
[void](New-Item -Path "C:\appetit" -ItemType Directory -Force)
# Move the created binary to C:\appetit
[void](Move-Item -Path "dist/appetit-$version.exe" -Destination "C:\appetit" -Force)
# Add C:\appetit to the PATH
[Environment]::SetEnvironmentVariable("Path", $env:Path + ";C:\appetit", [System.EnvironmentVariableTarget]::User)
# Remove the dist/ folder
[void](Remove-Item -Path "dist" -Recurse -Force -ErrorAction SilentlyContinue)

Write-Output "Done! You should now be able to run appetit-$version.exe from a PowerShell window."