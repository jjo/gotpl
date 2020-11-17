FROM golang:1.15-alpine as builder
WORKDIR /go/src/github.com/tsg/gotpl
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o gotpl .

FROM alpine:3.7
COPY --from=builder /go/src/github.com/tsg/gotpl/gotpl /usr/local/bin/gotpl
WORKDIR /work
ENTRYPOINT ["/usr/local/bin/gotpl"]
