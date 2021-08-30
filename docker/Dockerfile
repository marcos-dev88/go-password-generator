FROM golang

ADD . passgen-app/src/go-api-password-generator

COPY . passgen-app/src/go-api-password-generator

WORKDIR passgen-app/src/go-api-password-generator

ENV GOPATH=/passgen-app

RUN go build -ldflags "-s -w" cmd/*.go

ENTRYPOINT ["./main"]

EXPOSE 8095