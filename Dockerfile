FROM golang:alpine AS dev

WORKDIR /app

COPY . .

RUN go mod download

RUN go build .

EXPOSE 8080