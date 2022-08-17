FROM golang:1.17-alpine3.15 AS builder

ARG APOD_API_KEY

COPY . /github.com/apod/
WORKDIR /github.com/apod/

RUN go mod download
RUN GOOS=linux go build -o ./.bin/service ./cmd/main.go

FROM alpine:3.15.3

WORKDIR /root/

COPY --from=builder /github.com/apod/.bin/service .

EXPOSE 8080
CMD ["./service"]