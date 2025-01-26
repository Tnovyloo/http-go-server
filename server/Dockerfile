# FROM golang:1.22.5

# # Set the working directory
# WORKDIR /app

# # Copy the Go module files and download dependencies
# COPY /server/go.mod ./
# RUN go mod download

# # Copy the application code
# COPY ./server .

# # Copy static folder
# COPY ./static ./static

# # Build the application
# RUN go build -o server

# # Expose the application port
# EXPOSE 8080

# # Run the server
# CMD ["./server"]


FROM golang:1.23-alpine as builder

# Set up the working directory in the Docker image
WORKDIR /server/

# Copy the go.mod file only (to leverage Docker layer caching)
COPY ./go.mod ./

# Download dependencies (this will generate a go.sum file)
RUN go mod download

# Copy the remaining project files
COPY . .

# Install the `air` tool for hot reloading
RUN go install github.com/air-verse/air@latest

EXPOSE 8080

# Start the server using Air with the provided configuration
CMD ["air", "-c", ".air.toml"]




# FROM golang:1.23-alpine as builder

# # Copy everything from current directory to /app direcotry (in dockerimage)
# COPY . .

# # Working directory
# WORKDIR /server/

# RUN go install github.com/air-verse/air@latest

# RUN go mod download
# CMD ["air", "-c", ".air.toml"]