# Stage 1: Build
FROM golang:1.22.5 AS builder

# Set the working directory
WORKDIR /app

# Copy Go mod files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy all other files
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o gateway .

# Stage 2: Create the final image
FROM alpine:latest

# Install necessary packages
RUN apk add --no-cache ca-certificates

# Set the working directory
WORKDIR /app

# Copy the built binary and other necessary files
COPY --from=builder /app/gateway .
COPY --from=builder /app/internal/http/casbin/model.conf ./internal/http/casbin/
COPY --from=builder /app/internal/http/casbin/policy.csv ./internal/http/casbin/
COPY --from=builder /app/internal/app/logs/info.log ./internal/app/logs/

# Copy the environment file
COPY .env .env 

# Set the executable permissions
RUN chmod +x gateway

# Expose the port
EXPOSE 5050

# Command to run the application
CMD ["./gateway"]
