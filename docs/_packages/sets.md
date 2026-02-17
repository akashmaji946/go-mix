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
                <li><a href="#set">set</a></li>
                <li><a href="#insert_set">insert_set</a></li>
                <li><a href="#remove_set">remove_set</a></li>
                <li><a href="#contains_set">contains_set</a></li>
                <li><a href="#size_set">size_set</a></li>
                <li><a href="#union_set">union_set</a></li>
                <li><a href="#intersection_set">intersection_set</a></li>
                <li><a href="#difference_set">difference_set</a></li>
            </ul>
        </nav>
    </aside>
    
    <div class="content-body">
        <h1>Sets Package</h1>
        <p>Unique value collections with set operations.</p>
        
        <div class="function-card" id="set">
            <div class="function-header">
                <div class="function-name">set</div>
                <div class="function-signature">set(...elements) -> set</div>
            </div>
            <div class="function-body">
                <div class="function-description">Creates a new set with unique elements.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
set(1, 2, 3);                     // set{1, 2, 3}
set(1, 2, 2, 3, 3, 3);            // set{1, 2, 3} (duplicates removed)
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="insert_set">
            <div class="function-header">
                <div class="function-name">insert_set</div>
                <div class="function-signature">insert_set(s, elem) -> set</div>
            </div>
            <div class="function-body">
                <div class="function-description">Adds element to set (if not already present).</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var s = set(1, 2, 3);
insert_set(s, 4);          // s = set{1, 2, 3, 4}
insert_set(s, 2);          // s unchanged (2 already exists)
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
                <div class="function-description">Removes element from set. Returns true if element existed.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var s = set(1, 2, 3);
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
var s = set(1, 2, 3);
contains_set(s, 2);        // true
contains_set(s, 99);       // false
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
                <div class="function-description">Returns number of unique elements in set.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var s = set(1, 2, 3, 3, 3);
size_set(s);               // 3
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="union_set">
            <div class="function-header">
                <div class="function-name">union_set</div>
                <div class="function-signature">union_set(s1, s2) -> set</div>
            </div>
            <div class="function-body">
                <div class="function-description">Returns union of two sets (all unique elements from both).</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var s1 = set(1, 2, 3);
var s2 = set(3, 4, 5);
union_set(s1, s2);         // set{1, 2, 3, 4, 5}
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="intersection_set">
            <div class="function-header">
                <div class="function-name">intersection_set</div>
                <div class="function-signature">intersection_set(s1, s2) -> set</div>
            </div>
            <div class="function-body">
                <div class="function-description">Returns intersection of two sets (elements common to both).</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var s1 = set(1, 2, 3);
var s2 = set(3, 4, 5);
intersection_set(s1, s2);  // set{3}
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="difference_set">
            <div class="function-header">
                <div class="function-name">difference_set</div>
                <div class="function-signature">difference_set(s1, s2) -> set</div>
            </div>
            <div class="function-body">
                <div class="function-description">Returns difference of two sets (elements in s1 but not in s2).</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var s1 = set(1, 2, 3);
var s2 = set(3, 4, 5);
difference_set(s1, s2);  // set{1, 2}
{% endhighlight %}
                </div>
            </div>
        </div>
    </div>
</div>
