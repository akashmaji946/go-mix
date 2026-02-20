---
title: "Lists"
layout: default
parent: Standard Library
nav_order: 6
description: "List operations for mutable heterogeneous sequences"
permalink: /standard-library/lists/
---

# Lists Package
{: .no_toc }

List operations for mutable heterogeneous sequences
{: .fs-6 .fw-300 }

## Table of Contents
{: .no_toc .text-delta }

1. TOC
{:toc}

---

## Import

`import "lists"`
{: .fs-5 .fw-300 }

Import the lists package to use namespaced functions.

```go
// Standard import
import lists;
var l = lists.list(1, 2, 3)
lists.pushback_list(l, 4)

// With alias
import lists as lst;
var l = lst.list(1, 2, 3)
lst.pushback_list(l, 4)
```

---

## list

`list(...elements) -> list`
{: .fs-5 .fw-300 }

Creates a new list from elements.

```go
list();                           // list()
list(1, 2, 3);                    // list(1, 2, 3)
list(1, "two", 3.0, true);        // list(1, two, 3.0, true)
```

---

## pushback_list

`pushback_list(l, elem) -> list`
{: .fs-5 .fw-300 }

Adds element to the end of the list.

```go
var l = list(1, 2, 3);
pushback_list(l, 4);       // l = list(1, 2, 3, 4)
```

---

## pushfront_list

`pushfront_list(l, elem) -> list`
{: .fs-5 .fw-300 }

Adds element to the beginning of the list.

```go
var l = list(1, 2, 3);
pushfront_list(l, 0);      // l = list(0, 1, 2, 3)
```

---

## popback_list

`popback_list(l) -> any`
{: .fs-5 .fw-300 }

Removes and returns the last element.

```go
var l = list(1, 2, 3);
var last = popback_list(l); // last = 3, l = list(1, 2)
```

---

## popfront_list

`popfront_list(l) -> any`
{: .fs-5 .fw-300 }

Removes and returns the first element.

```go
var l = list(1, 2, 3);
var first = popfront_list(l); // first = 1, l = list(2, 3)
```

---

## insert_list

`insert_list(l, index, elem) -> list`
{: .fs-5 .fw-300 }

Inserts element at specified index.

```go
var l = list(1, 2, 3);
insert_list(l, 1, 99);     // l = list(1, 99, 2, 3)
```

---

## remove_list

`remove_list(l, index) -> any`
{: .fs-5 .fw-300 }

Removes and returns element at specified index.

```go
var l = list(1, 2, 3);
var removed = remove_list(l, 1); // removed = 2, l = list(1, 3)
```

---

## slice_list

`slice_list(l, start, end) -> list`
{: .fs-5 .fw-300 }

Returns a sublist from start to end index.

```go
var l = list(1, 2, 3, 4, 5);
var sub = slice_list(l, 1, 4); // sub = list(2, 3, 4)
```

---

## contains_list

`contains_list(l, elem) -> bool`
{: .fs-5 .fw-300 }

Checks if list contains the element.

```go
var l = list(1, 2, 3);
contains_list(l, 2);       // true
contains_list(l, 99);      // false
```

---

## index_list

`index_list(l, elem) -> int`
{: .fs-5 .fw-300 }

Returns index of first occurrence of element, or -1 if not found.

```go
var l = list(1, 2, 3, 2);
index_list(l, 2);          // 1
index_list(l, 99);         // -1
```

---

## peekback_list

`peekback_list(l) -> any`
{: .fs-5 .fw-300 }

Returns the last element without removing it.

```go
var l = list(1, 2, 3);
var last = peekback_list(l); // last = 3, l = list(1, 2, 3) unchanged
```

---

## peekfront_list

`peekfront_list(l) -> any`
{: .fs-5 .fw-300 }

Returns the first element without removing it.

```go
var l = list(1, 2, 3);
var first = peekfront_list(l); // first = 1, l = list(1, 2, 3) unchanged
```

---

## map_list

`map_list(l, func) -> list`
{: .fs-5 .fw-300 }

Applies a function to each element and returns a new list.

```go
var l = list(1, 2, 3);
var doubled = map_list(l, func(x) { return x * 2; });
// doubled = list(2, 4, 6)
```

---

## filter_list

`filter_list(l, func) -> list`
{: .fs-5 .fw-300 }

Returns elements that satisfy the predicate function.

```go
var l = list(1, 2, 3, 4, 5, 6);
var evens = filter_list(l, func(x) { return x % 2 == 0; });
// evens = list(2, 4, 6)
```

---

## reduce_list

`reduce_list(l, func, initial) -> any`
{: .fs-5 .fw-300 }

Reduces list to a single value using an accumulator function.

```go
var l = list(1, 2, 3, 4, 5);
var sum = reduce_list(l, func(acc, x) { return acc + x; }, 0);
// sum = 15
```

---

## find_list

`find_list(l, func) -> any`
{: .fs-5 .fw-300 }

Returns the first element matching the predicate, or nil if none found.

```go
var l = list(1, 2, 3, 4, 5);
var firstEven = find_list(l, func(x) { return x % 2 == 0; });
// firstEven = 2
```

---

## some_list

`some_list(l, func) -> bool`
{: .fs-5 .fw-300 }

Returns true if at least one element satisfies the predicate.

```go
var l = list(1, 2, 3, 4, 5);
some_list(l, func(x) { return x > 3; });  // true
some_list(l, func(x) { return x > 10; }); // false
```

---

## every_list

`every_list(l, func) -> bool`
{: .fs-5 .fw-300 }

Returns true if all elements satisfy the predicate.

```go
var l = list(2, 4, 6, 8);
every_list(l, func(x) { return x % 2 == 0; });  // true
every_list(l, func(x) { return x > 5; });       // false
```

---

## to_list

`to_list(iterable) -> list`
{: .fs-5 .fw-300 }

Converts an array or tuple to a list.

```go
var a = [1, 2, 3];
var t = tuple(4, 5, 6);
to_list(a);                // list(1, 2, 3)
to_list(t);                // list(4, 5, 6)
```

---

## size_list

`size_list(l) -> int`
{: .fs-5 .fw-300 }

Returns the number of elements in the list. Alias: length_list.

```go
var l = list(1, 2, 3, 4, 5);
size_list(l);              // 5
length_list(l);            // 5 (alias)
```

---
