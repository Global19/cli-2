#!/bin/sh
# Copyright 2019 ActiveState Software Inc. All rights reserved.
<#
.DESCRIPTION
Install the ActiveState state.exe tool.  Must be run as admin OR install state tool to
User profile folder.

.EXAMPLE
install.ps1 -b branchToInstall -t C:\dir\on\path
#>

param (
    [Parameter(Mandatory=$False)][string]$b = "unstable"
    ,[Parameter(Mandatory=$False)]
        [string]
        $t
    ,[Parameter(Mandatory=$False)][switch]$n
    ,[Parameter(Mandatory=$False)][switch]$f
    ,[Parameter(Mandatory=$False)][switch]$h
    ,[Parameter(Mandatory=$False)]
        [ValidateScript({[IO.Path]::GetExtension($_) -eq '.exe'})]
        [string]
        $e = "state.exe"
    ,[Parameter(Mandatory=$False)][string]$activate = ""
)

$script:NOPROMPT = $n
$script:FORCEOVERWRITE = $f
$script:TARGET = ($t).Trim()
$script:STATEEXE = ($e).Trim()
$script:STATE = $e.Substring(0, $e.IndexOf("."))
$script:BRANCH = ($b).Trim()
$script:ACTIVATE =($activate).Trim()

# Some cmd-lets throw exceptions that don't stop the script.  Force them to stop.
$ErrorActionPreference = "Stop"

# Helpers
function notifySettingChange(){
    $HWND_BROADCAST = [IntPtr] 0xffff;
    $WM_SETTINGCHANGE = 0x1a;
    $result = [UIntPtr]::Zero

    if (-not ("Win32.NativeMethods" -as [Type]))
    {
        # import sendmessagetimeout from win32
        Add-Type -Namespace Win32 -Name NativeMethods -MemberDefinition @"
        [DllImport("user32.dll", SetLastError = true, CharSet = CharSet.Auto)]
        public static extern IntPtr SendMessageTimeout(
        IntPtr hWnd, uint Msg, UIntPtr wParam, string lParam,
        uint fuFlags, uint uTimeout, out UIntPtr lpdwResult);
"@
    }
    # notify all windows of environment block change
    [Win32.Nativemethods]::SendMessageTimeout($HWND_BROADCAST, $WM_SETTINGCHANGE, [UIntPtr]::Zero, "Environment", 2, 5000, [ref] $result) | Out-Null;

}

function isAdmin
{
    $currentPrincipal = New-Object Security.Principal.WindowsPrincipal([Security.Principal.WindowsIdentity]::GetCurrent())
    $currentPrincipal.IsInRole([Security.Principal.WindowsBuiltInRole]::Administrator)
}

function promptYN([string]$msg)
{
    $response = Read-Host -Prompt $msg" [y/N]`n"

    if ( -Not ($response.ToLower() -eq "y") )
    {
        return $False
    }
    return $True
}

function errorOccured($suppress) {
    $errMsg = $Error[0]
    $Error.Clear()
    if($errMsg) {
        if (-Not $suppress){
            Write-Warning $errMsg
        }
        return $True, $errMsg
    }
    return $False, ""
}

function hasWritePermission([string] $path)
{
    # $user = "$env:userdomain\$env:username"
    # $acl = Get-Acl $path -ErrorAction 'silentlycontinue'
    # return (($acl.Access | Select-Object -ExpandProperty IdentityReference) -contains $user)
    $thefile = "activestate-perms"
    New-Item -Path (Join-Path $path $thefile) -ItemType File -ErrorAction 'silentlycontinue'
    $occurance = errorOccured $True
    #  If an error occurred and it's NOT and IOExpction error where the file already exists
    if( $occurance[0] -And -Not ($occurance[1].exception.GetType().fullname -eq "System.IO.IOException" -And (Test-Path $path))){
        return $False
    }
    Remove-Item -Path (Join-Path $path $thefile) -Force  -ErrorAction 'silentlycontinue'
    if((errorOccured $True)[0]){
        return $False
    }
    return $True
}

function checkPermsRecur([string] $path)
{
    $orig = $path
    # recurse up to the drive root if we have to
    while ($path -ne "") {
        if (Test-Path $path){
            if (-Not (hasWritePermission $path)){
                Write-Warning "You do not have permission to write to '$path'.  Are you running as admin?"
                return $False
            } else {
                return $True
            }
        }
        $path = split-path $path
    }
    Write-Warning "'$orig' is not a valid path"
    return $False
}

function isValidFolder([string] $path)
{      
    if($path[1] -ne ":"){
        Write-Warning "Must provide an absolute path."
        return $False
    }
    if(Test-Path $path){
        #it's a folder
        if (-Not (Test-Path $path -PathType 'Container')){
            Write-Warning "'$path' exists and is not a directory"
            return $False
        }
    }
    return checkPermsRecur $path
}

# isStateToolInstallationOnPath returns true if the state tool's installation directory is in the current PATH
function isStateToolInstallationOnPath($installDirectory) {
    $existing = getExistingOnPath
    $existing -eq $installDirectory
}

function getExistingOnPath(){
    $path = (get-command $script:STATEEXE -ErrorAction 'silentlycontinue').Source
    if ($null -eq $path) {
        ""
    } else {
       (Resolve-Path (split-path -Path $path -Parent)).Path
    }
}

function activateIfRequested() {
    if ( $script:ACTIVATE -ne "" ) {
        # This creates an interactive sub-shell.
        Write-Host "`nActivating project $script:ACTIVATE`n" -ForegroundColor Yellow
        &$script:STATEEXE activate $script:ACTIVATE
    }
}

function warningIfadmin() {
    if (IsAdmin) {
        Write-Warning "It's recommended that you close this command prompt and start a new one without admin privileges.`n"
    }
}

function fetchArtifacts($downloadDir, $statejson, $statepkg) {

    # State tool binary base dir
    $STATEURL="https://s3.ca-central-1.amazonaws.com/cli-update/update/state"
    
    Write-Host "Preparing for installation...`n"
    
    $downloader = new-object System.Net.WebClient

    # Get version and checksum
    $jsonurl = "$STATEURL/$script:BRANCH/$statejson"
    Write-Host "Determining latest version...`n"
    try{
        $branchJson = ConvertFrom-Json -InputObject $downloader.DownloadString($jsonurl)
        $latestVersion = $branchJson.Version
        $versionedJson = ConvertFrom-Json -InputObject $downloader.DownloadString("$STATEURL/$script:BRANCH/$latestVersion/$statejson")
    } catch [System.Exception] {
        Write-Warning "Unable to retrieve the latest version number"
        Write-Error $_.Exception.Message
        exit 1
    }
    $latestChecksum = $versionedJson.Sha256v2

    # Download pkg file
    $zipPath = Join-Path $downloadDir $statepkg
    # Clean it up to start but leave it behind when done 
    if(Test-Path $downloadDir){
        Remove-Item $downloadDir -Recurse
    }
    New-Item -Path $downloadDir -ItemType Directory | Out-Null # There is output from this command, don't show the user.
    $zipURL = "$STATEURL/$script:BRANCH/$latestVersion/$statepkg"
    Write-Host "Fetching the latest version: $latestVersion...`n"
    try{
        $downloader.DownloadFile($zipURL, $zipPath)
    } catch [System.Exception] {
        Write-Warning "Could not install state tool"
        Write-Warning "Could not access $zipURL"
        Write-Error $_.Exception.Message
        exit 1
    }

    # Check the sums
    Write-Host "Verifying checksums...`n"
    $hash = (Get-FileHash -Path $zipPath -Algorithm SHA256).Hash
    if ($hash -ne $latestChecksum){
        Write-Warning "SHA256 sum did not match:"
        Write-Warning "Expected: $latestChecksum"
        Write-Warning "Received: $hash"
        Write-Warning "Aborting installation"
        exit 1
    }

    # Extract binary from pkg and confirm checksum
    Write-Host "Extracting $statepkg...`n"
    Expand-Archive $zipPath $downloadDir
}

function install()
{
    $USAGE="install.ps1 [flags]
    
    Flags:
    -b <branch>          Default 'unstable'.  Specify an alternative branch to install from (eg. master)
    -n                   Don't prompt for anything, just install and override any existing executables
    -t <dir>             Install target dir
    -f <file>            Default 'state.exe'.  Binary filename to use
    -activate <project>  Activate a project when state tools is correctly installed
    -h                   Show usage information (what you're currently reading)"

    # Ensure errors from previously run commands are not reported during install
    $Error.Clear()

    if ($h) {
        Write-Host $USAGE
        exit 0
    }

    if ($script:NOPROMPT -and $script:ACTIVATE -ne "" ) {
        Write-Error "Flags -n and -activate cannot be set at the same time."
        exit 1
    }

    if ($script:FORCEOVERWRITE -and ( -not $script:NOPROMPT) ) {
        Write-Error "Flag -f also requires -n"
        exit 1
    }
    
    # $ENV:PROCESSOR_ARCHITECTURE == AMD64 | x86
    if ($ENV:PROCESSOR_ARCHITECTURE -eq "AMD64") {
        $statejson="windows-amd64.json"
        $statepkg="windows-amd64.zip"
        $stateexe="windows-amd64.exe"

    } else {
        Write-Warning "x86 processors are not supported at this time"
        Write-Warning "Contact ActiveState Support for assistance"
        Write-Warning "Aborting installation"
        exit 1
    }

    # Get the install directory and ensure we have permissions on it.
    # If the user provided an install dir we do no verification.
    if ($script:TARGET) {
        $installDir = $script:TARGET
    } else {
        $installDir = (Join-Path $Env:APPDATA (Join-Path "ActiveState" "bin"))
        if (-Not (hasWritePermission $Env:APPDATA)){
            Write-Error "Do not have write permissions to: '$Env:APPDATA'"
            Write-Error "Aborting installation"
            exit 1
        }
    }

    # stop if previous installation is detected, unless
    # - A. a different target directory has been specified
    # - B. FORCEOVERWRITE is true
    if (get-command $script:STATEEXE -ErrorAction 'silentlycontinue') {
        $existing = getExistingOnPath
    
        # check for A
        if (-not $script:TARGET -or ( $script:TARGET -eq $existing )) {
            # check for B
            if ($script:FORCEOVERWRITE) {
                Write-Warning "Overwriting previous installation."
            } else {
                Write-Host $("Previous install detected at '"+($existing)+"'") -ForegroundColor Yellow
                Write-Host "To update the state tool to the latest version, please run 'state update'."
                Write-Host "To install in a different location, please specify the installation directory with '-t TARGET_DIR'."
                exit 0
            }
        }
    }

    # Install binary
    Write-Host "`nInstalling to '$installDir'...`n" -ForegroundColor Yellow
    if ( -Not $script:NOPROMPT ) {
        if( -Not (promptYN "Continue?") ) {
            exit 1
        }
    }

    #  If the install dir doesn't exist
    $installPath = Join-Path $installDir $script:STATEEXE
    if( -Not (Test-Path $installDir)) {
        Write-host "NOTE: $installDir will be created`n"
        New-Item -Path $installDir -ItemType Directory | Out-Null
    } else {
        if(Test-Path $installPath -PathType Leaf) {
            Remove-Item $installPath -Erroraction 'silentlycontinue'
            $occurance = errorOccured $False
            if($occurance[0]){
                Write-Host "Aborting Installation" -ForegroundColor Yellow
                exit 1
            }
        }
    }

    $tmpParentPath = Join-Path $env:TEMP "ActiveState"
    fetchArtifacts $tmpParentPath $statejson $statepkg
    Move-Item (Join-Path $tmpParentPath $stateexe) $installPath

    # Check if installation is in $PATH
    if (isStateToolInstallationOnPath $installDir) {
        Write-Host "`nState tool installation complete." -ForegroundColor Yellow
        Write-Host "You may now start using the '$script:STATEEXE' program."
        warningIfAdmin
        activateIfRequested
        exit 0
    }

    # Update PATH for state tool installation directory
    $envTarget = [EnvironmentVariableTarget]::User
    $envTargetName = "user"
    if (isAdmin) {
        $envTarget = [EnvironmentVariableTarget]::Machine
	    $envTargetName = "system"
    }

    Write-Host "Updating environment...`n"
    Write-Host "Adding $installDir to $envTargetName PATH`n"
    # This only sets it in the registry and it will NOT be accessible in the current session
    [Environment]::SetEnvironmentVariable(
        'Path',
        $installDir + ";" + [Environment]::GetEnvironmentVariable(
            'Path', [EnvironmentVariableTarget]::Machine),
        $envTarget)

    notifySettingChange

    $env:Path = $installDir + ";" + $env:Path

    warningIfAdmin
    Write-Host "State tool successfully installed to: $installDir." -ForegroundColor Yellow
    Write-Host "Please restart your command prompt in order to start using the 'state.exe' program." -ForegroundColor Yellow
    activateIfRequested

}

install
