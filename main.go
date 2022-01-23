package main

import (
	"syscall/js"
)

// if using TinyGo for compiling: tinygo build -o wasm.wasm -target wasm ./main.go
// func sub(a, b float64) float64
func sub(this js.Value, inputs []js.Value) interface{} {
	return inputs[0].Float() - inputs[1].Float()
}

// if using TinyGo for compiling: tinygo build -o wasm.wasm -target wasm ./main.go
//export multiply
func multiply(x, y int) int {
	return x * y
}

func main() {
	c := make(chan int) // channel to keep the wasm running, it is not a library as in rust/c/c++, so we need to keep the binary running
	js.Global().Set("sub", js.FuncOf(sub))
	alert := js.Global().Get("alert")
	alert.Invoke("Hi")
	println("Hello wasm")

	num := js.Global().Call("add", 3, 4)
	println(num.Int())

	document := js.Global().Get("document")
	h1 := document.Call("createElement", "h1")
	h1.Set("innerText", "This is H1")
	document.Get("body").Call("appendChild", h1)

	<-c // pause the execution so that the resources we create for JS keep available
}

// compile to wasm:
// GOOS=js GOARCH=wasm go build -o www/wasm/main.wasm github.io/hajsf/wasm
// Copied the wasm_exec.js file to the same working folder as:
// cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" ./wasm
