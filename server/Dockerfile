# Use official Golang image for building
FROM golang:1.21 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

WORKDIR /app/server
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /server .

# Use a minimal base image to run
FROM gcr.io/distroless/static

COPY --from=builder /server /server

EXPOSE 50051
ENTRYPOINT ["/server"]
