package main

import (
	"crypto"
	_ "crypto/sha512"
	"encoding/hex"
	"fmt"
	"syscall/js"
)

func main() {
	fmt.Println("Go go wasm")
	done := make(chan struct{}, 0)
	js.Global().Set("wasmHash", js.FuncOf(hash))
	<-done
}

func hash(this js.Value, args []js.Value) interface{} {
	h := crypto.SHA512.New()
	h.Write([]byte(args[0].String()))
	return hex.EncodeToString(h.Sum(nil))
}
