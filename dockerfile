# syntax=docker/dockerfile:1
FROM golang:1.16 AS builder
WORKDIR /build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:3.11  
RUN apk --no-cache add ca-certificates && update-ca-certificates
RUN apk add --update tzdata
WORKDIR /root/
COPY --from=builder /build/app .
CMD ["./app"]  