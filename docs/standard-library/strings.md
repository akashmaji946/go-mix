---
title: "Strings"
layout: default
parent: Standard Library
nav_order: 3
description: "String manipulation functions including split, join, replace, and regex"
permalink: /standard-library/strings/
---

# Strings Package
{: .no_toc }

String manipulation functions including split, join, replace, and regex
{: .fs-6 .fw-300 }

## Table of Contents
{: .no_toc .text-delta }

1. TOC
{:toc}

---

## Import

`import "strings"`
{: .fs-5 .fw-300 }

Import the strings package to use namespaced functions.

```go
// Standard import
import strings;
var upper = strings.upper("hello")
var trimmed = strings.trim("  hello  ")

// With alias
import strings as str;
var upper = str.upper("hello")
var trimmed = str.trim("  hello  ")
```

---

## upper

`upper(str) -> string`
{: .fs-5 .fw-300 }

Converts string to uppercase.

```go
println(upper("hello")); // "HELLO"
```

---

## lower

`lower(str) -> string`
{: .fs-5 .fw-300 }

Converts string to lowercase.

```go
println(lower("HELLO")); // "hello"
```

---

## trim

`trim(str) -> string`
{: .fs-5 .fw-300 }

Removes leading and trailing whitespace.

```go
println(trim("  hello  ")); // "hello"
```

---

## ltrim

`ltrim(str) -> string`
{: .fs-5 .fw-300 }

Removes leading whitespace from the left side.

```go
println(ltrim("  hello  ")); // "hello  "
```

---

## rtrim

`rtrim(str) -> string`
{: .fs-5 .fw-300 }

Removes trailing whitespace from the right side.

```go
println(rtrim("  hello  ")); // "  hello"
```

---

## split

`split(str, sep) -> array`
{: .fs-5 .fw-300 }

Splits string by separator into an array.

```go
var parts = split("a,b,c", ","); 
// ["a", "b", "c"]
```

---

## join

`join(arr, sep) -> string`
{: .fs-5 .fw-300 }

Joins array elements with separator.

```go
var s = join(["a", "b", "c"], "-"); 
// "a-b-c"
```

---

## contains_string

`contains_string(str, sub) -> bool`
{: .fs-5 .fw-300 }

Checks if string contains substring.

```go
println(contains_string("hello world", "world")); // true
```

---

## starts_with

`starts_with(str, prefix) -> bool`
{: .fs-5 .fw-300 }

Checks if string starts with prefix.

```go
println(starts_with("go-mix", "go")); // true
```

---

## ends_with

`ends_with(str, suffix) -> bool`
{: .fs-5 .fw-300 }

Checks if string ends with suffix.

```go
println(ends_with("file.gm", ".gm")); // true
```

---

## index_string

`index_string(str, sub) -> int`
{: .fs-5 .fw-300 }

Finds first occurrence of substring (-1 if not found).

```go
println(index_string("hello", "l")); // 2
```

---

## substring

`substring(str, start, length) -> string`
{: .fs-5 .fw-300 }

Extracts substring starting at index with given length.

```go
println(substring("hello world", 0, 5)); // "hello"
```

---

## replace_string

`replace_string(str, old, new) -> string`
{: .fs-5 .fw-300 }

Replaces all occurrences of old substring with new.

```go
println(replace_string("hello world", "world", "go-mix")); 
// "hello go-mix"
```

---

## reverse_string

`reverse_string(str) -> string`
{: .fs-5 .fw-300 }

Reverses the characters in the string.

```go
println(reverse_string("abc")); // "cba"
```

---

## count

`count(str, sub) -> int`
{: .fs-5 .fw-300 }

Counts occurrences of substring.

```go
println(count("banana", "a")); // 3
```

---

## ord

`ord(char_or_string) -> int`
{: .fs-5 .fw-300 }

Returns the integer Unicode code point of a character.

```go
println(ord('A'));   // 65
println(ord("ABC")); // 65 (first character)
```

---

## chr

`chr(integer) -> char`
{: .fs-5 .fw-300 }

Returns a character from its integer Unicode code point.

```go
println(chr(65)); // 'A'
```

---

## strcmp

`strcmp(s1, s2) -> int`
{: .fs-5 .fw-300 }

Compares two strings lexicographically. Returns -1 if s1  s2.

```go
println(strcmp("apple", "banana")); // -1
println(strcmp("hello", "hello"));  // 0
println(strcmp("zoo", "zebra"));    // 1
```

---

## capitalize

`capitalize(str) -> string`
{: .fs-5 .fw-300 }

Capitalizes the first character and lowercases the rest.

```go
println(capitalize("gOMIX")); // "Gomix"
```

---

## repeat

`repeat(str, count) -> string`
{: .fs-5 .fw-300 }

Repeats a string n times.

```go
println(repeat("na", 2)); // "nana"
```

---

## is_digit

`is_digit(str) -> bool`
{: .fs-5 .fw-300 }

Checks if string contains only digits.

```go
println(is_digit("123")); // true
println(is_digit("12a")); // false
```

---

## is_alpha

`is_alpha(str) -> bool`
{: .fs-5 .fw-300 }

Checks if string contains only alphabetic characters.

```go
println(is_alpha("abc")); // true
println(is_alpha("a12")); // false
```

---

## size_string

`size_string(str) -> int`
{: .fs-5 .fw-300 }

Returns the length of a string (number of characters).

```go
println(size_string("hello")); // 5
```

---

## length_string

`length_string(str) -> int`
{: .fs-5 .fw-300 }

Returns the length of a string (alias for size_string).

```go
println(length_string("hello")); // 5
```

---
