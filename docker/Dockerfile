FROM golang:1.13.4-alpine3.10

ENV GOPROXY "https://goproxy.io"
ENV GOBIN   "/go/bin"

COPY . /go/src/ecovacs.com/corpus_platform

WORKDIR /go/src/ecovacs.com/corpus_platform

RUN go mod download

RUN apk add tzdata

RUN go install

WORKDIR /usr/local/bin

COPY /go/bin/corpus_platform corpus_platform
COPY /go/src/ecovacs.com/corpus_platform/docker/entrypoint.sh entrypoint.sh
COPY /go/src/ecovacs.com/corpus_platform/docker/env.template env.template

CMD ["./entrypoint.sh", "env.template"]
