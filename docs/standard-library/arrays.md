---
title: "Arrays"
layout: default
parent: Standard Library
nav_order: 2
description: "Array manipulation functions including push, pop, sort, map, filter, reduce"
permalink: /standard-library/arrays/
---

# Arrays Package
{: .no_toc }

Array manipulation functions including push, pop, sort, map, filter, reduce
{: .fs-6 .fw-300 }

## Table of Contents
{: .no_toc .text-delta }

1. TOC
{:toc}

---

## Import

`import arrays`
{: .fs-5 .fw-300 }

Import the arrays package to use namespaced functions.

```go
// Standard import
import arrays;
var arr = arrays.make_array(1, 2, 3)
arrays.push_array(arr, 4)

// With alias
import arrays as arr
var a = arr.make_array(1, 2, 3)
arr.push_array(a, 4)
```

---

## make_array

`make_array(...) -> array`
{: .fs-5 .fw-300 }

Creates a new array from arguments or converts an iterable to an array.

```go
make_array()                    // []
make_array(1, 2, 3)             // [1, 2, 3]
make_array([1, 2, 3])           // [1, 2, 3] (copy)
make_array(list(1, 2, 3))       // [1, 2, 3]
make_array(tuple(1, 2, 3))      // [1, 2, 3]
make_array(set{1, 2, 3})        // [1, 2, 3]
make_array(map{"a": 1, "b": 2}) // [1, 2] (values)
```

---

## push

`push(arr, elem) -> array`
{: .fs-5 .fw-300 }

Adds an element to the end of an array.

```go
var arr = [1, 2, 3];
push(arr, 4);              // arr is now [1, 2, 3, 4]
```

---

## pop

`pop(arr) -> any`
{: .fs-5 .fw-300 }

Removes and returns the last element of an array.

```go
var arr = [1, 2, 3, 4];
var last = pop(arr);       // last = 4, arr = [1, 2, 3]
```

---

## shift

`shift(arr) -> any`
{: .fs-5 .fw-300 }

Removes and returns the first element of an array.

```go
var arr = [1, 2, 3, 4];
var first = shift(arr);    // first = 1, arr = [2, 3, 4]
```

---

## unshift

`unshift(arr, elem) -> array`
{: .fs-5 .fw-300 }

Adds an element to the beginning of an array.

```go
var arr = [2, 3, 4];
unshift(arr, 1);           // arr is now [1, 2, 3, 4]
```

---

## sort

`sort(arr, [reverse]) -> array`
{: .fs-5 .fw-300 }

Sorts array elements in-place.

```go
var arr = [3, 1, 4, 1, 5];
sort(arr);                 // arr = [1, 1, 3, 4, 5]
sort(arr, true);            // arr = [5, 4, 3, 1, 1] (reverse)
```

---

## sorted

`sorted(arr, [reverse]) -> array`
{: .fs-5 .fw-300 }

Returns a new sorted array without modifying the original.

```go
var arr = [3, 1, 4, 1, 5];
var s = sorted(arr);       // s = [1, 1, 3, 4, 5], arr unchanged
```

---

## csort

`csort(arr, comparator) -> array`
{: .fs-5 .fw-300 }

Sorts array in-place using a custom comparator function.

```go
var arr = [3, 1, 4, 1, 5];
csort(arr, func(a, b) { return a > b; }); // Descending sort
```

---

## csorted

`csorted(arr, comparator) -> array`
{: .fs-5 .fw-300 }

Returns a new sorted array using a custom comparator function.

```go
var arr = [3, 1, 4, 1, 5];
var sorted = csorted(arr, func(a, b) { return a < b; });
```

---

## map

`map(arr, func) -> array`
{: .fs-5 .fw-300 }

Applies a function to each element and returns new array.

```go
var nums = [1, 2, 3, 4, 5];
var doubled = map(nums, func(x) { return x * 2; });
// doubled = [2, 4, 6, 8, 10]
```

---

## filter

`filter(arr, func) -> array`
{: .fs-5 .fw-300 }

Returns elements that satisfy the predicate function.

```go
var nums = [1, 2, 3, 4, 5, 6];
var evens = filter(nums, func(x) { return x % 2 == 0; });
// evens = [2, 4, 6]
```

---

## reduce

`reduce(arr, func, initial) -> any`
{: .fs-5 .fw-300 }

Reduces array to single value using accumulator function.

```go
var nums = [1, 2, 3, 4, 5];
var sum = reduce(nums, func(acc, x) { return acc + x; }, 0);
// sum = 15
```

---

## find

`find(arr, func) -> any`
{: .fs-5 .fw-300 }

Returns first element matching predicate, or nil if none found.

```go
var nums = [1, 2, 3, 4, 5];
var firstEven = find(nums, func(x) { return x % 2 == 0; });
// firstEven = 2
```

---

## some

`some(arr, func) -> bool`
{: .fs-5 .fw-300 }

Returns true if at least one element satisfies the predicate.

```go
var nums = [1, 2, 3, 4, 5];
some(nums, func(x) { return x > 3; });  // true
some(nums, func(x) { return x > 10; }); // false
```

---

## every

`every(arr, func) -> bool`
{: .fs-5 .fw-300 }

Returns true if all elements satisfy the predicate.

```go
var nums = [2, 4, 6, 8];
every(nums, func(x) { return x % 2 == 0; });  // true
every(nums, func(x) { return x > 5; });       // false
```

---

## clone

`clone(arr) -> array`
{: .fs-5 .fw-300 }

Returns a shallow copy of the array.

```go
var arr = [1, 2, 3];
var copy = clone(arr);     // copy = [1, 2, 3], independent of arr
```

---

## reverse

`reverse(arr) -> array`
{: .fs-5 .fw-300 }

Returns a new array with elements in reverse order.

```go
var arr = [1, 2, 3, 4, 5];
var rev = reverse(arr);    // rev = [5, 4, 3, 2, 1]
```

---

## contains

`contains(arr, value) -> bool`
{: .fs-5 .fw-300 }

Checks if a value exists in the array.

```go
var arr = [1, 2, 3, 4, 5];
contains(arr, 3);          // true
contains(arr, 10);         // false
```

---

## replace

`replace(arr, old_val, new_val) -> int`
{: .fs-5 .fw-300 }

Replaces first occurrence of old_val with new_val. Returns index or -1.

```go
var arr = [1, 2, 3];
replace(arr, 2, 42);       // arr = [1, 42, 3], returns 1
replace(arr, 99, 100);     // returns -1 (not found)
```

---

## index

`index(arr, value) -> int`
{: .fs-5 .fw-300 }

Returns index of first occurrence of value, or -1 if not found.

```go
var arr = [1, 2, 3, 2, 1];
index(arr, 2);             // 1
index(arr, 99);            // -1
```

---

## size

`size(arr) -> int`
{: .fs-5 .fw-300 }

Returns the number of elements in the array. Alias: length.

```go
var arr = [1, 2, 3, 4, 5];
size(arr);                 // 5
length(arr);               // 5 (alias)
```

---

## to_array

`to_array(iterable) -> array`
{: .fs-5 .fw-300 }

Converts a list or tuple to an array.

```go
var l = list(1, 2, 3);
var t = tuple(4, 5, 6);
to_array(l);               // [1, 2, 3]
to_array(t);               // [4, 5, 6]
```

---
