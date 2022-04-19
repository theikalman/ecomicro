FROM golang:1.18-alpine as builder
WORKDIR /go/src/github.com/theikalman/ecomicro/service-user/
COPY . ./
RUN go mod vendor
RUN go build -o service ./cmd/user/main.go

FROM alpine:3.15.4
WORKDIR /svc/
COPY --from=builder /go/src/github.com/theikalman/ecomicro/service-user/service ./
CMD ./service
