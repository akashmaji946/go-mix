---
title: Language Guide
layout: default
nav_order: 3
description: "Complete guide to Go-Mix programming language syntax, features, and constructs"
permalink: /language-guide/
---

# Language Guide
{: .no_toc }

Complete reference for Go-Mix programming language syntax, types, and features.
{: .fs-6 .fw-300 }

## Table of Contents
{: .no_toc .text-delta }

1. TOC
{:toc}

---

## Data Types

Go-Mix supports six primitive types plus `nil`:

| Type | Example | Description |
|:-----|:--------|:------------|
| `int` | `42`, `0xFF` | 64-bit signed integers; hex (0x) and octal (0o) support |
| `float` | `3.14`, `1.2e3` | 64-bit IEEE-754 double precision |
| `bool` | `true`, `false` | Boolean logic values |
| `string` | `"Hello"` | UTF-8 strings with escape sequences |
| `char` | `'A'`, `'\n'` | Single Unicode character |
| `nil` | `nil` | Represents absence of value |

### Type Checking

```go
typeof(42);                 // "int"
typeof(3.14);              // "float"
typeof("text");            // "string"
typeof(true);              // "bool"
typeof(nil);               // "nil"
typeof([1, 2, 3]);         // "array"
typeof(func() {});           // "func"
```

---

## Variables

Go-Mix provides three variable kinds with different semantics:

### Dynamic Variables (`var`)

Type can change during execution. Ideal for rapid prototyping.

```go
var x = 42;        // x is an integer
x = "hello";       // x is now a string
x = [1, 2, 3];     // x is now an array
```

### Static Variables (`let`)

Type is locked on first assignment. Provides safety without explicit type annotations.

```go
let count = 0;
count = count + 1;     // OK
count = "text";        // ERROR: type mismatch
```

### Immutable Constants (`const`)

Cannot be reassigned after initial binding.

```go
const PI = 3.14159;
const MAX_SIZE = 100;
MAX_SIZE = 200;        // ERROR: constants are immutable
```

---

## Operators

### Arithmetic Operators

| Operator | Description | Example |
|:---------|:------------|:--------|
| `+` | Addition | `5 + 3 // 8` |
| `-` | Subtraction | `5 - 3 // 2` |
| `*` | Multiplication | `5 * 3 // 15` |
| `/` | Division | `5 / 2 // 2.5` |
| `%` | Modulo | `5 % 2 // 1` |
| `**` | Power | `2 ** 3 // 8` |

### Comparison Operators

| Operator | Description |
|:---------|:------------|
| `==` | Equal to |
| `!=` | Not equal to |
| `<` | Less than |
| `>` | Greater than |
| `<=` | Less than or equal |
| `>=` | Greater than or equal |

### Logical Operators

| Operator | Description |
|:---------|:------------|
| `&&` | Logical AND (short-circuit) |
| `\|\|` | Logical OR (short-circuit) |
| `!` | Logical NOT |

---

## Control Flow

### If-Else Statements

```go
if (age >= 18) {
    println("Adult");
} else if (age >= 13) {
    println("Teenager");
} else {
    println("Child");
}

// Ternary-like expression
var status = if (x > 0) { "positive"; } else { "non-positive"; };
```

### For Loops

```go
// Standard for loop
for (var i = 0; i < 5; i = i + 1) {
    println(i);  // 0, 1, 2, 3, 4
}

// With break and continue
for (var i = 0; i < 10; i = i + 1) {
    if (i == 3) { continue; }    // Skip iteration
    if (i == 7) { break; }       // Exit loop
    println(i);
}
```

### While Loops

```go
var count = 0;
while (count < 5) {
    println(count);
    count = count + 1;
}

// Multiple conditions
var x = 10, y = 0;
while (x > 0, y < 10) {
    x = x - 1;
    y = y + 1;
}

// Infinite loop with break
var i = 0;
while (true) {
    if (i >= 100) { break; }
    i = i + 1;
}
```

### Foreach Loops

```go
// Range iteration (inclusive)
foreach i in 1...5 {      // 1, 2, 3, 4, 5
    println(i);
}

// Reverse range
foreach i in 5...1 {      // 5, 4, 3, 2, 1
    println(i);
}

// Array iteration with value
var arr = [10, 20, 30];
foreach val in arr {
    println(val);
}

// Array iteration with index and value
foreach idx, val in arr {
    println(idx, val);
}
```

---

## Switch Statements

Go-Mix supports switch statements for multi-way branching. Each case requires a `break` statement to exit.

### Basic Switch

```go
var dayNum = 3;
var dayName = "";

switch (dayNum) {
    case 1:
        dayName = "Monday";
        break;
    case 2:
        dayName = "Tuesday";
        break;
    case 3:
        dayName = "Wednesday";
        break;
    default:
        dayName = "Weekend";
}

println(dayName);  // Wednesday
```

### Switch on Enums

```go
enum TrafficLight {
    RED,
    YELLOW,
    GREEN
}

var light = TrafficLight.GREEN;
var action = "";

switch (light) {
    case TrafficLight.RED:
        action = "Stop";
        break;
    case TrafficLight.YELLOW:
        action = "Caution";
        break;
    case TrafficLight.GREEN:
        action = "Go";
        break;
    default:
        action = "Signal is broken!";
}

println(action);  // Go
```

### Switch with Fallthrough

```go
var num = 2;
switch (num) {
    case 1:
        println("One");
        fallthrough;
    case 2:
        println("Two");
        fallthrough;
    case 3:
        println("Three");
        break;
    default:
        println("Other");
}
// Output:
// Two
// Three
```

---

## Collections

### Arrays

Mutable sequences with homogeneous elements (typically).

```go
// Array literals
var empty = [];
var numbers = [1, 2, 3, 4, 5];
var mixed = [1, "hello", true, 3.14];

// Indexing (0-based)
var arr = [10, 20, 30, 40, 50];
println(arr[0]);           // 10 (first element)
println(arr[2]);           // 30 (third element)

// Negative indexing
println(arr[-1]);          // 50 (last element)
println(arr[-2]);          // 40 (second-to-last)

// Index assignment
arr[1] = 25;               // [10, 25, 30, 40, 50]

// Array methods
push(arr, 60);             // Add to end
pop(arr);                  // Remove from end
unshift(arr, 5);           // Add to start
shift(arr);                // Remove from start
sort(arr);                 // Sort in-place
```

### Maps (Dictionaries)

```go
// Map literals
var user = map{
    "name": "Alice",
    "age": 30,
    "city": "New York"
};

// Key access
println(user["name"]);     // "Alice"

// Adding/updating keys
user["email"] = "alice@example.com";
user["age"] = 31;
```

### Lists

Mutable, heterogeneous sequences.

```go
// Creating lists
var list = list(1, "two", 3.0, true);

// Access like arrays
println(list[0]);          // 1
println(list[1]);          // "two"
```

### Tuples

Immutable, fixed-size sequences.

```go
// Creating tuples
var coords = tuple(10, 20);
var rgb = tuple(255, 128, 64);

// Access elements
println(coords[0]);        // 10
println(coords[1]);        // 20

// Immutable - cannot be modified
coords[0] = 15;            // ERROR
```

---

## Path Operations

Go-Mix provides a comprehensive `path` package for file system operations.

### File I/O Functions

```go
// Reading and writing files
var content = read_file("data.txt");
write_file("output.txt", "Hello, World!");
append_file("log.txt", "New entry\n");

// File operations
file_exists("path.txt");     // Check if file exists
is_file("path.txt");         // Check if it's a file
is_dir("path.txt");          // Check if it's a directory

// Directory operations
mkdir("new_folder");         // Create directory
list_dir("folder");          // List directory contents
remove_file("old.txt");      // Remove file
```

### Path Manipulation Functions

```go
var fullPath = path_join("home", "user", "docs", "file.txt");
var fileName = path_base("/home/user/docs/file.txt");   // "file.txt"
var dirPath = path_dir("/home/user/docs/file.txt");      // "/home/user/docs"
var ext = path_ext("/home/user/docs/file.txt");          // ".txt"
var absPath = path_abs("relative/path");
var txtFiles = glob("*.txt");
```

---

## Enums

Go-Mix supports enumerated types (enums) for defining named constant values.

### Basic Enum

```go
enum Color {
    RED,
    GREEN,
    BLUE
}

println(Color.RED);    // 0
println(Color.GREEN);  // 1
println(Color.BLUE);   // 2
```

### Enum with Explicit Values

```go
enum Status {
    PENDING = 0,
    ACTIVE = 1,
    COMPLETED = 2,
    CANCELLED = 3
}
```

### Enum with Mixed Values

```go
enum Priority {
    LOW = 10,
    MEDIUM,      // Auto-assigned: 11
    HIGH = 50,
    CRITICAL     // Auto-assigned: 51
}
```

---

## Functions

### Function Definitions

```go
// Basic function
func greet(name) {
    return "Hello, " + name;
}

// Multiple parameters
func add(a, b) {
    return a + b;
}

// Recursive function
func factorial(n) {
    if (n <= 1) { return 1; }
    return n * factorial(n - 1);
}
```

### Function Expressions (Lambda)

```go
// Anonymous functions
var double = func(x) { return x * 2; };
var add = func(a, b) { return a + b; };

// Higher-order functions
var makeMultiplier = func(factor) {
    return func(x) { return x * factor; };
};

var times5 = makeMultiplier(5);
println(times5(3));        // 15
```

### Higher-Order Functions

```go
var numbers = [1, 2, 3, 4, 5];

// Map - transform each element
var doubled = map(numbers, func(x) { return x * 2; });
// [2, 4, 6, 8, 10]

// Filter - keep matching elements
var evens = filter(numbers, func(x) { return x % 2 == 0; });
// [2, 4]

// Reduce - accumulate to single value
var sum = reduce(numbers, func(acc, x) { return acc + x; }, 0);
// 15

// Find - get first matching element
var firstEven = find(numbers, func(x) { return x % 2 == 0; });
// 2

// Check predicates
var hasEven = some(numbers, func(x) { return x % 2 == 0; });     // true
var allPositive = every(numbers, func(x) { return x > 0; });       // true
```

---

## Closures

Functions are first-class citizens with full support for lexical closures:

```go
func makeCounter(start) {
    var count = start;
    func increment() {
        count = count + 1;
        return count;
    }
    return increment;
}

var counter = makeCounter(0);
println(counter()); // 1
println(counter()); // 2
println(counter()); // 3

// Each counter maintains its own 'count' variable
var counter2 = makeCounter(100);
println(counter2()); // 101
```

---

## Object-Oriented Programming

### Struct Definition

```go
struct Circle {
    var radius = 0;

    // Constructor
    func init(r) {
        this.radius = r;
    }

    // Instance methods
    func area() {
        return 3.14159 * this.radius * this.radius;
    }

    func circumference() {
        return 2 * 3.14159 * this.radius;
    }

    func scale(factor) {
        this.radius = this.radius * factor;
        return this;  // Method chaining
    }
}

// Creating instances
var c = new Circle(5);
println(c.area());         // 78.53975
println(c.circumference()); // 31.4159

// Method chaining
c.scale(2).scale(2);       // radius becomes 20
```

### Complex Example

```go
struct BankAccount {
    var balance = 0;
    var owner = "Unknown";

    func init(owner, initial) {
        this.owner = owner;
        this.balance = initial;
    }

    func deposit(amount) {
        if (amount <= 0) {
            panic("Deposit amount must be positive");
        }
        this.balance = this.balance + amount;
        return this.balance;
    }

    func withdraw(amount) {
        if (amount > this.balance) {
            panic("Insufficient funds");
        }
        this.balance = this.balance - amount;
        return this.balance;
    }

    func getBalance() {
        return this.balance;
    }

    func getInfo() {
        return this.owner + " has $" + this.balance;
    }
}

var account = new BankAccount("Alice", 1000);
account.deposit(500);      // 1500
account.withdraw(200);     // 1300
println(account.getInfo()); // Alice has $1300
```
