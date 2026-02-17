---
layout: default
title: Tuples Package - Go-Mix
description: Immutable fixed-size sequences with functional operations
---

<div class="content-page">
    <aside class="sidebar">
        <nav class="sidebar-nav">
            <div class="sidebar-title">Tuple Functions</div>
            <ul class="sidebar-menu">
                <li><a href="#tuple">tuple</a></li>
                <li><a href="#get_tuple">get_tuple</a></li>
                <li><a href="#size_tuple">size_tuple</a></li>
                <li><a href="#slice_tuple">slice_tuple</a></li>
                <li><a href="#concat_tuple">concat_tuple</a></li>
                <li><a href="#contains_tuple">contains_tuple</a></li>
                <li><a href="#index_tuple">index_tuple</a></li>
            </ul>
        </nav>
    </aside>
    
    <div class="content-body">
        <h1>Tuples Package</h1>
        <p>Immutable fixed-size sequences for data records and coordinates.</p>
        
        <div class="function-card" id="tuple">
            <div class="function-header">
                <div class="function-name">tuple</div>
                <div class="function-signature">tuple(...elements) -> tuple</div>
            </div>
            <div class="function-body">
                <div class="function-description">Creates a new immutable tuple.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
tuple();                          // tuple()
tuple(1, 2, 3);                   // tuple(1, 2, 3)
tuple("John", 25, true);          // tuple(John, 25, true)
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="get_tuple">
            <div class="function-header">
                <div class="function-name">get_tuple</div>
                <div class="function-signature">get_tuple(t, index) -> any</div>
            </div>
            <div class="function-body">
                <div class="function-description">Returns element at specified index.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var t = tuple(10, 20, 30);
get_tuple(t, 0);           // 10
get_tuple(t, 2);           // 30
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="size_tuple">
            <div class="function-header">
                <div class="function-name">size_tuple</div>
                <div class="function-signature">size_tuple(t) -> int</div>
            </div>
            <div class="function-body">
                <div class="function-description">Returns number of elements in tuple.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var t = tuple(1, 2, 3);
size_tuple(t);             // 3
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="slice_tuple">
            <div class="function-header">
                <div class="function-name">slice_tuple</div>
                <div class="function-signature">slice_tuple(t, start, end) -> tuple</div>
            </div>
            <div class="function-body">
                <div class="function-description">Returns subtuple from start to end index.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var t = tuple(1, 2, 3, 4, 5);
var sub = slice_tuple(t, 1, 4); // sub = tuple(2, 3, 4)
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="concat_tuple">
            <div class="function-header">
                <div class="function-name">concat_tuple</div>
                <div class="function-signature">concat_tuple(t1, t2) -> tuple</div>
            </div>
            <div class="function-body">
                <div class="function-description">Concatenates two tuples into one.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var t1 = tuple(1, 2);
var t2 = tuple(3, 4);
concat_tuple(t1, t2);      // tuple(1, 2, 3, 4)
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="contains_tuple">
            <div class="function-header">
                <div class="function-name">contains_tuple</div>
                <div class="function-signature">contains_tuple(t, elem) -> bool</div>
            </div>
            <div class="function-body">
                <div class="function-description">Checks if tuple contains the element.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var t = tuple(1, 2, 3);
contains_tuple(t, 2);      // true
contains_tuple(t, 99);     // false
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="index_tuple">
            <div class="function-header">
                <div class="function-name">index_tuple</div>
                <div class="function-signature">index_tuple(t, elem) -> int</div>
            </div>
            <div class="function-body">
                <div class="function-description">Returns index of first occurrence, or -1 if not found.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var t = tuple(1, 2, 3, 2);
index_tuple(t, 2);         // 1
index_tuple(t, 99);        // -1
{% endhighlight %}
                </div>
            </div>
        </div>
    </div>
</div>
