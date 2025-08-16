# Step 1: Build the Go application
FROM golang:1.18 AS build

WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./
RUN go mod tidy

# Copy the source code
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Step 2: Create the production image
FROM alpine:latest

WORKDIR /root/

# Install any dependencies you need in production (e.g., certificates, etc.)
RUN apk --no-cache add ca-certificates

# Copy the compiled binary from the build stage
COPY --from=build /app/main .

# Expose port 8080
EXPOSE 8080

# Command to run the app
CMD ["./main"]
