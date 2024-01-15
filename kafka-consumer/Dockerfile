FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

WORKDIR /app

# Build the application.
RUN go build -o main .

# Run the executable.
CMD ["./main"]