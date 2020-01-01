FROM golang:1.10.4

ENV APP kylin
ADD ./ /go/src/$APP
WORKDIR /go/src/$APP
ENTRYPOINT ["go", "run", "main.go"]
