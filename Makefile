build-assembly:
	@echo "Building assembly..."
	GOOS=linux GOARCH=wasm GOOS=js go build -o ./assets/assembly.wasm ./cmd/wasm/main.go
	@echo "Done."

download-wasm-exec:
	@echo "Downloading wasm_exec.js..."
	curl -o ./assets/wasm_exec.js https://raw.githubusercontent.com/golang/go/master/misc/wasm/wasm_exec.js
	@echo "Done."