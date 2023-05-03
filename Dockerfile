FROM golang:latest AS builder
WORKDIR /src
ENV CGO_ENABLED=0
COPY go.mod /src/
RUN go mod download
COPY . .
RUN go build -a -o go4hackers -trimpath

FROM alpine:latest

RUN apk add --no-cache ca-certificates \
    && rm -rf /var/cache/*

RUN mkdir -p /app \
    && adduser -D go4hackers \
    && chown -R go4hackers:go4hackers /app

USER go4hackers
WORKDIR /app

COPY --from=builder /src/go4hackers .

ENTRYPOINT [ "./go4hackers" ]