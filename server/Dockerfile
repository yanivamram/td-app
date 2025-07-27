# Stage 1: Build the Go binary
FROM golang:1.21 AS builder

WORKDIR /app

# Copy go mod files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the server binary
WORKDIR /app/server
RUN CGO_ENABLED=0 GOOS=linux go build -o /server

# Stage 2: Minimal runtime image
FROM gcr.io/distroless/static

COPY --from=builder /server /server

EXPOSE 50051
ENTRYPOINT ["/server"]
