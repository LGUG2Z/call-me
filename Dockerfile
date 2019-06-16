FROM golang:alpine

WORKDIR $GOPATH/src/github.com/LGUG2Z/call-me
COPY . .

RUN apk update && apk add --no-cache git

ENV GO111MODULE=on
RUN CGO_ENABLED=0 GOOS=linux go install ./cmd/call-me-server
RUN CGO_ENABLED=0 GOOS=linux go install ./cmd/call-me-client

ENTRYPOINT ["/go/bin/call-me-server"]
