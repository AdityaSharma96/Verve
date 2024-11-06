# Use an official Go runtime as a parent image
FROM golang:1.18-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy source code into the container
COPY . .

# Build the Go application
RUN go build -o verve

# Expose port 8080
EXPOSE 8080

# Run the executable
CMD ["./verve"]
