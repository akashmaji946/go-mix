---
title: "Sets"
layout: default
parent: Standard Library
nav_order: 8
description: "Set operations for unique value collections"
permalink: /standard-library/sets/
---

# Sets Package
{: .no_toc }

Set operations for unique value collections
{: .fs-6 .fw-300 }

## Table of Contents
{: .no_toc .text-delta }

1. TOC
{:toc}

---

## Import

`import "sets"`
{: .fs-5 .fw-300 }

Import the sets package to use namespaced functions. Sets are mutable collections that store unique values, automatically removing duplicates.

```go
// Standard import
import sets;
var s = sets.make_set(1, 2, 3)
sets.insert_set(s, 4)

// With alias
import sets as st;
var s = st.make_set(1, 2, 3);
st.insert_set(s, 4);
```

---

## make_set

`make_set(...elements) -> set`
{: .fs-5 .fw-300 }

Creates a new set from the provided elements. Duplicate values are automatically removed. Sets maintain insertion order for values.

```go
make_set();                           // set{}
make_set(1, 2, 3);                    // set{1, 2, 3}
make_set(1, 2, 2, 3, 3, 3);           // set{1, 2, 3} (duplicates removed)
make_set("a", 1, true);               // set{a, 1, true} (heterogeneous values)
```

---

## insert_set

`insert_set(s, elem) -> any`
{: .fs-5 .fw-300 }

Adds an element to the set if it doesn't already exist. Sets only store unique values, so inserting a duplicate has no effect. Returns the inserted element.

```go
var s = make_set(1, 2, 3);
insert_set(s, 4);          // Returns 4, s = set{1, 2, 3, 4}
insert_set(s, 2);          // Returns 2, s unchanged (2 already exists)
insert_set(s, "hello");    // Returns "hello", s = set{1, 2, 3, 4, hello}
```

---

## remove_set

`remove_set(s, elem) -> bool`
{: .fs-5 .fw-300 }

Removes an element from the set. Returns true if the element was found and removed, false if the element didn't exist in the set.

```go
var s = make_set(1, 2, 3);
remove_set(s, 2);          // true, s = set{1, 3}
remove_set(s, 99);         // false (element didn't exist)
remove_set(s, 1);          // true, s = set{3}
remove_set(s, 3);          // true, s = set{}
```

---

## contains_set

`contains_set(s, elem) -> bool`
{: .fs-5 .fw-300 }

Checks if the set contains a specific element. Returns true if the element exists in the set, false otherwise. Uses O(1) lookup time.

```go
var s = make_set(1, 2, 3);
contains_set(s, 2);        // true
contains_set(s, 99);         // false
contains_set(s, "2");      // false (type matters: string vs int)
```

---

## values_set

`values_set(s) -> array`
{: .fs-5 .fw-300 }

Returns an array containing all values in the set, preserving insertion order. Useful for iterating over set contents or converting to other data structures.

```go
var s = make_set(3, 1, 2);
values_set(s);             // [3, 1, 2] (insertion order preserved)

var empty = make_set();
values_set(empty);         // [] (empty array for empty set)
```

---

## size_set

`size_set(s) -> int`
{: .fs-5 .fw-300 }

Returns the number of unique elements in the set. This is an O(1) operation that returns the current cardinality of the set. Alias: length_set.

```go
var s = make_set(1, 2, 3, 3, 3);
size_set(s);               // 3 (duplicates not counted)

var empty = make_set();
size_set(empty);           // 0

// After modifications
insert_set(s, 4);
size_set(s);               // 4
remove_set(s, 1);
size_set(s);               // 3
```

---

## length_set

`length_set(s) -> int`
{: .fs-5 .fw-300 }

Returns the number of unique elements in the set. This is an alias for size_set() and provides the same functionality for those who prefer the length terminology.

```go
var s = make_set(1, 2, 3);
length_set(s);             // 3

// Equivalent to size_set
size_set(s) == length_set(s);  // true

var empty = make_set();
length_set(empty);         // 0
```

---
