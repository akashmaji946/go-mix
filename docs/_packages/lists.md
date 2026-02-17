---
layout: default
title: Lists Package - Go-Mix
description: List operations for mutable heterogeneous sequences
---

<div class="content-page">
    <aside class="sidebar">
        <nav class="sidebar-nav">
            <div class="sidebar-title">List Functions</div>
            <ul class="sidebar-menu">
                <li><a href="#list">list</a></li>
                <li><a href="#pushback_list">pushback_list</a></li>
                <li><a href="#pushfront_list">pushfront_list</a></li>
                <li><a href="#popback_list">popback_list</a></li>
                <li><a href="#popfront_list">popfront_list</a></li>
                <li><a href="#insert_list">insert_list</a></li>
                <li><a href="#remove_list">remove_list</a></li>
                <li><a href="#slice_list">slice_list</a></li>
                <li><a href="#contains_list">contains_list</a></li>
                <li><a href="#index_list">index_list</a></li>
                <li><a href="#peekback_list">peekback_list</a></li>
                <li><a href="#peekfront_list">peekfront_list</a></li>
                <li><a href="#map_list">map_list</a></li>
                <li><a href="#filter_list">filter_list</a></li>
                <li><a href="#reduce_list">reduce_list</a></li>
                <li><a href="#find_list">find_list</a></li>
                <li><a href="#some_list">some_list</a></li>
                <li><a href="#every_list">every_list</a></li>
                <li><a href="#to_list">to_list</a></li>
                <li><a href="#size_list">size_list</a></li>
            </ul>
        </nav>
    </aside>
    
    <div class="content-body">
        <h1>Lists Package</h1>
        <p>Mutable heterogeneous sequences with flexible operations.</p>
        
        <div class="function-card" id="list">
            <div class="function-header">
                <div class="function-name">list</div>
                <div class="function-signature">list(...elements) -> list</div>
            </div>
            <div class="function-body">
                <div class="function-description">Creates a new list from elements.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
list();                           // list()
list(1, 2, 3);                    // list(1, 2, 3)
list(1, "two", 3.0, true);        // list(1, two, 3.0, true)
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="pushback_list">
            <div class="function-header">
                <div class="function-name">pushback_list</div>
                <div class="function-signature">pushback_list(l, elem) -> list</div>
            </div>
            <div class="function-body">
                <div class="function-description">Adds element to the end of the list.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var l = list(1, 2, 3);
pushback_list(l, 4);       // l = list(1, 2, 3, 4)
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="pushfront_list">
            <div class="function-header">
                <div class="function-name">pushfront_list</div>
                <div class="function-signature">pushfront_list(l, elem) -> list</div>
            </div>
            <div class="function-body">
                <div class="function-description">Adds element to the beginning of the list.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var l = list(1, 2, 3);
pushfront_list(l, 0);      // l = list(0, 1, 2, 3)
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="popback_list">
            <div class="function-header">
                <div class="function-name">popback_list</div>
                <div class="function-signature">popback_list(l) -> any</div>
            </div>
            <div class="function-body">
                <div class="function-description">Removes and returns the last element.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var l = list(1, 2, 3);
var last = popback_list(l); // last = 3, l = list(1, 2)
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="popfront_list">
            <div class="function-header">
                <div class="function-name">popfront_list</div>
                <div class="function-signature">popfront_list(l) -> any</div>
            </div>
            <div class="function-body">
                <div class="function-description">Removes and returns the first element.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var l = list(1, 2, 3);
var first = popfront_list(l); // first = 1, l = list(2, 3)
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="insert_list">
            <div class="function-header">
                <div class="function-name">insert_list</div>
                <div class="function-signature">insert_list(l, index, elem) -> list</div>
            </div>
            <div class="function-body">
                <div class="function-description">Inserts element at specified index.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var l = list(1, 2, 3);
insert_list(l, 1, 99);     // l = list(1, 99, 2, 3)
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="remove_list">
            <div class="function-header">
                <div class="function-name">remove_list</div>
                <div class="function-signature">remove_list(l, index) -> any</div>
            </div>
            <div class="function-body">
                <div class="function-description">Removes and returns element at specified index.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var l = list(1, 2, 3);
var removed = remove_list(l, 1); // removed = 2, l = list(1, 3)
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="slice_list">
            <div class="function-header">
                <div class="function-name">slice_list</div>
                <div class="function-signature">slice_list(l, start, end) -> list</div>
            </div>
            <div class="function-body">
                <div class="function-description">Returns a sublist from start to end index.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var l = list(1, 2, 3, 4, 5);
var sub = slice_list(l, 1, 4); // sub = list(2, 3, 4)
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="contains_list">
            <div class="function-header">
                <div class="function-name">contains_list</div>
                <div class="function-signature">contains_list(l, elem) -> bool</div>
            </div>
            <div class="function-body">
                <div class="function-description">Checks if list contains the element.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var l = list(1, 2, 3);
contains_list(l, 2);       // true
contains_list(l, 99);      // false
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="index_list">
            <div class="function-header">
                <div class="function-name">index_list</div>
                <div class="function-signature">index_list(l, elem) -> int</div>
            </div>
            <div class="function-body">
                <div class="function-description">Returns index of first occurrence of element, or -1 if not found.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var l = list(1, 2, 3, 2);
index_list(l, 2);          // 1
index_list(l, 99);         // -1
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="peekback_list">
            <div class="function-header">
                <div class="function-name">peekback_list</div>
                <div class="function-signature">peekback_list(l) -> any</div>
            </div>
            <div class="function-body">
                <div class="function-description">Returns the last element without removing it.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var l = list(1, 2, 3);
var last = peekback_list(l); // last = 3, l = list(1, 2, 3) unchanged
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="peekfront_list">
            <div class="function-header">
                <div class="function-name">peekfront_list</div>
                <div class="function-signature">peekfront_list(l) -> any</div>
            </div>
            <div class="function-body">
                <div class="function-description">Returns the first element without removing it.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var l = list(1, 2, 3);
var first = peekfront_list(l); // first = 1, l = list(1, 2, 3) unchanged
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="map_list">
            <div class="function-header">
                <div class="function-name">map_list</div>
                <div class="function-signature">map_list(l, fn) -> list</div>
            </div>
            <div class="function-body">
                <div class="function-description">Applies a function to each element and returns a new list.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var l = list(1, 2, 3);
var doubled = map_list(l, fn(x) { return x * 2; });
// doubled = list(2, 4, 6)
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="filter_list">
            <div class="function-header">
                <div class="function-name">filter_list</div>
                <div class="function-signature">filter_list(l, fn) -> list</div>
            </div>
            <div class="function-body">
                <div class="function-description">Returns elements that satisfy the predicate function.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var l = list(1, 2, 3, 4, 5, 6);
var evens = filter_list(l, fn(x) { return x % 2 == 0; });
// evens = list(2, 4, 6)
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="reduce_list">
            <div class="function-header">
                <div class="function-name">reduce_list</div>
                <div class="function-signature">reduce_list(l, fn, initial) -> any</div>
            </div>
            <div class="function-body">
                <div class="function-description">Reduces list to a single value using an accumulator function.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var l = list(1, 2, 3, 4, 5);
var sum = reduce_list(l, fn(acc, x) { return acc + x; }, 0);
// sum = 15
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="find_list">
            <div class="function-header">
                <div class="function-name">find_list</div>
                <div class="function-signature">find_list(l, fn) -> any</div>
            </div>
            <div class="function-body">
                <div class="function-description">Returns the first element matching the predicate, or nil if none found.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var l = list(1, 2, 3, 4, 5);
var firstEven = find_list(l, fn(x) { return x % 2 == 0; });
// firstEven = 2
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="some_list">
            <div class="function-header">
                <div class="function-name">some_list</div>
                <div class="function-signature">some_list(l, fn) -> bool</div>
            </div>
            <div class="function-body">
                <div class="function-description">Returns true if at least one element satisfies the predicate.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var l = list(1, 2, 3, 4, 5);
some_list(l, fn(x) { return x > 3; });  // true
some_list(l, fn(x) { return x > 10; }); // false
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="every_list">
            <div class="function-header">
                <div class="function-name">every_list</div>
                <div class="function-signature">every_list(l, fn) -> bool</div>
            </div>
            <div class="function-body">
                <div class="function-description">Returns true if all elements satisfy the predicate.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var l = list(2, 4, 6, 8);
every_list(l, fn(x) { return x % 2 == 0; });  // true
every_list(l, fn(x) { return x > 5; });       // false
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="to_list">
            <div class="function-header">
                <div class="function-name">to_list</div>
                <div class="function-signature">to_list(iterable) -> list</div>
            </div>
            <div class="function-body">
                <div class="function-description">Converts an array or tuple to a list.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var a = [1, 2, 3];
var t = tuple(4, 5, 6);
to_list(a);                // list(1, 2, 3)
to_list(t);                // list(4, 5, 6)
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="size_list">
            <div class="function-header">
                <div class="function-name">size_list</div>
                <div class="function-signature">size_list(l) -> int</div>
            </div>
            <div class="function-body">
                <div class="function-description">Returns the number of elements in the list. Alias: length_list.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var l = list(1, 2, 3, 4, 5);
size_list(l);              // 5
length_list(l);            // 5 (alias)
{% endhighlight %}
                </div>
            </div>
        </div>
    </div>
</div>
