---
layout: default
title: JSON Package - Go-Mix
description: JSON parsing and serialization functions
---

<div class="content-page">
    <aside class="sidebar">
        <nav class="sidebar-nav">
            <div class="sidebar-title">JSON Functions</div>
            <ul class="sidebar-menu">
                <li><a href="#map_to_json_string">map_to_json_string</a></li>
                <li><a href="#json_string_to_map">json_string_to_map</a></li>
            </ul>
        </nav>
    </aside>
    
    <div class="content-body">
        <h1>JSON Package</h1>
        <p>JSON encoding and decoding functions.</p>
        
        <div class="function-card" id="map_to_json_string">
            <div class="function-header">
                <div class="function-name">map_to_json_string</div>
                <div class="function-signature">map_to_json_string(m) -> string</div>
            </div>
            <div class="function-body">
                <div class="function-description">Converts map to JSON string.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var user = map{
    "name": "Alice",
    "age": 30,
    "active": true
};
var json = map_to_json_string(user);
// {"name":"Alice","age":30,"active":true}
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="json_string_to_map">
            <div class="function-header">
                <div class="function-name">json_string_to_map</div>
                <div class="function-signature">json_string_to_map(json) -> map</div>
            </div>
            <div class="function-body">
                <div class="function-description">Parses JSON string into map.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var json = '{"name":"Bob","age":25}';
var data = json_string_to_map(json);
println(data["name"]);     // "Bob"
println(data["age"]);      // 25
{% endhighlight %}
                </div>
            </div>
        </div>
    </div>
</div>
