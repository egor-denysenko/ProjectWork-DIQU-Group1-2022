# syntax=docker/dockerfile:1
FROM golang:1.18-alpine AS builder
WORKDIR /go/src/
COPY . .
RUN mkdir build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -o ./build .

#FROM arm32v7/alpine:3.16
FROM scratch
WORKDIR /cmd/
# getting the compiled go excecutable from the builder stage container
COPY --from=builder /go/src/build .
CMD ["./serialReciever"]