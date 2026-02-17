---
layout: default
title: OS Package - Go-Mix
description: Operating system functions including getenv, exec, sleep, getpid
---

<div class="content-page">
    <aside class="sidebar">
        <nav class="sidebar-nav">
            <div class="sidebar-title">OS Functions</div>
            <ul class="sidebar-menu">
                <li><a href="#getenv">getenv</a></li>
                <li><a href="#setenv">setenv</a></li>
                <li><a href="#exec">exec</a></li>
                <li><a href="#getpid">getpid</a></li>
                <li><a href="#hostname">hostname</a></li>
                <li><a href="#user">user</a></li>
                <li><a href="#exit">exit</a></li>
                <li><a href="#args">args</a></li>
            </ul>
        </nav>
    </aside>
    
    <div class="content-body">
        <h1>OS Package</h1>
        <p>Operating system interaction functions.</p>
        
        <div class="function-card" id="getenv">
            <div class="function-header">
                <div class="function-name">getenv</div>
                <div class="function-signature">getenv(key) -> string</div>
            </div>
            <div class="function-body">
                <div class="function-description">Returns environment variable value.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var home = getenv("HOME");
var path = getenv("PATH");
println("Home: " + home);
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="setenv">
            <div class="function-header">
                <div class="function-name">setenv</div>
                <div class="function-signature">setenv(key, value) -> nil</div>
            </div>
            <div class="function-body">
                <div class="function-description">Sets environment variable.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
setenv("MY_VAR", "my_value");
var val = getenv("MY_VAR");
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="exec">
            <div class="function-header">
                <div class="function-name">exec</div>
                <div class="function-signature">exec(command, ...args) -> string</div>
            </div>
            <div class="function-body">
                <div class="function-description">Executes shell command and returns output.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var output = exec("echo", "Hello");
println(output);

var files = exec("ls", "-la");
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="getpid">
            <div class="function-header">
                <div class="function-name">getpid</div>
                <div class="function-signature">getpid() -> int</div>
            </div>
            <div class="function-body">
                <div class="function-description">Returns current process ID.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var pid = getpid();
println("Process ID: " + pid);
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="hostname">
            <div class="function-header">
                <div class="function-name">hostname</div>
                <div class="function-signature">hostname() -> string</div>
            </div>
            <div class="function-body">
                <div class="function-description">Returns system hostname.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var host = hostname();
println("Hostname: " + host);
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="user">
            <div class="function-header">
                <div class="function-name">user</div>
                <div class="function-signature">user() -> string</div>
            </div>
            <div class="function-body">
                <div class="function-description">Returns current username.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var username = user();
println("User: " + username);
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="exit">
            <div class="function-header">
                <div class="function-name">exit</div>
                <div class="function-signature">exit([code]) -> nil</div>
            </div>
            <div class="function-body">
                <div class="function-description">Terminates program with optional exit code.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
if (error) {
    println("Error occurred");
    exit(1);
}
exit(0);  // Success
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="args">
            <div class="function-header">
                <div class="function-name">args</div>
                <div class="function-signature">args() -> array</div>
            </div>
            <div class="function-body">
                <div class="function-description">Returns command-line arguments array.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var arguments = args();
foreach arg in arguments {
    println(arg);
}
{% endhighlight %}
                </div>
            </div>
        </div>
    </div>
</div>
