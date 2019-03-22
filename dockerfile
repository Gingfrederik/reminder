FROM golang:1.11.5 AS builder
RUN mkdir /bot
WORKDIR /bot
ENV GO111MODULE=on 

COPY go.mod . 
COPY go.sum .

RUN go mod download
COPY . .

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags "-s -w" -o main

FROM alpine:latest
RUN apk update && \
    apk upgrade && \
    apk add --no-cache ca-certificates tzdata &&\
    mkdir -p /bot
RUN cp /usr/share/zoneinfo/Asia/Taipei /etc/localtime
RUN echo "Asia/Taipei" >  /etc/timezone
COPY --from=builder /bot/main /bot/main
RUN addgroup -g 1000 appuser && \
    adduser -D -u 1000 -G appuser appuser && \
    chown -R appuser:appuser /bot
USER appuser

WORKDIR /bot
CMD ./main
