# syntax=docker/dockerfile:1
# NOTE: Don't run portician in a container on Windows (see README)

FROM golang:1.19-alpine AS base

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY main.go .
COPY portician ./portician
COPY porticianconfig ./porticianconfig

RUN go build -o /app/portician

FROM alpine:latest
WORKDIR /app
COPY --from=base /app/portician .
CMD [ "./portician", "config.json" ]