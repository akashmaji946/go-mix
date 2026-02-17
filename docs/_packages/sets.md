---
layout: default
title: Sets Package - Go-Mix
description: Set operations for unique value collections
---

<div class="content-page">
    <aside class="sidebar">
        <nav class="sidebar-nav">
            <div class="sidebar-title">Set Functions</div>
            <ul class="sidebar-menu">
                <li><a href="#make_set">make_set</a></li>
                <li><a href="#insert_set">insert_set</a></li>
                <li><a href="#remove_set">remove_set</a></li>
                <li><a href="#contains_set">contains_set</a></li>
                <li><a href="#values_set">values_set</a></li>
                <li><a href="#size_set">size_set</a></li>
            </ul>
        </nav>
    </aside>
    
    <div class="content-body">
        <h1>Sets Package</h1>
        <p>Unique value collections with set operations.</p>
        
        <div class="function-card" id="make_set">
            <div class="function-header">
                <div class="function-name">make_set</div>
                <div class="function-signature">make_set(...elements) -> set</div>
            </div>
            <div class="function-body">
                <div class="function-description">Creates a new set with unique elements.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
make_set(1, 2, 3);                    // set{1, 2, 3}
make_set(1, 2, 2, 3, 3, 3);           // set{1, 2, 3} (duplicates removed)
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="insert_set">
            <div class="function-header">
                <div class="function-name">insert_set</div>
                <div class="function-signature">insert_set(s, elem) -> bool</div>
            </div>
            <div class="function-body">
                <div class="function-description">Adds element to set (if not already present). Returns true if added.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var s = make_set(1, 2, 3);
insert_set(s, 4);          // true, s = set{1, 2, 3, 4}
insert_set(s, 2);          // false (2 already exists)
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="remove_set">
            <div class="function-header">
                <div class="function-name">remove_set</div>
                <div class="function-signature">remove_set(s, elem) -> bool</div>
            </div>
            <div class="function-body">
                <div class="function-description">Removes element from set. Returns true if element was removed.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var s = make_set(1, 2, 3);
remove_set(s, 2);          // true, s = set{1, 3}
remove_set(s, 99);         // false (element didn't exist)
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="contains_set">
            <div class="function-header">
                <div class="function-name">contains_set</div>
                <div class="function-signature">contains_set(s, elem) -> bool</div>
            </div>
            <div class="function-body">
                <div class="function-description">Checks if set contains the element.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var s = make_set(1, 2, 3);
contains_set(s, 2);        // true
contains_set(s, 99);       // false
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="values_set">
            <div class="function-header">
                <div class="function-name">values_set</div>
                <div class="function-signature">values_set(s) -> array</div>
            </div>
            <div class="function-body">
                <div class="function-description">Returns an array of all values in the set.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var s = make_set(1, 2, 3);
values_set(s);             // [1, 2, 3]
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="size_set">
            <div class="function-header">
                <div class="function-name">size_set</div>
                <div class="function-signature">size_set(s) -> int</div>
            </div>
            <div class="function-body">
                <div class="function-description">Returns number of unique elements in set. Alias: length_set.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var s = make_set(1, 2, 3, 3, 3);
size_set(s);               // 3
length_set(s);             // 3 (alias)
{% endhighlight %}
                </div>
            </div>
        </div>
    </div>
</div>
