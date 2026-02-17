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
                <li><a href="#push">push</a></li>
                <li><a href="#pop">pop</a></li>
                <li><a href="#shift">shift</a></li>
                <li><a href="#unshift">unshift</a></li>
                <li><a href="#sort">sort</a></li>
                <li><a href="#sorted">sorted</a></li>
                <li><a href="#map">map</a></li>
                <li><a href="#filter">filter</a></li>
                <li><a href="#reduce">reduce</a></li>
                <li><a href="#find">find</a></li>
                <li><a href="#some">some</a></li>
                <li><a href="#every">every</a></li>
                <li><a href="#clone">clone</a></li>
                <li><a href="#reverse">reverse</a></li>
            </ul>
        </nav>
    </aside>
    
    <div class="content-body">
        <h1>Arrays Package</h1>
        <p>Comprehensive array manipulation with functional programming support.</p>
        
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
push(arr, 5, 6);           // arr is now [1, 2, 3, 4, 5, 6]
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

var squared = map(nums, fn(x) { return x * x; });
// squared = [1, 4, 9, 16, 25]
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

var big = filter(nums, fn(x) { return x > 3; });
// big = [4, 5, 6]
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

var product = reduce(nums, fn(acc, x) { return acc * x; }, 1);
// product = 120
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

var firstBig = find(nums, fn(x) { return x > 10; });
// firstBig = nil
{% endhighlight %}
                </div>
            </div>
        </div>
    </div>
</div>
