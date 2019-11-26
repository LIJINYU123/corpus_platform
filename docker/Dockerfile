FROM golang:1.13.4-alpine3.10

ENV GOPROXY "https://goproxy.io"
ENV GOPATH  "/go/src"
ENV GOBIN   "/go/bin"

COPY . /go/src/ecovacs.com/corpus_platform

WORKDIR /go/src/ecovacs.com/corpus_platform

COPY docker/entrypoint.sh /usr/local/bin/entrypoint.sh
COPY docker/env.template /usr/local/bin/env.template

RUN go mod download

RUN apk add --no-cache tzdata



WORKDIR /usr/local/bin

CMD ["./entrypoint.sh", "env.template"]
