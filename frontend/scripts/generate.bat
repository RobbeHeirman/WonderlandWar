@echo off

set TARGET_DIR=.\src\proto
mkdir "%TARGET_DIR%"
if not exist "%TARGET_DIR%" (
    mkdir "%TARGET_DIR%"
)

for %%f in (..\protobuf\*.proto) do npx protoc --ts_out %TARGET_DIR% --proto_path ..\protobuf %%f