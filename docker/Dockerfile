FROM golang

ADD . passgen-app/src/go-api-password-generator

COPY . passgen-app/src/go-api-password-generator

WORKDIR passgen-app/src/go-api-password-generator

ENV GOPATH=/passgen-app

RUN go build cmd/main.go

ENTRYPOINT ["./main"]

EXPOSE 8095