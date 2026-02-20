---
title: "Path"
layout: default
parent: Standard Library
nav_order: 10
description: "File operations including read, write, mkdir, list_dir, and more"
permalink: /standard-library/path/
---

# Path Package
{: .no_toc }

File operations including read, write, mkdir, list_dir, and more
{: .fs-6 .fw-300 }

## Table of Contents
{: .no_toc .text-delta }

1. TOC
{:toc}

---

## Import

`import "file"`
{: .fs-5 .fw-300 }

Import the file package to use namespaced functions.

```go
// Standard import
import path;
var content = file.read_file("data.txt");
file.write_file("output.txt", "Hello");

// With alias
import path as f;
var content = f.read_file("data.txt");
f.write_file("output.txt", "Hello");
```

---

## read_file

`read_file(path) -> string`
{: .fs-5 .fw-300 }

Reads entire file content as string.

```go
var content = read_file("data.txt");
println(content);

var json = read_file("config.json");
```

---

## write_file

`write_file(path, content) -> nil`
{: .fs-5 .fw-300 }

Writes string content to file (creates or overwrites).

```go
write_file("output.txt", "Hello, World!");
write_file("data.json", '{"name": "John"}');
```

---

## append_file

`append_file(path, content) -> nil`
{: .fs-5 .fw-300 }

Appends content to end of file.

```go
append_file("log.txt", "New entry\n");
append_file("data.csv", "1,2,3\n");
```

---

## file_exists

`file_exists(path) -> bool`
{: .fs-5 .fw-300 }

Checks if file or directory exists.

```go
if (file_exists("config.txt")) {
    var content = read_file("config.txt");
} else {
    println("Config file not found");
}
```

---

## mkdir

`mkdir(path) -> nil`
{: .fs-5 .fw-300 }

Creates directory (recursive if needed).

```go
mkdir("data");             // Create single directory
mkdir("data/2024/jan");    // Create nested directories
```

---

## list_dir

`list_dir(path) -> array`
{: .fs-5 .fw-300 }

Returns array of files and directories in path.

```go
var files = list_dir(".");
foreach file in files {
    println(file);
}
```

---

## remove_file

`remove_file(path) -> nil`
{: .fs-5 .fw-300 }

Removes file or empty directory.

```go
remove_file("temp.txt");
remove_file("empty_dir");
```

---

## rename_file

`rename_file(old_path, new_path) -> nil`
{: .fs-5 .fw-300 }

Renames or moves file.

```go
rename_file("old.txt", "new.txt");
rename_file("file.txt", "backup/file.txt");
```

---

## is_dir

`is_dir(path) -> bool`
{: .fs-5 .fw-300 }

Checks if path is a directory.

```go
is_dir("data");            // true if data is a directory
is_dir("file.txt");        // false
```

---

## is_file

`is_file(path) -> bool`
{: .fs-5 .fw-300 }

Checks if path is a regular file.

```go
is_file("file.txt");       // true if file.txt exists and is a file
is_file("data");           // false if data is a directory
```

---

## touch

`touch(path) -> nil`
{: .fs-5 .fw-300 }

Creates an empty file or updates timestamps if it exists.

```go
touch("newfile.txt");      // Creates empty file
touch("existing.txt");     // Updates access/modification time
```

---

## pwd

`pwd() -> string`
{: .fs-5 .fw-300 }

Returns the current working directory.

```go
var current = pwd();
println("Current dir: " + current);
```

---

## home

`home() -> string`
{: .fs-5 .fw-300 }

Returns the user's home directory.

```go
var userHome = home();
println("Home: " + userHome);
```

---

## truncate_file

`truncate_file(path, size) -> nil`
{: .fs-5 .fw-300 }

Changes the size of a file. Use size 0 to clear file.

```go
truncate_file("data.log", 0);    // Clears the file
truncate_file("data.bin", 1024); // Resize to 1KB
```

---

## remove_all

`remove_all(path) -> nil`
{: .fs-5 .fw-300 }

Removes path and all its children (recursive delete, like rm -rf).

```go
remove_all("temp_dir");    // Removes directory and all contents
remove_all("old_files");   // Use with caution!
```

---

## chmod

`chmod(path, mode) -> nil`
{: .fs-5 .fw-300 }

Changes file permissions (Unix-style octal mode).

```go
chmod("script.gm", 0755);  // Make executable
chmod("file.txt", 0644);   // Owner read/write, others read-only
```

---

## cat

`cat(path, ...) -> nil`
{: .fs-5 .fw-300 }

Prints file contents to output (like Unix cat command).

```go
cat("file.txt");           // Print single file
cat("a.txt", "b.txt");     // Concatenate multiple files
```

---

## path_join

`path_join(elem1, elem2, ...) -> string`
{: .fs-5 .fw-300 }

Joins path elements with OS-specific separator.

```go
path_join("home", "user", "docs");  // "home/user/docs" (Unix)
path_join("C:", "Program Files");   // "C:\\Program Files" (Windows)
```

---

## path_base

`path_base(path) -> string`
{: .fs-5 .fw-300 }

Returns the last element of path (filename).

```go
path_base("/home/user/file.txt");  // "file.txt"
path_base("/home/user/docs/");     // "docs"
```

---

## path_dir

`path_dir(path) -> string`
{: .fs-5 .fw-300 }

Returns all but the last element of path (directory).

```go
path_dir("/home/user/file.txt");   // "/home/user"
path_dir("/home/user/docs/");      // "/home/user"
```

---

## path_ext

`path_ext(path) -> string`
{: .fs-5 .fw-300 }

Returns the file extension.

```go
path_ext("file.txt");      // ".txt"
path_ext("archive.tar.gz"); // ".gz"
path_ext("README");        // ""
```

---

## path_abs

`path_abs(path) -> string`
{: .fs-5 .fw-300 }

Returns absolute representation of path.

```go
path_abs("file.txt");      // "/current/dir/file.txt"
path_abs("../parent");     // "/parent"
```

---

## glob

`glob(pattern) -> array`
{: .fs-5 .fw-300 }

Returns files matching a glob pattern.

```go
glob("*.txt");             // ["a.txt", "b.txt", "notes.txt"]
glob("data/*.csv");        // All CSV files in data directory
```

---

## copy_file

`copy_file(src, dst) -> nil`
{: .fs-5 .fw-300 }

Copies a file from source to destination.

```go
copy_file("original.txt", "backup.txt");
copy_file("data.csv", "archive/data_backup.csv");
```

---
