---
title: "Tuples"
layout: default
parent: Standard Library
nav_order: 7
description: "Immutable fixed-size sequences with functional operations"
permalink: /standard-library/tuples/
---

# Tuples Package
{: .no_toc }

Immutable fixed-size sequences with functional operations
{: .fs-6 .fw-300 }

## Table of Contents
{: .no_toc .text-delta }

1. TOC
{:toc}

---

## Import

`import "tuples"`
{: .fs-5 .fw-300 }

Import the tuples package to use namespaced functions.

```go
// Standard import
import tuple;
var t = tuple.tuple(1, 2, 3)
var size = tuple.size_tuple(t)

// With alias
import tuple as tup;
var t = tup.tuple(1, 2, 3);
var size = tup.size_tuple(t);
```

---

## make_tuple

`make_tuple(...elements) -> tuple`
{: .fs-5 .fw-300 }

Creates a new immutable tuple from arguments. Tuples are heterogeneous and immutable, preventing modifications after creation.

```go
make_tuple();                          // tuple()
make_tuple(1, 2, 3);                   // tuple(1, 2, 3)
make_tuple("John", 25, true);          // tuple(John, 25, true)
```

---

## size_tuple

`size_tuple(t) -> int`
{: .fs-5 .fw-300 }

Returns the number of elements in a tuple.

```go
var t = make_tuple(1, 2, 3);
size_tuple(t);             // 3
```

---

## length_tuple

`length_tuple(t) -> int`
{: .fs-5 .fw-300 }

Returns the number of elements in a tuple (alias for size_tuple).

```go
var t = make_tuple(1, 2, 3);
length_tuple(t);           // 3
```

---

## peekback_tuple

`peekback_tuple(t) -> any`
{: .fs-5 .fw-300 }

Returns the last element of a tuple. Returns an error if the tuple is empty.

```go
var t = make_tuple(1, 2, 3);
peekback_tuple(t);         // 3
```

---

## peekfront_tuple

`peekfront_tuple(t) -> any`
{: .fs-5 .fw-300 }

Returns the first element of a tuple. Returns an error if the tuple is empty.

```go
var t = make_tuple(1, 2, 3);
peekfront_tuple(t);        // 1
```

---

## contains_tuple

`contains_tuple(t, elem) -> bool`
{: .fs-5 .fw-300 }

Checks if a value exists in the tuple. Comparison is done using both type and value.

```go
var t = make_tuple(1, 2, 3, 4);
contains_tuple(t, 3);      // true
contains_tuple(t, 5);      // false
contains_tuple(t, "2");    // false (type matters)
```

---

## find_tuple

`find_tuple(t, function) -> any`
{: .fs-5 .fw-300 }

Finds the first element that satisfies the provided testing function. Returns nil if no element matches.

```go
var t = make_tuple(1, 2, 3, 4, 5);
find_tuple(t, fn(x) { x > 3 });   // 4
find_tuple(t, fn(x) { x > 10 });  // nil
```

---

## some_tuple

`some_tuple(t, function) -> bool`
{: .fs-5 .fw-300 }

Checks if at least one element in the tuple passes the test (predicate function).

```go
var t = make_tuple(1, 2, 3);
some_tuple(t, fn(x) { x > 2 });   // true
some_tuple(t, fn(x) { x > 5 });   // false
```

---

## every_tuple

`every_tuple(t, function) -> bool`
{: .fs-5 .fw-300 }

Checks if all elements in the tuple pass the test (predicate function).

```go
var t = make_tuple(2, 4, 6);
every_tuple(t, fn(x) { x % 2 == 0 });  // true
every_tuple(t, fn(x) { x > 3 });       // false
```

---

## map_tuple

`map_tuple(t, function) -> list`
{: .fs-5 .fw-300 }

Applies a function to each element of the tuple and returns a list with the results.

```go
var t = make_tuple(1, 2, 3);
map_tuple(t, fn(x) { x * 2 });   // [2, 4, 6]
```

---

## filter_tuple

`filter_tuple(t, function) -> list`
{: .fs-5 .fw-300 }

Filters elements based on a predicate function and returns a list with matching elements.

```go
var t = make_tuple(1, 2, 3, 4, 5);
filter_tuple(t, fn(x) { x > 2 });   // [3, 4, 5]
```

---

## reduce_tuple

`reduce_tuple(t, function, initial) -> any`
{: .fs-5 .fw-300 }

Reduces the tuple to a single value using a binary function (accumulator, current) -> newAccumulator.

```go
var t = make_tuple(1, 2, 3, 4);
reduce_tuple(t, fn(acc, x) { acc + x }, 0);   // 10
reduce_tuple(t, fn(acc, x) { acc * x }, 1);   // 24
```

---

## to_tuple

`to_tuple(iterable) -> tuple`
{: .fs-5 .fw-300 }

Converts an array or list to a tuple. If the argument is already a tuple, returns it unchanged.

```go
var arr = [1, 2, 3];
to_tuple(arr);             // tuple(1, 2, 3)

var lst = list(4, 5, 6);
to_tuple(lst);             // tuple(4, 5, 6)
```

---
