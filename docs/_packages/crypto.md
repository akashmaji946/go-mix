---
layout: default
title: Crypto Package - Go-Mix
description: Cryptographic functions including hashing and random generation
---

<div class="content-page">
    <aside class="sidebar">
        <nav class="sidebar-nav">
            <div class="sidebar-title">Crypto Functions</div>
            <ul class="sidebar-menu">
                <li><a href="#import">Import</a></li>
                <li><a href="#md5">md5()</a></li>
                <li><a href="#sha1">sha1()</a></li>
                <li><a href="#sha256">sha256()</a></li>
                <li><a href="#base64_encode">base64_encode()</a></li>
                <li><a href="#base64_decode">base64_decode()</a></li>
                <li><a href="#uuid">uuid()</a></li>
                <li><a href="#random">random()</a></li>
                
            </ul>
        </nav>
    </aside>
    
    <div class="content-body">
        <h1>Crypto Package</h1>
        <p>Cryptographic hashing, encoding, and random generation functions.</p>
        
        <div class="function-card" id="import">
            <div class="function-header">
                <div class="function-name">Import</div>
                <div class="function-signature">import "crypto"</div>
            </div>
            <div class="function-body">
                <div class="function-description">Import the crypto package to use namespaced functions.</div>
                <div class="function-example">
                    <h4>Examples</h4>
{% highlight go %}
// Standard import
import crypto;
var hash = crypto.md5("hello");
var uuid = crypto.uuid();

// With alias
import "crypto" as c
var hash = c.md5("hello")
var uuid = c.uuid()
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="md5">
            <div class="function-header">
                <div class="function-name">md5</div>
                <div class="function-signature">md5(str) -> string</div>
            </div>
            <div class="function-body">
                <div class="function-description">Returns MD5 hash of string.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var hash = md5("hello");
// 5d41402abc4b2a76b9719d911017c592
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="sha1">
            <div class="function-header">
                <div class="function-name">sha1</div>
                <div class="function-signature">sha1(str) -> string</div>
            </div>
            <div class="function-body">
                <div class="function-description">Returns SHA1 hash of string.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var hash = sha1("hello");
// aaf4c61ddcc5e8a2dabede0f3b482cd9aea9434d
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="sha256">
            <div class="function-header">
                <div class="function-name">sha256</div>
                <div class="function-signature">sha256(str) -> string</div>
            </div>
            <div class="function-body">
                <div class="function-description">Returns SHA256 hash of string.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var hash = sha256("hello");
// 2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="base64_encode">
            <div class="function-header">
                <div class="function-name">base64_encode</div>
                <div class="function-signature">base64_encode(str) -> string</div>
            </div>
            <div class="function-body">
                <div class="function-description">Encodes string to Base64.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var encoded = base64_encode("hello");
// aGVsbG8=
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="base64_decode">
            <div class="function-header">
                <div class="function-name">base64_decode</div>
                <div class="function-signature">base64_decode(str) -> string</div>
            </div>
            <div class="function-body">
                <div class="function-description">Decodes Base64 string.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var decoded = base64_decode("aGVsbG8=");
// hello
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="uuid">
            <div class="function-header">
                <div class="function-name">uuid</div>
                <div class="function-signature">uuid() -> string</div>
            </div>
            <div class="function-body">
                <div class="function-description">Generates random UUID v4.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var id = uuid();
// 550e8400-e29b-41d4-a716-446655440000
{% endhighlight %}
                </div>
            </div>
        </div>
        
        <div class="function-card" id="random">
            <div class="function-header">
                <div class="function-name">random</div>
                <div class="function-signature">random() -> float</div>
            </div>
            <div class="function-body">
                <div class="function-description">Returns random float between 0 and 1.</div>
                <div class="function-example">
                    <h4>Example</h4>
{% highlight go %}
var r = random();          // 0.0 to 1.0
var scaled = random() * 100;  // 0.0 to 100.0
{% endhighlight %}
                </div>
            </div>
        </div>
    </div>
</div>
