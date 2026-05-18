# Stage 1: Builder
FROM golang:1.26-alpine AS builder

WORKDIR /app

# Install git for fetching dependencies if needed
RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -o ./bin/server ./cmd/main.go

# Stage 2: Runner
FROM alpine:latest AS runner

RUN apk add --no-cache ca-certificates tzdata

# Create a non-root user
RUN adduser -D -g '' nonroot
USER nonroot

WORKDIR /app

COPY --from=builder /app/bin/server ./server

EXPOSE 8080

CMD ["./server"]
