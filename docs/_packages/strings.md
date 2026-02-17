---
layout: default
title: Strings Package - Go-Mix
description: String manipulation functions including split, join, replace, and regex
---

<div class="content-page">
    <aside class="sidebar">
        <nav class="sidebar-nav">
            <div class="sidebar-title">String Functions</div>
            <ul class="sidebar-menu">
                <li><a href="#upper">upper</a></li>
                <li><a href="#lower">lower</a></li>
                <li><a href="#trim">trim</a></li>
                <li><a href="#split">split</a></li>
                <li><a href="#join">join</a></li>
                <li><a href="#contains">contains</a></li>
                <li><a href="#starts_with">starts_with</a></li>
                <li><a href="#ends_with">ends_with</a></li>
                <li><a href="#index">index</a></li>
                <li><a href="#substring">substring</a></li>
                <li><a href="#replace">replace</a></li>
                <li><a href="#reverse">reverse</a></li>
                <li><a href="#count">count</a></li>
                <li><a href="#is_digit">is_digit</a></li>
                <li><a href="#is_alpha">is_alpha</a></li>
            </ul>
        </nav>
    </aside>
    
    <div class="content-body">
        <h1>Strings Package</h1>
        <p>Comprehensive string manipulation and analysis utilities.</p>
        
        <div class="function-card" id="upper">
            <div class="function-header">
                <div class="function-name">upper</div>
                <div class="function-signature">upper(str) -> string</div>
            </div>
            <div class="function-body">
                <div class="function-description">Converts string to uppercase.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
println(upper("hello")); // "HELLO"
{% endhighlight %}
                </div>
            </div>
        </div>

        <div class="function-card" id="lower">
            <div class="function-header">
                <div class="function-name">lower</div>
                <div class="function-signature">lower(str) -> string</div>
            </div>
            <div class="function-body">
                <div class="function-description">Converts string to lowercase.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
println(lower("HELLO")); // "hello"
{% endhighlight %}
                </div>
            </div>
        </div>

        <div class="function-card" id="trim">
            <div class="function-header">
                <div class="function-name">trim</div>
                <div class="function-signature">trim(str) -> string</div>
            </div>
            <div class="function-body">
                <div class="function-description">Removes leading and trailing whitespace.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
println(trim("  hello  ")); // "hello"
{% endhighlight %}
                </div>
            </div>
        </div>

        <div class="function-card" id="split">
            <div class="function-header">
                <div class="function-name">split</div>
                <div class="function-signature">split(str, sep) -> array</div>
            </div>
            <div class="function-body">
                <div class="function-description">Splits string by separator into an array.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var parts = split("a,b,c", ","); 
// ["a", "b", "c"]
{% endhighlight %}
                </div>
            </div>
        </div>

        <div class="function-card" id="join">
            <div class="function-header">
                <div class="function-name">join</div>
                <div class="function-signature">join(arr, sep) -> string</div>
            </div>
            <div class="function-body">
                <div class="function-description">Joins array elements with separator.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var s = join(["a", "b", "c"], "-"); 
// "a-b-c"
{% endhighlight %}
                </div>
            </div>
        </div>

        <div class="function-card" id="contains">
            <div class="function-header">
                <div class="function-name">contains</div>
                <div class="function-signature">contains(str, sub) -> bool</div>
            </div>
            <div class="function-body">
                <div class="function-description">Checks if string contains substring.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
println(contains("hello world", "world")); // true
{% endhighlight %}
                </div>
            </div>
        </div>

        <div class="function-card" id="starts_with">
            <div class="function-header">
                <div class="function-name">starts_with</div>
                <div class="function-signature">starts_with(str, prefix) -> bool</div>
            </div>
            <div class="function-body">
                <div class="function-description">Checks if string starts with prefix.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
println(starts_with("go-mix", "go")); // true
{% endhighlight %}
                </div>
            </div>
        </div>

        <div class="function-card" id="ends_with">
            <div class="function-header">
                <div class="function-name">ends_with</div>
                <div class="function-signature">ends_with(str, suffix) -> bool</div>
            </div>
            <div class="function-body">
                <div class="function-description">Checks if string ends with suffix.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
println(ends_with("file.gm", ".gm")); // true
{% endhighlight %}
                </div>
            </div>
        </div>

        <div class="function-card" id="index">
            <div class="function-header">
                <div class="function-name">index</div>
                <div class="function-signature">index(str, sub) -> int</div>
            </div>
            <div class="function-body">
                <div class="function-description">Finds first occurrence of substring (-1 if not found).</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
println(index("hello", "l")); // 2
{% endhighlight %}
                </div>
            </div>
        </div>

        <div class="function-card" id="substring">
            <div class="function-header">
                <div class="function-name">substring</div>
                <div class="function-signature">substring(str, start, length) -> string</div>
            </div>
            <div class="function-body">
                <div class="function-description">Extracts substring starting at index with given length.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
println(substring("hello world", 0, 5)); // "hello"
{% endhighlight %}
                </div>
            </div>
        </div>

        <div class="function-card" id="replace">
            <div class="function-header">
                <div class="function-name">replace</div>
                <div class="function-signature">replace(str, old, new) -> string</div>
            </div>
            <div class="function-body">
                <div class="function-description">Replaces all occurrences of old substring with new.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
println(replace("hello world", "world", "go-mix")); 
// "hello go-mix"
{% endhighlight %}
                </div>
            </div>
        </div>

        <div class="function-card" id="reverse">
            <div class="function-header">
                <div class="function-name">reverse</div>
                <div class="function-signature">reverse(str) -> string</div>
            </div>
            <div class="function-body">
                <div class="function-description">Reverses the characters in the string.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
println(reverse("abc")); // "cba"
{% endhighlight %}
                </div>
            </div>
        </div>

        <div class="function-card" id="count">
            <div class="function-header">
                <div class="function-name">count</div>
                <div class="function-signature">count(str, sub) -> int</div>
            </div>
            <div class="function-body">
                <div class="function-description">Counts occurrences of substring.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
println(count("banana", "a")); // 3
{% endhighlight %}
                </div>
            </div>
        </div>
    </div>
</div>