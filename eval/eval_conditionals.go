/*
File    : go-mix/eval/eval_conditionals.go
Author  : Akash Maji
Contact : akashmaji(@iisc.ac.in)
*/
package eval

import (
	"github.com/akashmaji946/go-mix/parser"
	"github.com/akashmaji946/go-mix/std"
)

// evalIfExpression evaluates if-else conditional expressions.
//
// This method implements conditional branching by:
// 1. Evaluating the condition expression
// 2. Validating that the condition is a boolean type
// 3. Executing the then-block if the condition is true
// 4. Executing the else-block if the condition is false
//
// Both branches are represented as BlockStatementNodes, and the method returns
// the result of whichever branch is executed. If the condition is not a boolean,
// an error is returned.
//
// Parameters:
//   - n: An IfExpressionNode containing the condition and both branch blocks
//
// Returns:
//   - objects.GoMixObject: The result of the executed branch (then or else),
//     or an Error object if the condition is not a boolean type
//
// Example:
//
//	if (x > 10) {
//	    return "big";
//	} else {
//	    return "small";
//	}
func (e *Evaluator) evalIfExpression(n *parser.IfExpressionNode) std.GoMixObject {
	condition := e.Eval(n.Condition)
	if IsError(condition) {
		return condition
	}

	if condition.GetType() != std.BooleanType {
		return e.CreateError("ERROR: conditional expression must be (bool)")
	}
	if condition.(*std.Boolean).Value {
		return e.Eval(&n.ThenBlock)
	}
	return e.Eval(&n.ElseBlock)
}

// evalSwitchStatement evaluates a switch statement.
// It evaluates the switch expression, then matches it against each case value.
// The first matching case's body is executed. If no case matches and a default
// clause exists, the default body is executed.
func (e *Evaluator) evalSwitchStatement(node parser.SwitchStatementNode) std.GoMixObject {
	switchValue := e.Eval(node.Expression)
	if IsError(switchValue) {
		return switchValue
	}

	// Find the starting case
	startCase := -1
	for i, caseNode := range node.Cases {
		caseValue := e.Eval(caseNode.Value)
		if IsError(caseValue) {
			return caseValue
		}
		if switchValuesEqual(switchValue, caseValue) {
			startCase = i
			break
		}
	}

	// If no case matched, jump to default if it exists
	if startCase == -1 {
		if node.Default != nil {
			result := e.evalBlockStatement(&node.Default.Body)
			// A break in default is meaningless but shouldn't crash.
			// Only propagate return/error.
			if result != nil && (result.GetType() == std.ReturnValueType || IsError(result)) {
				return result
			}
		}
		// If no match and no default, we're done.
		return &std.Nil{}
	}

	// Execute cases from the matched one, handling fallthrough
	for i := startCase; i < len(node.Cases); i++ {
		result := e.evalBlockStatement(&node.Cases[i].Body)
		if result != nil {
			if result.GetType() == std.BreakType {
				return &std.Nil{} // Normal exit from switch
			}
			if result.GetType() == std.ReturnValueType || IsError(result) {
				return result // Propagate return/error up
			}
		}
	}

	// If we fell through all cases, execute default if it exists
	if node.Default != nil {
		result := e.evalBlockStatement(&node.Default.Body)
		if result != nil && (result.GetType() == std.ReturnValueType || IsError(result)) {
			return result
		}
	}

	return &std.Nil{}
}
