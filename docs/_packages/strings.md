---
layout: default
title: Strings Package - Go-Mix
description: String manipulation functions including upper, lower, split, join, and more
---

<div class="content-page">
    <aside class="sidebar">
        <nav class="sidebar-nav">
            <div class="sidebar-title">String Functions</div>
            <ul class="sidebar-menu">
                <li><a href="#length">length</a></li>
                <li><a href="#upper">upper</a></li>
                <li><a href="#lower">lower</a></li>
                <li><a href="#trim">trim</a></li>
                <li><a href="#ltrim">ltrim</a></li>
                <li><a href="#rtrim">rtrim</a></li>
                <li><a href="#split">split</a></li>
                <li><a href="#join">join</a></li>
                <li><a href="#replace">replace</a></li>
                <li><a href="#contains">contains</a></li>
                <li><a href="#index">index</a></li>
                <li><a href="#substring">substring</a></li>
                <li><a href="#reverse">reverse</a></li>
                <li><a href="#ord">ord</a></li>
                <li><a href="#chr">chr</a></li>
            </ul>
        </nav>
    </aside>
    
    <div class="content-body">
        <h1>Strings Package</h1>
        <p>Comprehensive string manipulation with 21 functions.</p>
        
        <div class="function-card" id="length">
            <div class="function-header">
                <div class="function-name">length</div>
                <div class="function-signature">length(str) -> int</div>
            </div>
            <div class="function-body">
                <div class="function-description">Returns string length.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
length("hello");           // 5
length("");                // 0
length("Go-Mix");          // 7
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="upper">
            <div class="function-header">
                <div class="function-name">upper</div>
                <div class="function-signature">upper(str) -> string</div>
            </div>
            <div class="function-body">
                <div class="function-description">Converts to uppercase.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
upper("hello");            // "HELLO"
upper("Go-Mix");           // "GO-MIX"
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
                <div class="function-description">Converts to lowercase.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
lower("HELLO");            // "hello"
lower("Go-Mix");           // "go-mix"
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
                <div class="function-description">Removes whitespace from both ends.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
trim("  hello  ");         // "hello"
trim("\t  world  \n");     // "world"
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="ltrim">
            <div class="function-header">
                <div class="function-name">ltrim</div>
                <div class="function-signature">ltrim(str) -> string</div>
            </div>
            <div class="function-body">
                <div class="function-description">Removes leading whitespace.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
ltrim("  hello  ");        // "hello  "
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="rtrim">
            <div class="function-header">
                <div class="function-name">rtrim</div>
                <div class="function-signature">rtrim(str) -> string</div>
            </div>
            <div class="function-body">
                <div class="function-description">Removes trailing whitespace.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
rtrim("  hello  ");        // "  hello"
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="split">
            <div class="function-header">
                <div class="function-name">split</div>
                <div class="function-signature">split(str, delimiter) -> array</div>
            </div>
            <div class="function-body">
                <div class="function-description">Splits string by delimiter.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
split("a,b,c", ",");       // ["a", "b", "c"]
split("hello world", " "); // ["hello", "world"]
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="join">
            <div class="function-header">
                <div class="function-name">join</div>
                <div class="function-signature">join(array, separator) -> string</div>
            </div>
            <div class="function-body">
                <div class="function-description">Joins array elements with separator.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
join(["a", "b", "c"], ",");  // "a,b,c"
join(["Go", "Mix"], "-");    // "Go-Mix"
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
                <div class="function-description">Replaces all occurrences.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
replace("banana", "a", "o");  // "bonono"
replace("hello", "l", "L");   // "heLLo"
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="contains">
            <div class="function-header">
                <div class="function-name">contains</div>
                <div class="function-signature">contains(str, substr) -> bool</div>
            </div>
            <div class="function-body">
                <div class="function-description">Checks if string contains substring.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
contains("hello", "ell");  // true
contains("hello", "xyz");  // false
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="index">
            <div class="function-header">
                <div class="function-name">index</div>
                <div class="function-signature">index(str, substr) -> int</div>
            </div>
            <div class="function-body">
                <div class="function-description">Returns first index of substring (-1 if not found).</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
index("hello", "e");       // 1
index("hello", "l");       // 2
index("hello", "z");       // -1
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="substring">
            <div class="function-header">
                <div class="function-name">substring</div>
                <div class="function-signature">substring(str, start, [end]) -> string</div>
            </div>
            <div class="function-body">
                <div class="function-description">Extracts substring from start to end.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
substring("hello", 1, 4);  // "ell"
substring("hello", 2);     // "llo" (to end)
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
                <div class="function-description">Reverses string.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
reverse("hello");          // "olleh"
reverse("Go-Mix");        // "xiM-oG"
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="ord">
            <div class="function-header">
                <div class="function-name">ord</div>
                <div class="function-signature">ord(str) -> int</div>
            </div>
            <div class="function-body">
                <div class="function-description">Returns ASCII code of first character.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
ord("A");                  // 65
ord("ABC");                // 65 (first char)
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="chr">
            <div class="function-header">
                <div class="function-name">chr</div>
                <div class="function-signature">chr(code) -> char</div>
            </div>
            <div class="function-body">
                <div class="function-description">Returns character from ASCII code.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
chr(65);                   // 'A'
chr(97);                   // 'a'
{% endhighlight %}
                </div>
            </div>
        </div>
    </div>
</div>
