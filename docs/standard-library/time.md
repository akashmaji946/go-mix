---
title: "Time"
layout: default
parent: Standard Library
nav_order: 9
description: "Time handling functions including now, format_time, parse_time, timezone"
permalink: /standard-library/time/
---

# Time Package
{: .no_toc }

Time handling functions including now, format_time, parse_time, timezone
{: .fs-6 .fw-300 }

## Table of Contents
{: .no_toc .text-delta }

1. TOC
{:toc}

---

## Import

`import "time"`
{: .fs-5 .fw-300 }

Import the time package to use namespaced functions.

```go
// Standard import
import time;
var now = time.now();
var formatted = time.format_time(now, "2006-01-02");

// With alias
import time as t;
var now = t.now();
var formatted = t.format_time(now, "2006-01-02");
```

---

## now

`now() -> int`
{: .fs-5 .fw-300 }

Returns current Unix timestamp in seconds.

```go
var timestamp = now();     // e.g., 1707542400
println(timestamp);
```

---

## now_ms

`now_ms() -> int`
{: .fs-5 .fw-300 }

Returns current Unix timestamp in milliseconds.

```go
var timestamp = now_ms();  // e.g., 1707542400000
println(timestamp);
```

---

## utc_now

`utc_now() -> int`
{: .fs-5 .fw-300 }

Returns current UTC Unix timestamp in seconds.

```go
var utc = utc_now();       // UTC timestamp
println(utc);
```

---

## format_time

`format_time(timestamp, layout) -> string`
{: .fs-5 .fw-300 }

Formats timestamp according to layout string.

```go
var t = now();
format_time(t, "2006-01-02");           // "2024-02-10"
format_time(t, "15:04:05");             // "12:00:00"
format_time(t, "2006-01-02 15:04:05"); // "2024-02-10 12:00:00"
```

---

## parse_time

`parse_time(value, layout) -> int`
{: .fs-5 .fw-300 }

Parses time string and returns Unix timestamp.

```go
var ts = parse_time("2024-02-10", "2006-01-02");
println(ts);               // Unix timestamp
```

---

## timezone

`timezone() -> string`
{: .fs-5 .fw-300 }

Returns current timezone name.

```go
var tz = timezone();       // e.g., "IST", "UTC", "PST"
println(tz);
```

---

## sleep

`sleep(milliseconds) -> nil`
{: .fs-5 .fw-300 }

Pauses execution for specified milliseconds.

```go
println("Starting...");
sleep(1000);               // Sleep 1 second
println("Done!");
```

---
