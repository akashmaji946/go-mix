/*
File    : go-mix/parser/switch_test.go
Author  : Akash Maji
Contact : akashmaji(@iisc.ac.in)
*/
package parser

import (
	"testing"

	"github.com/akashmaji946/go-mix/std"
	"github.com/stretchr/testify/assert"
)

// TestParseBasicSwitchStatement tests parsing of a basic switch statement with integer cases.
func TestParseBasicSwitchStatement(t *testing.T) {
	input := `
        switch (day) {
            case 1:
                dayName = "Monday";
                break;
            case 2:
                dayName = "Tuesday";
                break;
        }
    `
	par := NewParser(input)
	root := par.Parse()

	if par.HasErrors() {
		t.Fatalf("Parser has errors: %v", par.GetErrors())
	}

	if len(root.Statements) != 1 {
		t.Fatalf("Expected 1 statement, got %d", len(root.Statements))
	}

	switchNode, ok := root.Statements[0].(*SwitchStatementNode)
	if !ok {
		t.Fatalf("Expected SwitchStatementNode, got %T", root.Statements[0])
	}

	assert.NotNil(t, switchNode.Expression, "Switch expression should not be nil")
	ident, ok := switchNode.Expression.(*IdentifierExpressionNode)
	assert.True(t, ok, "Switch expression should be an identifier")
	assert.Equal(t, "day", ident.Name, "Switch expression identifier should be 'day'")

	assert.Equal(t, 2, len(switchNode.Cases), "Expected 2 case clauses")
	assert.Nil(t, switchNode.Default, "Expected no default clause")

	// Case 1
	case1 := switchNode.Cases[0]
	case1Val, ok := case1.Value.(*IntegerLiteralExpressionNode)
	assert.True(t, ok, "Case 1 value should be an integer literal")
	assert.Equal(t, int64(1), case1Val.Value.(*std.Integer).Value, "Case 1 value should be 1")
	assert.Equal(t, 2, len(case1.Body.Statements), "Case 1 body should have 2 statements")
	_, ok = case1.Body.Statements[0].(*AssignmentExpressionNode)
	assert.True(t, ok, "First statement in case 1 should be an assignment")
	_, ok = case1.Body.Statements[1].(*BreakStatementNode)
	assert.True(t, ok, "Second statement in case 1 should be a break")

	// Case 2
	case2 := switchNode.Cases[1]
	case2Val, ok := case2.Value.(*IntegerLiteralExpressionNode)
	assert.True(t, ok, "Case 2 value should be an integer literal")
	assert.Equal(t, int64(2), case2Val.Value.(*std.Integer).Value, "Case 2 value should be 2")
	assert.Equal(t, 2, len(case2.Body.Statements), "Case 2 body should have 2 statements")
}

// TestParseSwitchWithDefault tests parsing of a switch statement with string cases and a default clause.
func TestParseSwitchWithDefault(t *testing.T) {
	input := `
        switch (grade) {
            case "A":
                message = "Excellent!";
                break;
            default:
                message = "Invalid grade";
        }
    `
	par := NewParser(input)
	root := par.Parse()

	if par.HasErrors() {
		t.Fatalf("Parser has errors: %v", par.GetErrors())
	}

	if len(root.Statements) != 1 {
		t.Fatalf("Expected 1 statement, got %d", len(root.Statements))
	}

	switchNode, ok := root.Statements[0].(*SwitchStatementNode)
	if !ok {
		t.Fatalf("Expected SwitchStatementNode, got %T", root.Statements[0])
	}

	assert.Equal(t, 1, len(switchNode.Cases), "Expected 1 case clause")
	assert.NotNil(t, switchNode.Default, "Expected a default clause")

	// Case "A"
	caseA := switchNode.Cases[0]
	caseAVal, ok := caseA.Value.(*StringLiteralExpressionNode)
	assert.True(t, ok, "Case value should be a string literal")
	assert.Equal(t, "A", caseAVal.Value.(*std.String).Value, "Case value should be 'A'")
	assert.Equal(t, 2, len(caseA.Body.Statements), "Case 'A' body should have 2 statements")

	// Default
	defaultNode := switchNode.Default
	assert.Equal(t, 1, len(defaultNode.Body.Statements), "Default body should have 1 statement")
	_, ok = defaultNode.Body.Statements[0].(*AssignmentExpressionNode)
	assert.True(t, ok, "Statement in default should be an assignment")
}

// TestParseSwitchFallthrough tests parsing of a switch statement with fallthrough cases.
func TestParseSwitchFallthrough(t *testing.T) {
	input := `
        switch (month) {
            case 12:
            case 1:
            case 2:
                season = "Winter";
                break;
        }
    `
	par := NewParser(input)
	root := par.Parse()

	if par.HasErrors() {
		t.Fatalf("Parser has errors: %v", par.GetErrors())
	}

	switchNode, ok := root.Statements[0].(*SwitchStatementNode)
	if !ok {
		t.Fatalf("Expected SwitchStatementNode, got %T", root.Statements[0])
	}

	assert.Equal(t, 3, len(switchNode.Cases), "Expected 3 case clauses")

	// Case 12 (fallthrough)
	case12 := switchNode.Cases[0]
	assert.Equal(t, 0, len(case12.Body.Statements), "Case 12 body should be empty")

	// Case 1 (fallthrough)
	case1 := switchNode.Cases[1]
	assert.Equal(t, 0, len(case1.Body.Statements), "Case 1 body should be empty")

	// Case 2
	case2 := switchNode.Cases[2]
	assert.Equal(t, 2, len(case2.Body.Statements), "Case 2 body should have 2 statements")
}

// TestParseNestedSwitch tests parsing of a nested switch statement.
func TestParseNestedSwitch(t *testing.T) {
	input := `
        switch (category) {
            case "fruit":
                description = "It's a fruit.";
                switch (item) {
                    case "apple":
                        description += " an apple.";
                        break;
                    default:
                        description += " an unknown fruit.";
                }
                break;
        }
    `
	par := NewParser(input)
	root := par.Parse()

	if par.HasErrors() {
		t.Fatalf("Parser has errors: %v", par.GetErrors())
	}

	outerSwitch, ok := root.Statements[0].(*SwitchStatementNode)
	if !ok {
		t.Fatalf("Expected outer SwitchStatementNode, got %T", root.Statements[0])
	}

	assert.Equal(t, 1, len(outerSwitch.Cases), "Outer switch should have 1 case")
	outerCase := outerSwitch.Cases[0]
	assert.Equal(t, 3, len(outerCase.Body.Statements), "Outer case body should have 3 statements")

	// Check for nested switch
	nestedSwitch, ok := outerCase.Body.Statements[1].(*SwitchStatementNode)
	assert.True(t, ok, "Second statement in outer case should be a SwitchStatementNode")

	assert.Equal(t, 1, len(nestedSwitch.Cases), "Nested switch should have 1 case")
	assert.NotNil(t, nestedSwitch.Default, "Nested switch should have a default clause")
}

// TestParseSwitchWithExpression tests parsing a switch on a complex expression.
func TestParseSwitchWithExpression(t *testing.T) {
	input := `
        switch (a - b) {
            case 5:
                result = "is 5";
                break;
        }
    `
	par := NewParser(input)
	root := par.Parse()

	if par.HasErrors() {
		t.Fatalf("Parser has errors: %v", par.GetErrors())
	}

	switchNode, ok := root.Statements[0].(*SwitchStatementNode)
	if !ok {
		t.Fatalf("Expected SwitchStatementNode, got %T", root.Statements[0])
	}

	// Check switch expression
	_, ok = switchNode.Expression.(*BinaryExpressionNode)
	assert.True(t, ok, "Switch expression should be a BinaryExpressionNode")
}

// TestParseEmptySwitch tests parsing an empty switch statement.
func TestParseEmptySwitch(t *testing.T) {
	input := `switch (x) {}`
	par := NewParser(input)
	root := par.Parse()

	if par.HasErrors() {
		t.Fatalf("Parser has errors: %v", par.GetErrors())
	}

	switchNode, ok := root.Statements[0].(*SwitchStatementNode)
	if !ok {
		t.Fatalf("Expected SwitchStatementNode, got %T", root.Statements[0])
	}

	assert.Equal(t, 0, len(switchNode.Cases), "Expected 0 case clauses")
	assert.Nil(t, switchNode.Default, "Expected no default clause")
}

// TestParseSwitchWithOnlyDefault tests parsing a switch with only a default clause.
func TestParseSwitchWithOnlyDefault(t *testing.T) {
	input := `
        switch (x) {
            default:
                y = "default";
        }
    `
	par := NewParser(input)
	root := par.Parse()

	if par.HasErrors() {
		t.Fatalf("Parser has errors: %v", par.GetErrors())
	}

	switchNode, ok := root.Statements[0].(*SwitchStatementNode)
	if !ok {
		t.Fatalf("Expected SwitchStatementNode, got %T", root.Statements[0])
	}

	assert.Equal(t, 0, len(switchNode.Cases), "Expected 0 case clauses")
	assert.NotNil(t, switchNode.Default, "Expected a default clause")
	assert.Equal(t, 1, len(switchNode.Default.Body.Statements), "Default body should have 1 statement")
}

// TestParseSwitchInLoop tests parsing a switch statement inside a foreach loop.
func TestParseSwitchInLoop(t *testing.T) {
	input := `
        foreach cmd in commands {
            switch (cmd) {
                case "start":
                    break;
            }
        }
    `
	par := NewParser(input)
	root := par.Parse()

	if par.HasErrors() {
		t.Fatalf("Parser has errors: %v", par.GetErrors())
	}

	if len(root.Statements) != 1 {
		t.Fatalf("Expected 1 statement, got %d", len(root.Statements))
	}

	loopNode, ok := root.Statements[0].(*ForeachLoopStatementNode)
	if !ok {
		t.Fatalf("Expected ForeachLoopStatementNode, got %T", root.Statements[0])
	}

	if len(loopNode.Body.Statements) != 1 {
		t.Fatalf("Loop body should have 1 statement, got %d", len(loopNode.Body.Statements))
	}

	switchNode, ok := loopNode.Body.Statements[0].(*SwitchStatementNode)
	if !ok {
		t.Fatalf("Statement in loop body should be a SwitchStatementNode, got %T", loopNode.Body.Statements[0])
	}

	assert.Equal(t, 1, len(switchNode.Cases), "Switch in loop should have 1 case")
}
