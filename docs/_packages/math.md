---
layout: default
title: Math Package - Go-Mix
description: Mathematical functions including trigonometry, logarithms, and random numbers
---

<div class="content-page">
    <aside class="sidebar">
        <nav class="sidebar-nav">
            <div class="sidebar-title">Math Functions</div>
            <ul class="sidebar-menu">
                <li><a href="#abs">abs</a></li>
                <li><a href="#min">min/max</a></li>
                <li><a href="#floor">floor/ceil</a></li>
                <li><a href="#sqrt">sqrt</a></li>
                <li><a href="#pow">pow</a></li>
                <li><a href="#sin">sin/cos/tan</a></li>
                <li><a href="#log">log/log10</a></li>
                <li><a href="#rand">rand</a></li>
            </ul>
        </nav>
    </aside>
    
    <div class="content-body">
        <h1>Math Package</h1>
        <p>20 mathematical functions for numerical computing.</p>
        
        <div class="function-card" id="abs">
            <div class="function-header">
                <div class="function-name">abs</div>
                <div class="function-signature">abs(n) -> number</div>
            </div>
            <div class="function-body">
                <div class="function-description">Returns absolute value of a number.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
abs(-42);                  // 42
abs(3.14);                 // 3.14
abs(-3.14);                // 3.14
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="min">
            <div class="function-header">
                <div class="function-name">min / max</div>
                <div class="function-signature">min(a, b) -> number / max(a, b) -> number</div>
            </div>
            <div class="function-body">
                <div class="function-description">Returns minimum or maximum of two numbers.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
min(5, 3);                 // 3
max(5, 3);                 // 5
min(3.14, 2.71);           // 2.71
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="floor">
            <div class="function-header">
                <div class="function-name">floor / ceil / round</div>
                <div class="function-signature">floor(n) -> int / ceil(n) -> int / round(n, [precision]) -> float</div>
            </div>
            <div class="function-body">
                <div class="function-description">Rounding functions.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
floor(3.9);                // 3
floor(-3.1);               // -4
ceil(3.1);                 // 4
ceil(-3.9);                // -3
round(3.5);                // 4
round(3.14159, 2);         // 3.14
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="sqrt">
            <div class="function-header">
                <div class="function-name">sqrt</div>
                <div class="function-signature">sqrt(n) -> float</div>
            </div>
            <div class="function-body">
                <div class="function-description">Returns square root of a number.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
sqrt(16);                  // 4.0
sqrt(2);                   // 1.41421356...
sqrt(2.25);                // 1.5
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="pow">
            <div class="function-header">
                <div class="function-name">pow</div>
                <div class="function-signature">pow(base, exp) -> number</div>
            </div>
            <div class="function-body">
                <div class="function-description">Returns base raised to the power of exp.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
pow(2, 10);                // 1024
pow(3, 3);                 // 27
pow(2, -1);                // 0.5
pow(16, 0.5);              // 4.0 (same as sqrt)
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="sin">
            <div class="function-header">
                <div class="function-name">sin / cos / tan</div>
                <div class="function-signature">sin(rad) -> float / cos(rad) -> float / tan(rad) -> float</div>
            </div>
            <div class="function-body">
                <div class="function-description">Trigonometric functions (angles in radians).</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
sin(0);                    // 0.0
cos(0);                    // 1.0
tan(0);                    // 0.0
sin(3.14159 / 2);          // ~1.0 (π/2)
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="log">
            <div class="function-header">
                <div class="function-name">log / log10 / exp</div>
                <div class="function-signature">log(n) -> float / log10(n) -> float / exp(n) -> float</div>
            </div>
            <div class="function-body">
                <div class="function-description">Logarithmic and exponential functions.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
log(10);                   // 2.302585... (natural log)
log10(100);                // 2.0
log10(1000);               // 3.0
exp(1);                    // 2.71828... (e)
exp(2);                    // 7.38905... (e²)
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="rand">
            <div class="function-header">
                <div class="function-name">rand / rand_int</div>
                <div class="function-signature">rand() -> float / rand_int(min, max) -> int</div>
            </div>
            <div class="function-body">
                <div class="function-description">Random number generation.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
rand();                    // Random float [0.0, 1.0)
rand() * 100;              // Random float [0.0, 100.0)
rand_int(1, 7);            // Random int [1, 6] (dice roll)
rand_int(1, 101);          // Random int [1, 100]
{% endhighlight %}
                </div>
            </div>
        </div>
    </div>
</div>
