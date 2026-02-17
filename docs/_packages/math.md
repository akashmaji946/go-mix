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
                <li><a href="#min">min</a></li>
                <li><a href="#max">max</a></li>
                <li><a href="#floor">floor</a></li>
                <li><a href="#ceil">ceil</a></li>
                <li><a href="#round">round</a></li>
                <li><a href="#sqrt">sqrt</a></li>
                <li><a href="#pow">pow</a></li>
                <li><a href="#sin">sin</a></li>
                <li><a href="#cos">cos</a></li>
                <li><a href="#tan">tan</a></li>
                <li><a href="#log">log</a></li>
                <li><a href="#rand">rand</a></li>
                <li><a href="#rand_int">rand_int</a></li>
            </ul>
        </nav>
    </aside>
    
    <div class="content-body">
        <h1>Math Package</h1>
        <p>Mathematical functions and constants.</p>
        
        <div class="function-card" id="abs">
            <div class="function-header">
                <div class="function-name">abs</div>
                <div class="function-signature">abs(n) -> number</div>
            </div>
            <div class="function-body">
                <div class="function-description">Returns the absolute value of a number.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
println(abs(-5)); // 5
{% endhighlight %}
                </div>
            </div>
        </div>

        <div class="function-card" id="min">
            <div class="function-header">
                <div class="function-name">min</div>
                <div class="function-signature">min(a, b) -> number</div>
            </div>
            <div class="function-body">
                <div class="function-description">Returns the smaller of two numbers.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
println(min(10, 5)); // 5
{% endhighlight %}
                </div>
            </div>
        </div>

        <div class="function-card" id="max">
            <div class="function-header">
                <div class="function-name">max</div>
                <div class="function-signature">max(a, b) -> number</div>
            </div>
            <div class="function-body">
                <div class="function-description">Returns the larger of two numbers.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
println(max(10, 5)); // 10
{% endhighlight %}
                </div>
            </div>
        </div>

        <div class="function-card" id="floor">
            <div class="function-header">
                <div class="function-name">floor</div>
                <div class="function-signature">floor(n) -> int</div>
            </div>
            <div class="function-body">
                <div class="function-description">Returns the greatest integer less than or equal to n.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
println(floor(3.9)); // 3
{% endhighlight %}
                </div>
            </div>
        </div>

        <div class="function-card" id="ceil">
            <div class="function-header">
                <div class="function-name">ceil</div>
                <div class="function-signature">ceil(n) -> int</div>
            </div>
            <div class="function-body">
                <div class="function-description">Returns the smallest integer greater than or equal to n.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
println(ceil(3.1)); // 4
{% endhighlight %}
                </div>
            </div>
        </div>

        <div class="function-card" id="round">
            <div class="function-header">
                <div class="function-name">round</div>
                <div class="function-signature">round(n, [precision]) -> float</div>
            </div>
            <div class="function-body">
                <div class="function-description">Rounds a number to the nearest integer or specified precision.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
println(round(3.5)); // 4
println(round(3.14159, 2)); // 3.14
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
                <div class="function-description">Returns the square root of n.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
println(sqrt(16)); // 4
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
println(pow(2, 3)); // 8
{% endhighlight %}
                </div>
            </div>
        </div>

        <div class="function-card" id="sin">
            <div class="function-header">
                <div class="function-name">sin</div>
                <div class="function-signature">sin(rad) -> float</div>
            </div>
            <div class="function-body">
                <div class="function-description">Returns the sine of the angle (in radians).</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
println(sin(0)); // 0
{% endhighlight %}
                </div>
            </div>
        </div>

        <div class="function-card" id="rand">
            <div class="function-header">
                <div class="function-name">rand</div>
                <div class="function-signature">rand() -> float</div>
            </div>
            <div class="function-body">
                <div class="function-description">Returns a random float between 0.0 and 1.0.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
println(rand());
{% endhighlight %}
                </div>
            </div>
        </div>

        <div class="function-card" id="rand_int">
            <div class="function-header">
                <div class="function-name">rand_int</div>
                <div class="function-signature">rand_int(min, max) -> int</div>
            </div>
            <div class="function-body">
                <div class="function-description">Returns a random integer in range [min, max).</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
println(rand_int(1, 10));
{% endhighlight %}
                </div>
            </div>
        </div>
    </div>
</div>