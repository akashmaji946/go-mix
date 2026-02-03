package main

import (
	"os"

	"github.com/akashmaji946/go-mix/repl"
)

var VERSION = "v0.1"
var AUTHOR = "akashmaji(@iisc.ac.in)"
var LICENCE = "MIT"
var PROMPT = "GO-MIX >>> "
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
	repler := repl.NewRepl(BANNER, VERSION, AUTHOR, LINE, LICENCE, PROMPT)
	repler.Start(os.Stdin, os.Stdout)
}
