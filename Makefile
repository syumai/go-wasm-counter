GO_BIN = go1.11

.PHONY: build
build:
	GOOS=js GOARCH=wasm $(GO_BIN) build -o basic.wasm ./basic
	GOOS=js GOARCH=wasm $(GO_BIN) build -o vue.wasm ./vue
