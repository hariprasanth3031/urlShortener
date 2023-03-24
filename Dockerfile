FROM golang:1.19

WORKDIR /app

RUN mkdir -p /go/src/urlshortener

COPY . /go/src/urlshortener
WORKDIR /go/src/urlshortener

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN tail -1 .git/logs/HEAD > .git/version

RUN mkdir /go/logs && GO111MODULE=on go build -o /go/bin -mod vendor

ENTRYPOINT /go/bin/urlshortener

EXPOSE 8003

CMD [ "/urlshortener" ]