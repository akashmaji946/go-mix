package main

import (
	"os"

	"github.com/akashmaji946/go-mix/repl"
)

var VERSION = "v0.1"
var AUTHOR = "akashmaji(@iisc.ac.in)"
var LICENCE = "MIT"
var BANNER = `                                                        
    ▄▄▄▄                       ▄▄▄  ▄▄▄     ██              
  ██▀▀▀▀█                      ███  ███     ▀▀              
 ██         ▄████▄             ████████   ████     ▀██  ██▀ 
 ██  ▄▄▄▄  ██▀  ▀██   	       ██ ██ ██     ██       ████   
 ██  ▀▀██  ██    ██   █████    ██ ▀▀ ██     ██       ▄██▄   
  ██▄▄▄██  ▀██▄▄██▀            ██    ██  ▄▄▄██▄▄▄   ▄█▀▀█▄  
    ▀▀▀▀     ▀▀▀▀              ▀▀    ▀▀  ▀▀▀▀▀▀▀▀  ▀▀▀  ▀▀▀                                                       
`
var LINE = "----------------------------------------------------------------"

func main() {
	// This will only work for arithmetic, bitwise, boolean expressions
	// For now, it works only with binary and boolean expressions involving literals
	repler := repl.NewRepl(BANNER, VERSION, AUTHOR, LINE, LICENCE)
	repler.Start(os.Stdin, os.Stdout)
}
