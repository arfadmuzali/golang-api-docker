FROM golang:1.24.4 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

COPY .env .env

ENV AUTH_SECRET="secretkeyk"
ENV GOOGLE_APP_PASSWORD="rypd lmtm xeot dftf"


RUN go build -o /app/api ./cmd/api
RUN go build -o /app/worker ./cmd/worker

FROM debian:bookworm-slim

WORKDIR /app

RUN apt update && apt install -y ca-certificates

COPY --from=builder /app/api /app/api
COPY --from=builder /app/worker /app/worker
COPY --from=builder /app/.env /app/.env



CMD ["./api"]

