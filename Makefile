BINARY_NAME=server.exe

build: 
	go build -o ${BINARY_NAME} cmd/apiserver/main.go

run:
	go run cmd/apiserver/main.go

run-binary:
	./${BINARY_NAME}

clean:
	rm ${BINARY_NAME}

test:
	go test ./...

test-coverage:
	go test ./... -coverprofile=coverage.out