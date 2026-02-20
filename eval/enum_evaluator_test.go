/*
File    : go-mix/eval/enum_evaluator_test.go
Author  : Akash Maji
Contact : akashmaji(@iisc.ac.in)
*/
package eval

import (
	"testing"

	"github.com/akashmaji946/go-mix/parser"
	"github.com/akashmaji946/go-mix/std"
)

// TestEvalBasicEnum tests evaluation of basic enum declarations
func TestEvalBasicEnum(t *testing.T) {
	input := `enum Color { RED, GREEN, BLUE }`
	par := parser.NewParser(input)
	root := par.Parse()

	if par.HasErrors() {
		t.Fatalf("Parser has errors: %v", par.GetErrors())
	}

	ev := NewEvaluator()
	ev.SetParser(par)
	result := ev.Eval(root)

	// Enum declaration returns the enum object
	if _, ok := result.(*std.GoMixEnum); !ok {
		t.Errorf("Expected enum result for enum declaration, got %T", result)
	}

	// Check that the enum was registered in the evaluator's scope
	enumObj, exists := ev.Scp.LookUp("Color")
	if !exists {
		t.Errorf("Enum 'Color' was not registered in evaluator scope")
	}
	if _, ok := enumObj.(*std.GoMixEnum); !ok {
		t.Errorf("Expected GoMixEnum type, got %T", enumObj)
	}
}

// TestEvalEnumWithValues tests evaluation of enum with explicit values
func TestEvalEnumWithValues(t *testing.T) {
	input := `enum Status { PENDING = 0, ACTIVE = 1, COMPLETED = 2 }`
	par := parser.NewParser(input)
	root := par.Parse()

	if par.HasErrors() {
		t.Fatalf("Parser has errors: %v", par.GetErrors())
	}

	ev := NewEvaluator()
	ev.SetParser(par)
	result := ev.Eval(root)

	// Enum declaration returns the enum object
	if _, ok := result.(*std.GoMixEnum); !ok {
		t.Errorf("Expected enum result for enum declaration, got %T", result)
	}

	// Check that the enum was registered in the scope
	enumObj, exists := ev.Scp.LookUp("Status")
	if !exists {
		t.Errorf("Enum 'Status' was not registered in evaluator scope")
		return
	}

	// Check enum values
	enumType, ok := enumObj.(*std.GoMixEnum)
	if !ok {
		t.Errorf("Expected GoMixEnum type, got %T", enumObj)
		return
	}

	// Check member values
	pendingVal, exists := enumType.Members["PENDING"]
	if !exists {
		t.Errorf("PENDING member not found")
	} else if intVal, ok := pendingVal.(*std.Integer); !ok || intVal.Value != 0 {
		t.Errorf("Expected PENDING = 0, got %v", pendingVal)
	}

	activeVal, exists := enumType.Members["ACTIVE"]
	if !exists {
		t.Errorf("ACTIVE member not found")
	} else if intVal, ok := activeVal.(*std.Integer); !ok || intVal.Value != 1 {
		t.Errorf("Expected ACTIVE = 1, got %v", activeVal)
	}

	completedVal, exists := enumType.Members["COMPLETED"]
	if !exists {
		t.Errorf("COMPLETED member not found")
	} else if intVal, ok := completedVal.(*std.Integer); !ok || intVal.Value != 2 {
		t.Errorf("Expected COMPLETED = 2, got %v", completedVal)
	}
}

// TestEvalEnumAccessExpression tests evaluation of enum member access
func TestEvalEnumAccessExpression(t *testing.T) {
	input := `enum Color { RED, GREEN, BLUE }
Color.RED`
	par := parser.NewParser(input)
	root := par.Parse()

	if par.HasErrors() {
		t.Fatalf("Parser has errors: %v", par.GetErrors())
	}

	ev := NewEvaluator()
	ev.SetParser(par)
	result := ev.Eval(root)

	// The result should be the enum member value (0 for RED)
	intResult, ok := result.(*std.Integer)
	if !ok {
		t.Errorf("Expected Integer result for enum access, got %T", result)
		return
	}

	if intResult.Value != 0 {
		t.Errorf("Expected RED = 0, got %d", intResult.Value)
	}
}

// TestEvalEnumAccessMultipleMembers tests evaluation of multiple enum member accesses
func TestEvalEnumAccessMultipleMembers(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{`enum Color { RED, GREEN, BLUE } Color.RED`, 0},
		{`enum Color { RED, GREEN, BLUE } Color.GREEN`, 1},
		{`enum Color { RED, GREEN, BLUE } Color.BLUE`, 2},
	}

	for _, test := range tests {
		par := parser.NewParser(test.input)
		root := par.Parse()

		if par.HasErrors() {
			t.Errorf("Parser has errors for input '%s': %v", test.input, par.GetErrors())
			continue
		}

		ev := NewEvaluator()
		ev.SetParser(par)
		result := ev.Eval(root)

		intResult, ok := result.(*std.Integer)
		if !ok {
			t.Errorf("Expected Integer result, got %T for input '%s'", result, test.input)
			continue
		}

		if intResult.Value != test.expected {
			t.Errorf("Expected %d, got %d for input '%s'", test.expected, intResult.Value, test.input)
		}
	}
}

// TestEvalEnumWithExplicitValues tests evaluation of enum with explicit values
func TestEvalEnumWithExplicitValues(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{`enum Status { PENDING = 10, ACTIVE = 20, COMPLETED = 30 } Status.PENDING`, 10},
		{`enum Status { PENDING = 10, ACTIVE = 20, COMPLETED = 30 } Status.ACTIVE`, 20},
		{`enum Status { PENDING = 10, ACTIVE = 20, COMPLETED = 30 } Status.COMPLETED`, 30},
	}

	for _, test := range tests {
		par := parser.NewParser(test.input)
		root := par.Parse()

		if par.HasErrors() {
			t.Errorf("Parser has errors for input '%s': %v", test.input, par.GetErrors())
			continue
		}

		ev := NewEvaluator()
		ev.SetParser(par)
		result := ev.Eval(root)

		intResult, ok := result.(*std.Integer)
		if !ok {
			t.Errorf("Expected Integer result, got %T for input '%s'", result, test.input)
			continue
		}

		if intResult.Value != test.expected {
			t.Errorf("Expected %d, got %d for input '%s'", test.expected, intResult.Value, test.input)
		}
	}
}

// TestEvalEnumMixedValues tests evaluation of enum with mixed auto and explicit values
func TestEvalEnumMixedValues(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{`enum Priority { LOW = 10, MEDIUM, HIGH = 50, CRITICAL } Priority.LOW`, 10},
		{`enum Priority { LOW = 10, MEDIUM, HIGH = 50, CRITICAL } Priority.MEDIUM`, 11},
		{`enum Priority { LOW = 10, MEDIUM, HIGH = 50, CRITICAL } Priority.HIGH`, 50},
		{`enum Priority { LOW = 10, MEDIUM, HIGH = 50, CRITICAL } Priority.CRITICAL`, 51},
	}

	for _, test := range tests {
		par := parser.NewParser(test.input)
		root := par.Parse()

		if par.HasErrors() {
			t.Errorf("Parser has errors for input '%s': %v", test.input, par.GetErrors())
			continue
		}

		ev := NewEvaluator()
		ev.SetParser(par)
		result := ev.Eval(root)

		intResult, ok := result.(*std.Integer)
		if !ok {
			t.Errorf("Expected Integer result, got %T for input '%s'", result, test.input)
			continue
		}

		if intResult.Value != test.expected {
			t.Errorf("Expected %d, got %d for input '%s'", test.expected, intResult.Value, test.input)
		}
	}
}

// TestEvalEnumInVariableDeclaration tests using enum values in variable declarations
func TestEvalEnumInVariableDeclaration(t *testing.T) {
	input := `enum Color { RED, GREEN, BLUE }
var c = Color.GREEN
c`
	par := parser.NewParser(input)
	root := par.Parse()

	if par.HasErrors() {
		t.Fatalf("Parser has errors: %v", par.GetErrors())
	}

	ev := NewEvaluator()
	ev.SetParser(par)
	result := ev.Eval(root)

	intResult, ok := result.(*std.Integer)
	if !ok {
		t.Fatalf("Expected Integer result, got %T", result)
	}

	if intResult.Value != 1 {
		t.Errorf("Expected c = 1 (GREEN), got %d", intResult.Value)
	}
}

// TestEvalEnumInComparison tests using enum values in comparisons
func TestEvalEnumInComparison(t *testing.T) {
	input := `enum Status { PENDING = 0, ACTIVE = 1 }
Status.ACTIVE == 1`
	par := parser.NewParser(input)
	root := par.Parse()

	if par.HasErrors() {
		t.Fatalf("Parser has errors: %v", par.GetErrors())
	}

	ev := NewEvaluator()
	ev.SetParser(par)
	result := ev.Eval(root)

	boolResult, ok := result.(*std.Boolean)
	if !ok {
		t.Fatalf("Expected Boolean result, got %T", result)
	}

	if !boolResult.Value {
		t.Errorf("Expected Status.ACTIVE == 1 to be true")
	}
}

// TestEvalEnumInIfStatement tests using enum values in if statements
func TestEvalEnumInIfStatement(t *testing.T) {
	input := `enum Status { PENDING = 0, ACTIVE = 1 }
if (Status.ACTIVE == 1) { 100 } else { 200 }`
	par := parser.NewParser(input)
	root := par.Parse()

	if par.HasErrors() {
		t.Fatalf("Parser has errors: %v", par.GetErrors())
	}

	ev := NewEvaluator()
	ev.SetParser(par)
	result := ev.Eval(root)

	intResult, ok := result.(*std.Integer)
	if !ok {
		t.Fatalf("Expected Integer result, got %T", result)
	}

	if intResult.Value != 100 {
		t.Errorf("Expected 100 (if branch), got %d", intResult.Value)
	}
}

// TestEvalEnumInvalidMember tests accessing invalid enum member
func TestEvalEnumInvalidMember(t *testing.T) {
	input := `enum Color { RED, GREEN, BLUE }
Color.YELLOW`
	par := parser.NewParser(input)
	root := par.Parse()

	if par.HasErrors() {
		t.Fatalf("Parser has errors: %v", par.GetErrors())
	}

	ev := NewEvaluator()
	ev.SetParser(par)
	result := ev.Eval(root)

	// Should return an error
	if _, ok := result.(*std.Error); !ok {
		t.Errorf("Expected Error result for invalid enum member, got %T", result)
	}
}

// TestEvalEnumInvalidEnumName tests accessing non-existent enum
func TestEvalEnumInvalidEnumName(t *testing.T) {
	input := `NonExistent.MEMBER`
	par := parser.NewParser(input)
	root := par.Parse()

	if par.HasErrors() {
		t.Fatalf("Parser has errors: %v", par.GetErrors())
	}

	ev := NewEvaluator()
	ev.SetParser(par)
	result := ev.Eval(root)

	// Should return an error
	if _, ok := result.(*std.Error); !ok {
		t.Errorf("Expected Error result for invalid enum name, got %T", result)
	}
}

// TestEvalMultipleEnums tests declaring and using multiple enums
func TestEvalMultipleEnums(t *testing.T) {
	input := `enum Color { RED, GREEN }
enum Size { SMALL, MEDIUM, LARGE }
Color.GREEN + Size.MEDIUM`
	par := parser.NewParser(input)
	root := par.Parse()

	if par.HasErrors() {
		t.Fatalf("Parser has errors: %v", par.GetErrors())
	}

	ev := NewEvaluator()
	ev.SetParser(par)
	result := ev.Eval(root)

	intResult, ok := result.(*std.Integer)
	if !ok {
		t.Fatalf("Expected Integer result, got %T", result)
	}

	// Color.GREEN = 1, Size.MEDIUM = 1, so 1 + 1 = 2
	if intResult.Value != 2 {
		t.Errorf("Expected 2 (1 + 1), got %d", intResult.Value)
	}
}

// TestEvalEnumType tests that enum values have the correct type
func TestEvalEnumType(t *testing.T) {
	// First declare the enum and variable
	input := `enum Color { RED, GREEN }
var c = Color.RED`
	par := parser.NewParser(input)
	root := par.Parse()

	if par.HasErrors() {
		t.Fatalf("Parser has errors: %v", par.GetErrors())
	}

	ev := NewEvaluator()
	ev.SetParser(par)
	ev.Eval(root)

	// Look up the variable c
	val, ok := ev.Scp.LookUp("c")
	if !ok {
		t.Fatalf("Variable 'c' not found in scope")
	}

	// c should be an Integer (the enum value)
	if val.GetType() != std.IntegerType {
		t.Errorf("Expected enum value to be Integer, got %T", val)
	}
}
