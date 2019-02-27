.PHONY: build
build:
	GOOS=js GOARCH=wasm go build -o basic.wasm ./basic
	GOOS=js GOARCH=wasm go build -o vue.wasm ./vue
