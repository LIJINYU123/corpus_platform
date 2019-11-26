FROM golang:1.13.4-alpine3.10

ENV GOPROXY "https://goproxy.io"
ENV GOPATH  "/go"

COPY . /go/src/ecovacs.com/corpus_platform

WORKDIR /go/src/ecovacs.com/corpus_platform

RUN go mod download

RUN go build -o corpus_platform corpus_platform

RUN apk add --no-cache tzdata

WORKDIR /usr/local/bin

COPY /go/src/ecovacs.com/corpus_platform/corpus_platform corpus_platform
COPY /go/src/ecovacs.com/corpus_platform/docker/entrypoint.sh entrypoint.sh
COPY /go/src/ecovacs.com/corpus_platform/docker/env.template env.template

CMD ["./entrypoint.sh", "env.template"]
