package main

import (
	"github.com/syumai/go-web-example/counter"
)

func main() {
	counter.NewDOMCounter("app")
	select {}
}
