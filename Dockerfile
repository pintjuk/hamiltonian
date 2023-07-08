# Build stage
FROM golang:1.20.5-alpine3.18 AS build

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64


WORKDIR /build

COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY src src
COPY api api

WORKDIR /build/src
RUN go build -o /build/main

# Run stage
FROM alpine:3.18

WORKDIR /app

COPY --from=build /build/main .

# Export necessary port
EXPOSE 8080

# Command to run when starting the container
CMD ["./main"]