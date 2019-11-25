FROM golang:1.13.4-alpine3.10

COPY . /go/src/ecovacs.com/corpus_platform

RUN cd /go/src/ecovacs.com/corpus_platform && go-wrapper download && go-wrapper install

FROM alpine:latest
RUN apk add --no-cache tzdata

WORKDIR /go/src/ecovacs.com/corpus_platform/docker

ENTRYPOINT ["./entrypoint.sh"]
