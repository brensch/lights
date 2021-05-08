# syntax=docker/dockerfile:1
FROM golang:1.16 AS builder
WORKDIR /build
RUN go get -d -v github.com/brensch/lights  
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /build/app .
CMD ["./app"]  