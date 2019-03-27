FROM golang:1.10-stretch AS builder
WORKDIR /go/src/github.com/joerx/mock_http
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /bin/mock_http .

FROM debian:stretch
WORKDIR /bin
COPY --from=builder /bin/mock_http .
ENTRYPOINT ["/bin/mock_http"]
