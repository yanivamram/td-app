FROM golang:1.21 as builder

WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /client ./client/main.go

FROM gcr.io/distroless/static
COPY --from=builder /client /client
ENTRYPOINT ["/client"]
