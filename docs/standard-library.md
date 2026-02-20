---
layout: default
title: Standard Library - Go-Mix
description: Overview of Go-Mix standard library packages with 100+ built-in functions
permalink: /standard-library/
---

<div class="content-page">
    <aside class="sidebar">
        <nav class="sidebar-nav">
            <div class="sidebar-title">Packages</div>
            <ul class="sidebar-menu">
                <li><a href="{{ '/packages/common/' | relative_url }}">Common</a></li>
                <li><a href="{{ '/packages/arrays/' | relative_url }}">Arrays</a></li>
                <li><a href="{{ '/packages/strings/' | relative_url }}">Strings</a></li>
                <li><a href="{{ '/packages/math/' | relative_url }}">Math</a></li>
                <li><a href="{{ '/packages/maps/' | relative_url }}">Maps</a></li>
                <li><a href="{{ '/packages/lists/' | relative_url }}">Lists</a></li>
                <li><a href="{{ '/packages/tuples/' | relative_url }}">Tuples</a></li>
                <li><a href="{{ '/packages/sets/' | relative_url }}">Sets</a></li>
                <li><a href="{{ '/packages/time/' | relative_url }}">Time</a></li>
                <li><a href="{{ '/packages/file/' | relative_url }}">Path</a></li>
                <li><a href="{{ '/packages/io/' | relative_url }}">I/O</a></li>
                <li><a href="{{ '/packages/os/' | relative_url }}">OS</a></li>
                <li><a href="{{ '/packages/format/' | relative_url }}">Format</a></li>
                <li><a href="{{ '/packages/regex/' | relative_url }}">Regex</a></li>
                <li><a href="{{ '/packages/http/' | relative_url }}">HTTP</a></li>
                <li><a href="{{ '/packages/json/' | relative_url }}">JSON</a></li>
                <li><a href="{{ '/packages/crypto/' | relative_url }}">Crypto</a></li>
            </ul>
        </nav>
    </aside>
    <div class="content-body">
        <h1>Standard Library</h1>
        
        <p>Go-Mix provides a comprehensive standard library with <strong>100+ built-in functions</strong> organized into 17 packages. All functions are available globally and can also be imported as packages.</p>
        
        <div class="callout callout-info">
            <div class="callout-title">
                <i class="fas fa-info-circle"></i> Global Availability
            </div>
            <p>All builtin functions are available globally without explicit import. However, you can also import functions as packages for better organization:</p>
            <pre><code>import math;
import strings;</code></pre>
        </div>
        
        <h2>Package Overview</h2>
        
        <div class="packages-grid">
            <a href="{{ '/packages/common/' | relative_url }}" class="package-card">
                <div class="package-icon">
                    <i class="fas fa-cube"></i>
                </div>
                <h3>Common</h3>
                <p>Core functions: print, length, typeof, range, and type constructors.</p>
            </a>
            
            <a href="{{ '/packages/arrays/' | relative_url }}" class="package-card">
                <div class="package-icon">
                    <i class="fas fa-list-ol"></i>
                </div>
                <h3>Arrays</h3>
                <p>Array manipulation: push, pop, sort, map, filter, reduce, and more.</p>
            </a>
            
            <a href="{{ '/packages/strings/' | relative_url }}" class="package-card">
                <div class="package-icon">
                    <i class="fas fa-font"></i>
                </div>
                <h3>Strings</h3>
                <p>String operations: upper, lower, split, join, 21 functions total.</p>
            </a>
            
            <a href="{{ '/packages/math/' | relative_url }}" class="package-card">
                <div class="package-icon">
                    <i class="fas fa-calculator"></i>
                </div>
                <h3>Math</h3>
                <p>Mathematical functions: abs, sin, cos, pow, random, 20 functions.</p>
            </a>
            
            <a href="{{ '/packages/maps/' | relative_url }}" class="package-card">
                <div class="package-icon">
                    <i class="fas fa-map"></i>
                </div>
                <h3>Maps</h3>
                <p>Dictionary operations: keys, insert, remove, contain, enumerate.</p>
            </a>
            
            <a href="{{ '/packages/lists/' | relative_url }}" class="package-card">
                <div class="package-icon">
                    <i class="fas fa-list"></i>
                </div>
                <h3>Lists</h3>
                <p>Mutable heterogeneous sequences: pushback, popfront, insert, remove.</p>
            </a>
            
            <a href="{{ '/packages/tuples/' | relative_url }}" class="package-card">
                <div class="package-icon">
                    <i class="fas fa-stream"></i>
                </div>
                <h3>Tuples</h3>
                <p>Immutable fixed-size sequences with functional operations.</p>
            </a>
            
            <a href="{{ '/packages/sets/' | relative_url }}" class="package-card">
                <div class="package-icon">
                    <i class="fas fa-object-group"></i>
                </div>
                <h3>Sets</h3>
                <p>Unique value collections: insert, remove, contains, values.</p>
            </a>
            
            <a href="{{ '/packages/time/' | relative_url }}" class="package-card">
                <div class="package-icon">
                    <i class="fas fa-clock"></i>
                </div>
                <h3>Time</h3>
                <p>Time handling: now, format_time, parse_time, timezone.</p>
            </a>
            
            <a href="{{ '/packages/file/' | relative_url }}" class="package-card">
                <div class="package-icon">
                    <i class="fas fa-file-alt"></i>
                </div>
                <h3>Path</h3>
                <p>File operations: read_file, write_file, mkdir, list_dir, 17 functions.</p>
            </a>
            
            <a href="{{ '/packages/io/' | relative_url }}" class="package-card">
                <div class="package-icon">
                    <i class="fas fa-terminal"></i>
                </div>
                <h3>I/O</h3>
                <p>Input/output: scanln, scanf, input, getchar, sprintf.</p>
            </a>
            
            <a href="{{ '/packages/os/' | relative_url }}" class="package-card">
                <div class="package-icon">
                    <i class="fas fa-desktop"></i>
                </div>
                <h3>OS</h3>
                <p>System operations: getenv, exec, sleep, getpid, hostname.</p>
            </a>
            
            <a href="{{ '/packages/format/' | relative_url }}" class="package-card">
                <div class="package-icon">
                    <i class="fas fa-exchange-alt"></i>
                </div>
                <h3>Format</h3>
                <p>Type conversion: to_int, to_float, to_bool, to_string, to_char.</p>
            </a>
            
            <a href="{{ '/packages/regex/' | relative_url }}" class="package-card">
                <div class="package-icon">
                    <i class="fas fa-search"></i>
                </div>
                <h3>Regex</h3>
                <p>Pattern matching: match_regex, find_regex, replace_regex, split_regex.</p>
            </a>
            
            <a href="{{ '/packages/http/' | relative_url }}" class="package-card">
                <div class="package-icon">
                    <i class="fas fa-globe"></i>
                </div>
                <h3>HTTP</h3>
                <p>Web client/server: get_http, post_http, create_server, serve_static.</p>
            </a>
            
            <a href="{{ '/packages/json/' | relative_url }}" class="package-card">
                <div class="package-icon">
                    <i class="fas fa-code"></i>
                </div>
                <h3>JSON</h3>
                <p>JSON handling: map_to_json_string, json_string_to_map.</p>
            </a>
            
            <a href="{{ '/packages/crypto/' | relative_url }}" class="package-card">
                <div class="package-icon">
                    <i class="fas fa-lock"></i>
                </div>
                <h3>Crypto</h3>
                <p>Cryptography: md5, sha1, sha256, base64, uuid, random.</p>
            </a>
        </div>
        
        <h2>Quick Reference</h2>
        
        <h3>Common Functions</h3>
        
        <table>
            <thead>
                <tr>
                    <th>Function</th>
                    <th>Description</th>
                    <th>Example</th>
                </tr>
            </thead>
            <tbody>
                <tr>
                    <td><code>print(...)</code></td>
                    <td>Output without newline</td>
                    <td><code>print("Hello")</code></td>
                </tr>
                <tr>
                    <td><code>println(...)</code></td>
                    <td>Output with newline</td>
                    <td><code>println("World")</code></td>
                </tr>
                <tr>
                    <td><code>printf(fmt, ...)</code></td>
                    <td>Formatted output</td>
                    <td><code>printf("Value: %d", 42)</code></td>
                </tr>
                <tr>
                    <td><code>length(obj)</code></td>
                    <td>Length of collection</td>
                    <td><code>length("hello") // 5</code></td>
                </tr>
                <tr>
                    <td><code>typeof(obj)</code></td>
                    <td>Get type name</td>
                    <td><code>typeof(42) // "int"</code></td>
                </tr>
                <tr>
                    <td><code>range(start, end)</code></td>
                    <td>Create inclusive range</td>
                    <td><code>range(1, 5)</code></td>
                </tr>
            </tbody>
        </table>
        
        <h3>Import Syntax</h3>
        
        <div class="code-block">
            <div class="code-header">
                <span class="code-dot red"></span>
                <span class="code-dot yellow"></span>
                <span class="code-dot green"></span>
                <span class="code-title">imports.gm</span>
            </div>
{% highlight go %}
// Import packages
import arrays;
import strings;
import math;

// Use with package prefix (optional)
var arr = [3, 1, 4, 1, 5];
arrays.sort_array(arr);
var upper = strings.upper_string("hello");
var sqrt = math.sqrt(16);
{% endhighlight %}
</div>
        <div class="callout callout-success">
            <div class="callout-title">
                <i class="fas fa-book-open"></i> Explore Packages
            </div>
            <p>Click on any package card above to view detailed documentation for all functions in that package, including syntax, parameters, return values, and examples.</p>
        </div>
    </div>
</div>
