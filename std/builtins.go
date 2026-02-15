/*
File    : go-mix/std/builtins.go
Author  : Akash Maji
Contact : akashmaji(@iisc.ac.in)
*/

// Package std - builtins.go
// This file defines the builtin functions available in the Go-Mix language.
// It includes common functions like print, println, printf, length, and tostring,
// as well as utility functions for error creation and type conversion.
// These builtins are registered globally and can be called from Go-Mix code.
package std

import (
	"bufio"
	"io" // io.Writer is used for output operations in builtin functions
)

// Runtime defines the interface for the evaluator to allow builtins
// to call back into Go-Mix functions (e.g., for custom sorting).
type Runtime interface {
	CallFunction(fn GoMixObject, args ...GoMixObject) GoMixObject
	GetInputReader() *bufio.Reader
}

// CallbackFunc is the function signature for builtin functions.
// It takes an io.Writer for output (e.g., console) and a variadic list of GoMixObject arguments,
// returning a GoMixObject result (or an error if something goes wrong).
type CallbackFunc func(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject

// Builtin represents a builtin function with a name and its implementation callback.
// This struct is used to store and invoke builtin functions in the language.
type Builtin struct {
	Name     string       // The name of the builtin function (e.g., "print")
	Callback CallbackFunc // The function that implements the builtin behavior
}

// Builtins is a global slice of pointers to Builtin structs.
// It holds all the builtin functions available in the Go-Mix language.
// Functions are added to this slice during package initialization.
var Builtins = make([]*Builtin, 0)

// Package represents an imported package with its functions.
// It provides a way to organize builtins into namespaces.
// Packages can be imported with optional aliases (e.g., "import math as m").
// If an alias is provided, the alias is used as the namespace; otherwise, the original name is used.
type Package struct {
	Name      string              // The package name (e.g., "math")
	Alias     string              // Optional alias for the package (e.g., "m"); if empty, Name is used
	Functions map[string]*Builtin // Map of function name to builtin function
}

// GetType returns the type of the Package object
func (p *Package) GetType() GoMixType {
	return PackageType
}

// ToString returns the string representation of the package
func (p *Package) ToString() string {
	return "package:" + p.Name + ":" + p.Alias
}

// ToObject returns a detailed representation including type info
func (p *Package) ToObject() string {
	return "<package(" + p.Name + ":" + p.Alias + ")>"
}

// Packages is a global map of registered packages.
// Package registration happens during package initialization.
var Packages = make(map[string]*Package)

// RegisterPackage registers a package with the std package system.
// This allows packages to be imported and accessed via namespace (e.g., math.abs())
func RegisterPackage(pkg *Package) {
	Packages[pkg.Name] = pkg
}

// GetPackage returns a package by name, or nil if not found.
func GetPackage(name string) *Package {
	return Packages[name]
}
