---
layout: default
title: Format Package - Go-Mix
description: Type conversion functions including to_int, to_float, to_string, to_bool
---

<div class="content-page">
    <aside class="sidebar">
        <nav class="sidebar-nav">
            <div class="sidebar-title">Format Functions</div>
            <ul class="sidebar-menu">
                <li><a href="#to_int">to_int</a></li>
                <li><a href="#to_float">to_float</a></li>
                <li><a href="#to_string">to_string</a></li>
                <li><a href="#to_bool">to_bool</a></li>
                <li><a href="#to_char">to_char</a></li>
            </ul>
        </nav>
    </aside>
    
    <div class="content-body">
        <h1>Format Package</h1>
        <p>Type conversion and formatting functions.</p>
        
        <div class="function-card" id="to_int">
            <div class="function-header">
                <div class="function-name">to_int</div>
                <div class="function-signature">to_int(value) -> int</div>
            </div>
            <div class="function-body">
                <div class="function-description">Converts value to integer.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
to_int("42");              // 42
to_int("0xFF");            // 255 (hex)
to_int("0o77");            // 63 (octal)
to_int(3.14);              // 3 (truncated)
to_int(true);              // 1
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="to_float">
            <div class="function-header">
                <div class="function-name">to_float</div>
                <div class="function-signature">to_float(value) -> float</div>
            </div>
            <div class="function-body">
                <div class="function-description">Converts value to float.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
to_float("3.14");          // 3.14
to_float("2.5e10");        // 25000000000.0
to_float(42);              // 42.0
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="to_string">
            <div class="function-header">
                <div class="function-name">to_string</div>
                <div class="function-signature">to_string(value) -> string</div>
            </div>
            <div class="function-body">
                <div class="function-description">Converts value to string.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
to_string(42);             // "42"
to_string(3.14);           // "3.14"
to_string(true);           // "true"
to_string([1, 2, 3]);      // "[1, 2, 3]"
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="to_bool">
            <div class="function-header">
                <div class="function-name">to_bool</div>
                <div class="function-signature">to_bool(value) -> bool</div>
            </div>
            <div class="function-body">
                <div class="function-description">Converts value to boolean.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
to_bool(1);                // true
to_bool(0);                // false
to_bool("true");           // true
to_bool("false");          // false
to_bool(nil);              // false
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="to_char">
            <div class="function-header">
                <div class="function-name">to_char</div>
                <div class="function-signature">to_char(value) -> char</div>
            </div>
            <div class="function-body">
                <div class="function-description">Converts value to character.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
to_char(65);               // 'A'
to_char("hello");          // 'h' (first char)
{% endhighlight %}
                </div>
            </div>
        </div>
    </div>
</div>
