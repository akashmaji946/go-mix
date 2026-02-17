---
layout: default
title: Time Package - Go-Mix
description: Time handling functions including now, format_time, parse_time, timezone
---

<div class="content-page">
    <aside class="sidebar">
        <nav class="sidebar-nav">
            <div class="sidebar-title">Time Functions</div>
            <ul class="sidebar-menu">
                <li><a href="#import">Import</a></li>
                <li><a href="#now">now()</a></li>
                <li><a href="#now_ms">now_ms()</a></li>
                <li><a href="#utc_now">utc_now()</a></li>
                <li><a href="#format_time">format_time()</a></li>
                <li><a href="#parse_time">parse_time()</a></li>
                <li><a href="#timezone">timezone()</a></li>
                <li><a href="#sleep">sleep()</a></li>
            </ul>
        </nav>
    </aside>
    
    <div class="content-body">
        <h1>Time Package</h1>
        <p>Time handling and formatting functions.</p>
        
        <div class="function-card" id="import">
            <div class="function-header">
                <div class="function-name">Import</div>
                <div class="function-signature">import "time"</div>
            </div>
            <div class="function-body">
                <div class="function-description">Import the time package to use namespaced functions.</div>
                <div class="function-example">
                    <h4>Examples</h4>
{% highlight go %}
// Standard import
import time;
var now = time.now();
var formatted = time.format_time(now, "2006-01-02");

// With alias
import time as t;
var now = t.now();
var formatted = t.format_time(now, "2006-01-02");
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="now">
            <div class="function-header">
                <div class="function-name">now</div>
                <div class="function-signature">now() -> int</div>
            </div>
            <div class="function-body">
                <div class="function-description">Returns current Unix timestamp in seconds.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var timestamp = now();     // e.g., 1707542400
println(timestamp);
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="now_ms">
            <div class="function-header">
                <div class="function-name">now_ms</div>
                <div class="function-signature">now_ms() -> int</div>
            </div>
            <div class="function-body">
                <div class="function-description">Returns current Unix timestamp in milliseconds.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var timestamp = now_ms();  // e.g., 1707542400000
println(timestamp);
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="utc_now">
            <div class="function-header">
                <div class="function-name">utc_now</div>
                <div class="function-signature">utc_now() -> int</div>
            </div>
            <div class="function-body">
                <div class="function-description">Returns current UTC Unix timestamp in seconds.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var utc = utc_now();       // UTC timestamp
println(utc);
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="format_time">
            <div class="function-header">
                <div class="function-name">format_time</div>
                <div class="function-signature">format_time(timestamp, layout) -> string</div>
            </div>
            <div class="function-body">
                <div class="function-description">Formats timestamp according to layout string.</div>
                <div class="function-params">
                    <h4>Common Layouts</h4>
                    <ul class="param-list">
                        <li><span class="param-name">"2006-01-02"</span> <span class="param-type">Date only</span></li>
                        <li><span class="param-name">"15:04:05"</span> <span class="param-type">Time only</span></li>
                        <li><span class="param-name">"2006-01-02 15:04:05"</span> <span class="param-type">Date and time</span></li>
                        <li><span class="param-name">"Jan 02, 2006"</span> <span class="param-type">Pretty date</span></li>
                    </ul>
                </div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var t = now();
format_time(t, "2006-01-02");           // "2024-02-10"
format_time(t, "15:04:05");             // "12:00:00"
format_time(t, "2006-01-02 15:04:05"); // "2024-02-10 12:00:00"
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="parse_time">
            <div class="function-header">
                <div class="function-name">parse_time</div>
                <div class="function-signature">parse_time(value, layout) -> int</div>
            </div>
            <div class="function-body">
                <div class="function-description">Parses time string and returns Unix timestamp.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var ts = parse_time("2024-02-10", "2006-01-02");
println(ts);               // Unix timestamp
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="timezone">
            <div class="function-header">
                <div class="function-name">timezone</div>
                <div class="function-signature">timezone() -> string</div>
            </div>
            <div class="function-body">
                <div class="function-description">Returns current timezone name.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var tz = timezone();       // e.g., "IST", "UTC", "PST"
println(tz);
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="sleep">
            <div class="function-header">
                <div class="function-name">sleep</div>
                <div class="function-signature">sleep(milliseconds) -> nil</div>
            </div>
            <div class="function-body">
                <div class="function-description">Pauses execution for specified milliseconds.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
println("Starting...");
sleep(1000);               // Sleep 1 second
println("Done!");
{% endhighlight %}
                </div>
            </div>
        </div>
    </div>
</div>
