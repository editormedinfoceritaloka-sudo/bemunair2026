# Stage 1: Builder
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Install git for fetching dependencies if needed
RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -o ./bin/server ./cmd/main.go

# Stage 2: Runner
FROM alpine:3.22 AS runner

RUN apk add --no-cache ca-certificates tzdata

# Create a non-root user
RUN adduser -D -g '' nonroot
USER nonroot

WORKDIR /app

COPY --from=builder /app/bin/server ./server

EXPOSE 8080

HEALTHCHECK --interval=15s --timeout=5s --start-period=20s --retries=10 \
  CMD wget -qO- http://127.0.0.1:8080/ping || exit 1

CMD ["./server"]
