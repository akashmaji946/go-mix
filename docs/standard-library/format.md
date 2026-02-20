---
title: "Format"
layout: default
parent: Standard Library
nav_order: 13
description: "Type conversion functions including to_int, to_float, to_string, to_bool"
permalink: /standard-library/format/
---

# Format Package
{: .no_toc }

Type conversion functions including to_int, to_float, to_string, to_bool
{: .fs-6 .fw-300 }

## Table of Contents
{: .no_toc .text-delta }

1. TOC
{:toc}

---

## Import

`import "format"`
{: .fs-5 .fw-300 }

Import the format package to use namespaced functions.

```go
// Standard import
import format;
var num = format.to_int("42");
var str = format.to_str(123);

// With alias
import format as fmt;
var num = fmt.to_int("42")
var str = fmt.to_str(123)
```

---

## to_int

`to_int(value) -> int`
{: .fs-5 .fw-300 }

Converts value to integer.

```go
to_int("42");              // 42
to_int("0xFF");            // 255 (hex)
to_int("0o77");            // 63 (octal)
to_int(3.14);              // 3 (truncated)
to_int(true);              // 1
```

---

## to_float

`to_float(value) -> float`
{: .fs-5 .fw-300 }

Converts value to float.

```go
to_float("3.14");          // 3.14
to_float("2.5e10");        // 25000000000.0
to_float(42);              // 42.0
```

---

## to_string

`to_str(value) -> string`
{: .fs-5 .fw-300 }

Converts value to string.

```go
to_str(42);             // "42"
to_str(3.14);           // "3.14"
to_str(true);           // "true"
to_str([1, 2, 3]);      // "[1, 2, 3]"
```

---

## to_bool

`to_bool(value) -> bool`
{: .fs-5 .fw-300 }

Converts value to boolean.

```go
to_bool(1);                // true
to_bool(0);                // false
to_bool("true");           // true
to_bool("false");          // false
to_bool(nil);              // false
```

---

## to_char

`to_char(value) -> char`
{: .fs-5 .fw-300 }

Converts value to character.

```go
to_char(65);               // 'A'
to_char("hello");          // 'h' (first char)
```

---
