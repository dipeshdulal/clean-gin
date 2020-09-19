FROM golang:alpine

# Required because go requires gcc to build
RUN apk add build-base

RUN apk add inotify-tools

RUN echo $GOPATH

COPY . /clean_web

WORKDIR /clean_web

RUN go mod download

RUN go get github.com/go-delve/delve/cmd/dlv

CMD sh /clean_web/docker/run.sh