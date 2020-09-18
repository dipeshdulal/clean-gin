FROM golang:alpine

RUN echo $GOPATH

COPY . /clean_web

WORKDIR /clean_web

RUN go mod download

RUN go get github.com/githubnemo/CompileDaemon

CMD $GOPATH/bin/CompileDaemon --directory=/clean_web --command=/clean_web/clean-gin