FROM golang:1.20-alpine

RUN go install github.com/jackc/tern/v2@latest

COPY sql /sql

ENTRYPOINT ["tern"]