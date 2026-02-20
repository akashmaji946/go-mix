/*
File    : go-mix/eval/eval_loops.go
Author  : Akash Maji
Contact : akashmaji(@iisc.ac.in)
*/
package eval

import (
	"github.com/akashmaji946/go-mix/parser"
	"github.com/akashmaji946/go-mix/scope"
	"github.com/akashmaji946/go-mix/std"
)

// evalForLoop evaluates for loop statements with comprehensive scope management.
//
// This method implements the classic for loop with three parts:
// 1. Initializers: Executed once before the loop starts (e.g., var i = 0)
// 2. Condition: Checked before each iteration (e.g., i < 10)
// 3. Updates: Executed after each iteration (e.g., i = i + 1)
//
// Scope management (critical for correct variable scoping):
// - Loop scope: Created for the entire loop, contains initializer variables
// - Iteration scope: Created fresh for each iteration, contains body variables
// - This two-level scoping ensures:
//   - Initializer variables persist across iterations
//   - Body variables are fresh each iteration
//   - Updates can access and modify initializer variables
//
// Control flow:
// - Loop continues while condition evaluates to true
// - Stops immediately on error or return statement
// - If no condition is provided, loops indefinitely (until return/error)
//
// Parameters:
//   - n: A ForLoopStatementNode containing initializers, condition, updates, and body
//
// Returns:
//   - objects.GoMixObject: The result of the last iteration's body, a ReturnValue if
//     a return was encountered, or an Error if evaluation failed
//
// Example:
//
//	for (var i = 0; i < 5; i = i + 1) {
//	    print(i);  // Prints 0, 1, 2, 3, 4
//	}
func (e *Evaluator) evalForLoop(n *parser.ForLoopStatementNode) std.GoMixObject {
	// Create a new scope for the entire for loop (for initializers and loop variables)
	loopScope := scope.NewScope(e.Scp)
	oldScope := e.Scp
	e.Scp = loopScope

	// Evaluate initializers in the loop scope
	for _, init := range n.Initializers {
		result := e.Eval(init)
		if IsError(result) {
			e.Scp = oldScope
			return result
		}
	}

	// Loop execution
	var result std.GoMixObject = &std.Nil{}
	for {
		// Evaluate condition if present
		if n.Condition != nil {
			condition := e.Eval(n.Condition)
			if IsError(condition) {
				e.Scp = oldScope
				return condition
			}

			// Check if condition is false
			if condition.GetType() != std.BooleanType {
				e.Scp = oldScope
				return e.CreateError("ERROR: for loop condition must be (bool)")
			}
			if !condition.(*std.Boolean).Value {
				break
			}
		}

		// Create a new scope for each iteration of the loop body
		// This ensures variables declared in the body are scoped to that iteration
		iterationScope := scope.NewScope(loopScope)
		e.Scp = iterationScope

		// Execute loop body
		result = e.Eval(&n.Body)

		// Restore to loop scope after body execution
		e.Scp = loopScope

		if IsError(result) {
			e.Scp = oldScope
			return result
		}

		// Stop if we hit a return statement
		if _, isReturn := result.(*std.ReturnValue); isReturn {
			e.Scp = oldScope
			return result
		}

		if result.GetType() == std.BreakType {
			result = &std.Nil{}
			break
		}

		if result.GetType() == std.ContinueType {
			result = &std.Nil{}
			// continue to updates
		}

		// Evaluate updates in the loop scope (not iteration scope)
		for _, update := range n.Updates {
			updateResult := e.Eval(update)
			if IsError(updateResult) {
				e.Scp = oldScope
				return updateResult
			}

			if result.GetType() == std.BreakType {
				e.Scp = oldScope
				return &std.Nil{}
			}

			if result.GetType() == std.ContinueType {
				continue
			}
		}
	}

	// Restore the original scope
	e.Scp = oldScope
	return result
}

// evalWhileLoop evaluates while loop statements with multiple condition support.
//
// This method implements while loops with the following features:
// 1. Supports multiple conditions that are implicitly AND-ed together
// 2. Creates a loop scope for the entire while loop
// 3. Creates a fresh iteration scope for each loop iteration
// 4. Continues looping while all conditions evaluate to true
// 5. Stops on error, return statement, or when any condition becomes false
//
// Scope management (similar to for loops):
// - Loop scope: Created for the entire loop, persists across iterations
// - Iteration scope: Created fresh for each iteration, contains body variables
// - This ensures variables declared in the loop body don't leak between iterations
//
// Condition evaluation:
// - All conditions must be boolean expressions
// - Conditions are evaluated in order before each iteration
// - If any condition is false, the loop terminates
// - If all conditions are true, the body executes
//
// Parameters:
//   - n: A WhileLoopStatementNode containing the condition expressions and body
//
// Returns:
//   - objects.GoMixObject: The result of the last iteration's body, a ReturnValue if
//     a return was encountered, or an Error if evaluation failed
//
// Example:
//
//	var i = 0;
//	while (i < 5) {
//	    print(i);
//	    i = i + 1;
//	}
//
//	// Multiple conditions (AND-ed together):
//	while (x > 0, y < 10) {
//	    // Continues only while both conditions are true
//	}
func (e *Evaluator) evalWhileLoop(n *parser.WhileLoopStatementNode) std.GoMixObject {
	// Create a new scope for the entire while loop
	loopScope := scope.NewScope(e.Scp)
	oldScope := e.Scp
	e.Scp = loopScope

	var result std.GoMixObject = &std.Nil{}

	for {
		// Evaluate all conditions (they should be AND-ed together)
		allTrue := true
		for _, cond := range n.Conditions {
			condition := e.Eval(cond)
			if IsError(condition) {
				e.Scp = oldScope
				return condition
			}

			if condition.GetType() != std.BooleanType {
				e.Scp = oldScope
				return e.CreateError("ERROR: while loop condition must be (bool)")
			}

			if !condition.(*std.Boolean).Value {
				allTrue = false
				break
			}
		}

		if !allTrue {
			break
		}

		// Create a new scope for each iteration of the loop body
		// This ensures variables declared in the body are scoped to that iteration
		iterationScope := scope.NewScope(loopScope)
		e.Scp = iterationScope

		// Execute loop body
		result = e.Eval(&n.Body)

		// Restore to loop scope after body execution
		e.Scp = loopScope

		if IsError(result) {
			e.Scp = oldScope
			return result
		}

		// Stop if we hit a return statement
		if _, isReturn := result.(*std.ReturnValue); isReturn {
			e.Scp = oldScope
			return result
		}

		if result.GetType() == std.BreakType {
			result = &std.Nil{}
			break
		}

		if result.GetType() == std.ContinueType {
			result = &std.Nil{}
			continue
		}
	}

	// Restore the original scope
	e.Scp = oldScope
	return result
}

// evalForeachLoop evaluates foreach loop statements with support for ranges and arrays.
//
// This method implements foreach loops with the following features:
// 1. Supports iteration over Range objects (e.g., foreach i in 2...10)
// 2. Supports iteration over Array objects (e.g., foreach item in [1,2,3])
// 3. Creates a loop scope for the entire foreach loop
// 4. Creates a fresh iteration scope for each loop iteration
// 5. Binds the iterator variable to the current value in each iteration
// 6. Stops on error or return statement
//
// Scope management:
// - Loop scope: Created for the entire loop, persists across iterations
// - Iteration scope: Created fresh for each iteration, contains iterator and body variables
// - This ensures the iterator variable is fresh each iteration
//
// Parameters:
//   - n: A ForeachLoopStatementNode containing the iterator, iterable, and body
//
// Returns:
//   - objects.GoMixObject: The result of the last iteration's body, a ReturnValue if
//     a return was encountered, or an Error if evaluation failed
//
// Example:
//
//	foreach i in 2...5 {
//	    print(i);  // Prints 2, 3, 4, 5
//	}
//
//	foreach item in [10, 20, 30] {
//	    print(item);  // Prints 10, 20, 30
//	}
func (e *Evaluator) evalForeachLoop(n *parser.ForeachLoopStatementNode) std.GoMixObject {
	// Evaluate the iterable expression
	iterable := e.Eval(n.Iterable)
	if IsError(iterable) {
		return iterable
	}

	// Create a new scope for the entire foreach loop
	loopScope := scope.NewScope(e.Scp)
	oldScope := e.Scp
	e.Scp = loopScope

	var result std.GoMixObject = &std.Nil{}

	// Handle different iterable types
	switch iterable.GetType() {
	case std.RangeType:
		// Iterate over a range
		rangeObj := iterable.(*std.Range)
		start := rangeObj.Start
		end := rangeObj.End

		// Handle both ascending and descending ranges
		if start <= end {
			// Ascending range: iterate from start to end (inclusive)
			for i := start; i <= end; i++ {
				// Create a new scope for each iteration
				iterationScope := scope.NewScope(loopScope)
				e.Scp = iterationScope

				// Bind the iterator variable to the current value
				e.Scp.Bind(n.Iterator.Name, &std.Integer{Value: i})

				// Execute loop body
				result = e.Eval(&n.Body)

				// Restore to loop scope after body execution
				e.Scp = loopScope

				if IsError(result) {
					e.Scp = oldScope
					return result
				}

				// Stop if we hit a return statement
				if _, isReturn := result.(*std.ReturnValue); isReturn {
					e.Scp = oldScope
					return result
				}

				if result.GetType() == std.BreakType {
					e.Scp = oldScope
					return &std.Nil{}
				}

				if result.GetType() == std.ContinueType {
					continue
				}
			}
		} else {
			// Descending range: iterate from start down to end (inclusive)
			for i := start; i >= end; i-- {
				// Create a new scope for each iteration
				iterationScope := scope.NewScope(loopScope)
				e.Scp = iterationScope

				// Bind the iterator variable to the current value
				e.Scp.Bind(n.Iterator.Name, &std.Integer{Value: i})

				// Execute loop body
				result = e.Eval(&n.Body)

				// Restore to loop scope after body execution
				e.Scp = loopScope

				if IsError(result) {
					e.Scp = oldScope
					return result
				}

				// Stop if we hit a return statement
				if _, isReturn := result.(*std.ReturnValue); isReturn {
					e.Scp = oldScope
					return result
				}

				if result.GetType() == std.BreakType {
					e.Scp = oldScope
					return &std.Nil{}
				}

				if result.GetType() == std.ContinueType {
					continue
				}
			}
		}

	case std.ArrayType:
		// Iterate over an array
		arrayObj := iterable.(*std.Array)

		for _, elem := range arrayObj.Elements {
			// Create a new scope for each iteration
			iterationScope := scope.NewScope(loopScope)
			e.Scp = iterationScope

			// Bind the iterator variable to the current element
			e.Scp.Bind(n.Iterator.Name, elem)

			// Execute loop body
			result = e.Eval(&n.Body)

			// Restore to loop scope after body execution
			e.Scp = loopScope

			if IsError(result) {
				e.Scp = oldScope
				return result
			}

			// Stop if we hit a return statement
			if _, isReturn := result.(*std.ReturnValue); isReturn {
				e.Scp = oldScope
				return result
			}

			if result.GetType() == std.BreakType {
				e.Scp = oldScope
				return &std.Nil{}
			}

			if result.GetType() == std.ContinueType {
				continue
			}
		}

	case std.ListType:
		// Iterate over a list
		listObj := iterable.(*std.List)

		for _, elem := range listObj.Elements {
			// Create a new scope for each iteration
			iterationScope := scope.NewScope(loopScope)
			e.Scp = iterationScope

			// Bind the iterator variable to the current element
			e.Scp.Bind(n.Iterator.Name, elem)

			// Execute loop body
			result = e.Eval(&n.Body)

			// Restore to loop scope after body execution
			e.Scp = loopScope

			if IsError(result) {
				e.Scp = oldScope
				return result
			}

			// Stop if we hit a return statement
			if _, isReturn := result.(*std.ReturnValue); isReturn {
				e.Scp = oldScope
				return result
			}

			if result.GetType() == std.BreakType {
				e.Scp = oldScope
				return &std.Nil{}
			}

			if result.GetType() == std.ContinueType {
				continue
			}
		}

	case std.TupleType:
		// Iterate over a tuple
		tupleObj := iterable.(*std.Tuple)

		for _, elem := range tupleObj.Elements {
			// Create a new scope for each iteration
			iterationScope := scope.NewScope(loopScope)
			e.Scp = iterationScope

			// Bind the iterator variable to the current element
			e.Scp.Bind(n.Iterator.Name, elem)

			// Execute loop body
			result = e.Eval(&n.Body)

			// Restore to loop scope after body execution
			e.Scp = loopScope

			if IsError(result) {
				e.Scp = oldScope
				return result
			}

			// Stop if we hit a return statement
			if _, isReturn := result.(*std.ReturnValue); isReturn {
				e.Scp = oldScope
				return result
			}

			if result.GetType() == std.BreakType {
				e.Scp = oldScope
				return &std.Nil{}
			}

			if result.GetType() == std.ContinueType {
				continue
			}
		}

	default:
		e.Scp = oldScope
		return e.CreateError("ERROR: foreach requires an `iterable`, got `%s`", iterable.GetType())
	}

	// Restore the original scope
	e.Scp = oldScope
	return result
}
