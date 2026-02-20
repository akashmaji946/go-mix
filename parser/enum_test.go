/*
File    : go-mix/parser/enum_test.go
Author  : Akash Maji
Contact : akashmaji(@iisc.ac.in)
*/
package parser

import (
	"testing"

	"github.com/akashmaji946/go-mix/std"
)

// TestParseBasicEnum tests parsing of basic enum declarations
func TestParseBasicEnum(t *testing.T) {
	input := `enum Color { RED, GREEN, BLUE }`
	par := NewParser(input)

	root := par.Parse()

	if par.HasErrors() {
		t.Errorf("Parser has errors: %v", par.GetErrors())
	}

	if len(root.Statements) != 1 {
		t.Fatalf("Expected 1 statement, got %d", len(root.Statements))
	}

	enumNode, ok := root.Statements[0].(*EnumDeclarationNode)
	if !ok {
		t.Fatalf("Expected EnumDeclarationNode, got %T", root.Statements[0])
	}

	if enumNode.EnumName.Name != "Color" {
		t.Errorf("Expected enum name 'Color', got '%s'", enumNode.EnumName.Name)
	}

	if len(enumNode.Members) != 3 {
		t.Errorf("Expected 3 members, got %d", len(enumNode.Members))
	}

	expectedMembers := []string{"RED", "GREEN", "BLUE"}
	for i, member := range enumNode.Members {
		if member.Name != expectedMembers[i] {
			t.Errorf("Expected member '%s', got '%s'", expectedMembers[i], member.Name)
		}
		// Check auto-assigned values (0, 1, 2)
		intVal, ok := member.Value.(*std.Integer)
		if !ok {
			t.Errorf("Expected Integer value for member %s, got %T", member.Name, member.Value)
			continue
		}
		if intVal.Value != int64(i) {
			t.Errorf("Expected value %d for member %s, got %d", i, member.Name, intVal.Value)
		}
	}
}

// TestParseEnumWithValues tests parsing of enum with explicit values
func TestParseEnumWithValues(t *testing.T) {
	input := `enum Status { PENDING = 0, ACTIVE = 1, COMPLETED = 2 }`
	par := NewParser(input)

	root := par.Parse()

	if par.HasErrors() {
		t.Errorf("Parser has errors: %v", par.GetErrors())
	}

	enumNode, ok := root.Statements[0].(*EnumDeclarationNode)
	if !ok {
		t.Fatalf("Expected EnumDeclarationNode, got %T", root.Statements[0])
	}

	if enumNode.EnumName.Name != "Status" {
		t.Errorf("Expected enum name 'Status', got '%s'", enumNode.EnumName.Name)
	}

	if len(enumNode.Members) != 3 {
		t.Errorf("Expected 3 members, got %d", len(enumNode.Members))
	}

	expectedValues := map[string]int64{
		"PENDING":   0,
		"ACTIVE":    1,
		"COMPLETED": 2,
	}

	for _, member := range enumNode.Members {
		expectedVal, ok := expectedValues[member.Name]
		if !ok {
			t.Errorf("Unexpected member: %s", member.Name)
			continue
		}
		intVal, ok := member.Value.(*std.Integer)
		if !ok {
			t.Errorf("Expected Integer value for member %s, got %T", member.Name, member.Value)
			continue
		}
		if intVal.Value != expectedVal {
			t.Errorf("Expected value %d for member %s, got %d", expectedVal, member.Name, intVal.Value)
		}
	}
}

// TestParseEnumMixedValues tests parsing of enum with mixed auto and explicit values
func TestParseEnumMixedValues(t *testing.T) {
	input := `enum Priority { LOW = 10, MEDIUM, HIGH = 50, CRITICAL }`
	par := NewParser(input)

	root := par.Parse()

	if par.HasErrors() {
		t.Errorf("Parser has errors: %v", par.GetErrors())
	}

	enumNode, ok := root.Statements[0].(*EnumDeclarationNode)
	if !ok {
		t.Fatalf("Expected EnumDeclarationNode, got %T", root.Statements[0])
	}

	expectedValues := map[string]int64{
		"LOW":      10,
		"MEDIUM":   11, // Auto-assigned after 10
		"HIGH":     50,
		"CRITICAL": 51, // Auto-assigned after 50
	}

	for _, member := range enumNode.Members {
		expectedVal, ok := expectedValues[member.Name]
		if !ok {
			t.Errorf("Unexpected member: %s", member.Name)
			continue
		}
		intVal, ok := member.Value.(*std.Integer)
		if !ok {
			t.Errorf("Expected Integer value for member %s, got %T", member.Name, member.Value)
			continue
		}
		if intVal.Value != expectedVal {
			t.Errorf("Expected value %d for member %s, got %d", expectedVal, member.Name, intVal.Value)
		}
	}
}

// TestParseEmptyEnum tests parsing of empty enum
func TestParseEmptyEnum(t *testing.T) {
	input := `enum Empty { }`
	par := NewParser(input)

	root := par.Parse()

	if par.HasErrors() {
		t.Errorf("Parser has errors: %v", par.GetErrors())
	}

	enumNode, ok := root.Statements[0].(*EnumDeclarationNode)
	if !ok {
		t.Fatalf("Expected EnumDeclarationNode, got %T", root.Statements[0])
	}

	if len(enumNode.Members) != 0 {
		t.Errorf("Expected 0 members, got %d", len(enumNode.Members))
	}
}

// TestParseEnumAccessExpression tests parsing of enum member access
func TestParseEnumAccessExpression(t *testing.T) {
	input := `enum Color { RED, GREEN, BLUE }
var c = Color.RED`
	par := NewParser(input)

	root := par.Parse()

	if par.HasErrors() {
		t.Errorf("Parser has errors: %v", par.GetErrors())
	}

	if len(root.Statements) != 2 {
		t.Fatalf("Expected 2 statements, got %d", len(root.Statements))
	}

	// First statement should be enum declaration
	_, ok := root.Statements[0].(*EnumDeclarationNode)
	if !ok {
		t.Fatalf("Expected EnumDeclarationNode, got %T", root.Statements[0])
	}

	// Second statement should be variable declaration with enum access
	declNode, ok := root.Statements[1].(*DeclarativeStatementNode)
	if !ok {
		t.Fatalf("Expected DeclarativeStatementNode, got %T", root.Statements[1])
	}

	// The expression is parsed as a binary expression with DOT_OP
	// This is the expected behavior - the evaluator handles the enum access
	binExpr, ok := declNode.Expr.(*BinaryExpressionNode)
	if !ok {
		t.Fatalf("Expected BinaryExpressionNode for enum access, got %T", declNode.Expr)
	}

	// Check left side is the enum name (Color)
	leftIdent, ok := binExpr.Left.(*IdentifierExpressionNode)
	if !ok {
		t.Fatalf("Expected IdentifierExpressionNode on left, got %T", binExpr.Left)
	}
	if leftIdent.Name != "Color" {
		t.Errorf("Expected enum name 'Color', got '%s'", leftIdent.Name)
	}

	// Check right side is the member name (RED)
	rightIdent, ok := binExpr.Right.(*IdentifierExpressionNode)
	if !ok {
		t.Fatalf("Expected IdentifierExpressionNode on right, got %T", binExpr.Right)
	}
	if rightIdent.Name != "RED" {
		t.Errorf("Expected member name 'RED', got '%s'", rightIdent.Name)
	}
}
