# Implementação de web assembly em golang para gerar senhas parametrizaveis

## Como rodar

### Docker

```bash
docker build -t webassembly-go .
docker run -p 8080:8080 webassembly-go
```

### Local

```bash
make download-wasm-exec && make build-assembly
go run cmd/server/main.go
```

## Como usar

Acesse `http://localhost:8080` e preencha os campos com as informações desejadas.

## Testes

```bash
go test ./...
```

## Benchmark

```
=== RUN   BenchmarkGeneratePassword
BenchmarkGeneratePassword
BenchmarkGeneratePassword-4        79638             14239 ns/op             239 B/op          4 allocs/op
PASS
ok      github.com/kauebonfimm/web-assembly-with-golang/internal/core/password_engine   1.294s
```

## Referências

- [WebAssembly](https://webassembly.org/)
- [Golang WebAssembly](https://go.dev/blog/wasi)
- [Golang WebAssembly - Exemplo](https://golangbot.com/webassembly-using-go/)


