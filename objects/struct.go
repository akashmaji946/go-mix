// Package objects - struct.go
// This file defines the GoMixStruct type which represents user-defined struct types.
/*
File    : go-mix/objects/struct.go
Author  : Akash Maji
Contact : akashmaji(@iisc.ac.in)
*/
package objects

import "fmt"

// FunctionInterface defines the interface for function objects to avoid circular imports
type FunctionInterface interface {
	GetName() string
	GetType() GoMixType
	ToString() string
}

// GoMixStruct represents a user-defined struct type in GoMix.
// It stores the struct name and a list of methods associated with it.
type GoMixStruct struct {
	Name    string              // Name of the struct type
	Methods []FunctionInterface // Slice of method objects (using interface to avoid circular imports)
}

// Add adds a method to the struct, checking for duplicates.
func (g *GoMixStruct) Add(fn FunctionInterface) error {
	methodName := fn.GetName()

	// Check for duplicates
	for _, m := range g.Methods {
		if m.GetName() == methodName {
			return fmt.Errorf("method with name '%s' already exists in struct '%s'", methodName, g.Name)
		}
	}

	g.Methods = append(g.Methods, fn)
	return nil
}

// GetType returns the type of the struct, which is "STRUCT".
func (g *GoMixStruct) GetType() GoMixType {
	return STRUCT_TYPE
}

// ToString returns the string representation of the struct in the format "struct(name)".
func (g *GoMixStruct) ToString() string {
	return fmt.Sprintf("struct(%s)", g.Name)
}

// ToObject returns the detailed string representation of the struct including methods.
func (g *GoMixStruct) ToObject() string {
	methodStr := ""
	for i, method := range g.Methods {
		if i > 0 {
			methodStr += ", "
		}
		methodStr += method.GetName()
	}
	return fmt.Sprintf("struct(%s) {%s}", g.Name, methodStr)
}
