---
layout: default
title: Language Guide - Go-Mix
description: Complete guide to Go-Mix programming language syntax, features, and constructs
permalink: /language-guide/
---

<div class="content-page">
    <aside class="sidebar">
        <nav class="sidebar-nav">
            <div class="sidebar-title">On This Page</div>
            <ul class="sidebar-menu">
                <li><a href="#types">Data Types</a></li>
                <li><a href="#variables">Variables</a></li>
                <li><a href="#operators">Operators</a></li>
                <li><a href="#control-flow">Control Flow</a></li>
                <li><a href="#collections">Collections</a></li>
                <li><a href="#functions">Functions</a></li>
                <li><a href="#closures">Closures</a></li>
                <li><a href="#oop">Object-Oriented</a></li>
            </ul>
        </nav>
    </aside>
    
    <div class="content-body">
        <h1>Language Guide</h1>
        
        <p>Complete reference for Go-Mix programming language syntax, types, and features.</p>
        
        <h2 id="types">Data Types</h2>
        
        <p>Go-Mix supports six primitive types plus <code>nil</code>:</p>
        
        <table>
            <thead>
                <tr>
                    <th>Type</th>
                    <th>Example</th>
                    <th>Description</th>
                </tr>
            </thead>
            <tbody>
                <tr>
                    <td><code>int</code></td>
                    <td><code>42</code>, <code>0xFF</code></td>
                    <td>64-bit signed integers; hex (0x) and octal (0o) support</td>
                </tr>
                <tr>
                    <td><code>float</code></td>
                    <td><code>3.14</code>, <code>1.2e3</code></td>
                    <td>64-bit IEEE-754 double precision</td>
                </tr>
                <tr>
                    <td><code>bool</code></td>
                    <td><code>true</code>, <code>false</code></td>
                    <td>Boolean logic values</td>
                </tr>
                <tr>
                    <td><code>string</code></td>
                    <td><code>"Hello"</code></td>
                    <td>UTF-8 strings with escape sequences</td>
                </tr>
                <tr>
                    <td><code>char</code></td>
                    <td><code>'A'</code>, <code>'\n'</code></td>
                    <td>Single Unicode character</td>
                </tr>
                <tr>
                    <td><code>nil</code></td>
                    <td><code>nil</code></td>
                    <td>Represents absence of value</td>
                </tr>
            </tbody>
        </table>
        
        <h3>Type Checking</h3>
        
        <div class="code-block">
            <div class="code-header">
                <span class="code-dot red"></span>
                <span class="code-dot yellow"></span>
                <span class="code-dot green"></span>
                <span class="code-title">types.gm</span>
            </div>
{% highlight go %}
typeof(42);                 // "int"
typeof(3.14);              // "float"
typeof("text");            // "string"
typeof(true);              // "bool"
typeof(nil);               // "nil"
typeof([1, 2, 3]);         // "array"
typeof(fn() {});           // "func"
{% endhighlight %}
        </div>
        
        <h2 id="variables">Variables</h2>
        
        <p>Go-Mix provides three variable kinds with different semantics:</p>
        
        <h3>Dynamic Variables (<code>var</code>)</h3>
        
        <p>Type can change during execution. Ideal for rapid prototyping.</p>
        
        <div class="code-block">
            <div class="code-header">
                <span class="code-dot red"></span>
                <span class="code-dot yellow"></span>
                <span class="code-dot green"></span>
                <span class="code-title">dynamic.gm</span>
            </div>
{% highlight go %}
var x = 42;        // x is an integer
x = "hello";       // x is now a string
x = [1, 2, 3];     // x is now an array
{% endhighlight %}
        </div>
        
        <h3>Static Variables (<code>let</code>)</h3>
        
        <p>Type is locked on first assignment. Provides safety without explicit type annotations.</p>
        
        <div class="code-block">
            <div class="code-header">
                <span class="code-dot red"></span>
                <span class="code-dot yellow"></span>
                <span class="code-dot green"></span>
                <span class="code-title">static.gm</span>
            </div>
{% highlight go %}
let count = 0;
count = count + 1;     // OK
count = "text";        // ERROR: type mismatch
{% endhighlight %}
        </div>
        
        <h3>Immutable Constants (<code>const</code>)</h3>
        
        <p>Cannot be reassigned after initial binding.</p>
        
        <div class="code-block">
            <div class="code-header">
                <span class="code-dot red"></span>
                <span class="code-dot yellow"></span>
                <span class="code-dot green"></span>
                <span class="code-title">constants.gm</span>
            </div>
{% highlight go %}
const PI = 3.14159;
const MAX_SIZE = 100;
MAX_SIZE = 200;        // ERROR: constants are immutable
{% endhighlight %}
        </div>
        
        <h2 id="operators">Operators</h2>
        
        <h3>Arithmetic Operators</h3>
        
        <table>
            <thead>
                <tr>
                    <th>Operator</th>
                    <th>Description</th>
                    <th>Example</th>
                </tr>
            </thead>
            <tbody>
                <tr><td><code>+</code></td><td>Addition</td><td><code>5 + 3 // 8</code></td></tr>
                <tr><td><code>-</code></td><td>Subtraction</td><td><code>5 - 3 // 2</code></td></tr>
                <tr><td><code>*</code></td><td>Multiplication</td><td><code>5 * 3 // 15</code></td></tr>
                <tr><td><code>/</code></td><td>Division</td><td><code>5 / 2 // 2.5</code></td></tr>
                <tr><td><code>%</code></td><td>Modulo</td><td><code>5 % 2 // 1</code></td></tr>
                <tr><td><code>**</code></td><td>Power</td><td><code>2 ** 3 // 8</code></td></tr>
            </tbody>
        </table>
        
        <h3>Comparison Operators</h3>
        
        <table>
            <thead>
                <tr>
                    <th>Operator</th>
                    <th>Description</th>
                </tr>
            </thead>
            <tbody>
                <tr><td><code>==</code></td><td>Equal to</td></tr>
                <tr><td><code>!=</code></td><td>Not equal to</td></tr>
                <tr><td><code><</code></td><td>Less than</td></tr>
                <tr><td><code>></code></td><td>Greater than</td></tr>
                <tr><td><code><=</code></td><td>Less than or equal</td></tr>
                <tr><td><code>>=</code></td><td>Greater than or equal</td></tr>
            </tbody>
        </table>
        
        <h3>Logical Operators</h3>
        
        <table>
            <thead>
                <tr>
                    <th>Operator</th>
                    <th>Description</th>
                </tr>
            </thead>
            <tbody>
                <tr><td><code>&&</code></td><td>Logical AND (short-circuit)</td></tr>
                <tr><td><code>||</code></td><td>Logical OR (short-circuit)</td></tr>
                <tr><td><code>!</code></td><td>Logical NOT</td></tr>
            </tbody>
        </table>
        
        <h2 id="control-flow">Control Flow</h2>
        
        <h3>If-Else Statements</h3>
        
        <div class="code-block">
            <div class="code-header">
                <span class="code-dot red"></span>
                <span class="code-dot yellow"></span>
                <span class="code-dot green"></span>
                <span class="code-title">conditionals.gm</span>
            </div>
{% highlight go %}
if (age >= 18) {
    println("Adult");
} else if (age >= 13) {
    println("Teenager");
} else {
    println("Child");
}

// Ternary-like expression
var status = if (x > 0) { "positive"; } else { "non-positive"; };
{% endhighlight %}
        </div>
        
        <h3>For Loops</h3>
        
        <div class="code-block">
            <div class="code-header">
                <span class="code-dot red"></span>
                <span class="code-dot yellow"></span>
                <span class="code-dot green"></span>
                <span class="code-title">loops.gm</span>
            </div>
{% highlight go %}
// Standard for loop
for (var i = 0; i < 5; i = i + 1) {
    println(i);  // 0, 1, 2, 3, 4
}

// With break and continue
for (var i = 0; i < 10; i = i + 1) {
    if (i == 3) { continue; }    // Skip iteration
    if (i == 7) { break; }       // Exit loop
    println(i);
}
{% endhighlight %}
        </div>
        
        <h3>While Loops</h3>
        
        <div class="code-block">
            <div class="code-header">
                <span class="code-dot red"></span>
                <span class="code-dot yellow"></span>
                <span class="code-dot green"></span>
                <span class="code-title">while.gm</span>
            </div>
{% highlight go %}
var count = 0;
while (count < 5) {
    println(count);
    count = count + 1;
}

// Multiple conditions
var x = 10, y = 0;
while (x > 0, y < 10) {
    x = x - 1;
    y = y + 1;
}

// Infinite loop with break
var i = 0;
while (true) {
    if (i >= 100) { break; }
    i = i + 1;
}
{% endhighlight %}
        </div>
        
        <h3>Foreach Loops</h3>
        
        <div class="code-block">
            <div class="code-header">
                <span class="code-dot red"></span>
                <span class="code-dot yellow"></span>
                <span class="code-dot green"></span>
                <span class="code-title">foreach.gm</span>
            </div>
{% highlight go %}
// Range iteration (inclusive)
foreach i in 1...5 {      // 1, 2, 3, 4, 5
    println(i);
}

// Reverse range
foreach i in 5...1 {      // 5, 4, 3, 2, 1
    println(i);
}

// Array iteration with value
var arr = [10, 20, 30];
foreach val in arr {
    println(val);
}

// Array iteration with index and value
foreach idx, val in arr {
    println(idx, val);    // 0, 10; 1, 20; 2, 30
}
{% endhighlight %}
        </div>
        
        <h2 id="collections">Collections</h2>
        
        <h3>Arrays</h3>
        
        <p>Mutable sequences with homogeneous elements (typically).</p>
        
        <div class="code-block">
            <div class="code-header">
                <span class="code-dot red"></span>
                <span class="code-dot yellow"></span>
                <span class="code-dot green"></span>
                <span class="code-title">arrays.gm</span>
            </div>
{% highlight go %}
// Array literals
var empty = [];
var numbers = [1, 2, 3, 4, 5];
var mixed = [1, "hello", true, 3.14];

// Indexing (0-based)
var arr = [10, 20, 30, 40, 50];
println(arr[0]);           // 10 (first element)
println(arr[2]);           // 30 (third element)

// Negative indexing
println(arr[-1]);          // 50 (last element)
println(arr[-2]);          // 40 (second-to-last)

// Index assignment
arr[1] = 25;               // [10, 25, 30, 40, 50]

// Array methods
push(arr, 60);             // Add to end
pop(arr);                  // Remove from end
unshift(arr, 5);           // Add to start
shift(arr);                // Remove from start
sort(arr);                 // Sort in-place
{% endhighlight %}
        </div>
        
        <h3>Maps (Dictionaries)</h3>
        
        <div class="code-block">
            <div class="code-header">
                <span class="code-dot red"></span>
                <span class="code-dot yellow"></span>
                <span class="code-dot green"></span>
                <span class="code-title">maps.gm</span>
            </div>
{% highlight go %}
// Map literals
var user = map{
    "name": "Alice",
    "age": 30,
    "city": "New York"
};

// Key access
println(user["name"]);     // "Alice"

// Adding/updating keys
user["email"] = "alice@example.com";
user["age"] = 31;
{% endhighlight %}
        </div>
        
        <h3>Lists</h3>
        
        <p>Mutable, heterogeneous sequences.</p>
        
        <div class="code-block">
            <div class="code-header">
                <span class="code-dot red"></span>
                <span class="code-dot yellow"></span>
                <span class="code-dot green"></span>
                <span class="code-title">lists.gm</span>
            </div>
{% highlight go %}
// Creating lists
var list = list(1, "two", 3.0, true);

// Access like arrays
println(list[0]);          // 1
println(list[1]);          // "two"
{% endhighlight %}
        </div>
        
        <h3>Tuples</h3>
        
        <p>Immutable, fixed-size sequences.</p>
        
        <div class="code-block">
            <div class="code-header">
                <span class="code-dot red"></span>
                <span class="code-dot yellow"></span>
                <span class="code-dot green"></span>
                <span class="code-title">tuples.gm</span>
            </div>
{% highlight go %}
// Creating tuples
var coords = tuple(10, 20);
var rgb = tuple(255, 128, 64);

// Access elements
println(coords[0]);        // 10
println(coords[1]);        // 20

// Immutable - cannot be modified
coords[0] = 15;            // ERROR
{% endhighlight %}
        </div>
        
        <h2 id="functions">Functions</h2>
        
        <h3>Function Definitions</h3>
        
        <div class="code-block">
            <div class="code-header">
                <span class="code-dot red"></span>
                <span class="code-dot yellow"></span>
                <span class="code-dot green"></span>
                <span class="code-title">functions.gm</span>
            </div>
{% highlight go %}
// Basic function
fn greet(name) {
    return "Hello, " + name;
}

// Multiple parameters
fn add(a, b) {
    return a + b;
}

// No return (returns nil)
fn printInfo(x) {
    println("Value: " + x);
}

// Recursive function
fn factorial(n) {
    if (n <= 1) { return 1; }
    return n * factorial(n - 1);
}
{% endhighlight %}
        </div>
        
        <h3>Function Expressions (Lambda)</h3>
        
        <div class="code-block">
            <div class="code-header">
                <span class="code-dot red"></span>
                <span class="code-dot yellow"></span>
                <span class="code-dot green"></span>
                <span class="code-title">lambda.gm</span>
            </div>
{% highlight go %}
// Anonymous functions
var double = fn(x) { return x * 2; };
var add = fn(a, b) { return a + b; };

// Higher-order functions
var makeMultiplier = fn(factor) {
    return fn(x) { return x * factor; };
};

var times5 = makeMultiplier(5);
println(times5(3));        // 15
{% endhighlight %}
        </div>
        
        <h3>Higher-Order Functions</h3>
        
        <div class="code-block">
            <div class="code-header">
                <span class="code-dot red"></span>
                <span class="code-dot yellow"></span>
                <span class="code-dot green"></span>
                <span class="code-title">higher-order.gm</span>
            </div>
{% highlight go %}
var numbers = [1, 2, 3, 4, 5];

// Map - transform each element
var doubled = map(numbers, fn(x) { return x * 2; });
// [2, 4, 6, 8, 10]

// Filter - keep matching elements
var evens = filter(numbers, fn(x) { return x % 2 == 0; });
// [2, 4]

// Reduce - accumulate to single value
var sum = reduce(numbers, fn(acc, x) { return acc + x; }, 0);
// 15

// Find - get first matching element
var firstEven = find(numbers, fn(x) { return x % 2 == 0; });
// 2

// Check predicates
var hasEven = some(numbers, fn(x) { return x % 2 == 0; });     // true
var allPositive = every(numbers, fn(x) { return x > 0; });       // true
{% endhighlight %}
        </div>
        
        <h2 id="closures">Closures</h2>
        
        <p>Functions are first-class citizens with full support for lexical closures:</p>
        
        <div class="code-block">
            <div class="code-header">
                <span class="code-dot red"></span>
                <span class="code-dot yellow"></span>
                <span class="code-dot green"></span>
                <span class="code-title">closures.gm</span>
            </div>
{% highlight go %}
fn makeCounter(start) {
    var count = start;
    fn increment() { 
        count = count + 1; 
        return count; 
    }
    return increment;
}

var counter = makeCounter(0);
println(counter()); // 1
println(counter()); // 2
println(counter()); // 3

// Each counter maintains its own 'count' variable
var counter2 = makeCounter(100);
println(counter2()); // 101
{% endhighlight %}
        </div>
        
        <h2 id="oop">Object-Oriented Programming</h2>
        
        <h3>Struct Definition</h3>
        
        <div class="code-block">
            <div class="code-header">
                <span class="code-dot red"></span>
                <span class="code-dot yellow"></span>
                <span class="code-dot green"></span>
                <span class="code-title">structs.gm</span>
            </div>
{% highlight go %}
struct Circle {
    var radius = 0;

    // Constructor
    fn init(r) {
        this.radius = r;
    }

    // Instance methods
    fn area() {
        return 3.14159 * this.radius * this.radius;
    }

    fn circumference() {
        return 2 * 3.14159 * this.radius;
    }

    fn scale(factor) {
        this.radius = this.radius * factor;
        return this;  // Method chaining
    }
}

// Creating instances
var c = new Circle(5);
println(c.area());         // 78.53975
println(c.circumference()); // 31.4159

// Method chaining
c.scale(2).scale(2);       // radius becomes 20
{% endhighlight %}
        </div>
        
        <h3>Complex Example</h3>
        
        <div class="code-block">
            <div class="code-header">
                <span class="code-dot red"></span>
                <span class="code-dot yellow"></span>
                <span class="code-dot green"></span>
                <span class="code-title">bank-account.gm</span>
            </div>
{% highlight go %}
struct BankAccount {
    var balance = 0;
    var owner = "Unknown";

    fn init(owner, initial) {
        this.owner = owner;
        this.balance = initial;
    }

    fn deposit(amount) {
        if (amount <= 0) {
            panic("Deposit amount must be positive");
        }
        this.balance = this.balance + amount;
        return this.balance;
    }

    fn withdraw(amount) {
        if (amount > this.balance) {
            panic("Insufficient funds");
        }
        this.balance = this.balance - amount;
        return this.balance;
    }

    fn getBalance() {
        return this.balance;
    }

    fn getInfo() {
        return this.owner + " has $" + this.balance;
    }
}

var account = new BankAccount("Alice", 1000);
account.deposit(500);      // 1500
account.withdraw(200);     // 1300
println(account.getInfo()); // Alice has $1300
{% endhighlight %}
        </div>
    </div>
</div>
