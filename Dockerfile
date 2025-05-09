# Dockerfile

# ---- Build Stage ----
# Use an official Go image as a builder.
# Using alpine variant for smaller size.
# Pinning to a specific version is good practice, e.g., golang:1.22-alpine
FROM golang:alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the source code into the container
COPY main.go .

# Build the Go application
# -o /app/server: specifies the output file name and path
# CGO_ENABLED=0: disables Cgo to build a statically-linked binary (important for scratch/alpine images)
# GOOS=linux: specifies the target operating system as Linux
# -ldflags="-w -s": reduces the size of the binary by stripping debug information
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /app/server .

# ---- Run Stage ----
# Use a minimal image for the final stage. Alpine is small and has a shell.
# Scratch is even smaller but has no shell or any other binaries,
# which can make debugging harder.
FROM alpine:latest

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/server /app/server

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
# This is the command that will be executed when the container starts
CMD ["/app/server"]
