# Use the official Go 1.20 image as the base image
FROM golang:1.20

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . /app

# Build the main.go program
RUN go build -o main .

# Run the main.go program
CMD ["./main"]
