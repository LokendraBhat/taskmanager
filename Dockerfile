FROM golang:1.25-alpine AS builder

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/server


FROM alpine:3.22

RUN adduser -D -u 1000 appuser

WORKDIR /home/appuser

COPY --from=builder --chown=appuser:appuser /app/main .
COPY --from=builder --chown=appuser:appuser /app/web/templates .

USER appuser

EXPOSE 8080

CMD ["./main"]
