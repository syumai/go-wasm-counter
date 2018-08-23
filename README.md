# go-wasm-counter

* Simple counter app written in go.
* This app runs on browser.

## Example

* [Basic counter](https://syumai.github.io/go-wasm-counter/)
  - [Code](https://github.com/syumai/go-wasm-counter/tree/master/counter/basiccounter.go)
* [Vue.js counter](https://syumai.github.io/go-wasm-counter/vue.html)
  - [Code](https://github.com/syumai/go-wasm-counter/tree/master/counter/vuecounter.go)

## Environment

* go 1.11rc2

## Development

```console
go get go1.11rc2
go1.11rc2 download

make build
```

## How to run

```console
npm install -g serve
serve . # app served on localhost:5000
```

## References

* https://github.com/mattn/golang-wasm-example
* https://github.com/golang/go/tree/master/misc/wasm
* https://tip.golang.org/pkg/syscall/js

## License

MIT
