---
layout: default
title: Arrays Package - Go-Mix
description: Array manipulation functions including push, pop, sort, map, filter, reduce
---

<div class="content-page">
    <aside class="sidebar">
        <nav class="sidebar-nav">
            <div class="sidebar-title">Array Functions</div>
            <ul class="sidebar-menu">
                <li><a href="#import">Import</a></li>
                <li><a href="#make_array">make_array</a></li>
                <li><a href="#push">push</a></li>
                <li><a href="#pop">pop</a></li>
                <li><a href="#shift">shift</a></li>
                <li><a href="#unshift">unshift</a></li>
                <li><a href="#sort">sort</a></li>
                <li><a href="#sorted">sorted</a></li>
                <li><a href="#csort">csort</a></li>
                <li><a href="#csorted">csorted</a></li>
                <li><a href="#map">map</a></li>
                <li><a href="#filter">filter</a></li>
                <li><a href="#reduce">reduce</a></li>
                <li><a href="#find">find</a></li>
                <li><a href="#some">some</a></li>
                <li><a href="#every">every</a></li>
                <li><a href="#clone">clone</a></li>
                <li><a href="#reverse">reverse</a></li>
                <li><a href="#contains">contains</a></li>
                <li><a href="#replace">replace</a></li>
                <li><a href="#index">index</a></li>
                <li><a href="#size">size</a></li>
                <li><a href="#to_array">to_array</a></li>
            </ul>
        </nav>
    </aside>
    
    <div class="content-body">
        <h1>Arrays Package</h1>
        <p>Comprehensive array manipulation with functional programming support. All functions are available globally and can also be imported as a package.</p>
        
        <div class="function-card" id="import">
            <div class="function-header">
                <div class="function-name">Import</div>
                <div class="function-signature">import "arrays"</div>
            </div>
            <div class="function-body">
                <div class="function-description">Import the arrays package to use namespaced functions.</div>
                <div class="function-example">
                    <h4>Examples</h4>
{% highlight go %}
// Standard import
import "arrays"
var arr = arrays.make_array(1, 2, 3)
arrays.push_array(arr, 4)

// With alias
import "arrays" as arr
var a = arr.make_array(1, 2, 3)
arr.push_array(a, 4)
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="make_array">
            <div class="function-header">
                <div class="function-name">make_array</div>
                <div class="function-signature">make_array(...) -> array</div>
            </div>
            <div class="function-body">
                <div class="function-description">Creates a new array from arguments or converts an iterable to an array.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
make_array()                    // []
make_array(1, 2, 3)             // [1, 2, 3]
make_array([1, 2, 3])           // [1, 2, 3] (copy)
make_array(list(1, 2, 3))       // [1, 2, 3]
make_array(tuple(1, 2, 3))      // [1, 2, 3]
make_array(set{1, 2, 3})        // [1, 2, 3]
make_array(map{"a": 1, "b": 2}) // [1, 2] (values)
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="push">
            <div class="function-header">
                <div class="function-name">push</div>
                <div class="function-signature">push(arr, elem) -> array</div>
            </div>
            <div class="function-body">
                <div class="function-description">Adds an element to the end of an array.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var arr = [1, 2, 3];
push(arr, 4);              // arr is now [1, 2, 3, 4]
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="pop">
            <div class="function-header">
                <div class="function-name">pop</div>
                <div class="function-signature">pop(arr) -> any</div>
            </div>
            <div class="function-body">
                <div class="function-description">Removes and returns the last element of an array.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var arr = [1, 2, 3, 4];
var last = pop(arr);       // last = 4, arr = [1, 2, 3]
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="shift">
            <div class="function-header">
                <div class="function-name">shift</div>
                <div class="function-signature">shift(arr) -> any</div>
            </div>
            <div class="function-body">
                <div class="function-description">Removes and returns the first element of an array.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var arr = [1, 2, 3, 4];
var first = shift(arr);    // first = 1, arr = [2, 3, 4]
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="unshift">
            <div class="function-header">
                <div class="function-name">unshift</div>
                <div class="function-signature">unshift(arr, elem) -> array</div>
            </div>
            <div class="function-body">
                <div class="function-description">Adds an element to the beginning of an array.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var arr = [2, 3, 4];
unshift(arr, 1);           // arr is now [1, 2, 3, 4]
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="sort">
            <div class="function-header">
                <div class="function-name">sort</div>
                <div class="function-signature">sort(arr, [reverse]) -> array</div>
            </div>
            <div class="function-body">
                <div class="function-description">Sorts array elements in-place.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var arr = [3, 1, 4, 1, 5];
sort(arr);                 // arr = [1, 1, 3, 4, 5]
sort(arr, true);            // arr = [5, 4, 3, 1, 1] (reverse)
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="sorted">
            <div class="function-header">
                <div class="function-name">sorted</div>
                <div class="function-signature">sorted(arr, [reverse]) -> array</div>
            </div>
            <div class="function-body">
                <div class="function-description">Returns a new sorted array without modifying the original.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var arr = [3, 1, 4, 1, 5];
var s = sorted(arr);       // s = [1, 1, 3, 4, 5], arr unchanged
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="csort">
            <div class="function-header">
                <div class="function-name">csort</div>
                <div class="function-signature">csort(arr, comparator) -> array</div>
            </div>
            <div class="function-body">
                <div class="function-description">Sorts array in-place using a custom comparator function.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var arr = [3, 1, 4, 1, 5];
csort(arr, fn(a, b) { return a > b; }); // Descending sort
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="csorted">
            <div class="function-header">
                <div class="function-name">csorted</div>
                <div class="function-signature">csorted(arr, comparator) -> array</div>
            </div>
            <div class="function-body">
                <div class="function-description">Returns a new sorted array using a custom comparator function.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var arr = [3, 1, 4, 1, 5];
var sorted = csorted(arr, fn(a, b) { return a < b; });
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="map">
            <div class="function-header">
                <div class="function-name">map</div>
                <div class="function-signature">map(arr, fn) -> array</div>
            </div>
            <div class="function-body">
                <div class="function-description">Applies a function to each element and returns new array.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var nums = [1, 2, 3, 4, 5];
var doubled = map(nums, fn(x) { return x * 2; });
// doubled = [2, 4, 6, 8, 10]
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="filter">
            <div class="function-header">
                <div class="function-name">filter</div>
                <div class="function-signature">filter(arr, fn) -> array</div>
            </div>
            <div class="function-body">
                <div class="function-description">Returns elements that satisfy the predicate function.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var nums = [1, 2, 3, 4, 5, 6];
var evens = filter(nums, fn(x) { return x % 2 == 0; });
// evens = [2, 4, 6]
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="reduce">
            <div class="function-header">
                <div class="function-name">reduce</div>
                <div class="function-signature">reduce(arr, fn, initial) -> any</div>
            </div>
            <div class="function-body">
                <div class="function-description">Reduces array to single value using accumulator function.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var nums = [1, 2, 3, 4, 5];
var sum = reduce(nums, fn(acc, x) { return acc + x; }, 0);
// sum = 15
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="find">
            <div class="function-header">
                <div class="function-name">find</div>
                <div class="function-signature">find(arr, fn) -> any</div>
            </div>
            <div class="function-body">
                <div class="function-description">Returns first element matching predicate, or nil if none found.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var nums = [1, 2, 3, 4, 5];
var firstEven = find(nums, fn(x) { return x % 2 == 0; });
// firstEven = 2
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="some">
            <div class="function-header">
                <div class="function-name">some</div>
                <div class="function-signature">some(arr, fn) -> bool</div>
            </div>
            <div class="function-body">
                <div class="function-description">Returns true if at least one element satisfies the predicate.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var nums = [1, 2, 3, 4, 5];
some(nums, fn(x) { return x > 3; });  // true
some(nums, fn(x) { return x > 10; }); // false
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="every">
            <div class="function-header">
                <div class="function-name">every</div>
                <div class="function-signature">every(arr, fn) -> bool</div>
            </div>
            <div class="function-body">
                <div class="function-description">Returns true if all elements satisfy the predicate.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var nums = [2, 4, 6, 8];
every(nums, fn(x) { return x % 2 == 0; });  // true
every(nums, fn(x) { return x > 5; });       // false
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="clone">
            <div class="function-header">
                <div class="function-name">clone</div>
                <div class="function-signature">clone(arr) -> array</div>
            </div>
            <div class="function-body">
                <div class="function-description">Returns a shallow copy of the array.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var arr = [1, 2, 3];
var copy = clone(arr);     // copy = [1, 2, 3], independent of arr
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="reverse">
            <div class="function-header">
                <div class="function-name">reverse</div>
                <div class="function-signature">reverse(arr) -> array</div>
            </div>
            <div class="function-body">
                <div class="function-description">Returns a new array with elements in reverse order.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var arr = [1, 2, 3, 4, 5];
var rev = reverse(arr);    // rev = [5, 4, 3, 2, 1]
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="contains">
            <div class="function-header">
                <div class="function-name">contains</div>
                <div class="function-signature">contains(arr, value) -> bool</div>
            </div>
            <div class="function-body">
                <div class="function-description">Checks if a value exists in the array.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var arr = [1, 2, 3, 4, 5];
contains(arr, 3);          // true
contains(arr, 10);         // false
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="replace">
            <div class="function-header">
                <div class="function-name">replace</div>
                <div class="function-signature">replace(arr, old_val, new_val) -> int</div>
            </div>
            <div class="function-body">
                <div class="function-description">Replaces first occurrence of old_val with new_val. Returns index or -1.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var arr = [1, 2, 3];
replace(arr, 2, 42);       // arr = [1, 42, 3], returns 1
replace(arr, 99, 100);     // returns -1 (not found)
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="index">
            <div class="function-header">
                <div class="function-name">index</div>
                <div class="function-signature">index(arr, value) -> int</div>
            </div>
            <div class="function-body">
                <div class="function-description">Returns index of first occurrence of value, or -1 if not found.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var arr = [1, 2, 3, 2, 1];
index(arr, 2);             // 1
index(arr, 99);            // -1
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="size">
            <div class="function-header">
                <div class="function-name">size</div>
                <div class="function-signature">size(arr) -> int</div>
            </div>
            <div class="function-body">
                <div class="function-description">Returns the number of elements in the array. Alias: length.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var arr = [1, 2, 3, 4, 5];
size(arr);                 // 5
length(arr);               // 5 (alias)
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="to_array">
            <div class="function-header">
                <div class="function-name">to_array</div>
                <div class="function-signature">to_array(iterable) -> array</div>
            </div>
            <div class="function-body">
                <div class="function-description">Converts a list or tuple to an array.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var l = list(1, 2, 3);
var t = tuple(4, 5, 6);
to_array(l);               // [1, 2, 3]
to_array(t);               // [4, 5, 6]
{% endhighlight %}
                </div>
            </div>
        </div>
    </div>
</div>
