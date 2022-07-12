GOPATH=$(shell go env GOROOT)
FILE_NAME="cyber-kill-client.wasm"

build:
	cp "$(GOPATH)/misc/wasm/wasm_exec.js" web-client/public/wasm/wasm_exec.js \
	&& GOOS=js GOARCH=wasm go build -o web-client/public/wasm/$(FILE_NAME) web-client.go

clean:
	rm web-client/public/wasm/$(FILE_NAME) \
	&& rm web-client/public/wasm/wasm_exec.js