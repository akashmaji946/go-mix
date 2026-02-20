/*
File    : go-mix/eval/eval_statements.go
Author  : Akash Maji
Contact : akashmaji(@iisc.ac.in)
*/
package eval

import (
	"github.com/akashmaji946/go-mix/lexer"
	"github.com/akashmaji946/go-mix/parser"
	"github.com/akashmaji946/go-mix/std"
)

// evalBlockStatement evaluates a sequence of statements within a block.
//
// This method processes statement blocks (code between { and }) by delegating to
// evalStatements. Blocks are used in function bodies, if-else branches, and loops.
// The method returns the result of the last statement in the block, or stops early
// if a return statement or error is encountered.
//
// Note: This method does NOT create a new scope - scope creation is handled by
// the constructs that use blocks (functions, loops, etc.).
//
// Parameters:
//   - n: A BlockStatementNode containing a list of statements to evaluate
//
// Returns:
//   - objects.GoMixObject: The result of the last statement, a ReturnValue if a
//     return statement was encountered, or an Error if evaluation failed
//
// Example:
//
//	{
//	    var x = 10;
//	    var y = 20;
//	    x + y;  // Block returns 30
//	}
func (e *Evaluator) evalBlockStatement(n *parser.BlockStatementNode) std.GoMixObject {
	return e.evalStatements(n.Statements)
}

// evalDeclarativeStatement handles variable declarations with var, const, and let keywords.
//
// This method processes variable declarations by:
// 1. Evaluating the initialization expression to get the initial value
// 2. Checking for redeclaration conflicts in the current scope
// 3. Binding the variable to its value in the current scope
// 4. Recording special properties based on the declaration keyword:
//   - 'const': Marks the variable as immutable (stored in Consts map)
//   - 'let': Marks the variable as type-safe and records its type (stored in LetVars and LetTypes)
//   - 'var': Standard mutable variable with no type restrictions
//
// The distinction between declaration types affects later assignment operations:
// - const variables cannot be reassigned
// - let variables can only be assigned values of the same type
// - var variables can be reassigned to any type
//
// Parameters:
//   - n: A DeclarativeStatementNode containing the keyword, identifier, and initialization expression
//
// Returns:
//   - objects.GoMixObject: The initialized value on success, or an Error object if:
//   - The initialization expression fails
//   - The variable is already declared in the current scope
//
// Example:
//
//	var x = 10;      // Mutable, any type
//	const PI = 3.14; // Immutable
//	let name = "Go"; // Type-safe (must remain string)
func (e *Evaluator) evalDeclarativeStatement(n *parser.DeclarativeStatementNode) std.GoMixObject {
	// fmt.Printf("DEBUG: evalDeclarativeStatement for '%s', expr type=%T\n", n.Identifier.Name, n.Expr)
	val := e.Eval(n.Expr)
	// fmt.Printf("DEBUG: evalDeclarativeStatement result type=%s\n", val.GetType())
	if IsError(val) {
		return val
	}

	// redeclared?
	_, has := e.Scp.Bind(n.Identifier.Name, val)
	if has {
		return e.CreateError("ERROR: identifier redeclaration found: (%s)", n.Identifier.Name)
	}

	if n.VarToken.Type == lexer.CONST_KEY {
		e.Scp.Consts[n.Identifier.Name] = true
	} else if n.VarToken.Type == lexer.LET_KEY {
		e.Scp.LetVars[n.Identifier.Name] = true
		e.Scp.LetTypes[n.Identifier.Name] = val.GetType()
	}
	e.Scp.Bind(n.Identifier.Name, val)
	return val
}

// evalStatements evaluates a sequence of statements in order, with early termination support.
//
// This method processes a list of statements sequentially and implements two important
// control flow behaviors:
//  1. Error propagation: If any statement produces an error, evaluation stops immediately
//     and the error is returned
//  2. Return handling: If any statement produces a ReturnValue, evaluation stops and the
//     return value is propagated (used to exit from functions early)
//
// For normal execution, the method continues through all statements and returns the
// result of the last one. If the list is empty, returns Nil.
//
// Parameters:
//   - stmts: A slice of StatementNode objects to evaluate in sequence
//
// Returns:
//   - objects.GoMixObject: The result of the last statement, a ReturnValue if a return
//     was encountered, an Error if any statement failed, or Nil for an empty list
//
// Example:
//
//	var x = 10;
//	var y = 20;
//	return x + y;  // Stops here, returns 30
//	var z = 30;    // Never executed
func (e *Evaluator) evalStatements(stmts []parser.StatementNode) std.GoMixObject {
	var result std.GoMixObject = &std.Nil{}
	for _, stmt := range stmts {
		result = e.Eval(stmt)

		if IsError(result) {
			return result
		}
		// Stop evaluation if we hit a return statement
		if _, isReturn := result.(*std.ReturnValue); isReturn {
			return result
		}
		// Stop evaluation if we hit break or continue
		if result.GetType() == std.BreakType || result.GetType() == std.ContinueType {
			return result
		}
	}
	return result
}
