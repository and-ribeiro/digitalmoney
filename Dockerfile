FROM golang:latest
LABEL authors="mrand"

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download the Go dependencies
RUN go mod download

# Copy the rest of the project files
COPY . .

# Build the Go binary
RUN go build -o app .

# Expose a port for the application
EXPOSE 8080

# Define the command to run the application
CMD ["./app"]