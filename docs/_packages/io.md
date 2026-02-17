---
layout: default
title: I/O Package - Go-Mix
description: Input/output operations including scanln, scanf, input, sprintf
---

<div class="content-page">
    <aside class="sidebar">
        <nav class="sidebar-nav">
            <div class="sidebar-title">I/O Functions</div>
            <ul class="sidebar-menu">
                <li><a href="#scanln">scanln</a></li>
                <li><a href="#scanf">scanf</a></li>
                <li><a href="#input">input</a></li>
                <li><a href="#scan">scan</a></li>
                <li><a href="#getchar">getchar</a></li>
                <li><a href="#putchar">putchar</a></li>
                <li><a href="#sprintf">sprintf</a></li>
                <li><a href="#flush">flush</a></li>
            </ul>
        </nav>
    </aside>
    
    <div class="content-body">
        <h1>I/O Package</h1>
        <p>Input and output operations for user interaction.</p>
        
        <div class="function-card" id="scanln">
            <div class="function-header">
                <div class="function-name">scanln</div>
                <div class="function-signature">scanln() -> string</div>
            </div>
            <div class="function-body">
                <div class="function-description">Reads a line from standard input.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
println("Enter your name:");
var name = scanln();
println("Hello, " + name);
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="scanf">
            <div class="function-header">
                <div class="function-name">scanf</div>
                <div class="function-signature">scanf(format) -> array</div>
            </div>
            <div class="function-body">
                <div class="function-description">Reads formatted input and returns array of values.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
printf("Enter age and name: ");
var data = scanf("%d %s");
var age = data[0];
var name = data[1];
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="input">
            <div class="function-header">
                <div class="function-name">input</div>
                <div class="function-signature">input(prompt) -> string</div>
            </div>
            <div class="function-body">
                <div class="function-description">Displays prompt and reads user input.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var name = input("Enter your name: ");
var age = input("Enter your age: ");
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="sprintf">
            <div class="function-header">
                <div class="function-name">sprintf</div>
                <div class="function-signature">sprintf(format, ...args) -> string</div>
            </div>
            <div class="function-body">
                <div class="function-description">Returns formatted string without printing.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var msg = sprintf("Hello, %s! You are %d years old.", "Alice", 30);
// msg = "Hello, Alice! You are 30 years old."
{% endhighlight %}
                </div>
            </div>
        </div>
    </div>
</div>
