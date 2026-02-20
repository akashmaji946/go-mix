---
title: "I/O"
layout: default
parent: Standard Library
nav_order: 11
description: "Input/output operations including scanln, scanf, input, sprintf"
permalink: /standard-library/io/
---

# I/O Package
{: .no_toc }

Input/output operations including scanln, scanf, input, sprintf
{: .fs-6 .fw-300 }

## Table of Contents
{: .no_toc .text-delta }

1. TOC
{:toc}

---

## Import

`import "io"`
{: .fs-5 .fw-300 }

Import the io package to use namespaced functions for reading input, formatted output, and error stream operations.

```go
// Standard import
import io;
var name = io.scanln();
io.printf("Hello %s", name);

// With alias
import io as i;
var name = i.scanln();
i.printf("Hello %s", name);
```

---

## scanln

`scanln() -> string`
{: .fs-5 .fw-300 }

Reads a complete line from standard input, returning the text without the trailing newline. Blocks until the user presses Enter. Returns an empty string if EOF is encountered.

```go
println("Enter your name:");
var name = scanln();
println("Hello, " + name);

// Reading multiple lines
println("Enter first line:");
var line1 = scanln();
println("Enter second line:");
var line2 = scanln();
```

---

## gets

`gets() -> string`
{: .fs-5 .fw-300 }

Alias for scanln(). Reads a complete line from standard input. Provided for C-style familiarity.

```go
// gets is equivalent to scanln
var name = gets();         // Same as scanln()
println("Hello, " + name);
```

---

## scanf

`scanf(format) -> array`
{: .fs-5 .fw-300 }

Reads formatted input from stdin according to the format string and returns an array of parsed values. Supports %d (int), %f (float), %s (string), and other standard format specifiers.

```go
printf("Enter age and name: ");
var data = scanf("%d %s");
var age = data[0];         // Integer
var name = data[1];        // String

// Multiple values
printf("Enter x y z: ");
var coords = scanf("%f %f %f");
var x = coords[0];
var y = coords[1];
var z = coords[2];
```

---

## input

`input(prompt) -> string`
{: .fs-5 .fw-300 }

Displays a prompt message and reads a line of user input. Combines output and input in one convenient function. The prompt is displayed without a trailing newline.

```go
var name = input("Enter your name: ");
var age = input("Enter your age: ");
var confirm = input("Continue? (y/n): ");

// Building interactive menus
var choice = input("Select option (1-3): ");
```

---

## scan

`scan(delimiter) -> string`
{: .fs-5 .fw-300 }

Reads from standard input until the specified delimiter string is encountered. The delimiter is NOT consumed from the input buffer. Useful for parsing custom-terminated input.

```go
var text = scan(";;");
// Reads until ";;" is encountered
// Delimiter remains in buffer for next read

// Multi-character delimiters
var paragraph = scan("\n\n");
// Reads until double newline
```

---

## getchar

`getchar() -> string`
{: .fs-5 .fw-300 }

Reads a single character from standard input and returns it as a one-character string. Returns nil if EOF is encountered. Useful for character-by-character processing or single-key input.

```go
var ch = getchar();        // Read one character
println("You typed: " + ch);

// Character-by-character processing
while (ch != "q") {
    println("Char: " + ch);
    ch = getchar();
}
```

---

## putchar

`putchar(char_or_int) -> nil`
{: .fs-5 .fw-300 }

Outputs a single character to standard output. Accepts either a character literal, a string (outputs first character), or an integer (outputs corresponding Unicode character). Output is immediately flushed.

```go
putchar('A');              // Prints: A
putchar(65);               // Prints: A (ASCII 65)
putchar("Hello");          // Prints: H (first char only)

// Building strings character by character
putchar('H');
putchar('i');
// Output: Hi
```

---

## puts

`puts(...args) -> nil`
{: .fs-5 .fw-300 }

Alias for println(). Prints arguments with a trailing newline. Provided for C-style familiarity.

```go
// puts is equivalent to println
puts("Hello World");       // Same as println("Hello World")
puts("Value:", 42);        // Same as println("Value:", 42)
```

---

## sprintf

`sprintf(format, ...args) -> string`
{: .fs-5 .fw-300 }

Returns a formatted string using C-style format specifiers without printing it. Useful for building strings dynamically, creating messages, or formatting data for storage. Supports %s (string), %d (integer), %f (float), %x (hex), and more.

```go
var msg = sprintf("Hello, %s! You are %d years old.", "Alice", 30);
// msg = "Hello, Alice! You are 30 years old."

// Building complex messages
var filename = "data.txt";
var size = 1024;
var info = sprintf("File: %s, Size: %d bytes", filename, size);

// Formatting numbers
var hex = sprintf("0x%x", 255);     // "0xff"
var pi = sprintf("%.2f", 3.14159); // "3.14"
```

---

## eprintln

`eprintln(...args) -> nil`
{: .fs-5 .fw-300 }

Prints arguments to standard error (stderr) with a trailing newline. Useful for logging errors, warnings, or debug information separately from normal program output. Output is unbuffered and appears immediately.

```go
// Error reporting
if (error_occurred) {
    eprintln("ERROR: Failed to open file");
}

// Debug logging
eprintln("DEBUG: Variable x =", x);

// Progress indicators on stderr while output goes to stdout
eprintln("Processing...");  // Goes to stderr
println("Result: done");     // Goes to stdout
```

---

## eprintf

`eprintf(format, ...args) -> nil`
{: .fs-5 .fw-300 }

Prints a formatted string to standard error (stderr). Combines the formatting capabilities of sprintf with stderr output. Ideal for error messages with variable data, formatted warnings, or structured logging to error streams.

```go
// Formatted error messages
var filename = "config.txt";
eprintf("ERROR: Cannot read file '%s'\n", filename);

// Status reporting with codes
var code = 404;
eprintf("HTTP Error %d: Not Found\n", code);

// Debug with formatting
eprintf("DEBUG: Array has %d elements\n", size_array(arr));
```

---

## flush

`flush() -> nil`
{: .fs-5 .fw-300 }

Clears the input buffer and flushes the output writer. Discards any buffered characters waiting to be read from stdin, ensuring the next input operation reads fresh user input. Also forces any buffered output to be written immediately. Useful before critical input operations or when switching between input modes.

```go
// Clear pending input before important read
flush();
var password = input("Enter password: ");

// Ensure output appears before long operation
printf("Processing");
flush();
// ... long operation ...

// Reset input state
flush();
println("Input buffer cleared");
```

---
