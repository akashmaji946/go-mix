---
layout: default
title: Getting Started - Go-Mix
description: Installation guide and first steps with Go-Mix programming language
permalink: /getting-started/
---

<div class="content-page">
    <aside class="sidebar">
        <nav class="sidebar-nav">
            <div class="sidebar-title">On This Page</div>
            <ul class="sidebar-menu">
                <li><a href="#installation">Installation</a></li>
                <li><a href="#quick-start">Quick Start</a></li>
                <li><a href="#first-program">Your First Program</a></li>
                <li><a href="#repl">Using the REPL</a></li>
                <li><a href="#next-steps">Next Steps</a></li>
            </ul>
        </nav>
    </aside>
    
    <div class="content-body">
        <h1>Getting Started</h1>
        
        <p>Welcome to Go-Mix! This guide will help you install the language and write your first program in minutes.</p>
        
        <h2 id="installation">Installation</h2>
        
        <h3>Prerequisites</h3>
        <ul>
            <li><strong>Go 1.18 or higher</strong> - Required for building from source</li>
            <li><strong>Git</strong> - For cloning the repository</li>
        </ul>
        
        <h3>Method 1: Automated Installation (Recommended)</h3>
        
        <p>The easiest way to install Go-Mix is using the provided install script:</p>
        
        <div class="code-block">
            <div class="code-header">
                <span class="code-dot red"></span>
                <span class="code-dot yellow"></span>
                <span class="code-dot green"></span>
                <span class="code-title">Terminal</span>
            </div>
{% highlight bash %}
# Clone the repository
git clone https://github.com/akashmaji946/go-mix.git
cd go-mix

# Install system-wide (requires sudo)
sudo ./install.sh

# Or install locally to /usr/local/bin
./install.sh
{% endhighlight %}
        </div>
        
        <div class="callout callout-success">
            <div class="callout-title">
                <i class="fas fa-check-circle"></i> What the installer does
            </div>
            <p>Verifies Go installation, builds the binary, installs to <code>/usr/local/bin</code>, and makes <code>go-mix</code> available from anywhere on your system.</p>
        </div>
        
        <h3>Method 2: Manual Build</h3>
        
        <p>If you prefer to build manually or want to customize the installation:</p>
        
        <div class="code-block">
            <div class="code-header">
                <span class="code-dot red"></span>
                <span class="code-dot yellow"></span>
                <span class="code-dot green"></span>
                <span class="code-title">Terminal</span>
            </div>
{% highlight bash %}
# Clone the repository
git clone https://github.com/akashmaji946/go-mix.git
cd go-mix

# Build the binary
go build -o go-mix ./main

# Or use the build script
./build.sh

# Run tests to verify
go test ./...
{% endhighlight %}
        </div>
        
        <h3>Method 3: Docker Installation</h3>
        
        <p>Run Go-Mix in a container without installing locally:</p>
        
        <div class="code-block">
            <div class="code-header">
                <span class="code-dot red"></span>
                <span class="code-dot yellow"></span>
                <span class="code-dot green"></span>
                <span class="code-title">Terminal</span>
            </div>
{% highlight bash %}
# Pull from DockerHub
docker pull akashmaji946/go-mix:latest

# Run a program
docker run -v $(pwd)/samples:/samples akashmaji946/go-mix /samples/algo/05_factorial.gm

# Interactive REPL
docker run -it akashmaji946/go-mix
{% endhighlight %}
        </div>
        
        <h3>Verify Installation</h3>
        
        <p>Check that Go-Mix is installed correctly:</p>
        
        <div class="code-block">
            <div class="code-header">
                <span class="code-dot red"></span>
                <span class="code-dot yellow"></span>
                <span class="code-dot green"></span>
                <span class="code-title">Terminal</span>
            </div>
{% highlight bash %}
# Check version
go-mix --version

# Display help
go-mix --help
{% endhighlight %}
        </div>
        
        <h2 id="quick-start">Quick Start</h2>
        
        <p>Once installed, you can use Go-Mix in three ways:</p>
        
        <table>
            <thead>
                <tr>
                    <th>Mode</th>
                    <th>Command</th>
                    <th>Use Case</th>
                </tr>
            </thead>
            <tbody>
                <tr>
                    <td><strong>REPL</strong></td>
                    <td><code>go-mix</code></td>
                    <td>Interactive experimentation</td>
                </tr>
                <tr>
                    <td><strong>File</strong></td>
                    <td><code>go-mix file.gm</code></td>
                    <td>Run complete programs</td>
                </tr>
                <tr>
                    <td><strong>Script</strong></td>
                    <td><code>#!/usr/bin/env go-mix</code></td>
                    <td>Executable scripts</td>
                </tr>
            </tbody>
        </table>
        
        <h2 id="first-program">Your First Program</h2>
        
        <p>Create a file named <code>hello.gm</code>:</p>
        
        <div class="code-block">
            <div class="code-header">
                <span class="code-dot red"></span>
                <span class="code-dot yellow"></span>
                <span class="code-dot green"></span>
                <span class="code-title">hello.gm</span>
            </div>
{% highlight go %}
// My first Go-Mix program

// Variables
var name = "World";
let year = 2024;

// Function definition
fn greet(name, year) {
    return "Hello, " + name + "! Welcome to " + year;
}

// Main execution
var message = greet(name, year);
println(message);

// Working with arrays
var numbers = [1, 2, 3, 4, 5];
var doubled = map(numbers, fn(x) { return x * 2; });
println("Doubled: " + doubled);

// Calculate sum using reduce
var sum = reduce(numbers, fn(acc, x) { return acc + x; }, 0);
println("Sum: " + sum);
{% endhighlight %}
        </div>
        
        <p>Run the program:</p>
        
        <div class="code-block">
            <div class="code-header">
                <span class="code-dot red"></span>
                <span class="code-dot yellow"></span>
                <span class="code-dot green"></span>
                <span class="code-title">Terminal</span>
            </div>
{% highlight bash %}
go-mix hello.gm
{% endhighlight %}
        </div>
        
        <p>Expected output:</p>
        
        <div class="code-block">
            <div class="code-header">
                <span class="code-dot red"></span>
                <span class="code-dot yellow"></span>
                <span class="code-dot green"></span>
                <span class="code-title">Output</span>
            </div>
{% highlight %}
Hello, World! Welcome to 2024
Doubled: [2, 4, 6, 8, 10]
Sum: 15
{% endhighlight %}
        </div>
        
        <h2 id="repl">Using the REPL</h2>
        
        <p>The Read-Eval-Print Loop (REPL) is perfect for experimentation:</p>
        
        <div class="code-block">
            <div class="code-header">
                <span class="code-dot red"></span>
                <span class="code-dot yellow"></span>
                <span class="code-dot green"></span>
                <span class="code-title">REPL Session</span>
            </div>
{% highlight bash %}
$ go-mix
Go-Mix >>> var x = 42
Go-Mix >>> println(x)
42
Go-Mix >>> fn square(n) { return n * n; }
Go-Mix >>> square(5)
25
Go-Mix >>> var arr = [1, 2, 3, 4, 5]
Go-Mix >>> map(arr, fn(x) { return x * 2; })
[2, 4, 6, 8, 10]
Go-Mix >>> /help
Available commands:
  /exit    - Exit the REPL
  /scope   - Show current scope and variables
Go-Mix >>> /exit
{% endhighlight %}
        </div>
        
        <h3>REPL Commands</h3>
        
        <table>
            <thead>
                <tr>
                    <th>Command</th>
                    <th>Description</th>
                </tr>
            </thead>
            <tbody>
                <tr>
                    <td><code>/exit</code></td>
                    <td>Exit the REPL</td>
                </tr>
                <tr>
                    <td><code>/scope</code></td>
                    <td>Show current scope and variables</td>
                </tr>
                <tr>
                    <td><code>/help</code></td>
                    <td>Show available commands</td>
                </tr>
            </tbody>
        </table>
        
        <h2 id="next-steps">Next Steps</h2>
        
        <p>Now that you have Go-Mix installed, explore more:</p>
        
        <div class="callout callout-info">
            <div class="callout-title">
                <i class="fas fa-book"></i> Learn the Language
            </div>
            <p>Read the <a href="{{ '/language-guide/' | relative_url }}">Language Guide</a> to understand Go-Mix syntax, types, control flow, functions, and object-oriented programming.</p>
        </div>
        
        <div class="callout callout-info">
            <div class="callout-title">
                <i class="fas fa-cubes"></i> Explore the Standard Library
            </div>
            <p>Discover 100+ built-in functions across 17 packages in the <a href="{{ '/standard-library/' | relative_url }}">Standard Library</a> reference.</p>
        </div>
        
        <div class="callout callout-info">
            <div class="callout-title">
                <i class="fas fa-code"></i> Study Examples
            </div>
            <p>Browse 50+ sample programs in the <a href="{{ '/samples/' | relative_url }}">Samples</a> section, from algorithms to web applications.</p>
        </div>
        
        <div class="callout callout-success">
            <div class="callout-title">
                <i class="fas fa-lightbulb"></i> Pro Tip
            </div>
            <p>Install the <a href="https://github.com/akashmaji946/go-mix/tree/main/vscode-ext">VS Code Extension</a> for syntax highlighting and better development experience.</p>
        </div>
    </div>
</div>
