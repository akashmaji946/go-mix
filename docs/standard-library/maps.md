---
title: "Maps"
layout: default
parent: Standard Library
nav_order: 5
description: "Dictionary operations including keys, insert, remove, contain, enumerate"
permalink: /standard-library/maps/
---

# Maps Package
{: .no_toc }

Dictionary operations including keys, insert, remove, contain, enumerate
{: .fs-6 .fw-300 }

## Table of Contents
{: .no_toc .text-delta }

1. TOC
{:toc}

---

## Import

`import "maps"`
{: .fs-5 .fw-300 }

Import the maps package to use namespaced functions.

```go
// Standard import
import maps;
var m = maps.make_map("name", "John", "age", 25)
var keys = maps.keys_map(m)

// With alias
import maps as m;
var map = m.make_map("name", "John", "age", 25)
var keys = m.keys_map(map)
```

---

## make_map

`make_map(key1, val1, key2, val2, ...) -> map`
{: .fs-5 .fw-300 }

Creates a new map from key-value pairs.

```go
var m = make_map("name", "John", "age", 25);
// m = map{"name": "John", "age": 25}
```

---

## keys_map

`keys_map(m) -> array`
{: .fs-5 .fw-300 }

Returns array of all keys in the map.

```go
var m = map{"name": "John", "age": 25};
keys_map(m);               // ["name", "age"]
```

---

## values_map

`values_map(m) -> array`
{: .fs-5 .fw-300 }

Returns array of all values in the map.

```go
var m = map{"name": "John", "age": 25};
values_map(m);             // ["John", 25]
```

---

## insert_map

`insert_map(m, key, value) -> any`
{: .fs-5 .fw-300 }

Inserts or updates a key-value pair in the map.

```go
var m = map{"name": "John"};
insert_map(m, "age", 25);  // m = map{"name": "John", "age": 25}
insert_map(m, "name", "Jane");  // Updates existing key
```

---

## remove_map

`remove_map(m, key) -> any`
{: .fs-5 .fw-300 }

Removes a key from the map and returns its value.

```go
var m = map{"name": "John", "age": 25};
var age = remove_map(m, "age");  // age = 25, m = map{"name": "John"}
```

---

## contain_map

`contain_map(m, key) -> bool`
{: .fs-5 .fw-300 }

Checks if map contains the specified key.

```go
var m = map{"name": "John", "age": 25};
contain_map(m, "name");    // true
contain_map(m, "email");   // false
```

---

## enumerate_map

`enumerate_map(m) -> array`
{: .fs-5 .fw-300 }

Returns array of [key, value] pairs for iteration.

```go
var m = map{"name": "John", "age": 25};
enumerate_map(m);          // [["name", "John"], ["age", 25]]
```

---

## size_map

`size_map(m) -> int`
{: .fs-5 .fw-300 }

Returns the number of key-value pairs in the map.

```go
var m = map{"a": 1, "b": 2, "c": 3};
size_map(m);               // 3
```

---
