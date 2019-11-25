FROM golang:1.13.4-alpine3.10

ENV GOPROXY "https://goproxy.io"
WORKDIR /go/src/ecovacs.com

COPY . .

RUN go mod download

RUN cd /go/src/ecovacs.com/corpus_platform && go-wrapper download && go-wrapper install

RUN apk add --no-cache tzdata

WORKDIR /go/src/ecovacs.com/corpus_platform/docker

ENTRYPOINT ["./entrypoint.sh"]
