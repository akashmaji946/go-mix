# Go-Mix Docker Guide

This guide explains how to build and run Go-Mix using Docker.


## Quick Start

Build and run Go-Mix in Docker with a single command:

From the go-mix directory:

```bash
# Build the image
docker build -f docker/Dockerfile -t go-mix:v1.0.0 -t go-mix:latest .

# Run interactive REPL
docker run -it go-mix:latest

# Run a file
docker run -v $(pwd):/workspace go-mix:latest /workspace/samples/algo/05_factorial.gm

# Check version
docker run go-mix:latest --version

# Check help
docker run go-mix:latest --help
```
---

## Running Go-Mix in Docker

### Interactive REPL Mode

Start an interactive REPL session:

```bash
docker run -it go-mix:latest
```

Example session:
```
Go-Mix >>> var x = 42
Go-Mix >>> println(x * 2)
84
Go-Mix >>> /exit
```

### Execute a File

Run a Go-Mix script from your local machine:

```bash
# From your project directory
docker run -v $(pwd):/workspace go-mix:latest /workspace/samples/algo/05_factorial.gm
```

Or specify the absolute path:

```bash
docker run -v /path/to/your/files:/workspace go-mix:latest /workspace/script.gm
```


### Mount Working Directory

For interactive development:

```bash
docker run -it -v $(pwd):/workspace -w /workspace go-mix:latest

# Now inside container:
Go-Mix >>> # Can access files in /workspace
```

### Check Version Inside Container

```bash
docker run go-mix:latest --version
```

Output:
```
Go-Mix version v1.0.0
License: MIT
Author: akashmaji(@iisc.ac.in)
```

### Get Help

```bash
docker run go-mix:latest --help
```

---

## Docker Commands Reference

### Image Management

```bash
# List images
docker images | grep go-mix

# Remove an image
docker rmi go-mix:v1.0.0

# View image details
docker inspect go-mix:latest

# Tag an existing image
docker tag go-mix:latest go-mix:v1.0.0
```

### Container Management

```bash
# List running containers
docker ps

# List all containers (including stopped)
docker ps -a

# Remove a container
docker rm <container_id>

# View container logs
docker logs <container_id>

# Stop a running container
docker stop <container_id>
```
