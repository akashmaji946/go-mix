/*
File    : go-mix/parser/switch_node.go
Author  : Akash Maji
Contact : akashmaji(@iisc.ac.in)
*/
package parser

import (
	"github.com/akashmaji946/go-mix/lexer"
	"github.com/akashmaji946/go-mix/std"
)

// SwitchCaseNode represents a single case clause in a switch statement.
// It contains the value to match against and the block of statements to execute.
type SwitchCaseNode struct {
	// Value is the expression that the switch value is compared against.
	// For case clauses, this is the literal or expression after "case".
	Value ExpressionNode

	// Body is the block of statements to execute when this case matches.
	Body BlockStatementNode

	// Token is the "case" keyword token for error reporting.
	Token lexer.Token
}

// Literal returns a string representation of the case clause for debugging.
func (scn SwitchCaseNode) Literal() string {
	return "case"
}

// Accept allows the visitor pattern to process this node.
func (scn SwitchCaseNode) Accept(visitor NodeVisitor) {
	// SwitchCaseNode is not visited directly, its components are visited separately
}

// SwitchDefaultNode represents the default clause in a switch statement.
// It contains the block of statements to execute when no case matches.
type SwitchDefaultNode struct {
	// Body is the block of statements to execute when no case matches.
	Body BlockStatementNode

	// Token is the "default" keyword token for error reporting.
	Token lexer.Token
}

// Literal returns a string representation of the default clause for debugging.
func (sdn SwitchDefaultNode) Literal() string {
	return "default"
}

// Accept allows the visitor pattern to process this node.
func (sdn SwitchDefaultNode) Accept(visitor NodeVisitor) {
	// SwitchDefaultNode is not visited directly, its components are visited separately
}

// SwitchStatementNode represents a complete switch statement.
// It contains the expression to evaluate, case clauses, and an optional default clause.
type SwitchStatementNode struct {
	// Expression is the value being switched on.
	Expression ExpressionNode

	// Cases is a slice of case clauses to match against.
	Cases []SwitchCaseNode

	// Default is the optional default clause (nil if not present).
	Default *SwitchDefaultNode

	// Token is the "switch" keyword token for error reporting.
	Token lexer.Token

	// Value stores the result of evaluating the switch statement.
	Value std.GoMixObject
}

// Literal returns a string representation of the switch statement for debugging.
func (ssn SwitchStatementNode) Literal() string {
	return "switch"
}

// Statement marks this as a statement node.
func (ssn SwitchStatementNode) Statement() {}

// Accept allows the visitor pattern to process this node.
func (ssn SwitchStatementNode) Accept(visitor NodeVisitor) {
	visitor.VisitSwitchStatementNode(ssn)
}

// GetValue returns the evaluated result of the switch statement.
func (ssn SwitchStatementNode) GetValue() std.GoMixObject {
	return ssn.Value
}

// SetValue sets the evaluated result of the switch statement.
func (ssn *SwitchStatementNode) SetValue(value std.GoMixObject) {
	ssn.Value = value
}
