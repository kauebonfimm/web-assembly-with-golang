FROM golang:1.21-alpine as builder

WORKDIR /app

COPY . .

RUN apk add curl make

RUN make download-wasm-exec && make build-assembly

RUN go mod tidy && go mod download

RUN go build -o /main ./cmd/server/main.go

FROM alpine:latest

WORKDIR /app


COPY --from=builder /main /app/main
COPY --from=builder /app/assets /app/assets

CMD [ "/app/main" ]