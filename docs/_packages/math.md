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
                <li><a href="#import">Import</a></li>
                <li><a href="#abs">abs()</a></li>
                <li><a href="#min">min()</a></li>
                <li><a href="#max">max()</a></li>
                <li><a href="#floor">floor()</a></li>
                <li><a href="#ceil">ceil()</a></li>
                <li><a href="#round">round()</a></li>
                <li><a href="#sqrt">sqrt()</a></li>
                <li><a href="#pow">pow()</a></li>
                <li><a href="#sin">sin()</a></li>
                <li><a href="#cos">cos()</a></li>
                <li><a href="#tan">tan()</a></li>
                <li><a href="#log">log()</a></li>
                <li><a href="#rand">rand()</a></li>
                <li><a href="#rand_int">rand_int()</a></li>
                <li><a href="#fabs">fabs()</a></li>
                <li><a href="#asin">asin()</a></li>
                <li><a href="#acos">acos()</a></li>
                <li><a href="#atan">atan()</a></li>
                <li><a href="#atan2">atan2()</a></li>
                <li><a href="#log10">log10()</a></li>
                <li><a href="#exp">exp()</a></li>
            </ul>
        </nav>
    </aside>
    
    <div class="content-body">
        <h1>Math Package</h1>
        <p>Mathematical functions and constants.</p>

        <div class="function-card" id="import">
            <div class="function-header">
                <div class="function-name">Import</div>
                <div class="function-signature">import "math"</div>
            </div>
            <div class="function-body">
                <div class="function-description">Import the math package to use namespaced functions. Provides mathematical operations including trigonometry, logarithms, exponentiation, rounding, and random number generation.</div>
                <div class="function-example">
                    <h4>Examples</h4>
{% highlight go %}
// Standard import
import math;
var x = math.abs(-10);
println(x); // 10

// With alias
import math as m;
var y = m.abs(-20);
println(y); // 20

// Calculate hypotenuse
var hyp = m.sqrt(m.pow(3, 2) + m.pow(4, 2));  // 5.0
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="abs">
            <div class="function-header">
                <div class="function-name">abs</div>
                <div class="function-signature">abs(n) -> int</div>
            </div>
            <div class="function-body">
                <div class="function-description">Returns the absolute value of an integer. Converts negative integers to their positive equivalent. For floating-point absolute values, use fabs().</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
abs(-5);                    // 5
abs(10);                    // 10
abs(0);                     // 0
abs(-100);                  // 100
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
                <div class="function-description">Returns the smaller of two numbers. Accepts both integers and floats. If types are mixed, the result is promoted appropriately. Useful for clamping values or finding minimums in comparisons.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
min(10, 5);                 // 5
min(3.5, 2.1);              // 2.1
min(-10, -5);               // -10
min(5, 5.5);                // 5 (integer preserved if smaller)

// Clamp value to maximum
var maxValue = 100;
var input = 150;
var clamped = min(input, maxValue);  // 100
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
                <div class="function-description">Returns the larger of two numbers. Accepts both integers and floats. If types are mixed, the result is promoted appropriately. Useful for finding maximums or ensuring minimum thresholds.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
max(10, 5);                 // 10
max(3.5, 2.1);              // 3.5
max(-10, -5);               // -5
max(5, 5.5);                // 5.5 (float preserved if larger)

// Ensure minimum value
var minValue = 0;
var input = -5;
var clamped = max(input, minValue);  // 0
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
                <div class="function-description">Returns the greatest integer less than or equal to the input float. Rounds down towards negative infinity. Useful for integer division effects or rounding down measurements.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
floor(3.9);                 // 3
floor(3.1);                 // 3
floor(-3.1);                // -4 (rounds down, more negative)
floor(5.0);                 // 5 (already whole number)

// Calculate full pages needed
var items = 23;
var perPage = 10;
var pages = floor(items / perPage) + 1;  // 3
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
                <div class="function-description">Returns the smallest integer greater than or equal to the input float. Rounds up towards positive infinity. Useful for calculating required resources or rounding up measurements.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
ceil(3.1);                  // 4
ceil(3.9);                  // 4
ceil(-3.9);                 // -3 (rounds up, less negative)
ceil(5.0);                  // 5 (already whole number)

// Calculate containers needed
var items = 23;
var perContainer = 10;
var containers = ceil(items / perContainer);  // 3
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
                <div class="function-description">Rounds a float to the nearest value with optional decimal precision. Uses standard rounding rules (0.5 rounds up). Default precision is 0 (round to integer). Useful for formatting numbers or limiting decimal places.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
round(3.5);                 // 4.0
round(3.4);                 // 3.0
round(3.14159, 2);          // 3.14
round(3.14159, 4);          // 3.1416
round(2.5);                 // 2.0 (banker's rounding: to even)

// Format currency
var price = 19.999;
round(price, 2);            // 20.0
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
                <div class="function-description">Returns the square root of a non-negative number. Accepts both integers and floats. Returns an error for negative inputs. Commonly used in geometry, distance calculations, and statistical formulas.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
sqrt(16);                   // 4.0
sqrt(2.25);                 // 1.5
sqrt(2);                    // 1.4142135623730951

// Calculate distance between points
var dx = 3.0;
var dy = 4.0;
var distance = sqrt(pow(dx, 2) + pow(dy, 2));  // 5.0
{% endhighlight %}
                </div>
            </div>
        </div>

        <div class="function-card" id="pow">
            <div class="function-header">
                <div class="function-name">pow</div>
                <div class="function-signature">pow(base, exp) -> float</div>
            </div>
            <div class="function-body">
                <div class="function-description">Returns the base raised to the power of the exponent (base^exp). Supports both integer and floating-point arguments. Useful for exponential growth calculations, compound interest, and geometric formulas.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
pow(2, 3);                  // 8.0
pow(9, 0.5);                // 3.0 (square root)
pow(2, -1);                 // 0.5 (reciprocal)
pow(10, 6);                 // 1000000.0

// Compound interest: P * (1 + r)^t
var principal = 1000;
var rate = 0.05;
var years = 10;
var amount = principal * pow(1 + rate, years);  // 1628.89
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
                <div class="function-description">Returns the trigonometric sine of an angle specified in radians. The result is in the range [-1, 1]. Essential for wave functions, circular motion, and geometric calculations.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
sin(0);                     // 0.0
sin(3.14159 / 2);           // 1.0 (approximately, 90 degrees)
sin(3.14159);               // 0.0 (approximately, 180 degrees)

// Generate sine wave
for (var i = 0; i < 10; i = i + 1) {
    var y = sin(i * 0.5);
    println(y);
}
{% endhighlight %}
                </div>
            </div>
        </div>

        <div class="function-card" id="cos">
            <div class="function-header">
                <div class="function-name">cos</div>
                <div class="function-signature">cos(rad) -> float</div>
            </div>
            <div class="function-body">
                <div class="function-description">Returns the trigonometric cosine of an angle specified in radians. The result is in the range [-1, 1]. Complementary to sine, used for x-coordinates in circular motion and phase-shifted wave functions.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
cos(0);                     // 1.0
cos(3.14159 / 2);           // 0.0 (approximately, 90 degrees)
cos(3.14159);               // -1.0 (approximately, 180 degrees)

// Calculate x, y on unit circle
var angle = 3.14159 / 4;    // 45 degrees
var x = cos(angle);         // 0.707
var y = sin(angle);         // 0.707
{% endhighlight %}
                </div>
            </div>
        </div>

        <div class="function-card" id="tan">
            <div class="function-header">
                <div class="function-name">tan</div>
                <div class="function-signature">tan(rad) -> float</div>
            </div>
            <div class="function-body">
                <div class="function-description">Returns the trigonometric tangent of an angle specified in radians. Equal to sin(rad)/cos(rad). Useful for slope calculations and angle measurements. Undefined at π/2, π/2 + π, etc.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
tan(0);                     // 0.0
tan(3.14159 / 4);           // 1.0 (approximately, 45 degrees)
tan(3.14159);               // 0.0 (approximately, 180 degrees)

// Calculate slope from angle
var angle = 3.14159 / 6;    // 30 degrees
var slope = tan(angle);     // 0.577 (1/sqrt(3))
{% endhighlight %}
                </div>
            </div>
        </div>

        <div class="function-card" id="asin">
            <div class="function-header">
                <div class="function-name">asin</div>
                <div class="function-signature">asin(value) -> float</div>
            </div>
            <div class="function-body">
                <div class="function-description">Returns the arcsine (inverse sine) of a value in radians. The input must be in the range [-1, 1]. Returns the angle whose sine is the given value. Useful for converting ratios back to angles.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
asin(0);                    // 0.0
asin(1);                    // 1.5708 (π/2, 90 degrees)
asin(-1);                   // -1.5708 (-π/2, -90 degrees)
asin(0.5);                  // 0.5236 (π/6, 30 degrees)

// Find angle from opposite/hypotenuse
var opposite = 3.0;
var hypotenuse = 6.0;
var angle = asin(opposite / hypotenuse);  // 0.5236 radians (30 degrees)
{% endhighlight %}
                </div>
            </div>
        </div>

        <div class="function-card" id="acos">
            <div class="function-header">
                <div class="function-name">acos</div>
                <div class="function-signature">acos(value) -> float</div>
            </div>
            <div class="function-body">
                <div class="function-description">Returns the arccosine (inverse cosine) of a value in radians. The input must be in the range [-1, 1]. Returns the angle whose cosine is the given value. Useful for calculating angles in triangles.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
acos(1);                    // 0.0
acos(0);                    // 1.5708 (π/2, 90 degrees)
acos(-1);                   // 3.1416 (π, 180 degrees)
acos(0.5);                  // 1.0472 (π/3, 60 degrees)

// Find angle from adjacent/hypotenuse
var adjacent = 3.0;
var hypotenuse = 6.0;
var angle = acos(adjacent / hypotenuse);  // 1.0472 radians (60 degrees)
{% endhighlight %}
                </div>
            </div>
        </div>

        <div class="function-card" id="atan">
            <div class="function-header">
                <div class="function-name">atan</div>
                <div class="function-signature">atan(value) -> float</div>
            </div>
            <div class="function-body">
                <div class="function-description">Returns the arctangent (inverse tangent) of a value in radians. Accepts any real number. Returns the angle whose tangent is the given value. The result is in the range (-π/2, π/2).</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
atan(0);                    // 0.0
atan(1);                    // 0.7854 (π/4, 45 degrees)
atan(-1);                   // -0.7854 (-π/4, -45 degrees)
atan(1000000);              // 1.5708 (approaches π/2)

// Find angle from slope
var slope = 1.0;            // rise/run = 1
var angle = atan(slope);    // 0.7854 radians (45 degrees)
{% endhighlight %}
                </div>
            </div>
        </div>

        <div class="function-card" id="atan2">
            <div class="function-header">
                <div class="function-name">atan2</div>
                <div class="function-signature">atan2(y, x) -> float</div>
            </div>
            <div class="function-body">
                <div class="function-description">Returns the arctangent of y/x in radians, considering the signs of both arguments to determine the correct quadrant. Returns the angle between the positive x-axis and the point (x, y). More robust than atan(y/x) as it handles all quadrants correctly.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
atan2(0, 1);                // 0.0 (positive x-axis)
atan2(1, 0);                // 1.5708 (π/2, positive y-axis)
atan2(0, -1);               // 3.1416 (π, negative x-axis)
atan2(-1, 0);               // -1.5708 (-π/2, negative y-axis)
atan2(1, 1);                // 0.7854 (π/4, 45 degrees)

// Calculate angle of vector from origin
var dx = -3.0;
var dy = 4.0;
var angle = atan2(dy, dx);  // 2.2143 radians (about 126.87 degrees)
{% endhighlight %}
                </div>
            </div>
        </div>

        <div class="function-card" id="log">
            <div class="function-header">
                <div class="function-name">log</div>
                <div class="function-signature">log(n) -> float</div>
            </div>
            <div class="function-body">
                <div class="function-description">Returns the natural logarithm (base e) of a positive number. The natural logarithm is the inverse of the exponential function (exp). Essential for exponential decay, growth calculations, and information theory.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
log(1);                     // 0.0
log(2.71828);               // 1.0 (approximately, ln(e))
log(10);                    // 2.3026
log(100);                   // 4.6052

// Calculate doubling time
var growthRate = 0.05;      // 5% growth
var doublingTime = log(2) / log(1 + growthRate);  // 14.21 periods
{% endhighlight %}
                </div>
            </div>
        </div>

        <div class="function-card" id="log10">
            <div class="function-header">
                <div class="function-name">log10</div>
                <div class="function-signature">log10(n) -> float</div>
            </div>
            <div class="function-body">
                <div class="function-description">Returns the common logarithm (base 10) of a positive number. Useful for calculating orders of magnitude, decibel levels, pH values, and any logarithmic scale measurements.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
log10(1);                   // 0.0
log10(10);                  // 1.0
log10(100);                 // 2.0
log10(1000);                // 3.0
log10(0.001);               // -3.0

// Calculate order of magnitude
var value = 50000;
var magnitude = floor(log10(value));  // 4 (10^4 = 10000)
{% endhighlight %}
                </div>
            </div>
        </div>

        <div class="function-card" id="exp">
            <div class="function-header">
                <div class="function-name">exp</div>
                <div class="function-signature">exp(n) -> float</div>
            </div>
            <div class="function-body">
                <div class="function-description">Returns e (Euler's number, approximately 2.71828) raised to the power of n. The exponential function is the inverse of the natural logarithm. Essential for exponential growth models, compound interest, and probability distributions.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
exp(0);                     // 1.0
exp(1);                     // 2.71828 (e)
exp(2);                     // 7.38906 (e^2)
exp(-1);                    // 0.36788 (1/e)

// Continuous compound interest: P * e^(rt)
var principal = 1000;
var rate = 0.05;
var time = 10;
var amount = principal * exp(rate * time);  // 1648.72
{% endhighlight %}
                </div>
            </div>
        </div>

        <div class="function-card" id="fabs">
            <div class="function-header">
                <div class="function-name">fabs</div>
                <div class="function-signature">fabs(n) -> float</div>
            </div>
            <div class="function-body">
                <div class="function-description">Returns the absolute value of a floating-point number. This is the float equivalent of abs() for integers. Converts negative floats to their positive equivalent while preserving the floating-point type.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
fabs(-3.14);                // 3.14
fabs(2.5);                  // 2.5
fabs(-0.001);               // 0.001
fabs(0.0);                  // 0.0

// Calculate magnitude
var x = -3.5;
var y = 4.2;
var magnitude = sqrt(fabs(x) + fabs(y));  // 2.768
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
                <div class="function-description">Returns a random floating-point number in the range [0.0, 1.0). The value is uniformly distributed and can be used for probabilistic calculations, simulations, or generating random values within a specific range.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
rand();                     // 0.374 (example, different each call)
rand();                     // 0.891 (example, different each call)

// Generate random value in range [min, max]
var min = 10.0;
var max = 20.0;
var randomValue = min + rand() * (max - min);  // e.g., 15.73
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
                <div class="function-description">Returns a random integer in the inclusive range [min, max]. Both endpoints are included in the possible values. Useful for dice rolls, random selections, and generating test data. Returns an error if min > max.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
rand_int(1, 6);             // Dice roll: 1-6
rand_int(0, 100);           // Random percentage: 0-100
rand_int(-10, 10);          // Random signed value: -10 to 10

// Simulate dice roll
var roll = rand_int(1, 6);
println("You rolled: " + roll);
{% endhighlight %}
                </div>
            </div>
        </div>
    </div>
</div>
