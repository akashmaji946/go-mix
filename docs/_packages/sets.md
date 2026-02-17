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
                <li><a href="#import">Import</a></li>
                <li><a href="#make_set">make_set()</a></li>
                <li><a href="#insert_set">insert_set()</a></li>
                <li><a href="#remove_set">remove_set()</a></li>
                <li><a href="#contains_set">contains_set()</a></li>
                <li><a href="#values_set">values_set()</a></li>
                <li><a href="#size_set">size_set()</a></li>
                <li><a href="#length_set">length_set()</a></li>
            </ul>

        </nav>
    </aside>
    
    <div class="content-body">
        <h1>Sets Package</h1>
        <p>Unique value collections with set operations for efficient membership testing and deduplication.</p>
        
        <div class="function-card" id="import">
            <div class="function-header">
                <div class="function-name">Import</div>
                <div class="function-signature">import "sets"</div>
            </div>
            <div class="function-body">
                <div class="function-description">Import the sets package to use namespaced functions. Sets are mutable collections that store unique values, automatically removing duplicates.</div>
                <div class="function-example">
                    <h4>Examples</h4>
{% highlight go %}
// Standard import
import sets;
var s = sets.make_set(1, 2, 3)
sets.insert_set(s, 4)

// With alias
import sets as st;
var s = st.make_set(1, 2, 3);
st.insert_set(s, 4);
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="make_set">
            <div class="function-header">
                <div class="function-name">make_set</div>
                <div class="function-signature">make_set(...elements) -> set</div>
            </div>
            <div class="function-body">
                <div class="function-description">Creates a new set from the provided elements. Duplicate values are automatically removed. Sets maintain insertion order for values.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
make_set();                           // set{}
make_set(1, 2, 3);                    // set{1, 2, 3}
make_set(1, 2, 2, 3, 3, 3);           // set{1, 2, 3} (duplicates removed)
make_set("a", 1, true);               // set{a, 1, true} (heterogeneous values)
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="insert_set">
            <div class="function-header">
                <div class="function-name">insert_set</div>
                <div class="function-signature">insert_set(s, elem) -> any</div>
            </div>
            <div class="function-body">
                <div class="function-description">Adds an element to the set if it doesn't already exist. Sets only store unique values, so inserting a duplicate has no effect. Returns the inserted element.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var s = make_set(1, 2, 3);
insert_set(s, 4);          // Returns 4, s = set{1, 2, 3, 4}
insert_set(s, 2);          // Returns 2, s unchanged (2 already exists)
insert_set(s, "hello");    // Returns "hello", s = set{1, 2, 3, 4, hello}
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
                <div class="function-description">Removes an element from the set. Returns true if the element was found and removed, false if the element didn't exist in the set.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var s = make_set(1, 2, 3);
remove_set(s, 2);          // true, s = set{1, 3}
remove_set(s, 99);         // false (element didn't exist)
remove_set(s, 1);          // true, s = set{3}
remove_set(s, 3);          // true, s = set{}
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
                <div class="function-description">Checks if the set contains a specific element. Returns true if the element exists in the set, false otherwise. Uses O(1) lookup time.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var s = make_set(1, 2, 3);
contains_set(s, 2);        // true
contains_set(s, 99);         // false
contains_set(s, "2");      // false (type matters: string vs int)
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
                <div class="function-description">Returns an array containing all values in the set, preserving insertion order. Useful for iterating over set contents or converting to other data structures.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var s = make_set(3, 1, 2);
values_set(s);             // [3, 1, 2] (insertion order preserved)

var empty = make_set();
values_set(empty);         // [] (empty array for empty set)
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
                <div class="function-description">Returns the number of unique elements in the set. This is an O(1) operation that returns the current cardinality of the set. Alias: length_set.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var s = make_set(1, 2, 3, 3, 3);
size_set(s);               // 3 (duplicates not counted)

var empty = make_set();
size_set(empty);           // 0

// After modifications
insert_set(s, 4);
size_set(s);               // 4
remove_set(s, 1);
size_set(s);               // 3
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="length_set">
            <div class="function-header">
                <div class="function-name">length_set</div>
                <div class="function-signature">length_set(s) -> int</div>
            </div>
            <div class="function-body">
                <div class="function-description">Returns the number of unique elements in the set. This is an alias for size_set() and provides the same functionality for those who prefer the length terminology.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var s = make_set(1, 2, 3);
length_set(s);             // 3

// Equivalent to size_set
size_set(s) == length_set(s);  // true

var empty = make_set();
length_set(empty);         // 0
{% endhighlight %}
                </div>
            </div>
        </div>
    </div>
</div>
