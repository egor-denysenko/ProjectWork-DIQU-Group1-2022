# syntax=docker/dockerfile:1
FROM golang:1.18-alpine AS builder
WORKDIR /go/src/
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM scratch
WORKDIR /cmd/
# getting the compiled go excecutable from the builder stage container
COPY --from=builder /go/./
CMD ["./app"]