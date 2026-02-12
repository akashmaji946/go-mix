/*
File    : go-mix/std/struct.go
Author  : Akash Maji
Contact : akashmaji(@iisc.ac.in)
*/

// Package std - struct.go
// This file defines the GoMixStruct type which represents user-defined struct types.
package std

import "fmt"

// FunctionInterface defines the interface for function objects to avoid circular imports
type FunctionInterface interface {
	GetName() string
	GetType() GoMixType
	GetParameters() []string
	GetBody() string
	ToString() string
}

// GoMixStruct represents a user-defined struct type in Go-Mix.
// It stores the struct name and a list of methods associated with it.
type GoMixStruct struct {
	Name        string                       // Name of the struct type
	Methods     map[string]FunctionInterface // Slice of method objects (using interface to avoid circular imports)
	FieldNodes  []interface{}                // AST nodes for field declarations (interface{} to avoid import cycle)
	ClassFields map[string]GoMixObject       // Map of class fields (if needed)
	ConstFields map[string]bool              // Set of constant field names
	LetFields   map[string]bool              // Set of let field names
	LetTypes    map[string]GoMixType         // Map of let field types
}

// GetConstructor returns the constructor function for the struct instance,
// which is the "init" method if it exists.
func (o *GoMixStruct) GetConstructor() (FunctionInterface, bool) {
	return o.GetMethod("init")
}

// GetMethod retrieves a method by name from the struct's methods.
// It returns the method and a boolean indicating if it was found.
func (g *GoMixStruct) GetMethod(name string) (FunctionInterface, bool) {
	method, found := g.Methods[name]
	return method, found
}

// GetName returns the name of the struct type.
func (g *GoMixStruct) GetName() string {
	return g.Name
}

// Add adds a method to the struct, checking for duplicates.
func (g *GoMixStruct) Add(fn FunctionInterface) error {
	methodName := fn.GetName()

	_, found := g.Methods[methodName]

	if found {
		return fmt.Errorf("method '%s' already exists in struct '%s'", methodName, g.Name)
	}

	g.Methods[methodName] = fn
	return nil
}

// GetType returns the type of the struct, which is "STRUCT".
func (g *GoMixStruct) GetType() GoMixType {
	return StructType
}

// ToString returns the string representation of the struct in the format "struct(name)".
func (g *GoMixStruct) ToString() string {
	fieldStr := ""
	for name, val := range g.ClassFields {
		fieldStr += fmt.Sprintf("%s:%s; ", name, val.ToString())
	}

	methodStr := ""
	for name := range g.Methods {
		args := ""
		method, _ := g.GetMethod(name)
		for i, param := range method.GetParameters() {
			if i > 0 {
				args += ","
			}
			args += param
		}
		methodStr += fmt.Sprintf("%s(%s); ", name, args)
	}
	return fmt.Sprintf("struct(%s) {\nstatic fields: %s\nmethods: %s\n}", g.Name, fieldStr, methodStr)
}

// ToObject returns the detailed string representation of the struct including methods.
func (g *GoMixStruct) ToObject() string {
	methodStr := ""
	for name := range g.Methods {
		methodStr += fmt.Sprintf("\n  %s", name)
	}
	return fmt.Sprintf("struct(%s) {%s}", g.Name, methodStr)
}

// GoMixObjectInstance represents an instance of a struct type, holding field values and a reference to its struct definition.
type GoMixObjectInstance struct {
	Struct         *GoMixStruct           // Reference to the struct definition
	InstanceFields map[string]GoMixObject // Map of field names to their values
	ClassFields    map[string]GoMixObject // Map of class fields (if needed)
}

// NewStructInstance creates a new instance of a struct type given the struct definition.
func NewStructInstance(s *GoMixStruct) *GoMixObjectInstance {
	return &GoMixObjectInstance{
		Struct:         s,
		InstanceFields: make(map[string]GoMixObject), // Initialize fields map (if needed)
	}
}

// GetType returns the type of the struct instance, which is "OBJECT".
func (o *GoMixObjectInstance) GetType() GoMixType {
	return ObjectType
}

// ToString returns the string representation of the struct instance in the format "object(structName)".
func (o *GoMixObjectInstance) ToObject() string {
	return fmt.Sprintf("object(%s)", o.Struct.Name)
}

// ToObject returns the detailed string representation of the struct instance including its struct type and fields.
func (o *GoMixObjectInstance) ToString() string {
	// include all fields and their values in the output
	// include all methods and their params
	fieldStr := ""

	for name, val := range o.InstanceFields {
		fieldStr += fmt.Sprintf("%s:%s; ", name, val.ToString())
	}

	classFieldStr := ""
	if o.Struct != nil {
		for name, val := range o.Struct.ClassFields {
			classFieldStr += fmt.Sprintf("%s:%s; ", name, val.ToString())
		}
	}

	methodStr := ""
	for name := range o.Struct.Methods {
		args := ""
		method, _ := o.Struct.GetMethod(name)
		for i, param := range method.GetParameters() {
			if i > 0 {
				args += ","
			}
			args += param
		}
		methodStr += fmt.Sprintf("%s(%s); ", name, args)
	}
	res := fmt.Sprintf("object(%s) {\nfields: %s\n", o.Struct.Name, fieldStr)
	if classFieldStr != "" {
		res += fmt.Sprintf("static fields: %s\n", classFieldStr)
	}
	res += fmt.Sprintf("methods: %s\n}", methodStr)
	return res
}
