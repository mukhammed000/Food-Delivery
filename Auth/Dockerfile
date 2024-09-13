FROM golang:1.23.1-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main ./cmd

FROM alpine:3.16

WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /app/api/helper/format.html .
COPY --from=builder /app/.env .

EXPOSE 8082

CMD ["./main"]
