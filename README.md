# Introduction 

Golang and WebAssembly example based on https://binx.io/2022/04/22/golang-webassembly.

## Compiling Golang to WebAssembly

```bash
env GOOS=js GOARCH=wasm go build -o static/main.wasm cmd/wasm/main.go
```