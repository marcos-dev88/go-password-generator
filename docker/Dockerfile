FROM golang

WORKDIR /passgen-app/src/go-api-password-generator

COPY . .

ENV GOPATH=/passgen-app

RUN go build -ldflags "-s -w" -o bin/passgen cmd/*.go

ENTRYPOINT ["./bin/passgen", "server"]

EXPOSE 8095