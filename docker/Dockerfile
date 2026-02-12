# Stage 1: Builder
FROM golang:1.18-alpine AS builder

# Install build dependencies
RUN apk add --no-cache git ca-certificates

# Set working directory
WORKDIR /build

# Copy go.mod and go.sum
COPY go.mod ./

# Download dependencies
RUN go mod download

# Copy the entire source code
COPY . .

# Build the go-mix executable
RUN go build -o go-mix ./main

# Stage 2: Runtime
FROM alpine:latest

# Install runtime dependencies
RUN apk add --no-cache ca-certificates

# Set working directory
WORKDIR /app

# Copy the built executable from builder stage
COPY --from=builder /build/go-mix /usr/local/bin/go-mix

# Copy sample files and standard library
COPY --from=builder /build/samples /app/samples

# Make go-mix executable accessible from anywhere
RUN chmod +x /usr/local/bin/go-mix

# Set the default command to REPL
ENTRYPOINT ["go-mix"]

# Default to REPL mode when no arguments provided
CMD []

# Labels for image metadata
LABEL maintainer="akashmaji(@iisc.ac.in)"
LABEL author="Akash Maji"
LABEL version="1.0.0"
LABEL license="MIT"
LABEL homepage="https://github.com/akashmaji946/go-mix"
LABEL description="Go-Mix - An Interpreted Programming Language"
LABEL repository="https://github.com/akashmaji946/go-mix"
