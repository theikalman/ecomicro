FROM golang:1.18-alpine as builder
WORKDIR /go/src/github.com/theikalman/ecomicro/service-product/
COPY . ./
RUN go mod vendor
RUN go build -o service ./cmd/product/main.go

FROM alpine:3.15.4
WORKDIR /svc/
COPY --from=builder /go/src/github.com/theikalman/ecomicro/service-product/service ./
CMD ./service
