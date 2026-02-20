---
layout: default
title: Path Package - Go-Mix
description: File operations including read, write, mkdir, list_dir, and more
---

<div class="content-page">
    <aside class="sidebar">
        <nav class="sidebar-nav">
            <div class="sidebar-title">File Functions</div>
            <ul class="sidebar-menu">
                <li><a href="#import">Import</a></li>
                <li><a href="#read_file">read_file()</a></li>
                <li><a href="#write_file">write_file()</a></li>
                <li><a href="#append_file">append_file()</a></li>
                <li><a href="#file_exists">file_exists()</a></li>
                <li><a href="#mkdir">mkdir()</a></li>
                <li><a href="#list_dir">list_dir()</a></li>
                <li><a href="#remove_file">remove_file()</a></li>
                <li><a href="#rename_file">rename_file()</a></li>
                <li><a href="#is_dir">is_dir()</a></li>
                <li><a href="#is_file">is_file()</a></li>
                <li><a href="#touch">touch()</a></li>
                <li><a href="#pwd">pwd()</a></li>
                <li><a href="#home">home()</a></li>
                <li><a href="#truncate_file">truncate_file()</a></li>
                <li><a href="#remove_all">remove_all()</a></li>
                <li><a href="#chmod">chmod()</a></li>
                <li><a href="#cat">cat()</a></li>
                <li><a href="#path_join">path_join()</a></li>
                <li><a href="#path_base">path_base()</a></li>
                <li><a href="#path_dir">path_dir()</a></li>
                <li><a href="#path_ext">path_ext()</a></li>
                <li><a href="#path_abs">path_abs()</a></li>
                <li><a href="#glob">glob()</a></li>
                <li><a href="#copy_file">copy_file()</a></li>
            </ul>
        </nav>
    </aside>
    
    <div class="content-body">
        <h1>Path Package</h1>
        <p>Comprehensive file system operations (17 functions).</p>
        
        <div class="function-card" id="import">
            <div class="function-header">
                <div class="function-name">Import</div>
                <div class="function-signature">import "file"</div>
            </div>
            <div class="function-body">
                <div class="function-description">Import the file package to use namespaced functions.</div>
                <div class="function-example">
                    <h4>Examples</h4>
{% highlight go %}
// Standard import
import path;
var content = file.read_file("data.txt");
file.write_file("output.txt", "Hello");

// With alias
import path as f;
var content = f.read_file("data.txt");
f.write_file("output.txt", "Hello");
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="read_file">
            <div class="function-header">
                <div class="function-name">read_file</div>
                <div class="function-signature">read_file(path) -> string</div>
            </div>
            <div class="function-body">
                <div class="function-description">Reads entire file content as string.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var content = read_file("data.txt");
println(content);

var json = read_file("config.json");
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="write_file">
            <div class="function-header">
                <div class="function-name">write_file</div>
                <div class="function-signature">write_file(path, content) -> nil</div>
            </div>
            <div class="function-body">
                <div class="function-description">Writes string content to file (creates or overwrites).</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
write_file("output.txt", "Hello, World!");
write_file("data.json", '{"name": "John"}');
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="append_file">
            <div class="function-header">
                <div class="function-name">append_file</div>
                <div class="function-signature">append_file(path, content) -> nil</div>
            </div>
            <div class="function-body">
                <div class="function-description">Appends content to end of file.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
append_file("log.txt", "New entry\n");
append_file("data.csv", "1,2,3\n");
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="file_exists">
            <div class="function-header">
                <div class="function-name">file_exists</div>
                <div class="function-signature">file_exists(path) -> bool</div>
            </div>
            <div class="function-body">
                <div class="function-description">Checks if file or directory exists.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
if (file_exists("config.txt")) {
    var content = read_file("config.txt");
} else {
    println("Config file not found");
}
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="mkdir">
            <div class="function-header">
                <div class="function-name">mkdir</div>
                <div class="function-signature">mkdir(path) -> nil</div>
            </div>
            <div class="function-body">
                <div class="function-description">Creates directory (recursive if needed).</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
mkdir("data");             // Create single directory
mkdir("data/2024/jan");    // Create nested directories
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="list_dir">
            <div class="function-header">
                <div class="function-name">list_dir</div>
                <div class="function-signature">list_dir(path) -> array</div>
            </div>
            <div class="function-body">
                <div class="function-description">Returns array of files and directories in path.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var files = list_dir(".");
foreach file in files {
    println(file);
}
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="remove_file">
            <div class="function-header">
                <div class="function-name">remove_file</div>
                <div class="function-signature">remove_file(path) -> nil</div>
            </div>
            <div class="function-body">
                <div class="function-description">Removes file or empty directory.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
remove_file("temp.txt");
remove_file("empty_dir");
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="rename_file">
            <div class="function-header">
                <div class="function-name">rename_file</div>
                <div class="function-signature">rename_file(old_path, new_path) -> nil</div>
            </div>
            <div class="function-body">
                <div class="function-description">Renames or moves file.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
rename_file("old.txt", "new.txt");
rename_file("file.txt", "backup/file.txt");
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="is_dir">
            <div class="function-header">
                <div class="function-name">is_dir</div>
                <div class="function-signature">is_dir(path) -> bool</div>
            </div>
            <div class="function-body">
                <div class="function-description">Checks if path is a directory.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
is_dir("data");            // true if data is a directory
is_dir("file.txt");        // false
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="is_file">
            <div class="function-header">
                <div class="function-name">is_file</div>
                <div class="function-signature">is_file(path) -> bool</div>
            </div>
            <div class="function-body">
                <div class="function-description">Checks if path is a regular file.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
is_file("file.txt");       // true if file.txt exists and is a file
is_file("data");           // false if data is a directory
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="touch">
            <div class="function-header">
                <div class="function-name">touch</div>
                <div class="function-signature">touch(path) -> nil</div>
            </div>
            <div class="function-body">
                <div class="function-description">Creates an empty file or updates timestamps if it exists.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
touch("newfile.txt");      // Creates empty file
touch("existing.txt");     // Updates access/modification time
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="pwd">
            <div class="function-header">
                <div class="function-name">pwd</div>
                <div class="function-signature">pwd() -> string</div>
            </div>
            <div class="function-body">
                <div class="function-description">Returns the current working directory.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var current = pwd();
println("Current dir: " + current);
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="home">
            <div class="function-header">
                <div class="function-name">home</div>
                <div class="function-signature">home() -> string</div>
            </div>
            <div class="function-body">
                <div class="function-description">Returns the user's home directory.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var userHome = home();
println("Home: " + userHome);
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="truncate_file">
            <div class="function-header">
                <div class="function-name">truncate_file</div>
                <div class="function-signature">truncate_file(path, size) -> nil</div>
            </div>
            <div class="function-body">
                <div class="function-description">Changes the size of a file. Use size 0 to clear file.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
truncate_file("data.log", 0);    // Clears the file
truncate_file("data.bin", 1024); // Resize to 1KB
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="remove_all">
            <div class="function-header">
                <div class="function-name">remove_all</div>
                <div class="function-signature">remove_all(path) -> nil</div>
            </div>
            <div class="function-body">
                <div class="function-description">Removes path and all its children (recursive delete, like rm -rf).</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
remove_all("temp_dir");    // Removes directory and all contents
remove_all("old_files");   // Use with caution!
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="chmod">
            <div class="function-header">
                <div class="function-name">chmod</div>
                <div class="function-signature">chmod(path, mode) -> nil</div>
            </div>
            <div class="function-body">
                <div class="function-description">Changes file permissions (Unix-style octal mode).</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
chmod("script.gm", 0755);  // Make executable
chmod("file.txt", 0644);   // Owner read/write, others read-only
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="cat">
            <div class="function-header">
                <div class="function-name">cat</div>
                <div class="function-signature">cat(path, ...) -> nil</div>
            </div>
            <div class="function-body">
                <div class="function-description">Prints file contents to output (like Unix cat command).</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
cat("file.txt");           // Print single file
cat("a.txt", "b.txt");     // Concatenate multiple files
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="path_join">
            <div class="function-header">
                <div class="function-name">path_join</div>
                <div class="function-signature">path_join(elem1, elem2, ...) -> string</div>
            </div>
            <div class="function-body">
                <div class="function-description">Joins path elements with OS-specific separator.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
path_join("home", "user", "docs");  // "home/user/docs" (Unix)
path_join("C:", "Program Files");   // "C:\\Program Files" (Windows)
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="path_base">
            <div class="function-header">
                <div class="function-name">path_base</div>
                <div class="function-signature">path_base(path) -> string</div>
            </div>
            <div class="function-body">
                <div class="function-description">Returns the last element of path (filename).</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
path_base("/home/user/file.txt");  // "file.txt"
path_base("/home/user/docs/");     // "docs"
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="path_dir">
            <div class="function-header">
                <div class="function-name">path_dir</div>
                <div class="function-signature">path_dir(path) -> string</div>
            </div>
            <div class="function-body">
                <div class="function-description">Returns all but the last element of path (directory).</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
path_dir("/home/user/file.txt");   // "/home/user"
path_dir("/home/user/docs/");      // "/home/user"
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="path_ext">
            <div class="function-header">
                <div class="function-name">path_ext</div>
                <div class="function-signature">path_ext(path) -> string</div>
            </div>
            <div class="function-body">
                <div class="function-description">Returns the file extension.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
path_ext("file.txt");      // ".txt"
path_ext("archive.tar.gz"); // ".gz"
path_ext("README");        // ""
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="path_abs">
            <div class="function-header">
                <div class="function-name">path_abs</div>
                <div class="function-signature">path_abs(path) -> string</div>
            </div>
            <div class="function-body">
                <div class="function-description">Returns absolute representation of path.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
path_abs("file.txt");      // "/current/dir/file.txt"
path_abs("../parent");     // "/parent"
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="glob">
            <div class="function-header">
                <div class="function-name">glob</div>
                <div class="function-signature">glob(pattern) -> array</div>
            </div>
            <div class="function-body">
                <div class="function-description">Returns files matching a glob pattern.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
glob("*.txt");             // ["a.txt", "b.txt", "notes.txt"]
glob("data/*.csv");        // All CSV files in data directory
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="copy_file">
            <div class="function-header">
                <div class="function-name">copy_file</div>
                <div class="function-signature">copy_file(src, dst) -> nil</div>
            </div>
            <div class="function-body">
                <div class="function-description">Copies a file from source to destination.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
copy_file("original.txt", "backup.txt");
copy_file("data.csv", "archive/data_backup.csv");
{% endhighlight %}
                </div>
            </div>
        </div>
    </div>
</div>
