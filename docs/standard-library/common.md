---
title: "Common"
layout: default
parent: Standard Library
nav_order: 1
description: "Core functions including print, length, typeof, range, and type constructors"
permalink: /standard-library/common/
---

# Common Package
{: .no_toc }

Core functions including print, length, typeof, range, and type constructors
{: .fs-6 .fw-300 }

## Table of Contents
{: .no_toc .text-delta }

1. TOC
{:toc}

---

## Import

`import "common"`
{: .fs-5 .fw-300 }

Import the common package to use namespaced functions.

```go
// Standard import
import common;
common.print("Hello");
common.println("World");

// With alias
import "common" as c
c.print("Hello")
c.println("World")
```

---

## print

`print(...args) -> nil`
{: .fs-5 .fw-300 }

Outputs the string representations of arguments without a trailing newline.

```go
print("Hello");              // Output: Hello (no newline)
print("World");            // Output: World
print("Value:", 42);       // Output: Value: 42
```

---

## println

`println(...args) -> nil`
{: .fs-5 .fw-300 }

Outputs the string representations of arguments with a trailing newline.

```go
println("Hello");          // Output: Hello\n
println("World");          // Output: World\n
println("Value:", 42);     // Output: Value: 42\n
```

---

## printf

`printf(format, ...args) -> nil`
{: .fs-5 .fw-300 }

Outputs a formatted string using C-style format specifiers.

```go
printf("Number: %d\n", 42);           // Number: 42
printf("String: %s\n", "hello");       // String: hello
printf("Float: %.2f\n", 3.14159);     // Float: 3.14
printf("Hex: %x\n", 255);              // Hex: ff
printf("Multiple: %d %s\n", 1, "a"); // Multiple: 1 a
```

---

## length / size

`length(obj) -> int`
{: .fs-5 .fw-300 }

Returns the length of strings, arrays, maps, sets, lists, or tuples.

```go
length("hello");           // 5
length([1, 2, 3]);         // 3
length(map{"a": 1, "b": 2}); // 2
length(set{1, 2, 3});      // 3
size("hello");             // 5 (alias)
```

---

## typeof

`typeof(obj) -> string`
{: .fs-5 .fw-300 }

Returns the type name of any object as a string.

```go
typeof(42);                // "int"
typeof(3.14);              // "float"
typeof("hello");           // "string"
typeof(true);              // "bool"
typeof(nil);               // "nil"
typeof([1, 2, 3]);         // "array"
typeof(range(1, 5));       // "range"
typeof(func() {});           // "func"
```

---

## range

`range(start, end) -> range`
{: .fs-5 .fw-300 }

Creates an inclusive range object for iteration.

```go
// Using range() function
foreach i in range(1, 5) {
    println(i);            // 1, 2, 3, 4, 5
}

// Using ... operator (equivalent)
foreach i in 1...5 {
    println(i);            // 1, 2, 3, 4, 5
}

// Reverse range
foreach i in range(5, 1) {
    println(i);            // 5, 4, 3, 2, 1
}
```

---

## to_string

`to_string(obj) -> string`
{: .fs-5 .fw-300 }

Converts any object to its string representation.

```go
to_string(123);            // "123"
to_string(3.14);           // "3.14"
to_string(true);           // "true"
to_string([1, 2, 3]);       // "[1, 2, 3]"
to_string(map{"a": 1});    // "map{a: 1}"
```

---

## array

`array(...args) -> array`
{: .fs-5 .fw-300 }

Creates a new array from arguments or converts an iterable to an array.

```go
array();                          // []
array(1, 2, 3);                   // [1, 2, 3]
array(list(1, 2, 3));             // [1, 2, 3]
array(tuple(1, 2, 3));            // [1, 2, 3]
array(map{"a": 1, "b": 2});       // [1, 2] (values)
array(range(1, 3));               // [1, 2, 3]
```

---

## list

`list(...args) -> list`
{: .fs-5 .fw-300 }

Creates a new mutable list from elements.

```go
list();                           // list()
list(1, 2, 3);                    // list(1, 2, 3)
list(1, "two", 3.0, true);        // list(1, two, 3.0, true)
```

---

## tuple

`tuple(...args) -> tuple`
{: .fs-5 .fw-300 }

Creates a new immutable tuple from elements.

```go
tuple();                          // tuple()
tuple(1, 2, 3);                   // tuple(1, 2, 3)
tuple(10, 20);                    // tuple(10, 20) - for coordinates
tuple("John", 25, true);          // tuple(John, 25, true) - for records
```

---

## addr

`addr(obj) -> int`
{: .fs-5 .fw-300 }

Returns the memory address of an object as an integer. Useful for reference checking.

```go
var a = [1, 2];
println(addr(a));              // Prints memory address

var b = a;
println(addr(b));              // Same address as a
```

---

## is_same_ref

`is_same_ref(obj1, obj2) -> bool`
{: .fs-5 .fw-300 }

Checks if two objects point to the same memory address (same reference).

```go
var a = [1];
var b = a;
var c = [1];

println(is_same_ref(a, b));    // true (same reference)
println(is_same_ref(a, c));    // false (different objects)
```

---
