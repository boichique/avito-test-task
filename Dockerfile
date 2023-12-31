FROM golang:1.20-alpine AS builder

WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o service .

FROM alpine:3.14 AS final

COPY --from=builder /build/service /bin/service

ENTRYPOINT ["/bin/service"]