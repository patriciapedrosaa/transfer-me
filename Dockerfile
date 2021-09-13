FROM golang:1.17.0-alpine as builder

LABEL maintainer="Patrícia Pedrosa"

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o main ./cmd

EXPOSE 8000

CMD ["./main"]