FROM golang:1.13-alpine AS backend

RUN apk --no-cache add build-base libxml2-dev

WORKDIR /go/src/github.com/profbiss/go-xsd-validate

COPY ./ ./

RUN go build ./...