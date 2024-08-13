set TARGET_DIR=.\src\protobuf
mkdir "%TARGET_DIR%"
if not exist "%TARGET_DIR%" (
    mkdir "%TARGET_DIR%"
)

for %%f in (..\protobuff\*.proto) do protoc --plugin=protoc-gen-ts=.\node_modules\.bin\protoc-gen-ts.cmd --ts_opt=esModuleInterop=true --ts_out=".\src\protobuf" -I ..\protobuff %%f
