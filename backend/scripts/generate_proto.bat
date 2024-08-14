@echo off

set CURRENT_DIR=%cd%
set SCRIPT_PATH=%~dp0

cd %SCRIPT_PATH%\..\..\
set PROTO_PATH=.\protobuf
set TARGET_DIR=.\backend\src

for %%f in ("%PROTO_PATH%\*.proto") do protoc --proto_path=%PROTO_PATH%  --go_out=%TARGET_DIR% %%f

cd %CURRENT_DIR%
