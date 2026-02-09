/*
File    : go-mix/eval/evaluator_range_foreach.go
Author  : Akash Maji
Contact : akashmaji(@iisc.ac.in)
*/
package eval

import (
	"github.com/akashmaji946/go-mix/objects"
	"github.com/akashmaji946/go-mix/parser"
	"github.com/akashmaji946/go-mix/scope"
)

// evalRangeExpression evaluates range expressions to create Range objects.
//
// This method processes range expressions (e.g., 2...5) by:
// 1. Evaluating the start expression
// 2. Evaluating the end expression
// 3. Validating both are integers
// 4. Creating a Range object with the start and end values
//
// Ranges are inclusive on both ends, meaning 2...5 includes 2, 3, 4, and 5.
//
// Parameters:
//   - n: A RangeExpressionNode containing the start and end expressions
//
// Returns:
//   - objects.GoMixObject: A Range object, or an Error if:
//   - Start expression evaluation fails
//   - End expression evaluation fails
//   - Either operand is not an integer
//
// Example:
//
//	2...5        // Returns Range{Start: 2, End: 5}
//	var x = 1...10  // Creates a range from 1 to 10
func (e *Evaluator) evalRangeExpression(n *parser.RangeExpressionNode) objects.GoMixObject {
	// Evaluate start expression
	start := e.Eval(n.Start)
	if IsError(start) {
		return start
	}

	// Evaluate end expression
	end := e.Eval(n.End)
	if IsError(end) {
		return end
	}

	// Validate both are integers
	if start.GetType() != objects.IntegerType {
		return e.CreateError("ERROR: range start must be an integer, got '%s'", start.GetType())
	}
	if end.GetType() != objects.IntegerType {
		return e.CreateError("ERROR: range end must be an integer, got '%s'", end.GetType())
	}

	// Create and return the Range object
	startVal := start.(*objects.Integer).Value
	endVal := end.(*objects.Integer).Value

	return &objects.Range{
		Start: startVal,
		End:   endVal,
	}
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
func (e *Evaluator) evalForeachLoop(n *parser.ForeachLoopStatementNode) objects.GoMixObject {
	// Evaluate the iterable expression
	iterable := e.Eval(n.Iterable)
	if IsError(iterable) {
		return iterable
	}

	// Create a new scope for the entire foreach loop
	loopScope := scope.NewScope(e.Scp)
	oldScope := e.Scp
	e.Scp = loopScope

	var result objects.GoMixObject = &objects.Nil{}

	// Handle different iterable types
	switch iterable.GetType() {
	case objects.RangeType:
		// Iterate over a range
		rangeObj := iterable.(*objects.Range)
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
				e.Scp.Bind(n.Iterator.Name, &objects.Integer{Value: i})

				// Execute loop body
				result = e.Eval(&n.Body)

				// Restore to loop scope after body execution
				e.Scp = loopScope

				if IsError(result) {
					e.Scp = oldScope
					return result
				}

				// Stop if we hit a return statement
				if _, isReturn := result.(*objects.ReturnValue); isReturn {
					e.Scp = oldScope
					return result
				}

				if result.GetType() == objects.BreakType {
					e.Scp = oldScope
					return &objects.Nil{}
				}

				if result.GetType() == objects.ContinueType {
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
				e.Scp.Bind(n.Iterator.Name, &objects.Integer{Value: i})

				// Execute loop body
				result = e.Eval(&n.Body)

				// Restore to loop scope after body execution
				e.Scp = loopScope

				if IsError(result) {
					e.Scp = oldScope
					return result
				}

				// Stop if we hit a return statement
				if _, isReturn := result.(*objects.ReturnValue); isReturn {
					e.Scp = oldScope
					return result
				}

				if result.GetType() == objects.BreakType {
					e.Scp = oldScope
					return &objects.Nil{}
				}

				if result.GetType() == objects.ContinueType {
					continue
				}
			}
		}

	case objects.ArrayType:
		// Iterate over an array
		arrayObj := iterable.(*objects.Array)

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
			if _, isReturn := result.(*objects.ReturnValue); isReturn {
				e.Scp = oldScope
				return result
			}

			if result.GetType() == objects.BreakType {
				e.Scp = oldScope
				return &objects.Nil{}
			}

			if result.GetType() == objects.ContinueType {
				continue
			}
		}

	case objects.ListType:
		// Iterate over a list
		listObj := iterable.(*objects.List)

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
			if _, isReturn := result.(*objects.ReturnValue); isReturn {
				e.Scp = oldScope
				return result
			}

			if result.GetType() == objects.BreakType {
				e.Scp = oldScope
				return &objects.Nil{}
			}

			if result.GetType() == objects.ContinueType {
				continue
			}
		}

	case objects.TupleType:
		// Iterate over a tuple
		tupleObj := iterable.(*objects.Tuple)

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
			if _, isReturn := result.(*objects.ReturnValue); isReturn {
				e.Scp = oldScope
				return result
			}

			if result.GetType() == objects.BreakType {
				e.Scp = oldScope
				return &objects.Nil{}
			}

			if result.GetType() == objects.ContinueType {
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
