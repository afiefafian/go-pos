# Builder image
FROM golang:1.18.1-alpine3.15 as builder
RUN apk add git
RUN mkdir /build
ADD . /build
WORKDIR /build
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Generate clean, final image for end users
FROM alpine:latest
COPY --from=builder /build/main .

# executable
CMD ["./main"]
