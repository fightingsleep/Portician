# syntax=docker/dockerfile:1

FROM golang:1.19-alpine AS base

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY main.go .
COPY portforwarder ./portforwarder
COPY portforwarderconfig ./portforwarderconfig

RUN go build -o /app/portforwarder

FROM alpine:latest
WORKDIR /app
COPY --from=base /app/portforwarder .
CMD [ "./portforwarder", "config.json" ]