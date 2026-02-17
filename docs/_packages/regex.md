---
layout: default
title: Regex Package - Go-Mix
description: Regular expression functions for pattern matching
---

<div class="content-page">
    <aside class="sidebar">
        <nav class="sidebar-nav">
            <div class="sidebar-title">Regex Functions</div>
            <ul class="sidebar-menu">
                <li><a href="#import">Import</a></li>
                <li><a href="#match_regex">match_regex()</a></li>
                <li><a href="#find_regex">find_regex()</a></li>
                <li><a href="#findall_regex">findall_regex()</a></li>
                <li><a href="#replace_regex">replace_regex()</a></li>
                <li><a href="#split_regex">split_regex()</a></li>
            </ul>
        </nav>
    </aside>
    
    <div class="content-body">
        <h1>Regex Package</h1>
        <p>Regular expression pattern matching and manipulation.</p>
        
        <div class="function-card" id="import">
            <div class="function-header">
                <div class="function-name">Import</div>
                <div class="function-signature">import "regex"</div>
            </div>
            <div class="function-body">
                <div class="function-description">Import the regex package to use namespaced functions.</div>
                <div class="function-example">
                    <h4>Examples</h4>
{% highlight go %}
// Standard import
import regex;
var match = regex.match_regex("\\d+", "123")
var found = regex.find_regex("\\d+", "abc123")

// With alias
import regex as re;
var match = re.match_regex("\\d+", "123")
var found = re.find_regex("\\d+", "abc123")
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="match_regex">
            <div class="function-header">
                <div class="function-name">match_regex</div>
                <div class="function-signature">match_regex(pattern, str) -> bool</div>
            </div>
            <div class="function-body">
                <div class="function-description">Checks if string matches pattern.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
match_regex("\\d+", "123");        // true
match_regex("^[a-z]+$", "hello");  // true
match_regex("^[a-z]+$", "Hello"); // false (uppercase)
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="find_regex">
            <div class="function-header">
                <div class="function-name">find_regex</div>
                <div class="function-signature">find_regex(pattern, str) -> string</div>
            </div>
            <div class="function-body">
                <div class="function-description">Returns first match or empty string.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
find_regex("\\d+", "abc123def");   // "123"
find_regex("[a-z]+", "ABCdef");    // "def"
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="findall_regex">
            <div class="function-header">
                <div class="function-name">findall_regex</div>
                <div class="function-signature">findall_regex(pattern, str, [n]) -> array</div>
            </div>
            <div class="function-body">
                <div class="function-description">Returns all matches as array.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
findall_regex("\\d+", "a1b2c3");   // ["1", "2", "3"]
findall_regex("[a-z]+", "abc def", 1); // ["abc"] (limit 1)
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="replace_regex">
            <div class="function-header">
                <div class="function-name">replace_regex</div>
                <div class="function-signature">replace_regex(pattern, str, replacement) -> string</div>
            </div>
            <div class="function-body">
                <div class="function-description">Replaces all matches with replacement.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
replace_regex("\\d", "a1b2c3", "X"); // "aXbXcX"
replace_regex(" +", "hello   world", " "); // "hello world"
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="split_regex">
            <div class="function-header">
                <div class="function-name">split_regex</div>
                <div class="function-signature">split_regex(pattern, str, [n]) -> array</div>
            </div>
            <div class="function-body">
                <div class="function-description">Splits string by pattern.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
split_regex("\\s+", "hello   world"); // ["hello", "world"]
split_regex(",", "a,b,c", 2);         // ["a", "b,c"] (limit 2)
{% endhighlight %}
                </div>
            </div>
        </div>
    </div>
</div>
