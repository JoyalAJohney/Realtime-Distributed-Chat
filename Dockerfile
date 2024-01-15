# Use the official Golang image to create a build artifact.
FROM golang:1.19-alpine

# Set the working directory inside the container.
WORKDIR /app

# Copy go.mod and go.sum files to the working directory.
COPY go.mod go.sum ./

# Download all dependencies.
RUN go mod download

# Copy the source code into the container.
COPY src/ ./src

# Set the working directory to the source code directory.
WORKDIR /app/src

# Build the application.
RUN go build -o main .

# Run the executable.
CMD ["./main"]
