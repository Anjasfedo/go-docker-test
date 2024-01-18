# Use the specified Go version as the base image
FROM golang:1.21.5-alpine3.19

# Set the working directory inside the container
WORKDIR /app

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o main .

# Command to run the executable
CMD ["/app/main"]
