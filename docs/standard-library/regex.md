---
title: "Regex"
layout: default
parent: Standard Library
nav_order: 14
description: "Regular expression functions for pattern matching"
permalink: /standard-library/regex/
---

# Regex Package
{: .no_toc }

Regular expression functions for pattern matching
{: .fs-6 .fw-300 }

## Table of Contents
{: .no_toc .text-delta }

1. TOC
{:toc}

---

## Import

`import "regex"`
{: .fs-5 .fw-300 }

Import the regex package to use namespaced functions.

```go
// Standard import
import regex;
var match = regex.match_regex("\\d+", "123")
var found = regex.find_regex("\\d+", "abc123")

// With alias
import regex as re;
var match = re.match_regex("\\d+", "123")
var found = re.find_regex("\\d+", "abc123")
```

---

## match_regex

`match_regex(pattern, str) -> bool`
{: .fs-5 .fw-300 }

Checks if string matches pattern.

```go
match_regex("\\d+", "123");        // true
match_regex("^[a-z]+$", "hello");  // true
match_regex("^[a-z]+$", "Hello"); // false (uppercase)
```

---

## find_regex

`find_regex(pattern, str) -> string`
{: .fs-5 .fw-300 }

Returns first match or empty string.

```go
find_regex("\\d+", "abc123def");   // "123"
find_regex("[a-z]+", "ABCdef");    // "def"
```

---

## findall_regex

`findall_regex(pattern, str, [n]) -> array`
{: .fs-5 .fw-300 }

Returns all matches as array.

```go
findall_regex("\\d+", "a1b2c3");   // ["1", "2", "3"]
findall_regex("[a-z]+", "abc def", 1); // ["abc"] (limit 1)
```

---

## replace_regex

`replace_regex(pattern, str, replacement) -> string`
{: .fs-5 .fw-300 }

Replaces all matches with replacement.

```go
replace_regex("\\d", "a1b2c3", "X"); // "aXbXcX"
replace_regex(" +", "hello   world", " "); // "hello world"
```

---

## split_regex

`split_regex(pattern, str, [n]) -> array`
{: .fs-5 .fw-300 }

Splits string by pattern.

```go
split_regex("\\s+", "hello   world"); // ["hello", "world"]
split_regex(",", "a,b,c", 2);         // ["a", "b,c"] (limit 2)
```

---
