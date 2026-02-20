---
title: Getting Started
layout: default
nav_order: 2
description: "Installation guide and first steps with Go-Mix programming language"
permalink: /getting-started/
---

# Getting Started
{: .no_toc }

Welcome to Go-Mix! This guide will help you install the language and write your first program in minutes.
{: .fs-6 .fw-300 }

## Table of Contents
{: .no_toc .text-delta }

1. TOC
{:toc}

---

## Installation

### Prerequisites

- **Go 1.18 or higher** — Required for building from source
- **Git** — For cloning the repository

### Method 1: Automated Installation (Recommended)

The easiest way to install Go-Mix is using the provided install script:

```bash
# Clone the repository
git clone https://github.com/akashmaji946/go-mix.git
cd go-mix

# Install system-wide (requires sudo)
sudo ./install.sh

# Or install locally to /usr/local/bin
./install.sh
```

{: .note }
> **What the installer does:** Verifies Go installation, builds the binary, installs to `/usr/local/bin`, and makes `go-mix` available from anywhere on your system.

### Method 2: Manual Build

If you prefer to build manually or want to customize the installation:

```bash
# Clone the repository
git clone https://github.com/akashmaji946/go-mix.git
cd go-mix

# Build the binary
go build -o go-mix ./main

# Or use the build script
./build.sh

# Run tests to verify
go test ./...
```

### Method 3: Docker Installation

Run Go-Mix in a container without installing locally:

```bash
# Pull from DockerHub
docker pull akashmaji946/go-mix:latest

# Run a program
docker run -v $(pwd)/samples:/samples akashmaji946/go-mix /samples/algo/05_factorial.gm

# Interactive REPL
docker run -it akashmaji946/go-mix
```

### Verify Installation

Check that Go-Mix is installed correctly:

```bash
# Check version
go-mix --version

# Display help
go-mix --help
```

---

## Quick Start

Once installed, you can use Go-Mix in three ways:

| Mode | Command | Use Case |
|:-----|:--------|:---------|
| **REPL** | `go-mix` | Interactive experimentation |
| **File** | `go-mix file.gm` | Run complete programs |
| **Script** | `#!/usr/bin/env go-mix` | Executable scripts |

---

## Your First Program

Create a file named `hello.gm`:

```go
// My first Go-Mix program

// Variables
var name = "World";
let year = 2024;

// Function definition
func greet(name, year) {
    return "Hello, " + name + "! Welcome to " + year;
}

// Main execution
var message = greet(name, year);
println(message);

// Working with arrays
var numbers = [1, 2, 3, 4, 5];
var doubled = map(numbers, func(x) { return x * 2; });
println("Doubled: " + doubled);

// Calculate sum using reduce
var sum = reduce(numbers, func(acc, x) { return acc + x; }, 0);
println("Sum: " + sum);
```

Run the program:

```bash
go-mix hello.gm
```

Expected output:

```
Hello, World! Welcome to 2024
Doubled: [2, 4, 6, 8, 10]
Sum: 15
```

---

## Using the REPL

The Read-Eval-Print Loop (REPL) is perfect for experimentation:

```bash
$ go-mix
Go-Mix >>> var x = 42
Go-Mix >>> println(x)
42
Go-Mix >>> func square(n) { return n * n; }
Go-Mix >>> square(5)
25
Go-Mix >>> var arr = [1, 2, 3, 4, 5]
Go-Mix >>> map(arr, func(x) { return x * 2; })
[2, 4, 6, 8, 10]
Go-Mix >>> /help
Available commands:
  /exit    - Exit the REPL
  /scope   - Show current scope and variables
Go-Mix >>> /exit
```

### REPL Commands

| Command | Description |
|:--------|:------------|
| `/exit` | Exit the REPL |
| `/scope` | Show current scope and variables |
| `/help` | Show available commands |

---

## Next Steps

{: .note }
> **Learn the Language** — Read the [Language Guide]({{ site.baseurl }}/language-guide/) to understand Go-Mix syntax, types, control flow, functions, and object-oriented programming.

{: .note }
> **Explore the Standard Library** — Discover 100+ built-in functions across 17 packages in the [Standard Library]({{ site.baseurl }}/standard-library/) reference.

{: .note }
> **Study Examples** — Browse 50+ sample programs in the [Samples]({{ site.baseurl }}/samples/) section, from algorithms to web applications.

{: .tip }
> Install the [VS Code Extension](https://github.com/akashmaji946/go-mix/releases/tag/v1.0.0) for syntax highlighting and better development experience.
