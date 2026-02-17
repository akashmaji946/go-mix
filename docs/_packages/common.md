---
layout: default
title: Common Package - Go-Mix
description: Core functions including print, length, typeof, range, and type constructors
---

<div class="content-page">
    <aside class="sidebar">
        <nav class="sidebar-nav">
            <div class="sidebar-title">Common Functions</div>
            <ul class="sidebar-menu">
                <li><a href="#print">print</a></li>
                <li><a href="#println">println</a></li>
                <li><a href="#printf">printf</a></li>
                <li><a href="#length">length/size</a></li>
                <li><a href="#typeof">typeof</a></li>
                <li><a href="#range">range</a></li>
                <li><a href="#tostring">to_string</a></li>
                <li><a href="#array">array</a></li>
                <li><a href="#list">list</a></li>
                <li><a href="#tuple">tuple</a></li>
                <li><a href="#addr">addr</a></li>
                <li><a href="#is_same_ref">is_same_ref</a></li>
            </ul>
        </nav>
    </aside>
    
    <div class="content-body">
        <h1>Common Package</h1>
        
        <p>The Common package provides essential functions for output, type inspection, and object creation. These are the most frequently used functions in Go-Mix.</p>
        
        <div class="function-card" id="print">
            <div class="function-header">
                <div class="function-name">print</div>
                <div class="function-signature">print(...args) -> nil</div>
            </div>
            <div class="function-body">
                <div class="function-description">
                    Outputs the string representations of arguments without a trailing newline.
                </div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
print("Hello");              // Output: Hello (no newline)
print("World");            // Output: World
print("Value:", 42);       // Output: Value: 42
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="println">
            <div class="function-header">
                <div class="function-name">println</div>
                <div class="function-signature">println(...args) -> nil</div>
            </div>
            <div class="function-body">
                <div class="function-description">
                    Outputs the string representations of arguments with a trailing newline.
                </div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
println("Hello");          // Output: Hello\n
println("World");          // Output: World\n
println("Value:", 42);     // Output: Value: 42\n
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="printf">
            <div class="function-header">
                <div class="function-name">printf</div>
                <div class="function-signature">printf(format, ...args) -> nil</div>
            </div>
            <div class="function-body">
                <div class="function-description">
                    Outputs a formatted string using C-style format specifiers.
                </div>
                <div class="function-params">
                    <h4>Parameters</h4>
                    <ul class="param-list">
                        <li><span class="param-name">format</span> <span class="param-type">string - Format string with %d, %s, %f, %x, etc.</span></li>
                        <li><span class="param-name">args</span> <span class="param-type">any - Values to format</span></li>
                    </ul>
                </div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
printf("Number: %d\n", 42);           // Number: 42
printf("String: %s\n", "hello");       // String: hello
printf("Float: %.2f\n", 3.14159);     // Float: 3.14
printf("Hex: %x\n", 255);              // Hex: ff
printf("Multiple: %d %s\n", 1, "a"); // Multiple: 1 a
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="length">
            <div class="function-header">
                <div class="function-name">length / size</div>
                <div class="function-signature">length(obj) -> int</div>
            </div>
            <div class="function-body">
                <div class="function-description">
                    Returns the length of strings, arrays, maps, sets, lists, or tuples.
                </div>
                <div class="function-params">
                    <h4>Parameters</h4>
                    <ul class="param-list">
                        <li><span class="param-name">obj</span> <span class="param-type">string|array|map|set|list|tuple - Object to measure</span></li>
                    </ul>
                </div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
length("hello");           // 5
length([1, 2, 3]);         // 3
length(map{"a": 1, "b": 2}); // 2
length(set{1, 2, 3});      // 3
size("hello");             // 5 (alias)
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="typeof">
            <div class="function-header">
                <div class="function-name">typeof</div>
                <div class="function-signature">typeof(obj) -> string</div>
            </div>
            <div class="function-body">
                <div class="function-description">
                    Returns the type name of any object as a string.
                </div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
typeof(42);                // "int"
typeof(3.14);              // "float"
typeof("hello");           // "string"
typeof(true);              // "bool"
typeof(nil);               // "nil"
typeof([1, 2, 3]);         // "array"
typeof(range(1, 5));       // "range"
typeof(func() {});           // "func"
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="range">
            <div class="function-header">
                <div class="function-name">range</div>
                <div class="function-signature">range(start, end) -> range</div>
            </div>
            <div class="function-body">
                <div class="function-description">
                    Creates an inclusive range object for iteration.
                </div>
                <div class="function-params">
                    <h4>Parameters</h4>
                    <ul class="param-list">
                        <li><span class="param-name">start</span> <span class="param-type">int - Starting value</span></li>
                        <li><span class="param-name">end</span> <span class="param-type">int - Ending value (inclusive)</span></li>
                    </ul>
                </div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
// Using range() function
foreach i in range(1, 5) {
    println(i);            // 1, 2, 3, 4, 5
}

// Using ... operator (equivalent)
foreach i in 1...5 {
    println(i);            // 1, 2, 3, 4, 5
}

// Reverse range
foreach i in range(5, 1) {
    println(i);            // 5, 4, 3, 2, 1
}
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="tostring">
            <div class="function-header">
                <div class="function-name">to_string</div>
                <div class="function-signature">to_string(obj) -> string</div>
            </div>
            <div class="function-body">
                <div class="function-description">
                    Converts any object to its string representation.
                </div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
to_string(123);            // "123"
to_string(3.14);           // "3.14"
to_string(true);           // "true"
to_string([1, 2, 3]);       // "[1, 2, 3]"
to_string(map{"a": 1});    // "map{a: 1}"
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="array">
            <div class="function-header">
                <div class="function-name">array</div>
                <div class="function-signature">array(...args) -> array</div>
            </div>
            <div class="function-body">
                <div class="function-description">
                    Creates a new array from arguments or converts an iterable to an array.
                </div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
array();                          // []
array(1, 2, 3);                   // [1, 2, 3]
array(list(1, 2, 3));             // [1, 2, 3]
array(tuple(1, 2, 3));            // [1, 2, 3]
array(map{"a": 1, "b": 2});       // [1, 2] (values)
array(range(1, 3));               // [1, 2, 3]
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="list">
            <div class="function-header">
                <div class="function-name">list</div>
                <div class="function-signature">list(...args) -> list</div>
            </div>
            <div class="function-body">
                <div class="function-description">
                    Creates a new mutable list from elements.
                </div>
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
        
        <div class="function-card" id="tuple">
            <div class="function-header">
                <div class="function-name">tuple</div>
                <div class="function-signature">tuple(...args) -> tuple</div>
            </div>
            <div class="function-body">
                <div class="function-description">
                    Creates a new immutable tuple from elements.
                </div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
tuple();                          // tuple()
tuple(1, 2, 3);                   // tuple(1, 2, 3)
tuple(10, 20);                    // tuple(10, 20) - for coordinates
tuple("John", 25, true);          // tuple(John, 25, true) - for records
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="addr">
            <div class="function-header">
                <div class="function-name">addr</div>
                <div class="function-signature">addr(obj) -> int</div>
            </div>
            <div class="function-body">
                <div class="function-description">
                    Returns the memory address of an object as an integer. Useful for reference checking.
                </div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var a = [1, 2];
println(addr(a));              // Prints memory address

var b = a;
println(addr(b));              // Same address as a
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="is_same_ref">
            <div class="function-header">
                <div class="function-name">is_same_ref</div>
                <div class="function-signature">is_same_ref(obj1, obj2) -> bool</div>
            </div>
            <div class="function-body">
                <div class="function-description">
                    Checks if two objects point to the same memory address (same reference).
                </div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var a = [1];
var b = a;
var c = [1];

println(is_same_ref(a, b));    // true (same reference)
println(is_same_ref(a, c));    // false (different objects)
{% endhighlight %}
                </div>
            </div>
        </div>
    </div>
</div>
