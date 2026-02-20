/*
File    : go-mix/std/enum.go
Author  : Akash Maji
Contact : akashmaji(@iisc.ac.in)
*/
package std

import "fmt"

// GoMixEnum represents an enum type definition in Go-Mix.
// Enums are user-defined types with a fixed set of named constants.
//
// Example:
//
//	enum Color { RED, GREEN, BLUE }
//	enum Status { PENDING = 0, ACTIVE = 1, COMPLETED = 2 }
type GoMixEnum struct {
	Name    string                 // The name of the enum type
	Members map[string]GoMixObject // Map of member names to their values
}

// GetType returns the type of the GoMixEnum object
func (e *GoMixEnum) GetType() GoMixType {
	return EnumType
}

// ToString returns a string representation of the enum type
func (e *GoMixEnum) ToString() string {
	result := fmt.Sprintf("enum %s {", e.Name)
	first := true
	for name, value := range e.Members {
		if !first {
			result += ", "
		}
		result += name
		if value != nil {
			result += " = " + value.ToString()
		}
		first = false
	}
	result += "}"
	return result
}

// ToObject returns a detailed representation of the enum type
func (e *GoMixEnum) ToObject() string {
	return fmt.Sprintf("<enum(%s)>", e.Name)
}

// GetMember returns the value of a specific enum member
func (e *GoMixEnum) GetMember(name string) (GoMixObject, bool) {
	val, ok := e.Members[name]
	return val, ok
}
