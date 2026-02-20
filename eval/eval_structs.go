/*
File    : go-mix/eval/eval_structs.go
Author  : Akash Maji
Contact : akashmaji(@iisc.ac.in)
*/
package eval

import (
	"github.com/akashmaji946/go-mix/function"
	"github.com/akashmaji946/go-mix/lexer"
	"github.com/akashmaji946/go-mix/parser"
	"github.com/akashmaji946/go-mix/scope"
	"github.com/akashmaji946/go-mix/std"
)

// evalStructDeclaration evaluates a struct declaration statement.
//
// This method creates a new GoMixStruct type definition. It processes:
// - Fields: Evaluates initial values and registers them as static fields
// - Methods: Creates Function objects and registers them
// - Const/Let/Var modifiers: Records field properties
//
// The resulting struct type is bound to its name in the current scope.
//
// Parameters:
//   - n: The StructDeclarationNode from the AST
//
// Returns:
//   - objects.GoMixObject: The created GoMixStruct object
func (e *Evaluator) evalStructDeclaration(n *parser.StructDeclarationNode) std.GoMixObject {
	// Create a new struct type with the given name and fields
	s := &std.GoMixStruct{
		Name:        n.StructName.Name,
		Methods:     make(map[string]std.FunctionInterface),
		FieldNodes:  make([]interface{}, len(n.Fields)),
		ClassFields: make(map[string]std.GoMixObject),
		ConstFields: make(map[string]bool),
		LetFields:   make(map[string]bool),
		LetTypes:    make(map[string]std.GoMixType),
	}

	for i, f := range n.Fields {
		s.FieldNodes[i] = f
		val := e.Eval(f.Expr)
		if IsError(val) {
			return val
		}
		s.ClassFields[f.Identifier.Name] = val
		if f.VarToken.Type == lexer.CONST_KEY {
			s.ConstFields[f.Identifier.Name] = true
		} else if f.VarToken.Type == lexer.LET_KEY {
			s.LetFields[f.Identifier.Name] = true
			s.LetTypes[f.Identifier.Name] = val.GetType()
		}
	}

	for _, m := range n.Methods {
		method := &function.Function{
			Name:   m.FuncName.Name,
			Params: m.FuncParams,
			Body:   &m.FuncBody,
			Scp:    e.Scp, // Capture the current scope for closures
		}
		if err := s.Add(method); err != nil {
			return e.CreateError("ERROR: struct method '%s' already defined", method.Name)
		}
	}

	e.Types[s.Name] = s
	e.Scp.Bind(s.Name, s)
	return s
}

// evalNewCallExpression evaluates a 'new' expression to instantiate a struct.
//
// This method handles object creation:
// 1. Looks up the struct type
// 2. Creates a new instance
// 3. Calls the constructor ('init' method) if it exists
//
// Parameters:
//   - n: The NewCallExpressionNode
//
// Returns:
//   - objects.GoMixObject: The new struct instance, or an Error if the struct type is not found
func (e *Evaluator) evalNewCallExpression(n *parser.NewCallExpressionNode) std.GoMixObject {
	// Look up the struct type by name
	s, exists := e.Types[n.StructName.Name]
	if !exists {
		return e.CreateError("ERROR: struct type '%s' not defined", n.StructName.Name)
	}

	inst := std.NewStructInstance(s)

	// Initialize fields from struct definition
	initMethod, hasInit := s.GetConstructor()
	if hasInit {
		// Cast to Function to access Body and Params directly
		fn, ok := initMethod.(*function.Function)
		if !ok {
			return e.CreateError("ERROR: constructor method is not a valid function")
		}

		if len(n.Arguments) != len(fn.Params) {
			return e.CreateError("ERROR: constructor for struct '%s' expects %d arguments, got %d", s.Name, len(fn.Params), len(n.Arguments))
		}

		// Save the current scope before creating a new one
		oldScope := e.Scp

		// Create a new scope for the constructor call
		constructorScope := scope.NewScope(e.Scp)
		constructorScope.Bind("this", inst) // Set 'this' to the new instance

		// Evaluate the constructor with the given arguments
		for i, arg := range n.Arguments {
			argValue := e.Eval(arg)
			if IsError(argValue) {
				e.Scp = oldScope
				return argValue
			}
			constructorScope.Bind(fn.Params[i].Name, argValue)
		}

		// Switch to the constructor scope
		e.Scp = constructorScope

		// Execute the constructor body
		result := e.Eval(fn.Body)
		if IsError(result) {
			e.Scp = oldScope
			return result
		}

		// Restore the original scope
		e.Scp = oldScope
	}
	return inst
}

// callFunctionOnObject invokes a method on a struct instance.
//
// This method handles the mechanics of method dispatch:
// 1. Looks up the method in the struct definition
// 2. Creates a new scope for the method execution
// 3. Binds 'this' to the instance and 'self' to the struct type
// 4. Binds arguments to parameters
// 5. Evaluates the method body
//
// Parameters:
//   - name: The name of the method to call
//   - obj: The struct instance on which the method is called
//   - args: The arguments to pass to the method
//
// Returns:
//   - objects.GoMixObject: The return value of the method
func (e *Evaluator) callFunctionOnObject(name string, obj *std.GoMixObjectInstance, args ...NamedParameter) std.GoMixObject {

	initMethodInterface, exists := obj.Struct.Methods[name]
	if !exists {
		return e.CreateError("ERROR: method (%s) not found in struct (%s)", name, obj.Struct.GetName())
	}

	initMethod, ok := initMethodInterface.(*function.Function)
	if !ok {
		return e.CreateError("ERROR: method (%s) not found in struct (%s)", name, obj.Struct.GetName())
	}

	// Create a new scope for the method call with the struct instance's scope as parent
	methodScope := scope.NewScope(e.Scp)

	// Bind the struct instance to a special variable (e.g., "self") in the method scope
	methodScope.Bind("this", obj)
	methodScope.Bind("self", obj.Struct)
	for _, arg := range args {
		methodScope.Bind(arg.Name, arg.Value)
	}

	// Save the current scope and switch to the method scope for evaluation
	oldScope := e.Scp
	e.Scp = methodScope
	res := e.Eval(initMethod.Body)
	e.Scp = oldScope
	if res.GetType() == std.ErrorType {
		return res
	}
	return UnwrapReturnValue(res)
}

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
