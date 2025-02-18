@echo off

:: Run the make install command
make install
set OUT=%ERRORLEVEL%
echo The error level is %OUT%
echo APPDATA path is "%APPDATA%"

:: Check if the previous command was successful
if %OUT%==0 (
    :: Define the correct path for the provider installation
    set PLUGIN_DIR=%APPDATA%\.terraform.d\plugins\registry.terraform.io\site24x7\site24x7\1.0.0\windows_amd64\
    echo Plugin directory is "%PLUGIN_DIR%"

    :: Check if the directory exists, if not, create it
    if not exist "%PLUGIN_DIR%" (
        mkdir "%PLUGIN_DIR%"
    )
    
    :: Copy the terraform provider file to the correct directory
    copy /y terraform-provider-site24x7.exe "%PLUGIN_DIR%terraform-provider-site24x7_v1.0.0.exe"

    :: Remove the .terraform.lock.hcl file
    del /q .terraform.lock.hcl

    :: Initialize terraform
    terraform init
) else (
    echo Compilation Error
)
