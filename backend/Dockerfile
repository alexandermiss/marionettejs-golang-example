FROM golang:1.19-alpine3.16

LABEL version="1.0"
LABEL maintainer="mail@alexandermiss.com"

WORKDIR /app

COPY . ./

RUN go mod download \
    && go install github.com/cortesi/modd/cmd/modd@latest

EXPOSE 8599

ENTRYPOINT modd
