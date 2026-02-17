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
                <li><a href="#import">Import</a></li>
                <li><a href="#upper">upper()</a></li>
                <li><a href="#lower">lower()</a></li>
                <li><a href="#trim">trim()</a></li>
                <li><a href="#ltrim">ltrim()</a></li>
                <li><a href="#rtrim">rtrim()</a></li>
                <li><a href="#split">split()</a></li>
                <li><a href="#join">join()</a></li>
                <li><a href="#contains_string">contains_string()</a></li>
                <li><a href="#starts_with">starts_with()</a></li>
                <li><a href="#ends_with">ends_with()</a></li>
                <li><a href="#index_string">index_string()</a></li>
                <li><a href="#substring">substring()</a></li>
                <li><a href="#replace_string">replace_string()</a></li>
                <li><a href="#reverse_string">reverse_string()</a></li>
                <li><a href="#count">count()</a></li>
                <li><a href="#ord">ord()</a></li>
                <li><a href="#chr">chr()</a></li>
                <li><a href="#strcmp">strcmp()</a></li>
                <li><a href="#capitalize">capitalize()</a></li>
                <li><a href="#repeat">repeat()</a></li>
                <li><a href="#is_digit">is_digit()</a></li>
                <li><a href="#is_alpha">is_alpha()</a></li>
                <li><a href="#size_string">size_string()</a></li>
                <li><a href="#length_string">length_string()</a></li>
            </ul>
        </nav>
    </aside>
    
    <div class="content-body">
        <h1>Strings Package</h1>
        <p>Comprehensive string manipulation and analysis utilities.</p>
        
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
import strings;
var upper = strings.upper("hello")
var trimmed = strings.trim("  hello  ")

// With alias
import strings as str;
var upper = str.upper("hello")
var trimmed = str.trim("  hello  ")
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

        <div class="function-card" id="ltrim">
            <div class="function-header">
                <div class="function-name">ltrim</div>
                <div class="function-signature">ltrim(str) -> string</div>
            </div>
            <div class="function-body">
                <div class="function-description">Removes leading whitespace from the left side.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
println(ltrim("  hello  ")); // "hello  "
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
                <div class="function-description">Removes trailing whitespace from the right side.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
println(rtrim("  hello  ")); // "  hello"
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

        <div class="function-card" id="contains_string">
            <div class="function-header">
                <div class="function-name">contains_string</div>
                <div class="function-signature">contains_string(str, sub) -> bool</div>
            </div>
            <div class="function-body">
                <div class="function-description">Checks if string contains substring.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
println(contains_string("hello world", "world")); // true
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

        <div class="function-card" id="index_string">
            <div class="function-header">
                <div class="function-name">index_string</div>
                <div class="function-signature">index_string(str, sub) -> int</div>
            </div>
            <div class="function-body">
                <div class="function-description">Finds first occurrence of substring (-1 if not found).</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
println(index_string("hello", "l")); // 2
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

        <div class="function-card" id="replace_string">
            <div class="function-header">
                <div class="function-name">replace_string</div>
                <div class="function-signature">replace_string(str, old, new) -> string</div>
            </div>
            <div class="function-body">
                <div class="function-description">Replaces all occurrences of old substring with new.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
println(replace_string("hello world", "world", "go-mix")); 
// "hello go-mix"
{% endhighlight %}
                </div>
            </div>
        </div>

        <div class="function-card" id="reverse_string">
            <div class="function-header">
                <div class="function-name">reverse_string</div>
                <div class="function-signature">reverse_string(str) -> string</div>
            </div>
            <div class="function-body">
                <div class="function-description">Reverses the characters in the string.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
println(reverse_string("abc")); // "cba"
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

        <div class="function-card" id="ord">
            <div class="function-header">
                <div class="function-name">ord</div>
                <div class="function-signature">ord(char_or_string) -> int</div>
            </div>
            <div class="function-body">
                <div class="function-description">Returns the integer Unicode code point of a character.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
println(ord('A'));   // 65
println(ord("ABC")); // 65 (first character)
{% endhighlight %}
                </div>
            </div>
        </div>

        <div class="function-card" id="chr">
            <div class="function-header">
                <div class="function-name">chr</div>
                <div class="function-signature">chr(integer) -> char</div>
            </div>
            <div class="function-body">
                <div class="function-description">Returns a character from its integer Unicode code point.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
println(chr(65)); // 'A'
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
                <div class="function-description">Compares two strings lexicographically. Returns -1 if s1 < s2, 0 if s1 == s2, 1 if s1 > s2.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
println(strcmp("apple", "banana")); // -1
println(strcmp("hello", "hello"));  // 0
println(strcmp("zoo", "zebra"));    // 1
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
                <div class="function-description">Capitalizes the first character and lowercases the rest.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
println(capitalize("gOMIX")); // "Gomix"
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
                <div class="function-description">Repeats a string n times.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
println(repeat("na", 2)); // "nana"
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
println(is_digit("123")); // true
println(is_digit("12a")); // false
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
                <div class="function-description">Checks if string contains only alphabetic characters.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
println(is_alpha("abc")); // true
println(is_alpha("a12")); // false
{% endhighlight %}
                </div>
            </div>
        </div>

        <div class="function-card" id="size_string">
            <div class="function-header">
                <div class="function-name">size_string</div>
                <div class="function-signature">size_string(str) -> int</div>
            </div>
            <div class="function-body">
                <div class="function-description">Returns the length of a string (number of characters).</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
println(size_string("hello")); // 5
{% endhighlight %}
                </div>
            </div>
        </div>

        <div class="function-card" id="length_string">
            <div class="function-header">
                <div class="function-name">length_string</div>
                <div class="function-signature">length_string(str) -> int</div>
            </div>
            <div class="function-body">
                <div class="function-description">Returns the length of a string (alias for size_string).</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
println(length_string("hello")); // 5
{% endhighlight %}
                </div>
            </div>
        </div>
    </div>
</div>
