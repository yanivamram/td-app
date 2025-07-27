FROM golang:1.21 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /client ./client

FROM gcr.io/distroless/static

COPY --from=builder /app/client /client

ENTRYPOINT ["/client"]
