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
                <li><a href="#import">Import</a></li>
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
                <li><a href="#starts_with">starts_with</a></li>
                <li><a href="#ends_with">ends_with</a></li>
                <li><a href="#strcmp">strcmp</a></li>
                <li><a href="#capitalize">capitalize</a></li>
                <li><a href="#count">count</a></li>
                <li><a href="#repeat">repeat</a></li>
                <li><a href="#is_digit">is_digit</a></li>
                <li><a href="#is_alpha">is_alpha</a></li>
            </ul>
        </nav>
    </aside>
    
    <div class="content-body">
        <h1>Strings Package</h1>
        <p>Comprehensive string manipulation with 21 functions. All functions are available globally and can also be imported as a package.</p>
        
        <div class="function-card" id="import">
            <div class="function-header">
                <div class="function-name">Import</div>
                <div class="function-signature">import "strings"</div>
            </div>
            <div class="function-body">
                <div class="function-description">Import the strings package to use namespaced functions.</div>
                <div class="function-example">
                    <h4>Examples</h4>
{% highlight go %}
// Standard import
import "strings"
var s = strings.upper("hello")
var t = strings.split("a,b,c", ",")

// With alias
import "strings" as str
var s = str.upper("hello")
var t = str.split("a,b,c", ",")
{% endhighlight %}
                </div>
            </div>
        </div>
        
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
starts_with("hello", "he");  // true
starts_with("hello", "lo");  // false
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
ends_with("hello", "lo");    // true
ends_with("hello", "he");    // false
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="strcmp">
            <div class="function-header">
                <div class="function-name">strcmp</div>
                <div class="function-signature">strcmp(s1, s2) -> int</div>
            </div>
            <div class="function-body">
                <div class="function-description">Compares two strings lexicographically. Returns -1, 0, or 1.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
strcmp("apple", "banana");   // -1
strcmp("hello", "hello");    // 0
strcmp("zebra", "apple");    // 1
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="capitalize">
            <div class="function-header">
                <div class="function-name">capitalize</div>
                <div class="function-signature">capitalize(str) -> string</div>
            </div>
            <div class="function-body">
                <div class="function-description">Capitalizes first letter, lowercases rest.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
capitalize("hello");         // "Hello"
capitalize("HELLO");        // "Hello"
capitalize("gOMIX");        // "Gomix"
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="count">
            <div class="function-header">
                <div class="function-name">count</div>
                <div class="function-signature">count(str, substr) -> int</div>
            </div>
            <div class="function-body">
                <div class="function-description">Counts non-overlapping occurrences of substring.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
count("banana", "a");        // 3
count("hello", "l");         // 2
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="repeat">
            <div class="function-header">
                <div class="function-name">repeat</div>
                <div class="function-signature">repeat(str, count) -> string</div>
            </div>
            <div class="function-body">
                <div class="function-description">Repeats string n times.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
repeat("na", 2);             // "nana"
repeat("hi", 3);             // "hihihi"
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="is_digit">
            <div class="function-header">
                <div class="function-name">is_digit</div>
                <div class="function-signature">is_digit(str) -> bool</div>
            </div>
            <div class="function-body">
                <div class="function-description">Checks if string contains only digits.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
is_digit("12345");           // true
is_digit("12a45");          // false
is_digit("");               // false
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="is_alpha">
            <div class="function-header">
                <div class="function-name">is_alpha</div>
                <div class="function-signature">is_alpha(str) -> bool</div>
            </div>
            <div class="function-body">
                <div class="function-description">Checks if string contains only letters.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
is_alpha("hello");           // true
is_alpha("hello123");       // false
is_alpha("");               // false
{% endhighlight %}
                </div>
            </div>
        </div>
    </div>
</div>
