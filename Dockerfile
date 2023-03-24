FROM golang:1.18

RUN mkdir -p /urlshortener

COPY . /urlshortener
WORKDIR /urlshortener
RUN mv  .netrc ~/.netrc

RUN tail -1 .git/logs/HEAD > .git/version

RUN mkdir /go/logs && GO111MODULE=on go build -o /go/bin -mod vendor

ENTRYPOINT /go/bin/urlshortener

EXPOSE 8003

