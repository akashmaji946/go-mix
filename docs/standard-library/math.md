---
title: "Math"
layout: default
parent: Standard Library
nav_order: 4
description: "Mathematical functions including trigonometry, logarithms, and random numbers"
permalink: /standard-library/math/
---

# Math Package
{: .no_toc }

Mathematical functions including trigonometry, logarithms, and random numbers
{: .fs-6 .fw-300 }

## Table of Contents
{: .no_toc .text-delta }

1. TOC
{:toc}

---

## Import

`import "math"`
{: .fs-5 .fw-300 }

Import the math package to use namespaced functions. Provides mathematical operations including trigonometry, logarithms, exponentiation, rounding, and random number generation.

```go
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
```

---

## abs

`abs(n) -> int`
{: .fs-5 .fw-300 }

Returns the absolute value of an integer. Converts negative integers to their positive equivalent. For floating-point absolute values, use fabs().

```go
abs(-5);                    // 5
abs(10);                    // 10
abs(0);                     // 0
abs(-100);                  // 100
```

---

## min

`min(a, b) -> number`
{: .fs-5 .fw-300 }

Returns the smaller of two numbers. Accepts both integers and floats. If types are mixed, the result is promoted appropriately. Useful for clamping values or finding minimums in comparisons.

```go
min(10, 5);                 // 5
min(3.5, 2.1);              // 2.1
min(-10, -5);               // -10
min(5, 5.5);                // 5 (integer preserved if smaller)

// Clamp value to maximum
var maxValue = 100;
var input = 150;
var clamped = min(input, maxValue);  // 100
```

---

## max

`max(a, b) -> number`
{: .fs-5 .fw-300 }

Returns the larger of two numbers. Accepts both integers and floats. If types are mixed, the result is promoted appropriately. Useful for finding maximums or ensuring minimum thresholds.

```go
max(10, 5);                 // 10
max(3.5, 2.1);              // 3.5
max(-10, -5);               // -5
max(5, 5.5);                // 5.5 (float preserved if larger)

// Ensure minimum value
var minValue = 0;
var input = -5;
var clamped = max(input, minValue);  // 0
```

---

## floor

`floor(n) -> int`
{: .fs-5 .fw-300 }

Returns the greatest integer less than or equal to the input float. Rounds down towards negative infinity. Useful for integer division effects or rounding down measurements.

```go
floor(3.9);                 // 3
floor(3.1);                 // 3
floor(-3.1);                // -4 (rounds down, more negative)
floor(5.0);                 // 5 (already whole number)

// Calculate full pages needed
var items = 23;
var perPage = 10;
var pages = floor(items / perPage) + 1;  // 3
```

---

## ceil

`ceil(n) -> int`
{: .fs-5 .fw-300 }

Returns the smallest integer greater than or equal to the input float. Rounds up towards positive infinity. Useful for calculating required resources or rounding up measurements.

```go
ceil(3.1);                  // 4
ceil(3.9);                  // 4
ceil(-3.9);                 // -3 (rounds up, less negative)
ceil(5.0);                  // 5 (already whole number)

// Calculate containers needed
var items = 23;
var perContainer = 10;
var containers = ceil(items / perContainer);  // 3
```

---

## round

`round(n, [precision]) -> float`
{: .fs-5 .fw-300 }

Rounds a float to the nearest value with optional decimal precision. Uses standard rounding rules (0.5 rounds up). Default precision is 0 (round to integer). Useful for formatting numbers or limiting decimal places.

```go
round(3.5);                 // 4.0
round(3.4);                 // 3.0
round(3.14159, 2);          // 3.14
round(3.14159, 4);          // 3.1416
round(2.5);                 // 2.0 (banker's rounding: to even)

// Format currency
var price = 19.999;
round(price, 2);            // 20.0
```

---

## sqrt

`sqrt(n) -> float`
{: .fs-5 .fw-300 }

Returns the square root of a non-negative number. Accepts both integers and floats. Returns an error for negative inputs. Commonly used in geometry, distance calculations, and statistical formulas.

```go
sqrt(16);                   // 4.0
sqrt(2.25);                 // 1.5
sqrt(2);                    // 1.4142135623730951

// Calculate distance between points
var dx = 3.0;
var dy = 4.0;
var distance = sqrt(pow(dx, 2) + pow(dy, 2));  // 5.0
```

---

## pow

`pow(base, exp) -> float`
{: .fs-5 .fw-300 }

Returns the base raised to the power of the exponent (base^exp). Supports both integer and floating-point arguments. Useful for exponential growth calculations, compound interest, and geometric formulas.

```go
pow(2, 3);                  // 8.0
pow(9, 0.5);                // 3.0 (square root)
pow(2, -1);                 // 0.5 (reciprocal)
pow(10, 6);                 // 1000000.0

// Compound interest: P * (1 + r)^t
var principal = 1000;
var rate = 0.05;
var years = 10;
var amount = principal * pow(1 + rate, years);  // 1628.89
```

---

## sin

`sin(rad) -> float`
{: .fs-5 .fw-300 }

Returns the trigonometric sine of an angle specified in radians. The result is in the range [-1, 1]. Essential for wave functions, circular motion, and geometric calculations.

```go
sin(0);                     // 0.0
sin(3.14159 / 2);           // 1.0 (approximately, 90 degrees)
sin(3.14159);               // 0.0 (approximately, 180 degrees)

// Generate sine wave
for (var i = 0; i < 10; i = i + 1) {
    var y = sin(i * 0.5);
    println(y);
}
```

---

## cos

`cos(rad) -> float`
{: .fs-5 .fw-300 }

Returns the trigonometric cosine of an angle specified in radians. The result is in the range [-1, 1]. Complementary to sine, used for x-coordinates in circular motion and phase-shifted wave functions.

```go
cos(0);                     // 1.0
cos(3.14159 / 2);           // 0.0 (approximately, 90 degrees)
cos(3.14159);               // -1.0 (approximately, 180 degrees)

// Calculate x, y on unit circle
var angle = 3.14159 / 4;    // 45 degrees
var x = cos(angle);         // 0.707
var y = sin(angle);         // 0.707
```

---

## tan

`tan(rad) -> float`
{: .fs-5 .fw-300 }

Returns the trigonometric tangent of an angle specified in radians. Equal to sin(rad)/cos(rad). Useful for slope calculations and angle measurements. Undefined at π/2, π/2 + π, etc.

```go
tan(0);                     // 0.0
tan(3.14159 / 4);           // 1.0 (approximately, 45 degrees)
tan(3.14159);               // 0.0 (approximately, 180 degrees)

// Calculate slope from angle
var angle = 3.14159 / 6;    // 30 degrees
var slope = tan(angle);     // 0.577 (1/sqrt(3))
```

---

## asin

`asin(value) -> float`
{: .fs-5 .fw-300 }

Returns the arcsine (inverse sine) of a value in radians. The input must be in the range [-1, 1]. Returns the angle whose sine is the given value. Useful for converting ratios back to angles.

```go
asin(0);                    // 0.0
asin(1);                    // 1.5708 (π/2, 90 degrees)
asin(-1);                   // -1.5708 (-π/2, -90 degrees)
asin(0.5);                  // 0.5236 (π/6, 30 degrees)

// Find angle from opposite/hypotenuse
var opposite = 3.0;
var hypotenuse = 6.0;
var angle = asin(opposite / hypotenuse);  // 0.5236 radians (30 degrees)
```

---

## acos

`acos(value) -> float`
{: .fs-5 .fw-300 }

Returns the arccosine (inverse cosine) of a value in radians. The input must be in the range [-1, 1]. Returns the angle whose cosine is the given value. Useful for calculating angles in triangles.

```go
acos(1);                    // 0.0
acos(0);                    // 1.5708 (π/2, 90 degrees)
acos(-1);                   // 3.1416 (π, 180 degrees)
acos(0.5);                  // 1.0472 (π/3, 60 degrees)

// Find angle from adjacent/hypotenuse
var adjacent = 3.0;
var hypotenuse = 6.0;
var angle = acos(adjacent / hypotenuse);  // 1.0472 radians (60 degrees)
```

---

## atan

`atan(value) -> float`
{: .fs-5 .fw-300 }

Returns the arctangent (inverse tangent) of a value in radians. Accepts any real number. Returns the angle whose tangent is the given value. The result is in the range (-π/2, π/2).

```go
atan(0);                    // 0.0
atan(1);                    // 0.7854 (π/4, 45 degrees)
atan(-1);                   // -0.7854 (-π/4, -45 degrees)
atan(1000000);              // 1.5708 (approaches π/2)

// Find angle from slope
var slope = 1.0;            // rise/run = 1
var angle = atan(slope);    // 0.7854 radians (45 degrees)
```

---

## atan2

`atan2(y, x) -> float`
{: .fs-5 .fw-300 }

Returns the arctangent of y/x in radians, considering the signs of both arguments to determine the correct quadrant. Returns the angle between the positive x-axis and the point (x, y). More robust than atan(y/x) as it handles all quadrants correctly.

```go
atan2(0, 1);                // 0.0 (positive x-axis)
atan2(1, 0);                // 1.5708 (π/2, positive y-axis)
atan2(0, -1);               // 3.1416 (π, negative x-axis)
atan2(-1, 0);               // -1.5708 (-π/2, negative y-axis)
atan2(1, 1);                // 0.7854 (π/4, 45 degrees)

// Calculate angle of vector from origin
var dx = -3.0;
var dy = 4.0;
var angle = atan2(dy, dx);  // 2.2143 radians (about 126.87 degrees)
```

---

## log

`log(n) -> float`
{: .fs-5 .fw-300 }

Returns the natural logarithm (base e) of a positive number. The natural logarithm is the inverse of the exponential function (exp). Essential for exponential decay, growth calculations, and information theory.

```go
log(1);                     // 0.0
log(2.71828);               // 1.0 (approximately, ln(e))
log(10);                    // 2.3026
log(100);                   // 4.6052

// Calculate doubling time
var growthRate = 0.05;      // 5% growth
var doublingTime = log(2) / log(1 + growthRate);  // 14.21 periods
```

---

## log10

`log10(n) -> float`
{: .fs-5 .fw-300 }

Returns the common logarithm (base 10) of a positive number. Useful for calculating orders of magnitude, decibel levels, pH values, and any logarithmic scale measurements.

```go
log10(1);                   // 0.0
log10(10);                  // 1.0
log10(100);                 // 2.0
log10(1000);                // 3.0
log10(0.001);               // -3.0

// Calculate order of magnitude
var value = 50000;
var magnitude = floor(log10(value));  // 4 (10^4 = 10000)
```

---

## exp

`exp(n) -> float`
{: .fs-5 .fw-300 }

Returns e (Euler's number, approximately 2.71828) raised to the power of n. The exponential function is the inverse of the natural logarithm. Essential for exponential growth models, compound interest, and probability distributions.

```go
exp(0);                     // 1.0
exp(1);                     // 2.71828 (e)
exp(2);                     // 7.38906 (e^2)
exp(-1);                    // 0.36788 (1/e)

// Continuous compound interest: P * e^(rt)
var principal = 1000;
var rate = 0.05;
var time = 10;
var amount = principal * exp(rate * time);  // 1648.72
```

---

## fabs

`fabs(n) -> float`
{: .fs-5 .fw-300 }

Returns the absolute value of a floating-point number. This is the float equivalent of abs() for integers. Converts negative floats to their positive equivalent while preserving the floating-point type.

```go
fabs(-3.14);                // 3.14
fabs(2.5);                  // 2.5
fabs(-0.001);               // 0.001
fabs(0.0);                  // 0.0

// Calculate magnitude
var x = -3.5;
var y = 4.2;
var magnitude = sqrt(fabs(x) + fabs(y));  // 2.768
```

---

## rand

`rand() -> float`
{: .fs-5 .fw-300 }

Returns a random floating-point number in the range [0.0, 1.0). The value is uniformly distributed and can be used for probabilistic calculations, simulations, or generating random values within a specific range.

```go
rand();                     // 0.374 (example, different each call)
rand();                     // 0.891 (example, different each call)

// Generate random value in range [min, max]
var min = 10.0;
var max = 20.0;
var randomValue = min + rand() * (max - min);  // e.g., 15.73
```

---

## rand_int

`rand_int(min, max) -> int`
{: .fs-5 .fw-300 }

Returns a random integer in the inclusive range [min, max]. Both endpoints are included in the possible values. Useful for dice rolls, random selections, and generating test data. Returns an error if min > max.

```go
rand_int(1, 6);             // Dice roll: 1-6
rand_int(0, 100);           // Random percentage: 0-100
rand_int(-10, 10);          // Random signed value: -10 to 10

// Simulate dice roll
var roll = rand_int(1, 6);
println("You rolled: " + roll);
```

---
