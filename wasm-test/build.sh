
# copy wasm_exec.js from go install
#cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" ~/GolandProjects/wasm-test/assets/

# build wasm:
GOOS=js GOARCH=wasm go build -o ./assets/json.wasm ./wasm

# build server:
go build -o ./server/server.exe ./server

./server/server.exe