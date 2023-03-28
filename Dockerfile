FROM golang:alpine

RUN mkdir /urlshortener

WORKDIR /urlshortener

ADD go.mod .
ADD go.sum .

RUN go mod download
ADD . .
RUN mkdir /go/logs && GO111MODULE=on go build -o /go/bin -mod vendor

ENTRYPOINT /go/bin/urlshortener

EXPOSE 8080
