# Stage 1: Build the Go binary
FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
# Compile the application as a static binary
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Stage 2: Minimal secure runtime
FROM scratch
WORKDIR /app
# Copy the built binary and necessary certificates from the builder stage
COPY --from=builder /app/main .
# Need to copy SSL certificates if your app makes HTTPS calls
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Expose the application port
EXPOSE 8080

# Run the application
CMD ["./main"]
