# Using Golang image as a base image
FROM golang:1.22.4

# Set the Current Working Directory inside the container
WORKDIR /Users/user/Practice/proj

# Copy go mod and sum files
COPY go.mod ./

# Download all dependencies. 
RUN go mod download

# Copy the source code to the container
COPY . .

# Build the Go app
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
