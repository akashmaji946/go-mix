---
title: "JSON"
layout: default
parent: Standard Library
nav_order: 16
description: "JSON parsing and serialization functions"
permalink: /standard-library/json/
---

# JSON Package
{: .no_toc }

JSON parsing and serialization functions
{: .fs-6 .fw-300 }

## Table of Contents
{: .no_toc .text-delta }

1. TOC
{:toc}

---

## Import

`import "json"`
{: .fs-5 .fw-300 }

Import the json package to use namespaced functions.

```go
// Standard import
import json;
var data = json.json_string_to_map('{"name": "John"}')
var str = json.map_to_json_string(map{"age": 30})

// With alias
import json as j;
var data = j.json_string_to_map('{"name": "John"}')
var str = j.map_to_json_string(map{"age": 30})
```

---

## map_to_json_string

`map_to_json_string(m) -> string`
{: .fs-5 .fw-300 }

Converts map to JSON string.

```go
var user = map{
    "name": "Alice",
    "age": 30,
    "active": true
};
var json = map_to_json_string(user);
// {"name":"Alice","age":30,"active":true}
```

---

## json_string_to_map

`json_string_to_map(json) -> map`
{: .fs-5 .fw-300 }

Parses JSON string into map.

```go
var json = '{"name":"Bob","age":25}';
var data = json_string_to_map(json);
println(data["name"]);     // "Bob"
println(data["age"]);      // 25
```

---
