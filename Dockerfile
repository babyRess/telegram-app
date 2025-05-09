# Use the official Go image as the base
FROM golang:1.22

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum to the working directory
COPY go.mod go.sum ./

# Install dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the application with explicit output directory and name
RUN mkdir -p /app/bin && \
  go build -o /app/bin/bot_service ./bot_service/main.go && \
  chmod +x /app/bin/bot_service

# Expose the port for the health check server
EXPOSE 8080

# Command to run the bot
CMD ["/app/bin/bot_service"]