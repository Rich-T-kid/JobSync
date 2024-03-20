FROM golang:1.18

WORKDIR /app

# Copy only the go.mod and go.sum files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project (source code)
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-go-server .

# Expose port 8080
EXPOSE 8080

# Run the executable
CMD ["/docker-go-server"]
