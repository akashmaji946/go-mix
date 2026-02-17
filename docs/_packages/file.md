---
layout: default
title: File I/O Package - Go-Mix
description: File operations including read, write, mkdir, list_dir, and more
---

<div class="content-page">
    <aside class="sidebar">
        <nav class="sidebar-nav">
            <div class="sidebar-title">File Functions</div>
            <ul class="sidebar-menu">
                <li><a href="#read_file">read_file</a></li>
                <li><a href="#write_file">write_file</a></li>
                <li><a href="#append_file">append_file</a></li>
                <li><a href="#file_exists">file_exists</a></li>
                <li><a href="#mkdir">mkdir</a></li>
                <li><a href="#list_dir">list_dir</a></li>
                <li><a href="#remove_file">remove_file</a></li>
                <li><a href="#rename_file">rename_file</a></li>
            </ul>
        </nav>
    </aside>
    
    <div class="content-body">
        <h1>File I/O Package</h1>
        <p>Comprehensive file system operations (17 functions).</p>
        
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
    </div>
</div>
