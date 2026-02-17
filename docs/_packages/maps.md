---
layout: default
title: Maps Package - Go-Mix
description: Dictionary operations including keys, insert, remove, contain, enumerate
---

<div class="content-page">
    <aside class="sidebar">
        <nav class="sidebar-nav">
            <div class="sidebar-title">Map Functions</div>
            <ul class="sidebar-menu">
                <li><a href="#import">Import</a></li>
                <li><a href="#make_map">make_map()</a></li>
                <li><a href="#keys_map">keys_map()</a></li>
                <li><a href="#values_map">values_map()</a></li>
                <li><a href="#insert_map">insert_map()</a></li>
                <li><a href="#remove_map">remove_map()</a></li>
                <li><a href="#contain_map">contain_map()</a></li>
                <li><a href="#enumerate_map">enumerate_map()</a></li>
                <li><a href="#size_map">size_map()</a></li>
            </ul>
        </nav>
    </aside>
    
    <div class="content-body">
        <h1>Maps Package</h1>
        <p>Dictionary operations for key-value data structures.</p>
        
        <div class="function-card" id="import">
            <div class="function-header">
                <div class="function-name">Import</div>
                <div class="function-signature">import "maps"</div>
            </div>
            <div class="function-body">
                <div class="function-description">Import the maps package to use namespaced functions.</div>
                <div class="function-example">
                    <h4>Examples</h4>
{% highlight go %}
// Standard import
import maps;
var m = maps.make_map("name", "John", "age", 25)
var keys = maps.keys_map(m)

// With alias
import maps as m;
var map = m.make_map("name", "John", "age", 25)
var keys = m.keys_map(map)
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="make_map">
            <div class="function-header">
                <div class="function-name">make_map</div>
                <div class="function-signature">make_map(key1, val1, key2, val2, ...) -> map</div>
            </div>
            <div class="function-body">
                <div class="function-description">Creates a new map from key-value pairs.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var m = make_map("name", "John", "age", 25);
// m = map{"name": "John", "age": 25}
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="keys_map">
            <div class="function-header">
                <div class="function-name">keys_map</div>
                <div class="function-signature">keys_map(m) -> array</div>
            </div>
            <div class="function-body">
                <div class="function-description">Returns array of all keys in the map.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var m = map{"name": "John", "age": 25};
keys_map(m);               // ["name", "age"]
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="values_map">
            <div class="function-header">
                <div class="function-name">values_map</div>
                <div class="function-signature">values_map(m) -> array</div>
            </div>
            <div class="function-body">
                <div class="function-description">Returns array of all values in the map.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var m = map{"name": "John", "age": 25};
values_map(m);             // ["John", 25]
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="insert_map">
            <div class="function-header">
                <div class="function-name">insert_map</div>
                <div class="function-signature">insert_map(m, key, value) -> any</div>
            </div>
            <div class="function-body">
                <div class="function-description">Inserts or updates a key-value pair in the map.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var m = map{"name": "John"};
insert_map(m, "age", 25);  // m = map{"name": "John", "age": 25}
insert_map(m, "name", "Jane");  // Updates existing key
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="remove_map">
            <div class="function-header">
                <div class="function-name">remove_map</div>
                <div class="function-signature">remove_map(m, key) -> any</div>
            </div>
            <div class="function-body">
                <div class="function-description">Removes a key from the map and returns its value.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var m = map{"name": "John", "age": 25};
var age = remove_map(m, "age");  // age = 25, m = map{"name": "John"}
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="contain_map">
            <div class="function-header">
                <div class="function-name">contain_map</div>
                <div class="function-signature">contain_map(m, key) -> bool</div>
            </div>
            <div class="function-body">
                <div class="function-description">Checks if map contains the specified key.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var m = map{"name": "John", "age": 25};
contain_map(m, "name");    // true
contain_map(m, "email");   // false
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="enumerate_map">
            <div class="function-header">
                <div class="function-name">enumerate_map</div>
                <div class="function-signature">enumerate_map(m) -> array</div>
            </div>
            <div class="function-body">
                <div class="function-description">Returns array of [key, value] pairs for iteration.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var m = map{"name": "John", "age": 25};
enumerate_map(m);          // [["name", "John"], ["age", 25]]
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="size_map">
            <div class="function-header">
                <div class="function-name">size_map</div>
                <div class="function-signature">size_map(m) -> int</div>
            </div>
            <div class="function-body">
                <div class="function-description">Returns the number of key-value pairs in the map.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var m = map{"a": 1, "b": 2, "c": 3};
size_map(m);               // 3
{% endhighlight %}
                </div>
            </div>
        </div>
    </div>
</div>
