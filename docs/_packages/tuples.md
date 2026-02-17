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
                <li><a href="#import">Import</a></li>
                <li><a href="#make_tuple">make_tuple()</a></li>
                <li><a href="#size_tuple">size_tuple()</a></li>
                <li><a href="#length_tuple">length_tuple()</a></li>
                <li><a href="#peekback_tuple">peekback_tuple()</a></li>
                <li><a href="#peekfront_tuple">peekfront_tuple()</a></li>
                <li><a href="#contains_tuple">contains_tuple()</a></li>
                <li><a href="#find_tuple">find_tuple()</a></li>
                <li><a href="#some_tuple">some_tuple()</a></li>
                <li><a href="#every_tuple">every_tuple()</a></li>
                <li><a href="#map_tuple">map_tuple()</a></li>
                <li><a href="#filter_tuple">filter_tuple()</a></li>
                <li><a href="#reduce_tuple">reduce_tuple()</a></li>
                <li><a href="#to_tuple">to_tuple()</a></li>
            </ul>
        </nav>
    </aside>
    
    <div class="content-body">
        <h1>Tuples Package</h1>
        <p>Immutable fixed-size sequences for data records and coordinates.</p>
        
        <div class="function-card" id="import">
            <div class="function-header">
                <div class="function-name">Import</div>
                <div class="function-signature">import "tuples"</div>
            </div>
            <div class="function-body">
                <div class="function-description">Import the tuples package to use namespaced functions.</div>
                <div class="function-example">
                    <h4>Examples</h4>
{% highlight go %}
// Standard import
import tuple;
var t = tuple.tuple(1, 2, 3)
var size = tuple.size_tuple(t)

// With alias
import tuple as tup;
var t = tup.tuple(1, 2, 3);
var size = tup.size_tuple(t);
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="make_tuple">
            <div class="function-header">
                <div class="function-name">make_tuple</div>
                <div class="function-signature">make_tuple(...elements) -> tuple</div>
            </div>
            <div class="function-body">
                <div class="function-description">Creates a new immutable tuple from arguments. Tuples are heterogeneous and immutable, preventing modifications after creation.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
make_tuple();                          // tuple()
make_tuple(1, 2, 3);                   // tuple(1, 2, 3)
make_tuple("John", 25, true);          // tuple(John, 25, true)
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
                <div class="function-description">Returns the number of elements in a tuple.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var t = make_tuple(1, 2, 3);
size_tuple(t);             // 3
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="length_tuple">
            <div class="function-header">
                <div class="function-name">length_tuple</div>
                <div class="function-signature">length_tuple(t) -> int</div>
            </div>
            <div class="function-body">
                <div class="function-description">Returns the number of elements in a tuple (alias for size_tuple).</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var t = make_tuple(1, 2, 3);
length_tuple(t);           // 3
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="peekback_tuple">
            <div class="function-header">
                <div class="function-name">peekback_tuple</div>
                <div class="function-signature">peekback_tuple(t) -> any</div>
            </div>
            <div class="function-body">
                <div class="function-description">Returns the last element of a tuple. Returns an error if the tuple is empty.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var t = make_tuple(1, 2, 3);
peekback_tuple(t);         // 3
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="peekfront_tuple">
            <div class="function-header">
                <div class="function-name">peekfront_tuple</div>
                <div class="function-signature">peekfront_tuple(t) -> any</div>
            </div>
            <div class="function-body">
                <div class="function-description">Returns the first element of a tuple. Returns an error if the tuple is empty.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var t = make_tuple(1, 2, 3);
peekfront_tuple(t);        // 1
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
                <div class="function-description">Checks if a value exists in the tuple. Comparison is done using both type and value.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var t = make_tuple(1, 2, 3, 4);
contains_tuple(t, 3);      // true
contains_tuple(t, 5);      // false
contains_tuple(t, "2");    // false (type matters)
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="find_tuple">
            <div class="function-header">
                <div class="function-name">find_tuple</div>
                <div class="function-signature">find_tuple(t, function) -> any</div>
            </div>
            <div class="function-body">
                <div class="function-description">Finds the first element that satisfies the provided testing function. Returns nil if no element matches.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var t = make_tuple(1, 2, 3, 4, 5);
find_tuple(t, fn(x) { x > 3 });   // 4
find_tuple(t, fn(x) { x > 10 });  // nil
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="some_tuple">
            <div class="function-header">
                <div class="function-name">some_tuple</div>
                <div class="function-signature">some_tuple(t, function) -> bool</div>
            </div>
            <div class="function-body">
                <div class="function-description">Checks if at least one element in the tuple passes the test (predicate function).</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var t = make_tuple(1, 2, 3);
some_tuple(t, fn(x) { x > 2 });   // true
some_tuple(t, fn(x) { x > 5 });   // false
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="every_tuple">
            <div class="function-header">
                <div class="function-name">every_tuple</div>
                <div class="function-signature">every_tuple(t, function) -> bool</div>
            </div>
            <div class="function-body">
                <div class="function-description">Checks if all elements in the tuple pass the test (predicate function).</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var t = make_tuple(2, 4, 6);
every_tuple(t, fn(x) { x % 2 == 0 });  // true
every_tuple(t, fn(x) { x > 3 });       // false
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="map_tuple">
            <div class="function-header">
                <div class="function-name">map_tuple</div>
                <div class="function-signature">map_tuple(t, function) -> list</div>
            </div>
            <div class="function-body">
                <div class="function-description">Applies a function to each element of the tuple and returns a list with the results.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var t = make_tuple(1, 2, 3);
map_tuple(t, fn(x) { x * 2 });   // [2, 4, 6]
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="filter_tuple">
            <div class="function-header">
                <div class="function-name">filter_tuple</div>
                <div class="function-signature">filter_tuple(t, function) -> list</div>
            </div>
            <div class="function-body">
                <div class="function-description">Filters elements based on a predicate function and returns a list with matching elements.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var t = make_tuple(1, 2, 3, 4, 5);
filter_tuple(t, fn(x) { x > 2 });   // [3, 4, 5]
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="reduce_tuple">
            <div class="function-header">
                <div class="function-name">reduce_tuple</div>
                <div class="function-signature">reduce_tuple(t, function, initial) -> any</div>
            </div>
            <div class="function-body">
                <div class="function-description">Reduces the tuple to a single value using a binary function (accumulator, current) -> newAccumulator.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var t = make_tuple(1, 2, 3, 4);
reduce_tuple(t, fn(acc, x) { acc + x }, 0);   // 10
reduce_tuple(t, fn(acc, x) { acc * x }, 1);   // 24
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="to_tuple">
            <div class="function-header">
                <div class="function-name">to_tuple</div>
                <div class="function-signature">to_tuple(iterable) -> tuple</div>
            </div>
            <div class="function-body">
                <div class="function-description">Converts an array or list to a tuple. If the argument is already a tuple, returns it unchanged.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var arr = [1, 2, 3];
to_tuple(arr);             // tuple(1, 2, 3)

var lst = list(4, 5, 6);
to_tuple(lst);             // tuple(4, 5, 6)
{% endhighlight %}
                </div>
            </div>
        </div>
    </div>
</div>
