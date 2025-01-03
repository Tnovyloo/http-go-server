FROM golang:1.22.5

# Set the working directory
WORKDIR /app

# Copy the Go module files and download dependencies
COPY /server/go.mod ./
RUN go mod download

# Copy the application code
COPY ./server .

# Copy static folder
COPY ./static ./static

# Build the application
RUN go build -o server

# Expose the application port
EXPOSE 8080

# Run the server
CMD ["./server"]
