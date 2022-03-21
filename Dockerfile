FROM golang:1.17.6-alpine

RUN apk update && \
    apk add --no-cache git make gcc musl-dev && \
    rm -rf /var/cache/apk/*

WORKDIR /app
COPY . .

RUN go get -u github.com/cosmtrek/air && \
    go build -o /go/bin/air github.com/cosmtrek/air

EXPOSE 8080

CMD air -c .air.toml