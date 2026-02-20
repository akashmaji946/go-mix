---
title: "HTTP"
layout: default
parent: Standard Library
nav_order: 15
description: "HTTP client and server functions for web requests and APIs"
permalink: /standard-library/http/
---

# HTTP Package
{: .no_toc }

HTTP client and server functions for web requests and APIs
{: .fs-6 .fw-300 }

## Table of Contents
{: .no_toc .text-delta }

1. TOC
{:toc}

---

## Import

`import "http"`
{: .fs-5 .fw-300 }

Import the http package to use namespaced functions.

```go
// Standard import
import http;
var response = http.get_http("https://api.example.com")

// With alias
import http as h;
var response = h.get_http("https://api.example.com")
```

---

## get_http

`get_http(url) -> string`
{: .fs-5 .fw-300 }

Performs HTTP GET request.

```go
var response = get_http("https://api.example.com/data");
println(response);
```

---

## post_http

`post_http(url, body) -> string`
{: .fs-5 .fw-300 }

Performs HTTP POST request.

```go
var body = '{"name": "John", "age": 30}';
var response = post_http("https://api.example.com/users", body);
println(response);
```

---

## request_http

`request_http(method, url, [headers], [body]) -> map`
{: .fs-5 .fw-300 }

Generic HTTP request with full control.

```go
var headers = map{
    "Content-Type": "application/json",
    "Authorization": "Bearer token123"
};
var response = request_http("PUT", "https://api.example.com/data", headers, '{"key": "value"}');
println(response["status"]);
println(response["body"]);
```

---

## create_server

`create_server() -> server`
{: .fs-5 .fw-300 }

Creates new HTTP server instance.

```go
var srv = create_server();
```

---

## handle_server

`handle_server(server, path, handler) -> nil`
{: .fs-5 .fw-300 }

Registers route handler for server.

```go
handle_server(srv, "/", func(req) {
    return "Hello, World!";
});

handle_server(srv, "/api/users", func(req) {
    return '{"users": ["Alice", "Bob"]}';
});
```

---

## start_server

`start_server(server, address) -> nil`
{: .fs-5 .fw-300 }

Starts server on specified address.

```go
start_server(srv, ":8080");
println("Server running on http://localhost:8080");
```

---
