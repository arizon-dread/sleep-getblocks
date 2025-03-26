FROM golang:1.24-alpine AS build
LABEL MAINTAINER github.com/arizon-dread

WORKDIR /usr/local/go/src/github.com/arizon-dread/sleep-getblocks
COPY models ./models
COPY main.go go.mod go.sum ./

RUN apk update && apk add --no-cache git
RUN go build -v -o /usr/local/bin/sleep-getblocks/ ./...


FROM golang:1.24-alpine AS final
WORKDIR /go/bin
ENV GIN_MODE=release
RUN apk add --no-cache libc6-compat musl-dev
COPY --from=build /usr/local/bin/sleep-getblocks/ /go/bin/
EXPOSE 8080

ENTRYPOINT [ "sleep-getblocks" ]
