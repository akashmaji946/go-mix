/*
File    : go-mix/std/builtins.go
Author  : Akash Maji
Contact : akashmaji(@iisc.ac.in)
*/

// Package std - builtins.go
// This file defines the builtin functions available in the GoMix language.
// It includes common functions like print, println, printf, length, and tostring,
// as well as utility functions for error creation and type conversion.
// These builtins are registered globally and can be called from GoMix code.
package std

import (
	"io" // io.Writer is used for output operations in builtin functions
)

// CallbackFunc is the function signature for builtin functions.
// It takes an io.Writer for output (e.g., console) and a variadic list of GoMixObject arguments,
// returning a GoMixObject result (or an error if something goes wrong).
type CallbackFunc func(writer io.Writer, args ...GoMixObject) GoMixObject

// Builtin represents a builtin function with a name and its implementation callback.
// This struct is used to store and invoke builtin functions in the language.
type Builtin struct {
	Name     string       // The name of the builtin function (e.g., "print")
	Callback CallbackFunc // The function that implements the builtin behavior
}

// Builtins is a global slice of pointers to Builtin structs.
// It holds all the builtin functions available in the GoMix language.
// Functions are added to this slice during package initialization.
var Builtins = make([]*Builtin, 0)
