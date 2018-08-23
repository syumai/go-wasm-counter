GO_BIN = go1.11rc2

.PHONY: build
build:
	GOOS=js GOARCH=wasm $(GO_BIN) build -o index.wasm