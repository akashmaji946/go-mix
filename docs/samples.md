---
title: Samples
layout: default
nav_order: 5
description: "Example programs demonstrating Go-Mix language features"
permalink: /samples/
---

# Sample Programs
{: .no_toc }

Explore Go-Mix through practical examples organized by category.
{: .fs-6 .fw-300 }

## Table of Contents
{: .no_toc .text-delta }

1. TOC
{:toc}

---

## Algorithms

### Factorial

```go
func factorial(n) {
    if (n <= 1) {
        return 1;
    }
    return n * factorial(n - 1);
}

println(factorial(5));  // 120
```

### Fibonacci

```go
func fibonacci(n) {
    if (n <= 1) {
        return n;
    }
    return fibonacci(n - 1) + fibonacci(n - 2);
}

for (var i = 0; i < 10; i = i + 1) {
    print(fibonacci(i) + " ");
}
// 0 1 1 2 3 5 8 13 21 34
```

### Binary Search

```go
func binary_search(arr, target) {
    var left = 0;
    var right = length(arr) - 1;

    while (left <= right) {
        var mid = (left + right) / 2;
        if (arr[mid] == target) {
            return mid;
        }
        if (arr[mid] < target) {
            left = mid + 1;
        } else {
            right = mid - 1;
        }
    }
    return -1;
}

var sorted = [1, 3, 5, 7, 9, 11, 13];
println(binary_search(sorted, 7));  // 3
```

---

## Data Structures

### Stack Implementation

```go
struct Stack {
    var items;

    func init() {
        this.items = [];
    }

    func push(item) {
        push(this.items, item);
    }

    func pop() {
        return pop(this.items);
    }

    func peek() {
        return this.items[-1];
    }

    func is_empty() {
        return length(this.items) == 0;
    }
}

var stack = new Stack();
stack.push(10);
stack.push(20);
println(stack.pop());   // 20
```

### Linked List

```go
struct Node {
    var data;
    var next;

    func init(data) {
        this.data = data;
        this.next = nil;
    }
}

struct LinkedList {
    var head;

    func append(data) {
        var new_node = new Node(data);
        if (this.head == nil) {
            this.head = new_node;
            return;
        }
        var current = this.head;
        while (current.next != nil) {
            current = current.next;
        }
        current.next = new_node;
    }

    func print_list() {
        var current = this.head;
        while (current != nil) {
            print(current.data + " -> ");
            current = current.next;
        }
        println("nil");
    }
}

var list = new LinkedList();
list.append(1);
list.append(2);
list.append(3);
list.print_list();  // 1 -> 2 -> 3 -> nil
```

---

## Functions

### Higher-Order Functions

```go
func apply_twice(f, x) {
    return f(f(x));
}

func add_ten(x) {
    return x + 10;
}

var result = apply_twice(add_ten, 5);
println(result);  // 25 (5 + 10 + 10)

// Closure
func make_multiplier(factor) {
    return func(x) {
        return x * factor;
    };
}

var triple = make_multiplier(3);
println(triple(4));  // 12
```

---

## Collections

### Array Operations

```go
var numbers = [1, 2, 3, 4, 5];

// Map
var doubled = map(numbers, func(x) { return x * 2; });
println(doubled);  // [2, 4, 6, 8, 10]

// Filter
var evens = filter(numbers, func(x) { return x % 2 == 0; });
println(evens);    // [2, 4]

// Reduce
var sum = reduce(numbers, func(a, x) { return a + x; }, 0);
println(sum);      // 15

// Sort
var unsorted = [3, 1, 4, 1, 5, 9, 2, 6];
sort(unsorted);
println(unsorted); // [1, 1, 2, 3, 4, 5, 6, 9]
```

### Map Operations

```go
var user = map{
    "name": "Alice",
    "age": 30,
    "city": "New York"
};

println(user["name"]);           // Alice
insert_map(user, "email", "alice@example.com");

var keys = keys_map(user);
foreach key in keys {
    println(key + ": " + user[key]);
}
```

---

## Loops

### Loop Patterns

```go
// For loop
for (var i = 0; i < 5; i = i + 1) {
    print(i + " ");
}
// 0 1 2 3 4

// While loop
var n = 5;
while (n > 0) {
    print(n + " ");
    n = n - 1;
}
// 5 4 3 2 1

// Foreach
var fruits = ["apple", "banana", "cherry"];
foreach fruit in fruits {
    println(fruit);
}

// Foreach with index
foreach i, fruit in fruits {
    println(i + ": " + fruit);
}

// Nested loops with break
for (var i = 0; i < 3; i = i + 1) {
    for (var j = 0; j < 3; j = j + 1) {
        if (i == 1 && j == 1) {
            break;
        }
        print("(" + i + "," + j + ") ");
    }
}
```

---

## Object-Oriented Programming

### Class with Methods

```go
struct Rectangle {
    var width;
    var height;

    func init(w, h) {
        this.width = w;
        this.height = h;
    }

    func area() {
        return this.width * this.height;
    }

    func perimeter() {
        return 2 * (this.width + this.height);
    }

    func scale(factor) {
        this.width = this.width * factor;
        this.height = this.height * factor;
    }
}

var rect = new Rectangle(5, 3);
println("Area: " + rect.area());           // 15
println("Perimeter: " + rect.perimeter()); // 16

rect.scale(2);
println("New area: " + rect.area());       // 60
```

### Inheritance Pattern

```go
struct Dog {
    var name;
    var breed;

    func init(name, breed) {
        this.name = name;
        this.breed = breed;
    }

    func speak() {
        return "Woof!";
    }

    func fetch() {
        return this.name + " is fetching";
    }
}

var dog = new Dog("Buddy", "Golden Retriever");
println(dog.speak());   // Woof!
println(dog.fetch());   // Buddy is fetching
```

---

## File I/O

### File Operations

```go
// Write to file
write_file("data.txt", "Hello, World!");

// Read from file
var content = read_file("data.txt");
println(content);

// Append to file
append_file("log.txt", "New log entry\n");

// Check if exists
if (file_exists("config.json")) {
    var config = read_file("config.json");
    println("Config loaded");
}

// List directory
var files = list_dir(".");
foreach file in files {
    println(file);
}
```

---

## HTTP

### HTTP Client

```go
// GET request
var response = get_http("https://api.example.com/users");
println(response);

// POST request
var body = '{"name": "John", "age": 30}';
var result = post_http("https://api.example.com/users", body);
println(result);
```

### HTTP Server

```go
var srv = create_server();

handle_server(srv, "/", func(req) {
    return "Welcome to Go-Mix Server!";
});

handle_server(srv, "/api/hello", func(req) {
    return '{"message": "Hello, World!"}';
});

handle_server(srv, "/api/echo", func(req) {
    return '{"method": "' + req.method + '", "path": "' + req.path + '"}';
});

println("Server starting on http://localhost:8080");
start_server(srv, ":8080");
```

{: .tip }
> Check the `samples/` directory in the Go-Mix repository for 50+ complete example programs covering algorithms, data structures, and language features.
