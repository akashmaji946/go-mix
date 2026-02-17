---
layout: default
title: HTTP Package - Go-Mix
description: HTTP client and server functions for web requests and APIs
---

<div class="content-page">
    <aside class="sidebar">
        <nav class="sidebar-nav">
            <div class="sidebar-title">HTTP Functions</div>
            <ul class="sidebar-menu">
                <li><a href="#get_http">get_http</a></li>
                <li><a href="#post_http">post_http</a></li>
                <li><a href="#request_http">request_http</a></li>
                <li><a href="#create_server">create_server</a></li>
                <li><a href="#handle_server">handle_server</a></li>
                <li><a href="#start_server">start_server</a></li>
            </ul>
        </nav>
    </aside>
    
    <div class="content-body">
        <h1>HTTP Package</h1>
        <p>HTTP client and server functionality for web development.</p>
        
        <div class="function-card" id="get_http">
            <div class="function-header">
                <div class="function-name">get_http</div>
                <div class="function-signature">get_http(url) -> string</div>
            </div>
            <div class="function-body">
                <div class="function-description">Performs HTTP GET request.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var response = get_http("https://api.example.com/data");
println(response);
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="post_http">
            <div class="function-header">
                <div class="function-name">post_http</div>
                <div class="function-signature">post_http(url, body) -> string</div>
            </div>
            <div class="function-body">
                <div class="function-description">Performs HTTP POST request.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var body = '{"name": "John", "age": 30}';
var response = post_http("https://api.example.com/users", body);
println(response);
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="request_http">
            <div class="function-header">
                <div class="function-name">request_http</div>
                <div class="function-signature">request_http(method, url, [headers], [body]) -> map</div>
            </div>
            <div class="function-body">
                <div class="function-description">Generic HTTP request with full control.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var headers = map{
    "Content-Type": "application/json",
    "Authorization": "Bearer token123"
};
var response = request_http("PUT", "https://api.example.com/data", headers, '{"key": "value"}');
println(response["status"]);
println(response["body"]);
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="create_server">
            <div class="function-header">
                <div class="function-name">create_server</div>
                <div class="function-signature">create_server() -> server</div>
            </div>
            <div class="function-body">
                <div class="function-description">Creates new HTTP server instance.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var srv = create_server();
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="handle_server">
            <div class="function-header">
                <div class="function-name">handle_server</div>
                <div class="function-signature">handle_server(server, path, handler) -> nil</div>
            </div>
            <div class="function-body">
                <div class="function-description">Registers route handler for server.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
handle_server(srv, "/", fn(req) {
    return "Hello, World!";
});

handle_server(srv, "/api/users", fn(req) {
    return '{"users": ["Alice", "Bob"]}';
});
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="start_server">
            <div class="function-header">
                <div class="function-name">start_server</div>
                <div class="function-signature">start_server(server, address) -> nil</div>
            </div>
            <div class="function-body">
                <div class="function-description">Starts server on specified address.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
start_server(srv, ":8080");
println("Server running on http://localhost:8080");
{% endhighlight %}
                </div>
            </div>
        </div>
    </div>
</div>
