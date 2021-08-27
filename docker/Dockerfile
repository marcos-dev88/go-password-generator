FROM golang

ADD . app/src/go-api-password-generator

COPY . app/src/go-api-password-generator

WORKDIR app/src/go-api-password-generator

ENV GOPATH=/app

RUN go build cmd/main.go

ENTRYPOINT ["./main"]

EXPOSE 8095