FROM golang:1.10 as builder
WORKDIR /go/src/github.com/tsg/gotpl
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o gotpl .

FROM alpine:3.7
COPY --from=builder /go/src/github.com/tsg/gotpl/gotpl /usr/local/bin/gotpl
ENTRYPOINT ["/usr/local/bin/gotpl"]
