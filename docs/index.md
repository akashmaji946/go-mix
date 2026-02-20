---
title: Home
layout: default
nav_order: 1
description: "Go-Mix — A high-performance interpreted programming language built in Go"
permalink: /
---

# Go-Mix Programming Language
{: .fs-9 }

A high-performance, interpreted language implemented in Go. Features dynamic typing with static safety, lexical closures, first-class functions, and a comprehensive standard library with 100+ built-in functions.
{: .fs-6 .fw-300 }

[Get Started]({{ site.baseurl }}/getting-started){: .btn .btn-primary .fs-5 .mb-4 .mb-md-0 .mr-2 }
[View on GitHub](https://github.com/akashmaji946/go-mix){: .btn .fs-5 .mb-4 .mb-md-0 }

---

## Why Go-Mix?

Designed for rapid prototyping, education, and embedding as a scripting layer in Go applications.

| Feature | Description |
|:--------|:------------|
| **High Performance** | Implemented in Go with efficient Pratt parser and optimized evaluation engine |
| **Hybrid Type System** | Three variable kinds: `var` (dynamic), `let` (static), `const` (immutable) |
| **First-Class Functions** | Full support for higher-order functions, closures, and functional patterns |
| **Rich Collections** | Arrays, lists, tuples, maps, and sets with comprehensive manipulation functions |
| **100+ Built-in Functions** | 17 standard packages: strings, math, file I/O, HTTP, regex, JSON, and more |
| **Interactive REPL** | Built-in Read-Eval-Print Loop for rapid prototyping and experimentation |

---

## Quick Example

```go
// Variables and types
var name = "World";
let count = 42;
const PI = 3.14159;

// Functions
func greet(name) {
    return "Hello, " + name + "!";
}

// Arrays and functional operations
var numbers = [1, 2, 3, 4, 5];
var doubled = map(numbers, func(x) {
    return x * 2;
});

// Print results
println(greet(name));
println("Doubled: " + doubled);
```

```go
// Recursive Fibonacci with closures
func fibonacci(n) {
    if (n <= 1) {
        return n;
    }
    return fibonacci(n - 1) + fibonacci(n - 2);
}

// Higher-order function
func makeMultiplier(factor) {
    return func(x) { return x * factor; };
}

var triple = makeMultiplier(3);

for (var i = 0; i < 10; i = i + 1) {
    printf("fib(%d) = %d, triple = %d\n",
           i, fibonacci(i), triple(i));
}
```

---

## At a Glance

| | |
|:--|:--|
| **100+** Built-in Functions | **17** Standard Packages |
| **50+** Sample Programs | **MIT** Open Source License |

---

## Ready to Start?

Get up and running with Go-Mix in minutes. Install, write your first program, and explore the possibilities.

[Get Started Now]({{ site.baseurl }}/getting-started){: .btn .btn-primary .fs-5 .mb-4 .mb-md-0 .mr-2 }
[Read the Language Guide]({{ site.baseurl }}/language-guide){: .btn .fs-5 .mb-4 .mb-md-0 }
