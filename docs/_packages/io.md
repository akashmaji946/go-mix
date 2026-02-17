---
layout: default
title: I/O Package - Go-Mix
description: Input/output operations including scanln, scanf, input, sprintf
---

<div class="content-page">
    <aside class="sidebar">
        <nav class="sidebar-nav">
            <div class="sidebar-title">I/O Functions</div>
            <ul class="sidebar-menu">
                <li><a href="#import">Import</a></li>
                <li><a href="#scanln">scanln()</a></li>
                <li><a href="#gets">gets()</a></li>
                <li><a href="#scanf">scanf()</a></li>
                <li><a href="#input">input()</a></li>
                <li><a href="#scan">scan()</a></li>
                <li><a href="#getchar">getchar()</a></li>
                <li><a href="#putchar">putchar()</a></li>
                <li><a href="#puts">puts()</a></li>
                <li><a href="#sprintf">sprintf()</a></li>
                <li><a href="#eprintln">eprintln()</a></li>
                <li><a href="#eprintf">eprintf()</a></li>
                <li><a href="#flush">flush()</a></li>
            </ul>
        </nav>
    </aside>
    
    <div class="content-body">
        <h1>I/O Package</h1>
        <p>Input and output operations for user interaction, formatted printing, and error handling.</p>
        
        <div class="function-card" id="import">
            <div class="function-header">
                <div class="function-name">Import</div>
                <div class="function-signature">import "io"</div>
            </div>
            <div class="function-body">
                <div class="function-description">Import the io package to use namespaced functions for reading input, formatted output, and error stream operations.</div>
                <div class="function-example">
                    <h4>Examples</h4>
{% highlight go %}
// Standard import
import io;
var name = io.scanln();
io.printf("Hello %s", name);

// With alias
import io as i;
var name = i.scanln();
i.printf("Hello %s", name);
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="scanln">
            <div class="function-header">
                <div class="function-name">scanln</div>
                <div class="function-signature">scanln() -> string</div>
            </div>
            <div class="function-body">
                <div class="function-description">Reads a complete line from standard input, returning the text without the trailing newline. Blocks until the user presses Enter. Returns an empty string if EOF is encountered.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
println("Enter your name:");
var name = scanln();
println("Hello, " + name);

// Reading multiple lines
println("Enter first line:");
var line1 = scanln();
println("Enter second line:");
var line2 = scanln();
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="gets">
            <div class="function-header">
                <div class="function-name">gets</div>
                <div class="function-signature">gets() -> string</div>
            </div>
            <div class="function-body">
                <div class="function-description">Alias for scanln(). Reads a complete line from standard input. Provided for C-style familiarity.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
// gets is equivalent to scanln
var name = gets();         // Same as scanln()
println("Hello, " + name);
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="scanf">
            <div class="function-header">
                <div class="function-name">scanf</div>
                <div class="function-signature">scanf(format) -> array</div>
            </div>
            <div class="function-body">
                <div class="function-description">Reads formatted input from stdin according to the format string and returns an array of parsed values. Supports %d (int), %f (float), %s (string), and other standard format specifiers.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
printf("Enter age and name: ");
var data = scanf("%d %s");
var age = data[0];         // Integer
var name = data[1];        // String

// Multiple values
printf("Enter x y z: ");
var coords = scanf("%f %f %f");
var x = coords[0];
var y = coords[1];
var z = coords[2];
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="input">
            <div class="function-header">
                <div class="function-name">input</div>
                <div class="function-signature">input(prompt) -> string</div>
            </div>
            <div class="function-body">
                <div class="function-description">Displays a prompt message and reads a line of user input. Combines output and input in one convenient function. The prompt is displayed without a trailing newline.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var name = input("Enter your name: ");
var age = input("Enter your age: ");
var confirm = input("Continue? (y/n): ");

// Building interactive menus
var choice = input("Select option (1-3): ");
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="scan">
            <div class="function-header">
                <div class="function-name">scan</div>
                <div class="function-signature">scan(delimiter) -> string</div>
            </div>
            <div class="function-body">
                <div class="function-description">Reads from standard input until the specified delimiter string is encountered. The delimiter is NOT consumed from the input buffer. Useful for parsing custom-terminated input.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var text = scan(";;");
// Reads until ";;" is encountered
// Delimiter remains in buffer for next read

// Multi-character delimiters
var paragraph = scan("\n\n");
// Reads until double newline
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="getchar">
            <div class="function-header">
                <div class="function-name">getchar</div>
                <div class="function-signature">getchar() -> string</div>
            </div>
            <div class="function-body">
                <div class="function-description">Reads a single character from standard input and returns it as a one-character string. Returns nil if EOF is encountered. Useful for character-by-character processing or single-key input.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var ch = getchar();        // Read one character
println("You typed: " + ch);

// Character-by-character processing
while (ch != "q") {
    println("Char: " + ch);
    ch = getchar();
}
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="putchar">
            <div class="function-header">
                <div class="function-name">putchar</div>
                <div class="function-signature">putchar(char_or_int) -> nil</div>
            </div>
            <div class="function-body">
                <div class="function-description">Outputs a single character to standard output. Accepts either a character literal, a string (outputs first character), or an integer (outputs corresponding Unicode character). Output is immediately flushed.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
putchar('A');              // Prints: A
putchar(65);               // Prints: A (ASCII 65)
putchar("Hello");          // Prints: H (first char only)

// Building strings character by character
putchar('H');
putchar('i');
// Output: Hi
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="puts">
            <div class="function-header">
                <div class="function-name">puts</div>
                <div class="function-signature">puts(...args) -> nil</div>
            </div>
            <div class="function-body">
                <div class="function-description">Alias for println(). Prints arguments with a trailing newline. Provided for C-style familiarity.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
// puts is equivalent to println
puts("Hello World");       // Same as println("Hello World")
puts("Value:", 42);        // Same as println("Value:", 42)
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="sprintf">
            <div class="function-header">
                <div class="function-name">sprintf</div>
                <div class="function-signature">sprintf(format, ...args) -> string</div>
            </div>
            <div class="function-body">
                <div class="function-description">Returns a formatted string using C-style format specifiers without printing it. Useful for building strings dynamically, creating messages, or formatting data for storage. Supports %s (string), %d (integer), %f (float), %x (hex), and more.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var msg = sprintf("Hello, %s! You are %d years old.", "Alice", 30);
// msg = "Hello, Alice! You are 30 years old."

// Building complex messages
var filename = "data.txt";
var size = 1024;
var info = sprintf("File: %s, Size: %d bytes", filename, size);

// Formatting numbers
var hex = sprintf("0x%x", 255);     // "0xff"
var pi = sprintf("%.2f", 3.14159); // "3.14"
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="eprintln">
            <div class="function-header">
                <div class="function-name">eprintln</div>
                <div class="function-signature">eprintln(...args) -> nil</div>
            </div>
            <div class="function-body">
                <div class="function-description">Prints arguments to standard error (stderr) with a trailing newline. Useful for logging errors, warnings, or debug information separately from normal program output. Output is unbuffered and appears immediately.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
// Error reporting
if (error_occurred) {
    eprintln("ERROR: Failed to open file");
}

// Debug logging
eprintln("DEBUG: Variable x =", x);

// Progress indicators on stderr while output goes to stdout
eprintln("Processing...");  // Goes to stderr
println("Result: done");     // Goes to stdout
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="eprintf">
            <div class="function-header">
                <div class="function-name">eprintf</div>
                <div class="function-signature">eprintf(format, ...args) -> nil</div>
            </div>
            <div class="function-body">
                <div class="function-description">Prints a formatted string to standard error (stderr). Combines the formatting capabilities of sprintf with stderr output. Ideal for error messages with variable data, formatted warnings, or structured logging to error streams.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
// Formatted error messages
var filename = "config.txt";
eprintf("ERROR: Cannot read file '%s'\n", filename);

// Status reporting with codes
var code = 404;
eprintf("HTTP Error %d: Not Found\n", code);

// Debug with formatting
eprintf("DEBUG: Array has %d elements\n", size_array(arr));
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="flush">
            <div class="function-header">
                <div class="function-name">flush</div>
                <div class="function-signature">flush() -> nil</div>
            </div>
            <div class="function-body">
                <div class="function-description">Clears the input buffer and flushes the output writer. Discards any buffered characters waiting to be read from stdin, ensuring the next input operation reads fresh user input. Also forces any buffered output to be written immediately. Useful before critical input operations or when switching between input modes.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
// Clear pending input before important read
flush();
var password = input("Enter password: ");

// Ensure output appears before long operation
printf("Processing");
flush();
// ... long operation ...

// Reset input state
flush();
println("Input buffer cleared");
{% endhighlight %}
                </div>
            </div>
        </div>
    </div>
</div>
