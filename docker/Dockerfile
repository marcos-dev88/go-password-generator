FROM golang:latest

WORKDIR /passgen-app/src/go-api-password-generator

ENV GOPATH=/passgen-app

RUN go build -ldflags "-s -w" -o bin/passgen cmd/*.go

ENTRYPOINT ["./bin/passgen", "genstrong"]