all: build

run: build
	./remote-trackpad-wasm

build: wasm
	go build -o remote-trackpad-wasm

wasm: wasm-dep
	GOOS=js GOARCH=wasm go build -o assets/main.wasm ./wasm/main.go

wasm-dep:
	cp "$(shell go env GOROOT)/misc/wasm/wasm_exec.js" assets/

clean:
	rm remote-trackpad-wasm 2>/dev/null || exit 0
