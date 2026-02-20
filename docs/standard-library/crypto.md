---
title: "Crypto"
layout: default
parent: Standard Library
nav_order: 17
description: "Cryptographic functions including hashing and random generation"
permalink: /standard-library/crypto/
---

# Crypto Package
{: .no_toc }

Cryptographic functions including hashing and random generation
{: .fs-6 .fw-300 }

## Table of Contents
{: .no_toc .text-delta }

1. TOC
{:toc}

---

## Import

`import "crypto"`
{: .fs-5 .fw-300 }

Import the crypto package to use namespaced functions.

```go
// Standard import
import crypto;
var hash = crypto.md5("hello");
var uuid = crypto.uuid();

// With alias
import "crypto" as c
var hash = c.md5("hello")
var uuid = c.uuid()
```

---

## md5

`md5(str) -> string`
{: .fs-5 .fw-300 }

Returns MD5 hash of string.

```go
var hash = md5("hello");
// 5d41402abc4b2a76b9719d911017c592
```

---

## sha1

`sha1(str) -> string`
{: .fs-5 .fw-300 }

Returns SHA1 hash of string.

```go
var hash = sha1("hello");
// aaf4c61ddcc5e8a2dabede0f3b482cd9aea9434d
```

---

## sha256

`sha256(str) -> string`
{: .fs-5 .fw-300 }

Returns SHA256 hash of string.

```go
var hash = sha256("hello");
// 2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824
```

---

## base64_encode

`base64_encode(str) -> string`
{: .fs-5 .fw-300 }

Encodes string to Base64.

```go
var encoded = base64_encode("hello");
// aGVsbG8=
```

---

## base64_decode

`base64_decode(str) -> string`
{: .fs-5 .fw-300 }

Decodes Base64 string.

```go
var decoded = base64_decode("aGVsbG8=");
// hello
```

---

## uuid

`uuid() -> string`
{: .fs-5 .fw-300 }

Generates random UUID v4.

```go
var id = uuid();
// 550e8400-e29b-41d4-a716-446655440000
```

---

## random

`random() -> float`
{: .fs-5 .fw-300 }

Returns random float between 0 and 1.

```go
var r = random();          // 0.0 to 1.0
var scaled = random() * 100;  // 0.0 to 100.0
```

---
