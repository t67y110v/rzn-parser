BINARY_NAME=server.exe
USER_NAME=vova

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

swag:	
	/home/${USER_NAME}/go/bin/swag init -g cmd/apiserver/main.go
