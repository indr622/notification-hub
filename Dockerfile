FROM golang:1.22-alpine3.19 AS builder


RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o notification-hub main.go

FROM alpine:3.19

WORKDIR /root/

RUN apk update && apk upgrade && apk add --no-cache ca-certificates tzdata

COPY --from=builder /app/notification-hub .


CMD ["./notification-hub"]
