/*
File    : go-mix/eval/switch_evaluator.go
Author  : Akash Maji
Contact : akashmaji(@iisc.ac.in)
*/
package eval

import (
	"github.com/akashmaji946/go-mix/parser"
	"github.com/akashmaji946/go-mix/std"
)

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

// switchValuesEqual compares two GoMixObject values for equality in switch statements.
// It handles different types appropriately.
func switchValuesEqual(left, right std.GoMixObject) bool {
	// Handle nil values
	if left.GetType() == std.NilType && right.GetType() == std.NilType {
		return true
	}
	if left.GetType() == std.NilType || right.GetType() == std.NilType {
		return false
	}

	// Same type comparison
	if left.GetType() != right.GetType() {
		// Allow int/float comparison
		if (left.GetType() == std.IntegerType || left.GetType() == std.FloatType) &&
			(right.GetType() == std.IntegerType || right.GetType() == std.FloatType) {
			return switchToFloat64(left) == switchToFloat64(right)
		}
		return false
	}

	// Type-specific comparisons
	switch left.GetType() {
	case std.IntegerType:
		return left.(*std.Integer).Value == right.(*std.Integer).Value
	case std.FloatType:
		return left.(*std.Float).Value == right.(*std.Float).Value
	case std.BooleanType:
		return left.(*std.Boolean).Value == right.(*std.Boolean).Value
	case std.StringType:
		return left.(*std.String).Value == right.(*std.String).Value
	case std.CharType:
		return left.(*std.Char).Value == right.(*std.Char).Value
	default:
		// For other types, use string comparison as fallback
		return left.ToString() == right.ToString()
	}
}

// switchToFloat64 converts a GoMixObject to float64 for numeric comparisons.
func switchToFloat64(obj std.GoMixObject) float64 {
	if obj.GetType() == std.IntegerType {
		return float64(obj.(*std.Integer).Value)
	} else if obj.GetType() == std.FloatType {
		return obj.(*std.Float).Value
	}
	return 0
}
