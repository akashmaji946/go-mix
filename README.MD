# 🚀 Go-Mix Programming Language

![Go-Mix Logo](logo.PNG)

[![Go Version](https://img.shields.io/badge/Go-1.18+-00ADD8?style=flat-square&logo=go)](https://golang.org/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg?style=flat-square)](https://opensource.org/licenses/MIT)
[![Build Status](https://img.shields.io/badge/Build-Passing-brightgreen?style=flat-square)](https://github.com/akashmaji946/go-mix)

**Go-Mix** is a high-performance, interpreted programming language implemented in Go. It combines dynamic typing with static safety features, lexical closures, and a comprehensive standard library—perfect for prototyping, education, and embedding as a scripting layer in Go applications.

---

## Table of Contents

- [Quick Start](#-quick-start)
- [Core Architecture](#-core-architecture)
- [Language Features](#-language-features)
- [Data Types](#data-types)
- [Variables & Scoping](#variables--scoping)
- [Control Flow](#control-flow--iteration)
- [Functions](#functions--higher-order-programming)
- [Object-Oriented Programming](#object-oriented-programming)
- [Standard Library](#-standard-library--builtins)
- [Error Handling](#error-handling)
- [Sample Programs](#-sample-programs)
- [Development](#development--contributing)

---

## 🚀 Quick Start

### Installation

**Prerequisites**: Go 1.18 or higher

#### Option 1: Automated Installation (Recommended)
```bash
git clone https://github.com/akashmaji946/go-mix.git
cd go-mix
sudo ./install.sh
```

This builds and installs `go-mix` to `/usr/local/bin`, making it accessible from anywhere on your system. See [INSTALL.md](INSTALL.md) for detailed instructions.

This generates the `go-mix` executable in the current directory (local use only).

### Usage

**Display Version:**
```bash
go-mix --version    # or go-mix -v
```

```
Go-Mix version v1.0.0
License: MIT
Author: akashmaji(@iisc.ac.in)
```

**Display Help:**
```bash
go-mix --help       # or go-mix -h
```

**Interactive REPL (Read-Eval-Print Loop):**
```bash
go-mix
```

This launches an interactive session where you can type Go-Mix code line-by-line:
```bash
Go-Mix >>> var x = 42
Go-Mix >>> println(x)
42
Go-Mix >>> /help
/exit    - Exit the REPL
/scope   - Show current scope and variables
Go-Mix >>> 
```

**Run a Program File:**
```bash
go-mix samples/algo/05_factorial.gm
go-mix path/to/your/script.gm
```

#### Option 2: Manual Build
```bash
git clone https://github.com/akashmaji946/go-mix.git
cd go-mix
./build.sh
```

**Run Tests:**
```bash
go test ./...           # All tests
go test ./lexer -v      # Verbose
go test -coverage       # With coverage report
```

#### Option 3: Docker installation

**Manual Installation:**
Refer `DOCKER.MD` for detailed instructions.

**DockerHub Installation**
```bash
docker pull akashmaji946/go-mix:latest
```

---

## 📚 Core Architecture

### 1. **Pratt Parser (Top-Down Operator Precedence)**

Go-Mix uses a sophisticated Pratt Parser for expression parsing. Rather than relying on traditional recursive descent, the algorithm associates parsing logic directly with tokens and their precedence levels. This approach enables elegant handling of:

- **Complex operator hierarchies** — bitwise vs. arithmetic operators
- **Member access** — nested object property access (`obj.field.nested`)
- **Function calls** — argument parsing with proper precedence
- **Extensibility** — new operators can be added with minimal code changes

### 2. **Hybrid Type System**

Go-Mix implements three distinct variable kinds, each with different semantics:

```javascript
var x = 42;          // Dynamic: type can change at runtime
let y = 3.14;        // Static: type-locked after first assignment
const PI = 3.14159;  // Immutable: cannot be reassigned
```

**Dynamic Variables (`var`):**
- Type can change during execution
- Ideal for rapid prototyping and flexible code
- Behavior similar to Python or JavaScript

**Static Variables (`let`):**
- Type is locked on first assignment
- Provides safety without explicit type annotations
- Prevents accidental type changes

**Immutable Constants (`const`):**
- Cannot be reassigned after initial binding
- Useful for configuration values and constants
- Provides compile-time and runtime safety

### 3. **Lexical Scoping & Closures**

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

Variables are captured by reference in closures, allowing stateful function objects.

---

## 🎯 Language Features

### Data Types

Go-Mix supports six primitive types plus `nil`:

| Type | Example | Size | Notes |
|:-----|:--------|:-----|:------|
| `int` | `42`, `0xFF`, `0o77` | 64-bit | Signed integers; hex (0x) and octal (0o) support |
| `float` | `3.14`, `1.2e3`, `1e-5` | 64-bit | IEEE-754 double precision |
| `bool` | `true`, `false` | 1-bit | Boolean logic values |
| `string` | `"Hello"`, `"Line1\nLine2"` | Variable | UTF-8 strings with escape sequences |
| `char` | `'A'`, `'\n'`, `'\t'` | 32-bit | Single Unicode character |
| `nil` | `nil` | Pointer-size | Represents absence of value |

**Type Conversion Examples:**
```go
var s = "42";
var n = tostring(123);      // "123"
var i = 42;
var f = 42.0;               // Auto-conversion in numeric context

// Type checking
typeof(42);                 // "int"
typeof(3.14);              // "float"
typeof("text");            // "string"
typeof(true);              // "bool"
typeof(nil);               // "nil"
```

### Variables & Scoping

#### Declaration & Assignment

```go
// Variable declaration
var x = 10;
var y = "hello";
var z;                     // Uninitialized (nil)

// Static variables (type-locked)
let count = 0;
count = count + 1;         // OK
count = "text";            // ERROR: type mismatch

// Constant values
const MAX_SIZE = 100;
const PI = 3.14159;
MAX_SIZE = 200;            // ERROR: constants are immutable

// Multiple declaration (comma-separated)
var a = 1, b = 2, c = 3;
```

#### Scoping Rules

Go-Mix implements lexical (static) scoping with proper block-level scope:

```go
var x = "global";

func testScope() {
    var x = "function";     // Shadows outer x
    println(x);             // "function"
    
    if (true) {
        var x = "block";    // Shadows function-level x
        println(x);         // "block"
    }
    
    println(x);             // "function" (block x is gone)
}

println(x);                 // "global"
```

---

## Control Flow & Iteration

### Conditionals

#### If-Else Statements

```go
if (age >= 18) {
    println("Adult");
} else if (age >= 13) {
    println("Teenager");
} else {
    println("Child");
}

// Ternary operator (via if-else)
var status = if (x > 0) { "positive"; } else { "non-positive"; };
```

#### Boolean Logic (Short-Circuit Evaluation)

```go
var x = 5;

// AND operator (&&) - short circuits if first operand is false
if (x > 0 && x < 10) {
    println("Between 0 and 10");
}

// OR operator (||) - short circuits if first operand is true
if (x < 0 || x > 100) {
    println("Outside range");
}

// NOT operator (!)
if (!isError) {
    println("Success");
}

// Compound conditions
if ((x > 0 && x < 10) || (x > 100 && x < 110)) {
    println("Valid");
}
```

### Loops

#### C-Style For Loop

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

// Nested loops
for (var i = 0; i < 3; i = i + 1) {
    for (var j = 0; j < 3; j = j + 1) {
        println(i, j);
    }
}
```

#### While Loop

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

#### Range Iteration (Foreach)

```javascript
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
    println(idx, val);    // 0, 10; 1, 20; 2, 30
}
```

---

## Collections & Data Structures

### Arrays (Mutable Sequences)

```go
// Array literals
var empty = [];
var numbers = [1, 2, 3, 4, 5];
var mixed = [1, "hello", true, 3.14];
var nested = [[1, 2], [3, 4], [5, 6]];

// Array indexing (0-based)
var arr = [10, 20, 30, 40, 50];
println(arr[0]);           // 10 (first element)
println(arr[2]);           // 30 (third element)

// Negative indexing (Python-style)
println(arr[-1]);          // 50 (last element)
println(arr[-2]);          // 40 (second-to-last)

// Index assignment
arr[1] = 25;               // [10, 25, 30, 40, 50]
arr[-1] = 55;              // [10, 25, 30, 40, 55]

// Array length
println(length(arr));      // 5
println(size(arr));        // 5 (alias)

// Array methods
push(arr, 60);             // Add to end: [10, 25, 30, 40, 55, 60]
pop(arr);                  // Remove from end: [10, 25, 30, 40, 55]
unshift(arr, 5);           // Add to start: [5, 10, 25, 30, 40, 55]
shift(arr);                // Remove from start: [10, 25, 30, 40, 55]
```

### Maps (Dictionaries)

```go
// Map literals (string keys only)
var user = map{
    "name": "Alice",
    "age": 30,
    "city": "New York"
};

// Key access
println(user["name"]);     // "Alice"
println(user["age"]);      // 30

// Adding/updating keys
user["email"] = "alice@example.com";
user["age"] = 31;

// Checking keys (error if missing)
// Note: accessing non-existent key returns nil or error

// Map methods
println(length(user));     // Number of key-value pairs
```

### Lists (Heterogeneous)

```go
// Lists are explicitly mutable sequences
var list = list(1, "two", 3.0, true, nil);

// Access like arrays
println(list[0]);          // 1
println(list[1]);          // "two"

// Different from arrays in type flexibility
var mixed_array = [1, 2, 3];  // Homogeneous
var mixed_list = list(1, "2", 3.0);  // Heterogeneous
```

### Tuples (Immutable Sequences)

```go
// Tuples are fixed-size, immutable
var coords = tuple(10, 20);
var rgb = tuple(255, 128, 64);
var data = tuple("John", 25, true);

// Access like arrays
println(coords[0]);        // 10
println(coords[1]);        // 20

// Immutable - cannot be modified
coords[0] = 15;            // ERROR
```

### Sets (Unique Values)

```go
// Sets contain unique values
var primes = set{2, 3, 5, 7, 11};
var colors = set{"red", "green", "blue"};
var numbers = set{1, 2, 2, 3, 3, 3};  // Results in {1, 2, 3}

// Checking membership
// Note: Set operations depend on implementation

println(length(primes));   // 5
```

---

## Functions & Higher-Order Programming

### Function Definitions

```go
// Basic function
func greet(name) {
    return "Hello, " + name;
}

println(greet("Alice"));   // "Hello, Alice"

// Function with multiple parameters
func add(a, b) {
    return a + b;
}

println(add(3, 5));        // 8

// Function with no return (returns nil)
func printInfo(x) {
    println("Value: " + x);
}

// Recursive function
func factorial(n) {
    if (n <= 1) { return 1; }
    return n * factorial(n - 1);
}

println(factorial(5));     // 120
```

### Function Expressions (Lambda)

```go
// Anonymous functions assigned to variables
var double = func(x) { return x * 2; };
println(double(5));        // 10

var add = func(a, b) { return a + b; };
println(add(3, 4));        // 7

// Nested function definitions
var makeMultiplier = func(factor) {
    return func(x) { return x * factor; };
};

var times5 = makeMultiplier(5);
println(times5(3));        // 15
println(times5(7));        // 35
```

### Higher-Order Functions (Functional Programming)

#### Map

```go
// Apply function to each element
var nums = [1, 2, 3, 4, 5];
var doubled = map(nums, func(x) { return x * 2; });
println(doubled);          // [2, 4, 6, 8, 10]

var squared = map(nums, func(x) { return x * x; });
println(squared);          // [1, 4, 9, 16, 25]
```

#### Filter

```go
// Keep elements that satisfy predicate
var numbers = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10];
var evens = filter(numbers, func(x) { return x % 2 == 0; });
println(evens);            // [2, 4, 6, 8, 10]

var bigNums = filter(numbers, func(x) { return x > 5; });
println(bigNums);          // [6, 7, 8, 9, 10]
```

#### Reduce

```go
// Accumulate value across array
var nums = [1, 2, 3, 4, 5];

// Sum
var sum = reduce(nums, func(acc, x) { return acc + x; }, 0);
println(sum);              // 15

// Product
var product = reduce(nums, func(acc, x) { return acc * x; }, 1);
println(product);          // 120

// Building a string
var str = reduce(nums, func(acc, x) { return acc + x; }, "");
println(str);              // "12345"
```

#### Find, Some, Every

```go
var nums = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10];

// Find first matching element
var first_even = find(nums, func(x) { return x % 2 == 0; });
println(first_even);       // 2

// Check if at least one element matches
var has_negative = some(nums, func(x) { return x < 0; });
println(has_negative);     // false

var has_five = some(nums, func(x) { return x == 5; });
println(has_five);         // true

// Check if all elements match
var all_positive = every(nums, func(x) { return x > 0; });
println(all_positive);     // true

var all_even = every(nums, func(x) { return x % 2 == 0; });
println(all_even);         // false
```

#### Sorting with Custom Comparators

```go
var nums = [3, 1, 4, 1, 5, 9, 2, 6];

// In-place sort with default ordering
sort(nums);                // [1, 1, 2, 3, 4, 5, 6, 9]

// New sorted array (non-mutating)
var sorted = sorted(nums);

// Custom sort (descending)
csort(nums, func(a, b) { 
    if (a > b) { return -1; }
    if (a < b) { return 1; }
    return 0;
});

// Sort strings by length
var words = ["apple", "pi", "banana", "cat"];
csort(words, func(a, b) {
    var len_a = length(a);
    var len_b = length(b);
    if (len_a < len_b) { return -1; }
    if (len_a > len_b) { return 1; }
    return 0;
});
```

---

## Object-Oriented Programming

### Struct Definition & Methods

```cpp
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

    func toString() {
        return "Circle(radius=" + this.radius + ")";
    }
}

// Creating instances
var c = new Circle(5);
println(c.area());         // 78.53975
println(c.circumference()); // 31.4159

// Method chaining
c.scale(2).scale(2);       // radius becomes 20
```

### Complex OOP Example

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

---

## 📖 Standard Library & Builtins

Go-Mix provides a comprehensive standard library organized into functional modules. All functions are available as global builtins, and can also be imported as packages.

### Global Imports

All builtin functions are available globally without explicit import. However, you can also import functions as packages:

```java
import arrays;      // Access array functions
import strings;     // Access string functions
import math;        // Access math functions
import maps;        // Access map functions
import time;        // Access time functions
import http;        // Access HTTP client/server functions
import regex;       // Access regex functions
import file;        // Access file I/O functions
import os;          // Access OS functions
import io;          // Access input/output functions
import format;      // Access type conversion functions
```

### Common Functions

| Function | Parameters | Returns | Description |
|:---------|:-----------|:--------|:------------|
| `print(...)` | Variable | nil | Output without newline |
| `println(...)` | Variable | nil | Output with newline |
| `printf(fmt, ...)` | String + args | nil | Formatted output (%d, %s, %f, %x) |
| `length(obj)` | Array/String/Map | int | Length of collection |
| `size(obj)` | Array/String/Map | int | Alias for length |
| `typeof(obj)` | Any | string | Type name |
| `tostring(obj)` | Any | string | String conversion |
| `addr(obj)` | Any | int | Memory address (for reference checking) |
| `is_same_ref(obj1, obj2)` | Any, Any | bool | Check if same reference |
| `range(start, end)` | int, int | range | Create inclusive range |

**Common Function Examples:**
```javascript
print("Hello");             // Output: Hello (no newline)
println("World");           // Output: World (with newline)
printf("Number: %d\n", 42); // Output: Number: 42

length("hello");            // 5
length([1, 2, 3]);         // 3
length(map{"a": 1});       // 1

typeof(42);                 // "int"
typeof(3.14);              // "float"
typeof(true);              // "bool"

tostring(123);             // "123"
tostring([1, 2]);          // "[1, 2]"

var x = [1, 2, 3];
var y = x;
is_same_ref(x, y);         // true (same reference)

// Range creation
range(1, 5);               // Creates range 1 to 5 inclusive
foreach i in range(1, 5) { println(i); }  // 1, 2, 3, 4, 5
```

### String Functions

The `strings` package provides comprehensive string manipulation and analysis utilities.

| Function | Parameters | Returns | Description |
|:---------|:-----------|:--------|:------------|
| `upper(str)` | string | string | Convert to uppercase |
| `lower(str)` | string | string | Convert to lowercase |
| `trim(str)` | string | string | Remove leading/trailing whitespace |
| `ltrim(str)` | string | string | Remove leading whitespace |
| `rtrim(str)` | string | string | Remove trailing whitespace |
| `split(str, sep)` | string, string | array | Split by separator |
| `join(arr, sep)` | array, string | string | Join array with separator |
| `contains(str, sub)` | string, string | bool | Check if contains substring |
| `starts_with(str, prefix)` | string, string | bool | Check if starts with |
| `ends_with(str, suffix)` | string, string | bool | Check if ends with |
| `index(str, sub)` | string, string | int | Find first occurrence (-1 if not found) |
| `substring(str, start, length)` | string, int, int | string | Extract substring |
| `replace(str, old, new)` | string, string, string | string | Replace all occurrences |
| `reverse(str)` | string | string | Reverse string characters |
| `capitalize(str)` | string | string | Capitalize first letter |
| `ord(char)` | char/string | int | Get Unicode code point |
| `chr(code)` | int | char | Get character from code point |
| `count(str, sub)` | string, string | int | Count occurrences |
| `is_digit(str)` | string | bool | Check if all digits |
| `is_alpha(str)` | string | bool | Check if all letters |
| `strcmp(s1, s2)` | string, string | int | Compare strings (-1, 0, 1) |

**String Function Examples:**
```go
// Case conversion
upper("hello world");       // "HELLO WORLD"
lower("HELLO WORLD");       // "hello world"
capitalize("hello world");  // "Hello world"

// Whitespace handling
trim("  hello  ");          // "hello"
ltrim("  hello  ");         // "hello  "
rtrim("  hello  ");         // "  hello"

// Splitting and joining
split("a,b,c,d", ",");      // ["a", "b", "c", "d"]
split("hello world", " ");  // ["hello", "world"]
join(["apple", "banana", "cherry"], ", ");  // "apple, banana, cherry"

// Searching and matching
contains("hello world", "world");  // true
starts_with("go-mix", "go");       // true
ends_with("filename.gm", ".gm");   // true
index("hello world", "world");     // 6
count("mississippi", "s");         // 4

// Substring extraction
substring("hello world", 0, 5);    // "hello"
substring("hello world", 6, 5);    // "world"

// Character operations
ord('A');                   // 65
chr(65);                    // 'A'
ord("ABC");                 // 65 (first character)

// Type checking
is_digit("12345");          // true
is_alpha("hello");          // true
is_digit("123abc");         // false

// String comparison
strcmp("apple", "banana");  // -1 (apple < banana)
strcmp("hello", "hello");   // 0 (equal)
strcmp("zebra", "apple");   // 1 (zebra > apple)

// String reversal
reverse("hello");           // "olleh"
reverse("12345");           // "54321"
```

### Array Functions

The `arrays` package provides functions for manipulating sequences and functional operations on arrays.

| Function | Parameters | Returns | Description |
|:---------|:-----------|:--------|:------------|
| `push(arr, elem)` | array, any | array | Add element to end |
| `pop(arr)` | array | any | Remove & return last element |
| `shift(arr)` | array | any | Remove & return first element |
| `unshift(arr, elem)` | array, any | array | Add element to start |
| `sort(arr, [reverse])` | array, [bool] | array | In-place sort |
| `sorted(arr, [reverse])` | array, [bool] | array | New sorted array |
| `clone_array(arr)` | array | array | Shallow copy of array |
| `reverse(arr)` | array | array | Reverse array |
| `map_array(arr, func)` | array, function | array | Map function over elements |
| `filter_array(arr, func)` | array, function | array | Filter elements by predicate |
| `reduce_array(arr, func, init)` | array, function, any | any | Reduce array to single value |
| `find_array(arr, func)` | array, function | any | Find first matching element |
| `some_array(arr, func)` | array, function | bool | Check if any element matches |
| `every_array(arr, func)` | array, function | bool | Check if all elements match |
| `to_array(iterable)` | list/tuple | array | Convert to array |
| `csort(arr, comparator)` | array, function | array | Custom in-place sort |
| `csorted(arr, comparator)` | array, function | array | Custom sorted copy |

**Array Function Examples:**
```go
var arr = [1, 2, 3, 4, 5];

// Stack-like operations
push(arr, 6);               // [1, 2, 3, 4, 5, 6]
pop(arr);                   // Returns 6, arr = [1, 2, 3, 4, 5]
unshift(arr, 0);            // [0, 1, 2, 3, 4, 5]
shift(arr);                 // Returns 0, arr = [1, 2, 3, 4, 5]

// Sorting
sort([3, 1, 4, 1, 5]);      // [1, 1, 3, 4, 5] (in-place)
sorted([3, 1, 4, 1, 5]);    // [1, 1, 3, 4, 5] (new array)
sort([3, 1, 4, 1, 5], true); // [5, 4, 3, 1, 1] (reverse)

// Custom sorting
var numbers = [5, 2, 8, 1, 9];
csort(numbers, func(a, b) { return a > b; });  // Sort descending

// Array cloning
var original = [1, 2, 3];
var copy = clone_array(original);

// Functional operations
var nums = [1, 2, 3, 4, 5];

// Map - transform each element
var doubled = map(nums, func(x) { return x * 2; });
// doubled = [2, 4, 6, 8, 10]

var squared = map_array(nums, func(x) { return x * x; });
// squared = [1, 4, 9, 16, 25]

// Filter - keep matching elements
var evens = filter(nums, func(x) { return x % 2 == 0; });
// evens = [2, 4]

var large = filter_array(nums, func(x) { return x > 3; });
// large = [4, 5]

// Reduce - accumulate to single value
var sum = reduce(nums, func(acc, x) { return acc + x; }, 0);
// sum = 15

var product = reduce(nums, func(acc, x) { return acc * x; }, 1);
// product = 120

var concatenated = reduce(nums, func(acc, x) { return acc + tostring(x); }, "");
// concatenated = "12345"

// Find - get first matching element
var first_large = find(nums, func(x) { return x > 3; });
// first_large = 4

var first_even = find_array(nums, func(x) { return x % 2 == 0; });
// first_even = 2

// Check if any/all match
some(nums, func(x) { return x > 10; });        // false
some_array(nums, func(x) { return x > 3; });  // true

every(nums, func(x) { return x > 0; });       // true
every_array(nums, func(x) { return x < 10; }); // true

// Type conversion
var list_data = list(10, 20, 30);
var arr_data = to_array(list_data);  // [10, 20, 30]
```

### Math Functions

The `math` package provides mathematical operations and special functions.

| Function | Parameters | Returns | Description |
|:---------|:-----------|:--------|:------------|
| `abs(n)` | int/float | number | Absolute value |
| `fabs(n)` | float | float | Float absolute value |
| `min(a, b)` | number, number | number | Minimum of two numbers |
| `max(a, b)` | number, number | number | Maximum of two numbers |
| `floor(n)` | float | int | Greatest integer ≤ n |
| `ceil(n)` | float | int | Smallest integer ≥ n |
| `round(n, [precision])` | float, [int] | float | Round to precision |
| `sqrt(n)` | number | float | Square root |
| `pow(base, exp)` | number, number | number | Exponentiation (base^exp) |
| `sin(rad)` | float | float | Sine (radians) |
| `cos(rad)` | float | float | Cosine (radians) |
| `tan(rad)` | float | float | Tangent (radians) |
| `asin(n)` | float | float | Arc sine |
| `acos(n)` | float | float | Arc cosine |
| `atan(n)` | float | float | Arc tangent |
| `atan2(y, x)` | float, float | float | Two-argument arc tangent |
| `log(n)` | number | float | Natural logarithm |
| `log10(n)` | number | float | Base-10 logarithm |
| `exp(n)` | float | float | e^n |
| `rand()` | none | float | Random [0.0, 1.0) |
| `rand_int(min, max)` | int, int | int | Random in range [min, max) |

**Math Function Examples:**
```go
// Absolute value
abs(-42);                   // 42
fabs(-3.14);                // 3.14

// Min/Max
min(5, 3);                  // 3
max(5, 3);                  // 5
min(3.7, 2.1);              // 2.1

// Rounding
floor(3.9);                 // 3
ceil(3.1);                  // 4
round(3.5);                 // 4 (rounds to nearest)
round(3.14159, 2);          // 3.14 (2 decimal places)

// Power and roots
sqrt(16);                   // 4.0
sqrt(2);                    // 1.41421356...
pow(2, 10);                 // 1024
pow(3, 3);                  // 27
pow(2, -1);                 // 0.5

// Logarithms
log(10);                    // 2.302585... (natural log)
log10(100);                 // 2.0
log10(1000);                // 3.0
exp(1);                     // 2.71828... (e)

// Trigonometry (angles in radians)
sin(0);                     // 0.0
cos(0);                     // 1.0
tan(0);                     // 0.0

// Inverse trigonometry
asin(0);                    // 0.0
acos(1);                    // 0.0
atan(1);                    // 0.785398... (π/4)
atan2(1, 1);                // 0.785398... (π/4)

// Random numbers
var random_float = rand();           // 0.0 to 1.0
var random_percent = rand() * 100;   // 0 to 100
var dice = rand_int(1, 7);           // 1 to 6 (inclusive)
var card = rand_int(1, 53);          // 1 to 52 (card deck)

// Practical examples
var radius = 5;
var area = 3.14159 * pow(radius, 2);
var hypotenuse = sqrt(pow(3, 2) + pow(4, 2));  // Pythagorean theorem
```

### List Functions

The `list` package provides functions for mutable, heterogeneous list operations.

| Function | Parameters | Returns | Description |
|:---------|:-----------|:--------|:------------|
| `list(...)` | variadic | list | Create new list from elements |
| `pushback_list(list, elem)` | list, any | list | Add element to end |
| `pushfront_list(list, elem)` | list, any | list | Add element to beginning |
| `popback_list(list)` | list | any | Remove & return last element |
| `popfront_list(list)` | list | any | Remove & return first element |
| `size_list(list)` | list | int | Get number of elements |
| `peekback_list(list)` | list | any | Get last element (no remove) |
| `peekfront_list(list)` | list | any | Get first element (no remove) |
| `insert_list(list, index, elem)` | list, int, any | list | Insert at index |
| `remove_list(list, index)` | list, int | any | Remove at index |
| `contains_list(list, elem)` | list, any | bool | Check if contains element |
| `map_list(list, func)` | list, function | list | Apply function to each element |
| `filter_list(list, func)` | list, function | list | Filter by predicate |
| `reduce_list(list, func, init)` | list, function, any | any | Reduce to single value |
| `find_list(list, func)` | list, function | any | Find first matching element |
| `some_list(list, func)` | list, function | bool | Check if any element matches |
| `every_list(list, func)` | list, function | bool | Check if all elements match |
| `to_list(iterable)` | array/tuple | list | Convert to list |

**List Function Examples:**
```go
// Creating and modifying lists
var l = list(10, 20, 30);
pushback_list(l, 40);          // l = list(10, 20, 30, 40)
pushfront_list(l, 5);          // l = list(5, 10, 20, 30, 40)

popback_list(l);               // Returns 40, l = list(5, 10, 20, 30)
popfront_list(l);              // Returns 5, l = list(10, 20, 30)

size_list(l);                  // 3
peekback_list(l);              // 30 (without removing)
peekfront_list(l);             // 10 (without removing)

// Insertion and removal
insert_list(l, 1, 15);         // l = list(10, 15, 20, 30)
remove_list(l, 1);             // Returns 15, l = list(10, 20, 30)

// Checking membership
contains_list(l, 20);          // true
contains_list(l, 99);          // false

// Functional operations
var nums = list(1, 2, 3, 4, 5);
var doubled = map_list(nums, func(x) { return x * 2; });
var evens = filter_list(nums, func(x) { return x % 2 == 0; });
var sum = reduce_list(nums, func(a, x) { return a + x; }, 0);

// Search and predicates
var found = find_list(nums, func(x) { return x > 3; }); // 4
var has_even = some_list(nums, func(x) { return x % 2 == 0; }); // true
var all_pos = every_list(nums, func(x) { return x > 0; }); // true

// Type conversion
var arr = [1, 2, 3];
var list_from_arr = to_list(arr);
```

### Tuple Functions

The `tuple` package provides functions for immutable fixed-size sequences.

| Function | Parameters | Returns | Description |
|:---------|:-----------|:--------|:------------|
| `tuple(...)` | variadic | tuple | Create new tuple from elements |
| `size_tuple(tuple)` | tuple | int | Get number of elements |
| `peekback_tuple(tuple)` | tuple | any | Get last element |
| `peekfront_tuple(tuple)` | tuple | any | Get first element |
| `contains_tuple(tuple, elem)` | tuple, any | bool | Check if contains element |
| `map_tuple(tuple, func)` | tuple, function | tuple | Apply function to each element |
| `filter_tuple(tuple, func)` | tuple, function | tuple | Filter by predicate |
| `reduce_tuple(tuple, func, init)` | tuple, function, any | any | Reduce to single value |
| `find_tuple(tuple, func)` | tuple, function | any | Find first matching element |
| `some_tuple(tuple, func)` | tuple, function | bool | Check if any element matches |
| `every_tuple(tuple, func)` | tuple, function | bool | Check if all elements match |
| `to_tuple(iterable)` | array/list | tuple | Convert to tuple |

**Tuple Function Examples:**
```go
// Creating tuples (immutable)
var coords = tuple(10, 20);
var rgb = tuple(255, 128, 64);
var person = tuple("Alice", 30, true);

// Accessing elements
coords[0];                     // 10
coords[1];                     // 20
coords[-1];                    // 20 (last element)

// Tuple properties
size_tuple(coords);            // 2
peekback_tuple(person);        // true (last element)
peekfront_tuple(person);       // "Alice" (first element)

// Membership checking
contains_tuple(rgb, 128);      // true
contains_tuple(coords, 99);    // false

// Functional operations
var nums = tuple(1, 2, 3, 4, 5);
var doubled = map_tuple(nums, func(x) { return x * 2; });
var sum = reduce_tuple(nums, func(a, x) { return a + x; }, 0);

// Search and predicates
var found = find_tuple(nums, func(x) { return x > 3; }); // 4
var has_even = some_tuple(nums, func(x) { return x % 2 == 0; }); // true
var all_pos = every_tuple(nums, func(x) { return x > 0; }); // true

// Type conversion
var arr = [1, 2, 3];
var t = to_tuple(arr);         // tuple(1, 2, 3)
```

### Map Functions

The `maps` package provides functions for dictionary/hash map operations.

| Function | Parameters | Returns | Description |
|:---------|:-----------|:--------|:------------|
| `keys_map(map)` | map | array | Get all keys |
| `insert_map(map, key, value)` | map, any, any | any | Insert/update key-value pair |
| `remove_map(map, key)` | map, any | any | Remove key and return value |
| `contain_map(map, key)` | map, any | bool | Check if key exists |
| `enumerate_map(map)` | map | array | Get [key, value] pairs |
| `json_string_to_map(json)` | string | map | Parse JSON string to map |
| `map_to_json_string(map)` | map | string | Convert map to JSON string |

**Map Function Examples:**
```go
// Creating and modifying maps
var user = map{
    "name": "Alice",
    "age": 30,
    "email": "alice@example.com"
};

// Accessing values
user["name"];                  // "Alice"
user["age"];                   // 30

// Inserting/updating
insert_map(user, "city", "New York");
insert_map(user, "age", 31);

// Checking keys
contain_map(user, "name");     // true
contain_map(user, "phone");    // false

// Getting all keys
keys_map(user);                // ["name", "age", "email", "city"]

// Enumerating key-value pairs
enumerate_map(user);           // [["name", "Alice"], ["age", 31], ...]

// JSON conversion
var json_str = map_to_json_string(user);
// {"name":"Alice","age":31,"email":"alice@example.com","city":"New York"}

var parsed = json_string_to_map(json_str);
parsed["name"];                // "Alice"
```

### Set Functions

The `sets` package provides functions for unique value collections.

| Function | Parameters | Returns | Description |
|:---------|:-----------|:--------|:------------|
| `insert_set(set, value)` | set, any | any | Add value to set |
| `remove_set(set, value)` | set, any | bool | Remove value (returns if existed) |
| `contains_set(set, value)` | set, any | bool | Check if contains value |
| `values_set(set)` | set | array | Get all unique values |
| `size_set(set)` | set | int | Get number of unique values |

**Set Function Examples:**
```go
// Creating sets (unique values only)
var colors = set{"red", "green", "blue"};
var primes = set{2, 3, 5, 7, 11};

// Adding/removing values
insert_set(colors, "yellow");     // colors now has 4 elements
insert_set(colors, "red");        // No change (already exists)

remove_set(colors, "blue");       // true (existed and removed)
remove_set(colors, "orange");     // false (didn't exist)

// Checking membership
contains_set(colors, "red");      // true
contains_set(colors, "purple");   // false

// Getting all values
values_set(primes);               // [2, 3, 5, 7, 11]

// Set size
size_set(colors);                 // Number of unique colors
```

### Time Functions

The `time` package provides functions for time handling and formatting.

| Function | Parameters | Returns | Description |
|:---------|:-----------|:--------|:------------|
| `now()` | none | int | Current Unix timestamp (seconds) |
| `now_ms()` | none | int | Current Unix timestamp (milliseconds) |
| `utc_now()` | none | int | Current UTC Unix timestamp |
| `format_time(timestamp, layout)` | int, string | string | Format timestamp to string |
| `parse_time(value, layout)` | string, string | int | Parse time string to timestamp |
| `timezone()` | none | string | Get current timezone name |

**Time Function Examples:**
```go
// Get current time
var now_seconds = now();        // e.g., 1707542400
var now_millis = now_ms();      // e.g., 1707542400000

// Format time
var current = now();
var formatted = format_time(current, "2006-01-02 15:04:05");
// Output: "2024-02-10 12:00:00"

// Format variations
format_time(now(), "2006-01-02");           // "2024-02-10"
format_time(now(), "Jan 02, 2006");         // "Feb 10, 2024"
format_time(now(), "15:04:05");             // "12:00:00"
format_time(now(), "Monday, Jan 2 2006");   // "Saturday, Feb 10 2024"

// Parse time string
var timestamp = parse_time("2024-02-10", "2006-01-02");
var timestamp2 = parse_time("02/10/2024", "01/02/2006");

// Timezone information
var tz = timezone();            // e.g., "IST" or "UTC"

// Practical examples
var start_time = now();
sleep(2000);                    // Wait 2 seconds
var end_time = now();
var elapsed = end_time - start_time;  // ~2 seconds
```

### File I/O Functions

The `file` package provides comprehensive file system operations.

| Function | Parameters | Returns | Description |
|:---------|:-----------|:--------|:------------|
| `read_file(path)` | string | string | Read entire file content |
| `write_file(path, content)` | string, string | nil | Write string to file |
| `append_file(path, content)` | string, string | nil | Append string to file |
| `file_exists(path)` | string | bool | Check if file/directory exists |
| `is_dir(path)` | string | bool | Check if path is directory |
| `is_file(path)` | string | bool | Check if path is regular file |
| `mkdir(path)` | string | nil | Create directory (recursive) |
| `remove_file(path)` | string | nil | Remove file or empty directory |
| `remove_all(path)` | string | nil | Remove path and all children |
| `rename_file(old, new)` | string, string | nil | Rename/move file |
| `touch(path)` | string | nil | Create empty file or update timestamps |
| `list_dir(path)` | string | array | List directory contents |
| `pwd()` | none | string | Get current working directory |
| `home()` | none | string | Get user's home directory |
| `truncate_file(path, size)` | string, int | nil | Change file size |
| `chmod(path, mode)` | string, int | nil | Change file permissions |
| `cat(path, ...)` | string... | nil | Print file contents |

**File I/O Function Examples:**
```go
// Reading files
var content = read_file("config.txt");
println(content);

// Writing files
write_file("output.txt", "Hello, World!");

// Appending to files
append_file("log.txt", "New log entry\n");

// File existence checking
if (file_exists("data.json")) {
    var data = read_file("data.json");
}

// Directory operations
is_dir("./samples");           // true
is_file("./samples/file.gm");  // true

// Create directories
mkdir("./data");               // Create single directory
mkdir("./data/nested/deep");   // Recursive creation

// List directory contents
var files = list_dir(".");     // ["file1.gm", "file2.gm", ...]

// File manipulation
rename_file("old_name.txt", "new_name.txt");
remove_file("temp.txt");
remove_all("old_directory");   // Recursive delete

// Create empty file
touch("empty.txt");

// File truncation
truncate_file("huge_file.log", 0);  // Clear file

// Get directory info
var curr_dir = pwd();          // Get current working directory
var home_dir = home();         // Get home directory

// Display file contents
cat("README.md");              // Print file to stdout
cat("file1.txt", "file2.txt"); // Print multiple files
```

### Type Conversion Functions

The `format` package provides functions for converting between types.

| Function | Parameters | Returns | Description |
|:---------|:-----------|:--------|:------------|
| `to_int(value)` | any | int | Convert to integer |
| `to_float(value)` | any | float | Convert to float |
| `to_bool(value)` | any | bool | Convert to boolean |
| `to_string(value)` | any | string | Convert to string |
| `to_char(value)` | any | char | Convert to character |

**Type Conversion Examples:**
```go
// String to number
to_int("42");                  // 42
to_float("3.14");              // 3.14
to_int("0xFF");                // 255 (hex support)
to_int("0o77");                // 63 (octal support)

// Number to string
to_string(123);                // "123"
to_string(3.14);               // "3.14"

// Float to int (truncation)
to_int(3.9);                   // 3
to_int(-3.1);                  // -3

// Boolean conversions
to_bool(1);                    // true
to_bool(0);                    // false
to_bool("true");               // true
to_bool("false");              // false
to_bool(nil);                  // false

// Character operations
to_char(65);                   // 'A'
to_char("hello");              // 'h' (first character)

// Type checking before conversion
var x = "123";
if (typeof(x) == "string") {
    var num = to_int(x);
}
```

### Input/Output Functions

The `io` package provides functions for reading user input.

| Function | Parameters | Returns | Description |
|:---------|:-----------|:--------|:------------|
| `scanln()` | none | string | Read line from stdin |
| `scanf(format)` | string | array | Read formatted input |
| `input(prompt)` | string | string | Print prompt and read line |
| `scan(delimiter)` | string | string | Read until delimiter |
| `getchar()` | none | char | Read single character |
| `putchar(char)` | char | nil | Print single character |
| `sprintf(fmt, ...)` | string... | string | Format string without printing |
| `flush()` | none | nil | Flush input/output buffers |
| `eprintln(...)` | variadic | nil | Print to stderr with newline |
| `eprintf(fmt, ...)` | string... | nil | Print formatted to stderr |

**Input/Output Examples:**
```go
// Reading input
var name = input("Enter your name: ");
println("Hello, " + name);

// Reading a line
var line = scanln();

// Reading formatted input
var data = scanf("%d %s");     // Returns array [num, string]
var age = data[0];
var city = data[1];

// Reading until delimiter
var part = scan(";;");         // Read until ";;"

// Reading character
var ch = getchar();

// Formatted string without printing
var message = sprintf("Value: %d, Name: %s", 42, "Alice");

// Error output
eprintf("Error: %s\n", "something went wrong");
eprintln("An error occurred");
```

### OS & System Functions

The `os` package provides operating system interactions.

| Function | Parameters | Returns | Description |
|:---------|:-----------|:--------|:------------|
| `getenv(key)` | string | string | Get environment variable |
| `setenv(key, value)` | string, string | nil | Set environment variable |
| `unsetenv(key)` | string | nil | Unset environment variable |
| `exec(command, arg1, ...)` | string... | string | Execute shell command |
| `exit([code])` | [int] | nil | Terminate program |
| `args()` | none | array | Get command-line arguments |
| `sleep(milliseconds)` | int | nil | Sleep for duration |
| `getpid()` | none | int | Get process ID |
| `hostname()` | none | string | Get system hostname |
| `user()` | none | string | Get current username |

**OS & System Examples:**
```go
// Environment variables
var home = getenv("HOME");
var path = getenv("PATH");

setenv("MY_VAR", "my_value");
unsetenv("MY_VAR");

// Command execution
var output = exec("echo", "Hello from shell");
var dir_list = exec("ls", "-la");
var result = exec("python", "script.py");

// Process information
var pid = getpid();
println("Process ID: " + pid);

var current_user = user();
println("User: " + current_user);

var hostname = hostname();
println("Hostname: " + hostname);

// Command line arguments
var args_array = args();  // First element is program name
foreach arg in args() {
    println(arg);
}

// Program termination
if (error_condition) {
    exit(1);  // Exit with error code
}

// Timing
sleep(1000);               // Sleep for 1 second (1000 ms)
sleep(2500);               // Sleep for 2.5 seconds
```

### Regular Expression Functions

The `regex` package provides powerful pattern matching capabilities.

| Function | Parameters | Returns | Description |
|:---------|:-----------|:--------|:------------|
| `match_regex(pattern, str)` | string, string | bool | Check if pattern matches |
| `find_regex(pattern, str)` | string, string | string | Find first match |
| `findall_regex(pattern, str, [n])` | string, string, [int] | array | Find all matches |
| `replace_regex(pattern, str, repl)` | string, string, string | string | Replace matches |
| `split_regex(pattern, str, [n])` | string, string, [int] | array | Split by pattern |

**Regex Examples:**
```go
// Pattern matching
var text = "Contact: test@example.com";

match_regex("\\w+@\\w+\\.\\w+", text);  // true (matches email)
match_regex("^Contact:", text);         // true (starts with)

// Finding matches
find_regex("[A-Z][a-z]+", text);       // "Contact"

// Finding all matches
var numbers = "Items: 10, 25, 8, 100";
var nums = findall_regex("\\d+", numbers);  // ["10", "25", "8", "100"]

// Replace matches
var redacted = replace_regex("\\d+", numbers, "X");
// "Items: X, X, X, X"

// Split by pattern
split_regex("\\s+", "hello   world   go-mix");  // ["hello", "world", "go-mix"]

// Extract email addresses
var emails = findall_regex("\\b[\\w._%+-]+@[\\w.-]+\\.\\w+\\b", email_text);

// Phone number validation
if (match_regex("\\d{3}-\\d{3}-\\d{4}", phone)) {
    println("Valid US phone format");
}

// Common patterns
match_regex("\\b\\d{1,3}\\.\\d{1,3}\\.\\d{1,3}\\.\\d{1,3}\\b", "192.168.1.1");  // IP
match_regex("^https?://", url);        // URL protocol check
```

### HTTP Client & Server

The `http` package provides comprehensive HTTP client and server functionality for web requests and building HTTP servers.

| Function | Parameters | Returns | Description |
|:---------|:-----------|:--------|:------------|
| `get_http(url)` | string | string | Perform GET request |
| `post_http(url, body)` | string, string | string | Perform POST request |
| `put_http(url, body)` | string, string | string | Perform PUT request |
| `delete_http(url)` | string | string | Perform DELETE request |
| `request_http(method, url, [headers], [body])` | string, string, [map], [string] | map | Generic HTTP request |
| `create_server()` | none | server | Create HTTP server |
| `handle_server(server, path, handler)` | server, string, function | nil | Register route handler |
| `start_server(server, address)` | server, string | nil | Start server |
| `serve_static(server, prefix, root_dir)` | server, string, string | nil | Serve static files |
| `url_encode(str)` | string | string | URL encode string |
| `url_decode(str)` | string | string | URL decode string |
| `download_file(url, path)` | string, string | nil | Download file from URL |

**HTTP Client Examples:**
```go
// Simple GET request
var response = get_http("https://httpbin.org/get");
println(response);

// POST request
var post_body = "{\"name\": \"Alice\", \"age\": 30}";
var result = post_http("https://api.example.com/users", post_body);

// Generic request with headers
var headers = map{
    "Content-Type": "application/json",
    "Authorization": "Bearer token123"
};
var body = "{\"query\": \"search term\"}";

var response = request_http("POST", "https://api.example.com/search", headers, body);
println(response["status"]);     // Status code (200, 404, 500, etc.)
println(response["body"]);       // Response body
println(response["headers"]);    // Response headers map

// Other HTTP methods
put_http("https://api.example.com/users/123", "{\"age\": 31}");
delete_http("https://api.example.com/users/123");

// URL encoding/decoding
var encoded = url_encode("Hello World!");  // "Hello%20World%21"
var decoded = url_decode("Hello%20World%21");  // "Hello World!"

// Download file
download_file("https://example.com/data.zip", "./downloads/data.zip");

// API example with error handling
var response = request_http("GET", "https://api.example.com/data", nil, nil);
if (response["status"] == 200) {
    println("Success: " + response["body"]);
} else {
    println("Error: Status " + response["status"]);
}
```

**HTTP Server Examples:**
```go
// Create a basic HTTP server
var srv = create_server();

// Simple text response
handle_server(srv, "/", func(req) {
    return "Welcome to Go-Mix!";
});

// JSON API endpoint
handle_server(srv, "/api/users", func(req) {
    var users = [
        map{"id": 1, "name": "Alice"},
        map{"id": 2, "name": "Bob"}
    ];
    return map{
        "status": 200,
        "headers": map{"Content-Type": "application/json"},
        "body": map_to_json_string(map{"users": users})
    };
});

// Serve static files (CSS, JS, HTML, images, etc.)
serve_static(srv, "/static/", "./public/");
serve_static(srv, "/assets/", "./assets/");

// Request inspection
handle_server(srv, "/info", func(req) {
    var info = "Method: " + req["method"] + "\n" +
               "URL: " + req["url"] + "\n" +
               "Timestamp: " + now();
    return info;
});

// Start server on port 8080
start_server(srv, ":8080");
println("Server listening on http://localhost:8080");
```

### JSON Functions

JSON handling is integrated into the string and map functions.

| Function | Parameters | Returns | Description |
|:---------|:-----------|:--------|:------------|
| `map_to_json_string(map)` | map | string | Convert map to JSON string |
| `json_string_to_map(json_str)` | string | map | Parse JSON string to map |
| `json_encode(obj)` | any | string | Encode to JSON (alias) |

**JSON Examples:**
```go
// Create data structures
var user = map{
    "id": 123,
    "name": "Alice",
    "email": "alice@example.com",
    "age": 30,
    "active": true,
    "roles": ["admin", "developer"],
    "metadata": map{
        "created": "2024-01-01",
        "updated": "2024-02-10"
    }
};

// Encode to JSON string
var json_str = map_to_json_string(user);
// {"id":123,"name":"Alice","email":"alice@example.com","age":30,"active":true,"roles":["admin","developer"],"metadata":{"created":"2024-01-01","updated":"2024-02-10"}}

// Pretty print
println(json_str);

// Decode from JSON string
var parsed = json_string_to_map(json_str);
println(parsed["name"]);           // "Alice"
println(parsed["age"]);            // 30
println(parsed["roles"][0]);       // "admin"
println(parsed["metadata"]["created"]);  // "2024-01-01"

// Array of objects
var users_data = [
    map{"id": 1, "name": "Alice"},
    map{"id": 2, "name": "Bob"},
    map{"id": 3, "name": "Charlie"}
];

var users_map = map{"users": users_data};
var json = map_to_json_string(users_map);
var restored = json_string_to_map(json);

// Working with JSON arrays
foreach user in restored["users"] {
    println(user["name"]);
}
```

---

## Error Handling

### Panic & Recovery

```go
// Panic terminates execution with error message
func divide(a, b) {
    if (b == 0) {
        panic("Division by zero!");
    }
    return a / b;
}

// In normal execution:
println(divide(10, 2));     // 5
println(divide(10, 0));     // Panics with error
```

### Type Safety

```javascript
// Static variables prevent type errors
let count = 0;
count = count + 1;          // OK
count = "text";             // ERROR at runtime

// Dynamic variables allow flexibility
var data = 42;
data = "text";              // OK - type changed
data = [1, 2, 3];          // OK - type changed again
```

### Error Objects

```go
// Some functions return error objects
var result = someOperation();
if (result.GetType() == "error") {
    println("Error: " + result.ToString());
}
```

---

## 📦 Recent Updates & New Features

### V1.0.0 - Comprehensive Standard Library

#### New Standard Library Packages

**1. Arrays Package (`arrays.go`)**
- Comprehensive array manipulation with `push`, `pop`, `shift`, `unshift`
- Functional operations: `map_array`, `filter_array`, `reduce_array`
- Predicates: `find_array`, `some_array`, `every_array`
- Sorting: `sort`, `sorted`, `csort`, `csorted` with custom comparators
- Utilities: `clone_array`, `reverse`, `to_array`

**2. Lists Package (`list.go`)**
- Mutable heterogeneous list operations
- Operations: `pushback_list`, `pushfront_list`, `popback_list`, `popfront_list`
- Access: `peekback_list`, `peekfront_list`, `size_list`
- Modification: `insert_list`, `remove_list`, `contains_list`
- Functional methods: `map_list`, `filter_list`, `reduce_list`

**3. Tuples Package (`tuple.go`)**
- Immutable fixed-size sequences
- Creation: `tuple()` for creating tuples
- Access: `peekback_tuple`, `peekfront_tuple`, `size_tuple`
- Searching: `contains_tuple`
- Functional: `map_tuple`, `filter_tuple`, `reduce_tuple`

**4. Maps Package (`maps.go`)**
- Dictionary/hash map operations
- Functions: `keys_map`, `insert_map`, `remove_map`, `contain_map`, `enumerate_map`
- JSON integration: `json_string_to_map`, `map_to_json_string`

**5. Sets Package (`sets.go`)**
- Unique value collections
- Operations: `insert_set`, `remove_set`, `contains_set`, `values_set`, `size_set`

**6. Strings Package (`strings.go`)**
- Complete string manipulation suite (21 functions)
- Case: `upper`, `lower`, `capitalize`
- Searching: `contains`, `starts_with`, `ends_with`, `index`, `count`
- Manipulation: `split`, `join`, `replace`, `reverse`, `substring`
- Character operations: `ord`, `chr`
- Type checking: `is_digit`, `is_alpha`
- Comparison: `strcmp`

**7. Math Package (`math.go`)**
- Mathematical functions (20 functions)
- Basic: `abs`, `min`, `max`, `floor`, `ceil`, `round`
- Powers: `sqrt`, `pow`, `exp`
- Trigonometry: `sin`, `cos`, `tan`, `asin`, `acos`, `atan`, `atan2`
- Logarithms: `log`, `log10`
- Random: `rand`, `rand_int`

**8. Time Package (`time.go`)**
- Time handling and formatting
- Functions: `now`, `now_ms`, `utc_now`, `format_time`, `parse_time`, `timezone`
- Support for flexible time formatting with Go's reference layout

**9. File I/O Package (`file_io.go`)**
- Comprehensive file system operations (17 functions)
- Reading/Writing: `read_file`, `write_file`, `append_file`, `cat`
- File properties: `file_exists`, `is_dir`, `is_file`
- Directory operations: `mkdir`, `list_dir`, `pwd`, `home`
- File manipulation: `touch`, `remove_file`, `remove_all`, `rename_file`, `truncate_file`, `chmod`

**10. I/O Package (`io.go`)**
- User input and output operations (11 functions)
- Input: `scanln`, `scanf`, `input`, `scan`, `getchar`
- Output: `putchar`, `puts`, `sprintf`
- Error output: `eprintln`, `eprintf`
- Buffering: `flush`

**11. OS Package (`os.go`)**
- Operating system interactions (10 functions)
- Environment: `getenv`, `setenv`, `unsetenv`
- Process: `exec`, `exit`, `getpid`, `sleep`
- System info: `hostname`, `user`, `args`

**12. Format Package (`format.go`)**
- Type conversion functions (5 functions)
- Conversions: `to_int`, `to_float`, `to_bool`, `to_string`, `to_char`
- Support for base conversion (hex 0x, octal 0o)

**13. Regex Package (`regex.go`)**
- Regular expression support (5 functions)
- Pattern matching: `match_regex`, `find_regex`, `findall_regex`
- Manipulation: `replace_regex`, `split_regex`
- Full Go regex support

**14. HTTP Package (`http.go`)**
- HTTP client and server support (13 functions)
- Client: `get_http`, `post_http`, `put_http`, `delete_http`, `request_http`
- Server: `create_server`, `handle_server`, `start_server`, `serve_static`
- Utilities: `url_encode`, `url_decode`, `download_file`

#### Common Functions Enhanced
- `print`, `println`, `printf` - Output with proper formatting
- `length`/`size` - Works with arrays, strings, maps, sets, lists, tuples
- `typeof` - Type inspection
- `tostring` - String conversion
- `range` - Create inclusive ranges
- `addr`, `is_same_ref` - Reference checking and memory inspection
- Functional operations: `map`, `filter`, `reduce`, `find`, `some`, `every`

#### Total Functions Added
- **100+ builtin functions** across 14 packages
- **Fully documented** with syntax and examples
- **Package-based organization** for easy access
- **Backward compatible** - all functions available globally

### Enhanced Language Features

**1. Collections Support**
- Arrays: Mutable sequences with full manipulation support
- Lists: Heterogeneous mutable collections
- Tuples: Fixed-size immutable sequences
- Maps: Key-value dictionaries
- Sets: Unique value collections

**2. Functional Programming**
- `map`, `filter`, `reduce` for arrays and lists
- Higher-order functions with custom comparators
- First-class functions and lambda expressions
- Full closure support

**3. String Processing**
- 21 string manipulation functions
- Pattern matching with regex support
- Unicode-aware operations
- Format string support

**4. System Integration**
- File I/O operations
- Environment variables
- Process execution
- System information

**5. Web Development**
- HTTP client for making requests
- HTTP server for building services
- JSON encoding/decoding
- URL encoding/decoding
- Static file serving

**6. Mathematical Computing**
- Trigonometric functions
- Logarithmic functions
- Random number generation
- Type-safe number operations

### Documentation Updates
- Comprehensive function reference for each package
- Detailed syntax and parameter descriptions
- Real-world usage examples for each function
- Common patterns and best practices
- Error handling guidelines

### Code Quality Improvements
- Consistent error handling across all builtin functions
- Type checking and validation
- Informative error messages
- Support for variadic arguments where appropriate

---

The `samples/` directory contains 200+ example programs organized by topic:

### Algorithms (`samples/algo/`)

- **01_linear_search.gm** — Sequential search in array
- **02_binary_search.gm** — Logarithmic search in sorted array
- **03_insertion_sort.gm** — Insertion sort algorithm
- **04_selection_sort.gm** — Selection sort algorithm
- **05_factorial.gm** — Factorial (recursive and iterative)
- **06_fibonacci.gm** — Fibonacci sequence (multiple approaches)
- **07_gcd.gm** — Greatest common divisor
- **08_prime_check.gm** — Prime number checking
- **09_bubble_sort.gm** — Bubble sort implementation
- **10_heap_sort.gm** — Heap sort algorithm

**Run any sample:**
```bash
./go-mix samples/algo/05_factorial.gm
```

### Arrays & Collections (`samples/arrays/`, `samples/maps/`, `samples/lists/`)

- **arrays/01_test_array_features.gm** — Array operations and methods
- **maps/01_test_map_operations.gm** — Dictionary/map examples
- **lists/01_test_list_features.gm** — List manipulation

### Strings & Formatting (`samples/strings/`, `samples/format/`)

- **strings/01_test_string_operations.gm** — String manipulation examples
- **format/01_sprintf_demo.gm** — String formatting with printf

### File I/O & System (`samples/files/`, `samples/os/`)

- **files/04_file_basics.gm** — File reading and writing
- **os/01_environment_demo.gm** — Environment variables and process info

### Objects & Structs (`samples/structs/`)

Examples demonstrating object-oriented programming with structs and methods.

---

## Development & Contributing

### Building from Source

**Manual Build:**
```bash
cd /path/to/go-mix
go build -o go-mix ./main
```

**With Script:**
```bash
./build.sh
```

### Running Tests

```bash
# All tests
go test ./...

# Specific package
go test ./lexer
go test ./parser
go test ./eval

# Verbose output
go test -v ./...

# Coverage report
go test -coverage ./...

# Run with filtering
go test -run TestName ./package
```

### Project Directory Structure

```bash
go-mix/
├── main/                  # REPL & File Execution Entry Point
│   ├── main.go           # Main interpreter logic (REPL & file execution modes)
│   ├── print_visitor.go  # AST printing utilities for debugging
│   └── main_test.go      # Main function tests
│
├── lexer/                 # Lexical Analysis (Source → Tokens)
│   ├── lexer.go          # Core lexer implementation, tokenization logic
│   ├── token.go          # Token type definitions and constants
│   ├── lexer_utils.go    # Helper functions for lexing
│   └── lexer_test.go     # Lexer unit tests
│
├── parser/                # Parsing & AST (Tokens → Abstract Syntax Tree)
│   ├── parser.go         # Pratt parser main implementation
│   ├── parser_expressions.go   # Expression parsing logic
│   ├── parser_precedence.go    # Operator precedence definitions
│   ├── node.go           # AST node structures (all node types)
│   ├── test_visitor.go   # AST visitor for testing
│   └── parser_test.go    # Parser unit tests
│
├── eval/                  # Evaluation Engine (AST → Execution)
│   ├── evaluator.go      # Core evaluator & runtime state management
│   ├── evaluator_expressions.go  # Expression evaluation logic
│   ├── evaluator_helpers.go      # Helper utilities for evaluation
│   └── evaluator_test.go # Evaluator unit tests
│
├── repl/                  # Interactive REPL Loop
│   └── repl.go           # Read-Eval-Print-Loop implementation
│
├── std/                   # Standard Library & Builtins
│   ├── builtins.go       # Builtin function registry and interface
│   ├── common.go         # Common utilities (print, println, length, typeof, etc.)
│   ├── arrays.go         # Array manipulation (push, pop, slice, map, filter, reduce, sort, etc.)
│   ├── list.go           # List operations (heterogeneous mutable sequences)
│   ├── tuple.go          # Tuple operations (fixed immutable sequences)
│   ├── map.go            # Map/Dictionary operations (key-value pairs)
│   ├── set.go            # Set operations (unique value collections)
│   ├── strings.go        # String utilities (upper, lower, split, join, substring, etc.)
│   ├── math.go           # Mathematical functions (abs, sqrt, pow, floor, ceil, min, max, etc.)
│   ├── file_io.go        # File operations (read_file, write_file, append_file, file_exists, etc.)
│   ├── io.go             # Input/output utilities (scanf, input from stdin)
│   ├── os.go             # OS interactions (getenv, setenv, exec, getpid, sleep, user, etc.)
│   ├── time.go           # Time handling (now, now_ms, format_time, sleep, etc.)
│   ├── format.go         # Type conversion & formatting (sprintf, json_encode, etc.)
│   ├── struct.go         # Struct/object utilities
│   ├── regex.go          # Regular expression support
│   ├── http.go           # HTTP client and server implementation
│   └── types.go          # Type system utilities
│
├── file/                  # Stateful File I/O Handles
│   └── file.go           # File object implementation (file handles for advanced I/O)
│
├── scope/                 # Lexical Scoping & Variable Management
│   └── scope.go          # Scope management, variable binding, environment handling
│
├── function/              # Function Objects & Closures
│   └── function.go       # Function representation, closure implementation
│
├── samples/               # Example Programs (50+ .gm files)
│   ├── algo/             # Algorithms & Data Structures (50 examples)
│   │
│   ├── arrays/           # Array Operations & Features (11 examples)
│   │
│   ├── functions/        # Function Examples (14 examples)
│   │
│   ├── loops/            # Loop Constructs (20 examples)
│   │
│   ├── maps/             # Map/Dictionary Operations (8 examples)
│   │
│   ├── lists/            # List Operations (7 examples)
│   │
│   ├── sets/             # Set Operations (4 examples)
│   │
│   ├── tuples/           # Tuple Operations (6 examples)
│   │
│   ├── strings/          # String Manipulation (5 examples)
│   │
│   ├── format/           # Formatting Examples (5 examples)
│   │
│   ├── files/            # File I/O Examples (10 examples + test files)
│   │
│   ├── structs/          # Object-Oriented Examples (40 examples)
│   │
│   ├── os/               # OS & Environment Examples (5 examples)
│   │
│   ├── time/             # Time Handling Examples (5 examples)
│   │
│   ├── range/            # Range Iteration Examples (7 examples)
│   │
│   └── typeof/           # Type Checking Examples (3 examples)
│
├── vscode-ext/            # VS Code Extension Source
│   ├── package.json      # Extension manifest
│   ├── language-configuration.json # Language config
│   └── syntaxes/         # TextMate grammar files
│
├── go.mod                 # Go module definition (github.com/akashmaji946/go-mix)
├── go.sum                 # Go dependencies lock file
├── build.sh              # Build script (runs tests + compiles binary)
├── run.sh                # Test + Executable runner script
├── LICENSE               # MIT License
├── logo.PNG              # Project logo
├── README.md             # This comprehensive documentation
└── PROJECT_STRUCTURE.md  # Detailed project structure documentation
```

### Key Module Responsibilities

**Lexer Package** (`lexer/`)
- Converts source code characters into tokens
- Handles keywords, operators, identifiers, literals
- Supports comments (single-line `//` and multi-line `/* */`)
- Token types: Keywords, identifiers, numbers, strings, operators, delimiters

**Parser Package** (`parser/`)
- Implements Pratt Parser (Top-Down Operator Precedence)
- Creates Abstract Syntax Tree (AST) from tokens
- Handles operator precedence and associativity
- Supports all language constructs: declarations, expressions, statements, functions, structs
- Error collection and reporting

**Evaluator Package** (`eval/`)
- Walks the AST and executes code
- Manages runtime state and environments
- Handles variable scoping and closures
- Implements all language semantics
- Error handling and panic recovery

**Standard Library Package** (`std/`)
- Provides 100+ builtin functions
- Organized by type: arrays, strings, math, file I/O, OS, time, etc.
- Each file focuses on one collection type or functional area
- All builtins registered in `builtins.go`

### Adding New Features

1. **Add lexer support** (if new syntax needed)
   - Update `lexer/token.go` with new token type
   - Modify `lexer/lexer.go` tokenization logic

2. **Extend parser** (if new AST nodes needed)
   - Add node struct to `parser/node.go`
   - Implement parsing logic in `parser/parser*.go`

3. **Implement in evaluator**
   - Add evaluation logic to `eval/evaluator*.go`
   - Handle new node type in evaluation switch

4. **Add builtin functions** (easiest approach)
   - Implement function in appropriate `std/*.go` file
   - Register in builtins slice in init()

---

## �️ Installation & Tools

### System-Wide Installation

To install Go-Mix globally on your system:

```bash
sudo ./install.sh
```

This script:
- ✅ Verifies Go is installed (v1.18+)
- ✅ Builds the executable from source
- ✅ Installs to `/usr/local/bin`
- ✅ Makes `go-mix` accessible from anywhere

For detailed installation instructions, see [INSTALL.md](INSTALL.md).

#### Install Commands
```bash
sudo ./install.sh         # Install system-wide (recommended)
./install.sh rebuild      # Just build, don't install
sudo ./install.sh uninstall  # Remove installation
./install.sh clean        # Remove build artifacts
```

### VS Code Extension

A professional VS Code extension is available for Go-Mix development:

**Features:**
- 🎨 Syntax highlighting for `.gm` files


**Installation:**
Run this command:
```bash
code --install-extension go-mix-v1.0.0.vsix
```

Or install from the [VS Code Marketplace](https://marketplace.visualstudio.com/items?itemName=akashmaji946.go-mix)


---

## �📝 License & Attribution

**License**: MIT

**Author**: Akash Maji  
**Contact**: akashmaji@iisc.ac.in  
**Repository**: [github.com/akashmaji946/go-mix](https://github.com/akashmaji946/go-mix)

Go-Mix is designed for:
- **Education** — Learning language design and interpreter implementation
- **Prototyping** — Rapid algorithm and idea development
- **Embedding** — Adding scripting capabilities to JS like applications
- **Experimentation** — Testing language features and design patterns

Contributions, issues, and feedback are welcome!

---

## Quick Reference Card

```go
// Variables
var x = 10;      let y = 3.14;    const Z = 42;

// Arrays & Collections
var arr = [1,2,3];              var map = map{"key": "value"};
arr[0];          arr[-1];        map["key"];

// Functions
func add(a,b) { return a + b; }   var f = func(x){ return x*2; };

// Control Flow
if (cond) { } else if { } else { }
for (var i=0; i<n; i=i+1) { }   while (cond) { }
foreach x in arr { }            foreach i, v in arr { }

// Functional
map(arr, func(x){ return x*2; });
filter(arr, func(x){ return x>5; });
reduce(arr, func(a,x){ return a+x; }, 0);

// Classes & Objects
struct Name { var x; func method() { } }
var obj = new Name();

// Common Functions
print(), println(), printf(), length(), typeof(), tostring()

// String Methods
upper(), lower(), split(), join(), contains(), substring()

// Math Functions
abs(), min(), max(), floor(), ceil(), sqrt(), pow()

// File I/O
read_file(), write_file(), append_file(), file_exists()

// Collections
list(), tuple(), insert_set()

// Network & Web
get_http(), post_http(), create_server()

// System
getenv(), setenv(), exec(), getpid(), sleep(), now()
```

---

**Happy Coding with Go-Mix!** 🚀
---
