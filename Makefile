
build:
    go build -v ./cmd/apiserver



test:
    go test -v -race -timeout 30s ./ ...
run: 
    go run cmd/apiserver/main.go

