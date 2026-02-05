/*
File    : go-mix/function/function.go
Author  : Akash Maji
Contact : akashmaji(@iisc.ac.in)
*/
package function

import (
	"fmt"

	"github.com/akashmaji946/go-mix/objects"
	"github.com/akashmaji946/go-mix/parser"
	"github.com/akashmaji946/go-mix/scope"
)

// Function represents a user-defined function object in GoMix.
// It captures the function's name, parameters, body, and the scope
// in which it was defined (for closure support).
//
// Fields:
//   - Name: The name of the function as declared in the source code.
//     Used for identification and debugging purposes.
//   - Params: A slice of identifier nodes representing the function's
//     parameter names. These are bound to argument values when the
//     function is called.
//   - Body: A block statement node containing the function's executable
//     statements. This is evaluated when the function is invoked.
//   - Scp: A pointer to the scope in which the function was defined.
//     This enables closure behavior, allowing the function to access
//     variables from its enclosing scope even after that scope has
//     finished executing.
type Function struct {
	Name   string                             // Name of the function
	Params []*parser.IdentifierExpressionNode // Function parameter names
	Body   *parser.BlockStatementNode         // Function body (statements to execute)
	Scp    *scope.Scope                       // Captured scope for closures
}

// GetType returns the type identifier for this Function object.
// This implements the objects.GoMixObject interface.
// The function type is represented as "func" in the GoMix type system.
//
// Returns:
//   - objects.GoMixType: The string "func" indicating this is a function object
func (f *Function) GetType() objects.GoMixType {
	return "func"
}

// ToString returns a simple string representation of the function.
// This is used for basic display purposes and debugging.
// The format is: "func(functionName)"
//
// Example:
//
//	If f.Name = "add", this returns: "func(add)"
//
// Returns:
//   - string: A formatted string representation of the function
func (f *Function) ToString() string {
	return fmt.Sprintf("func(%s)", f.Name)
}

// ToObject returns a detailed string representation of the function,
// including its name and parameter names. This is useful for debugging,
// inspection, and displaying function information to users.
//
// The format is: "<func[name(param1, param2, ...)]>"
//
// Example:
//
//	If f.Name = "add" and Params = ["a", "b"], this returns:
//	"<func[add(a, b)]>"
//
// Returns:
//   - string: A detailed string representation including name and parameters
func (f *Function) ToObject() string {
	// Build a comma-separated list of parameter names
	args := ""
	for i, param := range f.Params {
		if i > 0 {
			args += ", " // Add comma between parameters
		}
		args += param.Name
	}
	// Return the formatted function representation
	return fmt.Sprintf("<func[%s(%s)]>", f.Name, args)
}
