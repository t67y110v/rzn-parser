FROM golang:1.19-alpine3.15 AS builder

COPY . /restApi/
WORKDIR /restApi/

RUN go mod download

RUN go build -o ./bin/server cmd/apiserver/main.go


FROM alpine:latest 

WORKDIR /root/

COPY --from=0 /restApi/bin/server .
COPY --from=0 /restApi/configs configs/ 

EXPOSE 80

CMD ["./server"]


#docker build -t server-api:v0.1 .
#docker run --name server -p 80:80 --env-file configs/apiserver.toml server-api:v0.1