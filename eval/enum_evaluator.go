/*
File    : go-mix/eval/enum_evaluator.go
Author  : Akash Maji
Contact : akashmaji(@iisc.ac.in)
*/
package eval

import (
	"github.com/akashmaji946/go-mix/parser"
	"github.com/akashmaji946/go-mix/std"
)

// evalEnumDeclaration evaluates an enum declaration statement.
//
// This method processes enum declarations by:
// 1. Creating an EnumType object to store the enum definition
// 2. Registering each enum member with its value
// 3. Binding the enum type to its name in the current scope
//
// Parameters:
//   - n: An EnumDeclarationNode containing the enum name and members
//
// Returns:
//   - objects.GoMixObject: The created EnumType object
//
// Example:
//
//	enum Color { RED, GREEN, BLUE }
//	enum Status { PENDING = 0, ACTIVE = 1, COMPLETED = 2 }
func (e *Evaluator) evalEnumDeclaration(n *parser.EnumDeclarationNode) std.GoMixObject {
	// Create a new enum type
	enumType := &std.GoMixEnum{
		Name:    n.EnumName.Name,
		Members: make(map[string]std.GoMixObject),
	}

	// Register each member
	for _, member := range n.Members {
		enumType.Members[member.Name] = member.Value
	}

	// Store the enum type in the evaluator's types map
	if e.Types == nil {
		e.Types = make(map[string]*std.GoMixStruct)
	}

	// Bind the enum type to its name in the current scope
	e.Scp.Bind(n.EnumName.Name, enumType)

	return enumType
}

// evalEnumAccessExpression evaluates enum member access expressions.
//
// This method handles accessing enum members like Color.RED or Status.ACTIVE
// by looking up the enum type and retrieving the member's value.
//
// Parameters:
//   - n: An EnumAccessExpressionNode containing the enum name and member name
//
// Returns:
//   - objects.GoMixObject: The enum member's value, or an Error if:
//   - The enum type is not found
//   - The member is not found in the enum
//
// Example:
//
//	Color.RED      // Returns the integer value of RED
//	Status.ACTIVE  // Returns the integer value of ACTIVE
func (e *Evaluator) evalEnumAccessExpression(n *parser.EnumAccessExpressionNode) std.GoMixObject {
	// Look up the enum type
	enumObj, ok := e.Scp.LookUp(n.EnumName.Name)
	if !ok {
		return e.CreateError("ERROR: enum type '%s' not found", n.EnumName.Name)
	}

	// Check if it's an enum type
	enumType, ok := enumObj.(*std.GoMixEnum)
	if !ok {
		return e.CreateError("ERROR: '%s' is not an enum type", n.EnumName.Name)
	}

	// Look up the member
	memberValue, ok := enumType.Members[n.MemberName.Name]
	if !ok {
		return e.CreateError("ERROR: enum member '%s' not found in enum '%s'", n.MemberName.Name, n.EnumName.Name)
	}

	return memberValue
}
