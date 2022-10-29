# syntax=docker/dockerfile:1
# NOTE: Don't run portician in a container on Windows (see README)

FROM golang:1.19-alpine AS base

RUN apk add --no-cache make

WORKDIR /build

COPY . .
RUN make build

FROM alpine:latest
WORKDIR /app
COPY --from=base /build/bin/portician .
CMD [ "./portician", "config.json" ]