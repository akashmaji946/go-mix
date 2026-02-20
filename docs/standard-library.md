---
title: Standard Library
layout: default
nav_order: 4
description: "Overview of Go-Mix standard library packages with 100+ built-in functions"
permalink: /standard-library/
has_children: true
---

# Standard Library
{: .no_toc }

Go-Mix provides a comprehensive standard library with **100+ built-in functions** organized into 17 packages.
{: .fs-6 .fw-300 }

## Table of Contents
{: .no_toc .text-delta }

1. TOC
{:toc}

---

{: .note }
> All builtin functions are available globally without explicit import. You can also import functions as packages for better organization: `import math;` or `import strings;`

## Package Overview

| Package | Description |
|:--------|:------------|
| [**Common**]({{ site.baseurl }}/standard-library/common/) | Core functions: print, length, typeof, range, and type constructors |
| [**Arrays**]({{ site.baseurl }}/standard-library/arrays/) | Array manipulation: push, pop, sort, map, filter, reduce, and more |
| [**Strings**]({{ site.baseurl }}/standard-library/strings/) | String operations: upper, lower, split, join, 21 functions total |
| [**Math**]({{ site.baseurl }}/standard-library/math/) | Mathematical functions: abs, sin, cos, pow, random, 20 functions |
| [**Maps**]({{ site.baseurl }}/standard-library/maps/) | Dictionary operations: keys, insert, remove, contain, enumerate |
| [**Lists**]({{ site.baseurl }}/standard-library/lists/) | Mutable heterogeneous sequences: pushback, popfront, insert, remove |
| [**Tuples**]({{ site.baseurl }}/standard-library/tuples/) | Immutable fixed-size sequences with functional operations |
| [**Sets**]({{ site.baseurl }}/standard-library/sets/) | Unique value collections: insert, remove, contains, values |
| [**Time**]({{ site.baseurl }}/standard-library/time/) | Time handling: now, format_time, parse_time, timezone |
| [**Path**]({{ site.baseurl }}/standard-library/path/) | File operations: read_file, write_file, mkdir, list_dir, 17 functions |
| [**I/O**]({{ site.baseurl }}/standard-library/io/) | Input/output: scanln, scanf, input, getchar, sprintf |
| [**OS**]({{ site.baseurl }}/standard-library/os/) | System operations: getenv, exec, sleep, getpid, hostname |
| [**Format**]({{ site.baseurl }}/standard-library/format/) | Type conversion: to_int, to_float, to_bool, to_string, to_char |
| [**Regex**]({{ site.baseurl }}/standard-library/regex/) | Pattern matching: match_regex, find_regex, replace_regex, split_regex |
| [**HTTP**]({{ site.baseurl }}/standard-library/http/) | Web client/server: get_http, post_http, create_server, serve_static |
| [**JSON**]({{ site.baseurl }}/standard-library/json/) | JSON handling: map_to_json_string, json_string_to_map |
| [**Crypto**]({{ site.baseurl }}/standard-library/crypto/) | Cryptography: md5, sha1, sha256, base64, uuid, random |

---

## Quick Reference

### Common Functions

| Function | Description | Example |
|:---------|:------------|:--------|
| `print(...)` | Output without newline | `print("Hello")` |
| `println(...)` | Output with newline | `println("World")` |
| `printf(fmt, ...)` | Formatted output | `printf("Value: %d", 42)` |
| `length(obj)` | Length of collection | `length("hello") // 5` |
| `typeof(obj)` | Get type name | `typeof(42) // "int"` |
| `range(start, end)` | Create inclusive range | `range(1, 5)` |

### Import Syntax

```go
// Import packages
import arrays;
import strings;
import math;

// Use with package prefix (optional)
var arr = [3, 1, 4, 1, 5];
arrays.sort_array(arr);
var upper = strings.upper_string("hello");
var sqrt = math.sqrt(16);
```

{: .tip }
> Click on any package in the table above to view detailed documentation for all functions in that package, including syntax, parameters, return values, and examples.
