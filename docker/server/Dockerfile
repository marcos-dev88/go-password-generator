FROM golang

WORKDIR /passgen-app/src/go-api-password-generator

#COPY . .

ENV GOPATH=/passgen-app

#ENV PORT 8095

RUN go build -ldflags "-s -w" -o bin/passgen cmd/*.go

ENTRYPOINT ["./bin/passgen", "genstrong"]

EXPOSE 8095
EXPOSE 8096