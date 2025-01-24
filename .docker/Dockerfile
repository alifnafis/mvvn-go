### STAGE 1: Build ###
FROM golang:1.23-alpine AS build

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum for dependency management
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the entire application source code
COPY . .

# Build the Go application
RUN go build -o /main

### STAGE 2: Run ###
FROM golang:1.23-alpine

# Set the working directory
WORKDIR /app

# Copy the compiled binary from the build stage
COPY --from=build /main ./main

# Copy the .env file into the container
COPY .env .env

# Expose the application port
EXPOSE 80

# Run the application
CMD [ "./main" ]
