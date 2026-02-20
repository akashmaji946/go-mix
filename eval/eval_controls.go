/*
File    : go-mix/eval/eval_controls.go
Author  : Akash Maji
Contact : akashmaji(@iisc.ac.in)
*/
package eval

import (
	"github.com/akashmaji946/go-mix/function"
	"github.com/akashmaji946/go-mix/parser"
	"github.com/akashmaji946/go-mix/scope"
	"github.com/akashmaji946/go-mix/std"
)

// evalCallExpression evaluates function call expressions for both builtin and user-defined functions.
//
// This method handles the complete function call process:
// 1. Checks if the function name is a builtin (print, len, push, etc.)
//   - If builtin: evaluates arguments and invokes the builtin directly
//
// 2. For user-defined functions:
//   - Looks up the function in the scope chain
//   - Validates it's actually a function object
//   - Checks argument count matches parameter count
//   - Creates a new call-site scope with the function's captured scope as parent
//   - Binds arguments to parameters in the call-site scope
//   - Evaluates the function body in the new scope
//   - Unwraps return values and handles closure scope updates
//
// The scope handling is critical for closures: functions capture their defining scope,
// and when they return other functions, those returned functions get an updated scope
// that includes variables from the call site.
//
// Parameters:
//   - n: A CallExpressionNode containing the function identifier and argument expressions
//
// Returns:
//   - objects.GoMixObject: The function's return value, or an Error object if:
//   - Function not found
//   - Identifier is not a function
//   - Wrong number of arguments provided
//
// Example:
//
//	print("Hello");           // Builtin function call
//	add(5, 3);                // User-defined function call
//	makeCounter()(10);        // Closure returning a function
func (e *Evaluator) evalCallExpression(n *parser.CallExpressionNode) std.GoMixObject {

	funcName := n.FunctionIdentifier.Name

	// Support package function calls: package.function()
	if dotIdx := IndexOfDot(funcName); dotIdx > 0 {
		objName := funcName[:dotIdx]
		methodName := funcName[dotIdx+1:]

		objVal, ok := e.Scp.LookUp(objName)
		if !ok {
			return e.createError(n.FunctionIdentifier.Token, "ERROR: object not found: (%s)", objName)
		}

		// Check if it's a package (imported module)
		if pkg, isPkg := objVal.(*std.Package); isPkg {
			// Look up the function in the package
			fn, exists := pkg.Functions[methodName]
			if !exists {
				return e.createError(n.FunctionIdentifier.Token, "ERROR: function '%s' not found in package '%s'", methodName, objName)
			}
			// Evaluate arguments
			args := make([]std.GoMixObject, len(n.Arguments))
			for i, arg := range n.Arguments {
				args[i] = e.Eval(arg)
				if IsError(args[i]) {
					return args[i]
				}
			}
			// Call the package function
			return fn.Callback(e, e.Writer, args...)
		}

		// Handle struct instance method calls
		inst, ok := objVal.(*std.GoMixObjectInstance)
		if !ok {
			return e.CreateError("ERROR: (%s) is not a struct instance or package", objName)
		}
		// Prepare named parameters for method call
		params := make([]NamedParameter, len(n.Arguments))
		for i, arg := range n.Arguments {
			evaluated := e.Eval(arg)
			if IsError(evaluated) {
				return evaluated
			}
			params[i] = NamedParameter{Name: "", Value: evaluated}
		}
		result := e.callFunctionOnObject(methodName, inst, params...)
		// fmt.Printf("DEBUG: Method call result type=%s\n", result.GetType())
		return result
	}

	// look for builtin name
	if ok := e.IsBuiltin(funcName); ok {
		args := make([]std.GoMixObject, len(n.Arguments))
		for i, arg := range n.Arguments {
			args[i] = e.Eval(arg)
			if IsError(args[i]) {
				return args[i]
			}
		}
		rv := e.InvokeBuiltin(funcName, args...)
		return rv
	}

	// lookup for function name
	obj, ok := e.Scp.LookUp(funcName)
	if !ok {
		return e.createError(n.FunctionIdentifier.Token, "ERROR: function not found: (%s)", funcName)
	}
	if obj.GetType() != std.FunctionType {
		return e.createError(n.FunctionIdentifier.Token, "ERROR: not a function: (%s)", funcName)
	}
	functionObject := obj.(*function.Function)

	// Validate argument count
	expectedArgs := len(functionObject.Params)
	actualArgs := len(n.Arguments)
	if actualArgs != expectedArgs {
		return e.CreateError("ERROR: wrong number of arguments: expected %d, got %d", expectedArgs, actualArgs)
	}

	// Create a new scope with the function's captured scope as parent
	var parentScope *scope.Scope
	if functionObject.Scp != nil {
		parentScope = functionObject.Scp
	} else {
		parentScope = e.Scp
	}
	callSiteScope := scope.NewScope(parentScope)

	for i, param := range functionObject.Params {
		val := e.Eval(n.Arguments[i])
		if IsError(val) {
			return val
		}
		callSiteScope.Bind(param.Name, val)
	}
	oldScope := e.Scp
	e.Scp = callSiteScope
	result := e.Eval(functionObject.Body)
	e.Scp = oldScope

	// Unwrap return value if present
	if retVal, isReturn := result.(*std.ReturnValue); isReturn {
		returnVal := retVal.Value
		// If returning a function, update its captured scope to the current scope
		// This is essential for closures to work correctly
		// Only copy if the call site scope has variables not in the function's existing scope
		if fn, isFunc := returnVal.(*function.Function); isFunc {
			if len(callSiteScope.Variables) > len(fn.Scp.Variables) {
				fn.Scp = callSiteScope.Copy()
			}
		}
		return returnVal
	}
	return result

}

// evalImportStatement evaluates an import statement to make a package available.
//
// This method processes import statements (e.g., import math;) by:
// 1. Looking up the package in the std.Packages registry
// 2. Binding the package name to a special Package object in the current scope
//
// Once imported, package functions can be called using the dot notation
// (e.g., math.abs(), strings.upper(), etc.)
//
// Parameters:
//   - n: An ImportStatementNode containing the package name to import
func (e *Evaluator) evalImportStatement(n *parser.ImportStatementNode) std.GoMixObject {
	// Look up the package by name
	pkg, exists := e.Imports[n.Name]
	if !exists {
		return e.CreateError("ERROR: package '%s' not found", n.Name)
	}

	// Bind the package to the scope using alias if provided, otherwise use package name
	// This allows access to the package via the dot operator
	bindName := n.Name
	if n.Alias != "" {
		bindName = n.Alias
	}
	e.Scp.Bind(bindName, pkg)

	return pkg
}

// evalReturnStatement evaluates a return statement and wraps the result for propagation.
//
// This method handles the 'return' keyword by:
// 1. Evaluating the return expression to get the value to return
// 2. Wrapping the value in a ReturnValue object for special handling
//
// The ReturnValue wrapper is used to signal that evaluation should stop and
// propagate the return value up through nested blocks and function calls.
// The wrapper is unwrapped by evalCallExpression when returning from a function.
//
// Parameters:
//   - n: A ReturnStatementNode containing the expression to return
//
// Returns:
//   - objects.GoMixObject: A ReturnValue wrapper containing the evaluated expression,
//     or an Error object if the expression evaluation failed
//
// Example:
//
//	func add(a, b) {
//	    return a + b;  // Evaluates a + b, wraps in ReturnValue
//	}
func (e *Evaluator) evalReturnStatement(n *parser.ReturnStatementNode) std.GoMixObject {
	val := e.Eval(n.Expr)
	if IsError(val) {
		return val
	}
	return &std.ReturnValue{Value: val}
}
