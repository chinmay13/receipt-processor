# Start by using the official Golang image
# Use an alpine image for smaller size
FROM golang:1.23.2

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files first to leverage Docker layer caching
COPY go.mod go.sum ./

# Download all dependencies; this is cached until go.mod/go.sum changes
RUN go mod download

# Copy the entire project into the container
COPY . .

# Build the Go application
RUN go build -o main .

# Expose the application on port 8080 (or whatever port your Gin app uses)
EXPOSE 8080

# Run the executable
CMD ["./main"]
