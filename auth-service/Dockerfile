FROM golang:1.23-alpine AS builder

WORKDIR /app

# Go mod files are directly in the current folder
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code (all of auth-service)
COPY . .

# Build the Go app
RUN go build -o main cmd/main.go

EXPOSE 8080
EXPOSE 50051

CMD ["./main"]
