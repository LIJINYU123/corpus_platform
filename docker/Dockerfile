FROM golang:1.13.4-alpine3.10

ENV GOPROXY "https://goproxy.io"

COPY . /go/src/ecovacs.com/corpus_platform

WORKDIR /go/src/ecovacs.com/corpus_platform

RUN go mod download

RUN apk add --no-cache tzdata

RUN go build

WORKDIR /usr/local/bin

COPY /go/src/ecovacs.com/corpus_platform/corpus-platform corpus-platform
COPY /go/src/ecovacs.com/corpus_platform/docker/entrypoint.sh entrypoint.sh
COPY /go/src/ecovacs.com/corpus_platform/docker/env.template env.template

CMD ["./entrypoint.sh", "env.template"]
