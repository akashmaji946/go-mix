---
layout: default
title: OS Package - Go-Mix
description: Operating system functions including environment variables, process control, and command execution
---

<div class="content-page">
    <aside class="sidebar">
        <nav class="sidebar-nav">
            <div class="sidebar-title">OS Functions</div>
            <ul class="sidebar-menu">
                <li><a href="#import">Import</a></li>
                <li><a href="#getenv">getenv()</a></li>
                <li><a href="#setenv">setenv()</a></li>
                <li><a href="#unsetenv">unsetenv()</a></li>
                <li><a href="#exec">exec()</a></li>
                <li><a href="#getcwd">getcwd()</a></li>
                <li><a href="#getpid">getpid()</a></li>
                <li><a href="#hostname">hostname()</a></li>
                <li><a href="#user">user()</a></li>
                <li><a href="#platform">platform()</a></li>
                <li><a href="#arch">arch()</a></li>
                <li><a href="#sleep">sleep()</a></li>
                <li><a href="#exit">exit()</a></li>
                <li><a href="#args">args()</a></li>
                <li><a href="#assert">assert()</a></li>
                <li><a href="#assert_equal">assert_equal()</a></li>
                <li><a href="#assert_true">assert_true()</a></li>
                <li><a href="#assert_false">assert_false()</a></li>
            </ul>
        </nav>
    </aside>
    
    <div class="content-body">
        <h1>OS Package</h1>
        <p>Operating system interaction functions for environment variables, process control, command execution, and system information.</p>
        
        <div class="function-card" id="import">
            <div class="function-header">
                <div class="function-name">Import</div>
                <div class="function-signature">import "os"</div>
            </div>
            <div class="function-body">
                <div class="function-description">Import the os package to use namespaced functions. Provides access to system-level operations including environment variables, process management, and command execution.</div>
                <div class="function-example">
                    <h4>Examples</h4>
{% highlight go %}
// Standard import
import os;
var home = os.getenv("HOME")
os.exec("echo", "Hello")

// With alias
import os as sys;
var home = sys.getenv("HOME")
sys.exec("echo", "Hello")
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="getenv">
            <div class="function-header">
                <div class="function-name">getenv</div>
                <div class="function-signature">getenv(key) -> string</div>
            </div>
            <div class="function-body">
                <div class="function-description">Retrieves the value of an environment variable. Returns an empty string if the variable does not exist. Environment variables are system-wide key-value pairs that configure program behavior.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var home = getenv("HOME");        // e.g., "/home/username"
var path = getenv("PATH");        // System executable search paths
var missing = getenv("NONEXISTENT");  // Returns "" (empty string)

println("Home directory: " + home);
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="setenv">
            <div class="function-header">
                <div class="function-name">setenv</div>
                <div class="function-signature">setenv(key, value) -> nil</div>
            </div>
            <div class="function-body">
                <div class="function-description">Sets or updates an environment variable with the specified key and value. The change affects the current process and any child processes spawned after the call. Returns nil on success or an error if the operation fails.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
setenv("MY_APP_MODE", "production");
setenv("DEBUG_LEVEL", "2");

var mode = getenv("MY_APP_MODE");  // Returns "production"

// Overwrite existing variable
setenv("PATH", "/usr/local/bin:" + getenv("PATH"));
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="unsetenv">
            <div class="function-header">
                <div class="function-name">unsetenv</div>
                <div class="function-signature">unsetenv(key) -> nil</div>
            </div>
            <div class="function-body">
                <div class="function-description">Removes an environment variable from the current process environment. Subsequent calls to getenv for this key will return an empty string. Returns nil on success or an error if the operation fails.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
setenv("TEMP_VAR", "temporary");
println(getenv("TEMP_VAR"));      // "temporary"

unsetenv("TEMP_VAR");
println(getenv("TEMP_VAR"));      // "" (empty, variable removed)
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="exec">
            <div class="function-header">
                <div class="function-name">exec</div>
                <div class="function-signature">exec(command, ...args) -> string</div>
            </div>
            <div class="function-body">
                <div class="function-description">Executes a shell command with the provided arguments and returns the combined standard output and standard error as a string. The command runs synchronously, blocking until completion. Returns an error if the command fails to execute or returns a non-zero exit code.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
// Simple command execution
var output = exec("echo", "Hello World");
println(output);                  // "Hello World\n"

// Command with multiple arguments
var files = exec("ls", "-la", "/tmp");
println(files);

// Get current date
var date = exec("date", "+%Y-%m-%d");
println("Today is: " + date);
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="getcwd">
            <div class="function-header">
                <div class="function-name">getcwd</div>
                <div class="function-signature">getcwd() -> string</div>
            </div>
            <div class="function-body">
                <div class="function-description">Returns the current working directory of the process. This is the directory from which the program was launched or to which it has navigated. Returns an error if the directory cannot be determined.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var cwd = getcwd();
println("Current directory: " + cwd);  // e.g., "/home/user/projects"

// Useful for constructing relative paths
var configPath = cwd + "/config.json";
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="getpid">
            <div class="function-header">
                <div class="function-name">getpid</div>
                <div class="function-signature">getpid() -> int</div>
            </div>
            <div class="function-body">
                <div class="function-description">Returns the process ID (PID) of the current running process. The PID is a unique identifier assigned by the operating system, useful for logging, creating unique temporary files, or process management.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var pid = getpid();
println("Process ID: " + pid);    // e.g., 12345

// Create unique filename using PID
var tempFile = "/tmp/myapp_" + pid + ".tmp";
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="hostname">
            <div class="function-header">
                <div class="function-name">hostname</div>
                <div class="function-signature">hostname() -> string</div>
            </div>
            <div class="function-body">
                <div class="function-description">Returns the system hostname (computer name) as configured in the operating system. Useful for identifying the machine in distributed systems, logging, or configuration management.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var host = hostname();
println("Running on: " + host);     // e.g., "my-laptop" or "server01"

// Include hostname in log messages
var logPrefix = "[" + hostname() + "] ";
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="user">
            <div class="function-header">
                <div class="function-name">user</div>
                <div class="function-signature">user() -> string</div>
            </div>
            <div class="function-body">
                <div class="function-description">Returns the username of the currently logged-in user running the process. Useful for personalization, access control logging, or creating user-specific file paths.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var username = user();
println("Current user: " + username);  // e.g., "john" or "admin"

// Create user-specific directory path
var userDir = "/home/" + user() + "/documents";
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="platform">
            <div class="function-header">
                <div class="function-name">platform</div>
                <div class="function-signature">platform() -> string</div>
            </div>
            <div class="function-body">
                <div class="function-description">Returns the name of the operating system platform. Common values include "linux", "darwin" (macOS), "windows", "freebsd", etc. Useful for writing platform-specific code or logging system information.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var os = platform();
println("Operating system: " + os);  // e.g., "linux", "darwin", "windows"

// Platform-specific behavior
if (platform() == "windows") {
    exec("dir");
} else {
    exec("ls", "-la");
}
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="arch">
            <div class="function-header">
                <div class="function-name">arch</div>
                <div class="function-signature">arch() -> string</div>
            </div>
            <div class="function-body">
                <div class="function-description">Returns the system architecture (CPU architecture) of the machine. Common values include "amd64" (x86_64), "arm64", "386", "arm", etc. Useful for determining hardware capabilities or selecting appropriate binaries.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var architecture = arch();
println("Architecture: " + architecture);  // e.g., "amd64", "arm64"

// Log system information
println("Running on " + platform() + "/" + arch());
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
                <div class="function-description">Pauses program execution for the specified number of milliseconds. Useful for adding delays, rate limiting, or waiting for external resources. The sleep is non-busy, meaning the CPU is freed for other tasks during the wait.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
println("Starting...");
sleep(1000);                      // Sleep for 1 second
println("1 second later");

// Retry with backoff
for (var i = 0; i < 3; i = i + 1) {
    var result = tryConnect();
    if (result) {
        break;
    }
    sleep(500 * (i + 1));         // 500ms, 1000ms, 1500ms
}
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="exit">
            <div class="function-header">
                <div class="function-name">exit</div>
                <div class="function-signature">exit([code]) -> nil</div>
            </div>
            <div class="function-body">
                <div class="function-description">Terminates the program immediately with an optional exit code. Exit code 0 typically indicates success, while non-zero values indicate errors. The default exit code is 0 if not specified. No further code in the program will execute after exit is called.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
// Normal termination
exit(0);                          // Success exit

// Error termination
if (configError) {
    println("Configuration error occurred");
    exit(1);                      // Error exit code
}

// Default is success
exit();                           // Same as exit(0)
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="args">
            <div class="function-header">
                <div class="function-name">args</div>
                <div class="function-signature">args() -> array</div>
            </div>
            <div class="function-body">
                <div class="function-description">Returns an array of strings containing all command-line arguments passed to the program. The first element (index 0) is always the program name/path. Subsequent elements are the arguments provided by the user.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
// Program invoked as: ./myapp input.txt --verbose

var arguments = args();
println("Program: " + arguments[0]);   // "./myapp"
println("Input file: " + arguments[1]);  // "input.txt"
println("Flag: " + arguments[2]);      // "--verbose"

// Iterate over all arguments
foreach arg in arguments {
    println("Arg: " + arg);
}
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="assert">
            <div class="function-header">
                <div class="function-name">assert</div>
                <div class="function-signature">assert(condition, message) -> nil</div>
            </div>
            <div class="function-body">
                <div class="function-description">Verifies that a condition is true and raises an error with the provided message if it is not. Useful for debugging and validating assumptions during development. The program continues only if the assertion passes.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var x = 10;
assert(x > 0, "x must be positive");   // Passes, continues

var y = -5;
assert(y > 0, "y must be positive");   // Fails, raises error

// Validate function preconditions
function divide(a, b) {
    assert(b != 0, "Cannot divide by zero");
    return a / b;
}
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="assert_equal">
            <div class="function-header">
                <div class="function-name">assert_equal</div>
                <div class="function-signature">assert_equal(value1, value2, message) -> nil</div>
            </div>
            <div class="function-body">
                <div class="function-description">Verifies that two values are equal and exits the program with an error message if they are not. Performs deep equality comparison for arrays, maps, lists, and tuples. Prints a pass message on success for test visibility.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
assert_equal(2 + 2, 4, "Basic math should work");        // [PASS]
assert_equal("hello", "hello", "Strings should match"); // [PASS]

var arr1 = [1, 2, 3];
var arr2 = [1, 2, 3];
assert_equal(arr1, arr2, "Arrays should be equal");      // [PASS]

// This would exit with error:
// assert_equal(1, 2, "Values should match");          // [FAIL] + exit(1)
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="assert_true">
            <div class="function-header">
                <div class="function-name">assert_true</div>
                <div class="function-signature">assert_true(condition, message) -> nil</div>
            </div>
            <div class="function-body">
                <div class="function-description">Verifies that a condition evaluates to true. Exits with an error message if the condition is false. Prints a pass message on success. Useful for test assertions where you expect a specific condition to hold.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
assert_true(5 > 3, "Five should be greater than three");  // [PASS]
assert_true(isValid, "Data should be valid");           // [PASS] if isValid is true

function isEven(n) {
    return n % 2 == 0;
}
assert_true(isEven(4), "Four should be even");          // [PASS]

// This would exit with error:
// assert_true(3 > 5, "This is false");                  // [FAIL] + exit(1)
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="assert_false">
            <div class="function-header">
                <div class="function-name">assert_false</div>
                <div class="function-signature">assert_false(condition, message) -> nil</div>
            </div>
            <div class="function-body">
                <div class="function-description">Verifies that a condition evaluates to false. Exits with an error message if the condition is true. The inverse of assert_true, useful for checking that certain conditions do not hold.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
assert_false(3 > 5, "Three should not be greater than five");  // [PASS]
assert_false(isEmpty, "List should not be empty");             // [PASS] if isEmpty is false

function isOdd(n) {
    return n % 2 != 0;
}
assert_false(isOdd(4), "Four should not be odd");              // [PASS]

// This would exit with error:
// assert_false(5 > 3, "This is true");                         // [FAIL] + exit(1)
{% endhighlight %}
                </div>
            </div>
        </div>
    </div>
</div>
