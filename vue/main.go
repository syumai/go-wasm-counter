package main

import (
	"github.com/syumai/go-wasm-counter/counter"
)

func main() {
	counter.NewVueCounter("app")
	select {}
}
