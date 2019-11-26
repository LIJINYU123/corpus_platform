FROM golang:1.13.4-alpine3.10

ENV GOPROXY "https://goproxy.io"

COPY . /go/src/ecovacs.com/corpus_platform

WORKDIR /go/src/ecovacs.com/corpus_platform

RUN go mod download

RUN apk add --no-cache tzdata

RUN chmod +x /go/src/ecovacs.com/corpus_platform/docker/entrypoint.sh

