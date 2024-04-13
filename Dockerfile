# Use the official Golang image as a base image
FROM golang:latest

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

# Build the Go app
RUN go build -o main ./cmd/web
EXPOSE 3000

CMD ["./main"]