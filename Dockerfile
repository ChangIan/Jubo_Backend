FROM golang:latest

WORKDIR $GOPATH/src/jubo

ADD . $GOPATH/src/jubo

RUN go env -w GO111MODULE=on

RUN export GIN_MODE=release

RUN go mod vendor

RUN go build -mod=mod

EXPOSE 8899

ENTRYPOINT  ["./jubo"]