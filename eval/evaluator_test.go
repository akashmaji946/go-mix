/*
File    : go-mix/eval/evaluator_test.go
Author  : Akash Maji
Contact : akashmaji(@iisc.ac.in)
*/
package eval

import (
	"fmt"
	"math"
	"strings"
	"testing"

	"github.com/akashmaji946/go-mix/objects"
	"github.com/akashmaji946/go-mix/parser"
)

func TestEvaluator_Integers(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"-1", -1},
		{"2", 2},
		{"-2", -2},
		{fmt.Sprintf("%d", math.MaxInt64), math.MaxInt64},
		{fmt.Sprintf("%d", math.MinInt64), math.MinInt64},
		{"1 + 1", 2},
		{"1 - 1", 0},
		{"2 * 15", 30},
		{"15 / 3", 5},
		{"1 + 2 * 3", 7},
		{"1 * -2", -2},
		{"1 + 2 * 3", 7},
		{"1 * -2", -2},
	}

	for _, tt := range tests {
		p := parser.NewParser(tt.input)
		rootNode := p.Parse()
		evaluator := NewEvaluator()
		evaluator.SetParser(p)
		result := evaluator.Eval(rootNode)
		if result.GetType() != objects.IntegerType {
			t.Errorf("expected %s, got %s", objects.IntegerType, result.GetType())
		}
		if result.(*objects.Integer).Value != tt.expected {
			t.Errorf("expected %d, got %d", tt.expected, result.(*objects.Integer).Value)
		}
	}
}

func TestEvaluator_Floats(t *testing.T) {
	tests := []struct {
		input    string
		expected float64
	}{
		{"-1.0", -1.0},
		{"2.0", 2.0},
		{"-2.0", -2.0},
		{fmt.Sprintf("%f", math.MaxFloat64), math.MaxFloat64},
		// {fmt.Sprintf("%f", math.MinFloat64), math.MinFloat64},
		{"1.0 + 2.0", 3.0},
		{"1.0 - 2.0", -1.0},
		{"2.0 * 15.0", 30.0},
		{"15.0 / 3.0", 5.0},
		{"1.0 + 2.0 * 3.0", 7.0},
		{"1.0 * -2.0", -2.0},
		{"1.0 + 2.0 * 3.0", 7.0},
		{"1.0 * -2.0", -2.0},
		{"2.2 / 2", 1.1},
	}

	for _, tt := range tests {
		p := parser.NewParser(tt.input)
		rootNode := p.Parse()
		evaluator := NewEvaluator()
		evaluator.SetParser(p)
		result := evaluator.Eval(rootNode)
		if result.GetType() != objects.FloatType {
			t.Errorf("expected %s, got %s", objects.FloatType, result.GetType())
		}
		if result.(*objects.Float).Value != tt.expected {
			t.Errorf("expected %f, got %f", tt.expected, result.(*objects.Float).Value)
		}
	}
}

func TestEvaluator_Booleans(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"true", true},
		{"false", false},
		{"true == true", true},
		{"true != false", true},
		{"true == false", false},
		{"true != false", true},
		{"1 == 1", true},
		{"1 != 1", false},
		{"1 == 2", false},
		{"1 != 2", true},
	}

	for _, tt := range tests {
		p := parser.NewParser(tt.input)
		rootNode := p.Parse()
		evaluator := NewEvaluator()
		evaluator.SetParser(p)
		result := evaluator.Eval(rootNode)
		if result.GetType() != objects.BooleanType {
			t.Errorf("expected %s, got %s", objects.BooleanType, result.GetType())
		}
		if result.(*objects.Boolean).Value != tt.expected {
			t.Errorf("expected %t, got %t", tt.expected, result.(*objects.Boolean).Value)
		}
	}
}

func TestEvaluator_Nil(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{"nil", nil},
		{``, nil},
	}

	for _, tt := range tests {
		p := parser.NewParser(tt.input)
		rootNode := p.Parse()
		evaluator := NewEvaluator()
		evaluator.SetParser(p)
		result := evaluator.Eval(rootNode)
		if result.GetType() != objects.NilType {
			t.Errorf("expected %s, got %s", objects.NilType, result.GetType())
		}

		if val, ok := result.(*objects.Nil); ok {
			if val.Value != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, val.Value)
			}
		} else {
			t.Errorf("expected result to be of type *Nil")
		}
	}
}

func TestEvaluator_Strings(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"\"hello\"", "hello"},
		{"\"hello world\"", "hello world"},
		{"\"\"", ""},
	}

	for _, tt := range tests {
		p := parser.NewParser(tt.input)
		rootNode := p.Parse()
		evaluator := NewEvaluator()
		evaluator.SetParser(p)
		result := evaluator.Eval(rootNode)
		if result.GetType() != objects.StringType {
			t.Errorf("expected %s, got %s", objects.StringType, result.GetType())
		}
		if result.(*objects.String).Value != tt.expected {
			t.Errorf("expected %s, got %s", tt.expected, result.(*objects.String).Value)
		}
	}
}

// evaluate expression
func TestEvaluator_ExpressionInteger(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		// arithemetic operations
		{"(1) + 2", 3},
		{"1 - 2", -1},
		{"1 * 2", 2},
		{"1 / 2", 0},
		{"1 % 2", 1},
		// bitwise operations
		{"1 & 2", 0},
		{"1 | 2", 3},
		{"1 ^ 2", 3},
		{"~1", -2},
		{"1 << 2", 4},
		{"1 >> 2", 0},
		// parenthesized expression
		{"(1 + 2) * 3", 9},
		{"(1 - 2) * 3", -3},
		{"(1 * 2) * 3", 6},
		{"(1 / 2) * 3", 0},
		{"(1 % 2) * 3", 3},
		{"(1 & 2) * 3", 0},
		{"(1 | 2) * 3", 9},
		{"(1 ^ 2) * 3", 9},
		{"(~1) * 3", -6},
		{"(1 << 2) * 3", 12},
		{"(((1 >> 2) * (3)))", 0},
		// unary operations
		{"-1", -1},
		{"+1", 1},
	}

	for _, tt := range tests {
		p := parser.NewParser(tt.input)
		rootNode := p.Parse()
		evaluator := NewEvaluator()
		evaluator.SetParser(p)
		result := evaluator.Eval(rootNode)
		if result.GetType() != objects.IntegerType {
			t.Errorf("expected %s, got %s", objects.IntegerType, result.GetType())
		}
		if result.(*objects.Integer).Value != tt.expected {
			t.Errorf("expected %d, got %d", tt.expected, result.(*objects.Integer).Value)
		}
	}
}

func TestEvaluator_ExpressionErrror(t *testing.T) {
	tests := []struct {
		Src              string
		ExpectedErrorMsg string
	}{
		{
			"1 + true",
			"ERROR: operator (+) not implemented for (int) and (bool)",
		},
		{
			"1 - true",
			"ERROR: operator (-) not implemented for (int) and (bool)",
		},
		{
			"1 * true",
			"ERROR: operator (*) not implemented for (int) and (bool)",
		},
		{
			"1 / true",
			"ERROR: operator (/) not implemented for (int) and (bool)",
		},
		{
			"1 % true",
			"ERROR: operator (%) not implemented for (int) and (bool)",
		},
		{
			"1 & true",
			"ERROR: operator (&) not implemented for (int) and (bool)",
		},
		{
			"1 | true",
			"ERROR: operator (|) not implemented for (int) and (bool)",
		},
		{
			"1 ^ true",
			"ERROR: operator (^) not implemented for (int) and (bool)",
		},
		{
			"~true",
			"ERROR: operator (~) not implemented for (bool)",
		},
		{
			"1 << true",
			"ERROR: operator (<<) not implemented for (int) and (bool)",
		},
		{
			"1 >> true",
			"ERROR: operator (>>) not implemented for (int) and (bool)",
		},
		{
			"true + true",
			"ERROR: operator (+) not implemented for (bool) and (bool)",
		},
		{
			"true - true",
			"ERROR: operator (-) not implemented for (bool) and (bool)",
		},
		{
			"true * true",
			"ERROR: operator (*) not implemented for (bool) and (bool)",
		},
		{
			"true / true",
			"ERROR: operator (/) not implemented for (bool) and (bool)",
		},
		{
			"true % true",
			"ERROR: operator (%) not implemented for (bool) and (bool)",
		},
		{
			"true & true",
			"ERROR: operator (&) not implemented for (bool) and (bool)",
		},
		{
			"true | true",
			"ERROR: operator (|) not implemented for (bool) and (bool)",
		},
		{
			"true ^ true",
			"ERROR: operator (^) not implemented for (bool) and (bool)",
		},
		{
			"~true",
			"ERROR: operator (~) not implemented for (bool)",
		},
		{
			"true << true",
			"ERROR: operator (<<) not implemented for (bool) and (bool)",
		},
		{
			"true >> true",
			"ERROR: operator (>>) not implemented for (bool) and (bool)",
		},

		{
			`
				if (true) {
					!1
					false
				}
				`,
			"ERROR: operator (!) not implemented for (int)",
		},
	}

	for _, tt := range tests {
		p := parser.NewParser(tt.Src)
		rootNode := p.Parse()
		evaluator := NewEvaluator()
		evaluator.SetParser(p)
		result := evaluator.Eval(rootNode)
		if result.GetType() != objects.ErrorType {
			t.Errorf("expected %s, got %s", objects.ErrorType, result.GetType())
		}
		if !strings.Contains(result.(*objects.Error).Message, tt.ExpectedErrorMsg) {
			t.Errorf("expected to contain %s, got %s", tt.ExpectedErrorMsg, result.(*objects.Error).Message)
		}

	}

}

func TestEvaluator_Eval_Conditionals(t *testing.T) {
	tests := []struct {
		Src      string
		Expected bool
	}{
		{
			`if(true) { true }`,
			true,
		},
		{
			`if(false) { true } else { false }`,
			false,
		},
		{
			`if(1 < 2) { true }`,
			true,
		},
		{
			`if(1 == 1) { true }`,
			true,
		},
	}
	for _, tt := range tests {
		p := parser.NewParser(tt.Src)
		rootNode := p.Parse()
		evaluator := NewEvaluator()
		evaluator.SetParser(p)
		result := evaluator.Eval(rootNode)
		AssertBoolean(t, result, tt.Expected)

	}

	errorTests := []struct {
		Src              string
		ExpectedErrorMsg string
	}{
		{
			`if(1) { true }`,
			"ERROR: conditional expression must be (bool)",
		},
		{
			`if(2 + 2 * 3) { true }`,
			"ERROR: conditional expression must be (bool)",
		},
	}
	for _, tt := range errorTests {
		p := parser.NewParser(tt.Src)
		rootNode := p.Parse()
		evaluator := NewEvaluator()
		evaluator.SetParser(p)
		result := evaluator.Eval(rootNode)
		if result.GetType() != objects.ErrorType {
			t.Errorf("expected %s, got %s", objects.ErrorType, result.GetType())
		}
		if !strings.Contains(result.(*objects.Error).Message, tt.ExpectedErrorMsg) {
			t.Errorf("expected to contain %s, got %s", tt.ExpectedErrorMsg, result.(*objects.Error).Message)
		}

	}
}

func TestEvaluator_Eval_ReturnStatement(t *testing.T) {
	tests := []struct {
		Src      string
		Expected int64
	}{
		{
			"return 10",
			10,
		},
		{
			`9 * 9
			return 10`,
			10,
		},
		{
			`9 * 9
				return 10
				8 + 10`,
			10,
		},
		{
			`if(true) {
					if (true) {
						return 10
					}
					return 1
				}`, 10,
		},
	}

	for _, tt := range tests {
		p := parser.NewParser(tt.Src)
		rootNode := p.Parse()
		evaluator := NewEvaluator()
		evaluator.SetParser(p)
		result := evaluator.Eval(rootNode)
		AssertInteger(t, result, tt.Expected)

	}
}

func TestEvaluator_Eval_Declarations(t *testing.T) {
	tests := []struct {
		Src      string
		Expected int64
	}{
		{
			`var a = 1
				a
				`,
			1,
		},
		{
			`var a = 1 * 2 + 1
				a
				`,
			3,
		},

		{
			`var a = 1 * 2 + 1
                 var c = 10
				 var d = a * c
				 d
				`,
			30,
		},
		{
			`var a = 1 * 2 + 1;
             var c = 10;
			 var d = a * c;
			`,
			30,
		},
	}

	for _, tt := range tests {
		p := parser.NewParser(tt.Src)
		rootNode := p.Parse()
		evaluator := NewEvaluator()
		evaluator.SetParser(p)
		result := evaluator.Eval(rootNode)
		AssertInteger(t, result, tt.Expected)

	}
}

func TestEvaluator_Eval_DeclarationError(t *testing.T) {
	errorTests := []struct {
		Src              string
		ExpectedErrorMsg string
	}{
		{
			`var a = b * 10;`,
			"ERROR: identifier not found: (b)",
		},
		{
			`var a = 1; var b = 2; var c = a + b + c;`,
			"ERROR: identifier not found: (c)",
		},
		{
			`var a = 1; var a = 2; var c = a;`,
			"ERROR: identifier redeclaration found: (a)",
		},
	}

	for _, tt := range errorTests {
		p := parser.NewParser(tt.Src)
		rootNode := p.Parse()
		evaluator := NewEvaluator()
		evaluator.SetParser(p)
		result := evaluator.Eval(rootNode)
		if result.GetType() != objects.ErrorType {
			t.Errorf("expected %s, got %s", objects.ErrorType, result.GetType())
		}
		if !strings.Contains(result.(*objects.Error).Message, tt.ExpectedErrorMsg) {
			t.Errorf("expected to contain %s, got %s", tt.ExpectedErrorMsg, result.(*objects.Error).Message)
		}

	}

}

// function declaration
func TestEvaluator_Eval_FunctionCall(t *testing.T) {
	tests := []struct {
		Src      string
		Expected int64
	}{
		{
			`var g = func(a) { return a + 1; }; g(1)`,
			2,
		},
		{
			`var add = func(a, b) { return a + b; }; add(2, 3)`,
			5,
		},
		{
			`var noArgs = func() { return 42; }; noArgs()`,
			42,
		},
	}

	for _, tt := range tests {
		p := parser.NewParser(tt.Src)
		rootNode := p.Parse()
		evaluator := NewEvaluator()
		evaluator.SetParser(p)
		result := evaluator.Eval(rootNode)
		AssertInteger(t, result, tt.Expected)
	}
}

func TestEvaluator_Eval_FunctionCallArgumentCountError(t *testing.T) {
	errorTests := []struct {
		Src              string
		ExpectedErrorMsg string
	}{
		{
			`var g = func(a) { return a + 1; }; g()`,
			"ERROR: wrong number of arguments: expected 1, got 0",
		},
		{
			`var g = func(a) { return a + 1; }; g(1, 2)`,
			"ERROR: wrong number of arguments: expected 1, got 2",
		},
		{
			`var g = func(a) { return a + 1; }; g(1, 2, 3, 4)`,
			"ERROR: wrong number of arguments: expected 1, got 4",
		},
		{
			`var add = func(a, b) { return a + b; }; add(1)`,
			"ERROR: wrong number of arguments: expected 2, got 1",
		},
		{
			`var add = func(a, b) { return a + b; }; add(1, 2, 3)`,
			"ERROR: wrong number of arguments: expected 2, got 3",
		},
		{
			`var noArgs = func() { return 42; }; noArgs(1)`,
			"ERROR: wrong number of arguments: expected 0, got 1",
		},
	}

	for _, tt := range errorTests {
		p := parser.NewParser(tt.Src)
		rootNode := p.Parse()
		evaluator := NewEvaluator()
		evaluator.SetParser(p)
		result := evaluator.Eval(rootNode)
		AssertError(t, result, tt.ExpectedErrorMsg)
	}
}

func TestParser_LetKeyword(t *testing.T) {
	tests := []struct {
		Src      string
		Expected int64
	}{
		{
			`let a = 1; a`,
			1,
		},
		{
			`let a = 1 * 2 + 1;`,
			3,
		},
		{
			`let a = 1 * 2 + 1;
				let c = 10;
				let d = a * c;`,
			30,
		},
	}
	for _, tt := range tests {
		p := parser.NewParser(tt.Src)
		rootNode := p.Parse()
		evaluator := NewEvaluator()
		evaluator.SetParser(p)
		result := evaluator.Eval(rootNode)
		AssertInteger(t, result, tt.Expected)
	}

}

func TestEvaluator_Eval_ConstDeclarationError(t *testing.T) {
	errorTests := []struct {
		Src              string
		ExpectedErrorMsg string
	}{
		{
			`const a = 1; const a = 2;`,
			"ERROR: identifier redeclaration found: (a)",
		},
		{
			`const a = 1; var a = 2;`,
			"ERROR: identifier redeclaration found: (a)",
		},
		{
			`var a = 1; const a = 2;`,
			"ERROR: identifier redeclaration found: (a)",
		},
	}

	for _, tt := range errorTests {
		p := parser.NewParser(tt.Src)
		rootNode := p.Parse()
		evaluator := NewEvaluator()
		evaluator.SetParser(p)
		result := evaluator.Eval(rootNode)
		if result.GetType() != objects.ErrorType {
			t.Errorf("expected %s, got %s", objects.ErrorType, result.GetType())
		}
		if !strings.Contains(result.(*objects.Error).Message, tt.ExpectedErrorMsg) {
			t.Errorf("expected to contain %s, got %s", tt.ExpectedErrorMsg, result.(*objects.Error).Message)
		}
	}
}

func TestEvaluator_Eval_ConstReassignmentError(t *testing.T) {
	errorTests := []struct {
		Src              string
		ExpectedErrorMsg string
	}{
		{
			`const a = 1; a = 2;`,
			"ERROR: can't assign to constant (a)",
		},
		{
			`const a = "hello"; a = "world";`,
			"ERROR: can't assign to constant (a)",
		},
		{
			`const a = 1.5; a = 2.5;`,
			"ERROR: can't assign to constant (a)",
		},
	}

	for _, tt := range errorTests {
		p := parser.NewParser(tt.Src)
		rootNode := p.Parse()
		evaluator := NewEvaluator()
		evaluator.SetParser(p)
		result := evaluator.Eval(rootNode)
		if result.GetType() != objects.ErrorType {
			t.Errorf("expected %s, got %s", objects.ErrorType, result.GetType())
		}
		if !strings.Contains(result.(*objects.Error).Message, tt.ExpectedErrorMsg) {
			t.Errorf("expected to contain %s, got %s", tt.ExpectedErrorMsg, result.(*objects.Error).Message)
		}
	}
}

func TestEvaluator_Eval_LetDeclarationError(t *testing.T) {
	errorTests := []struct {
		Src              string
		ExpectedErrorMsg string
	}{
		{
			`let a = 1; let a = 2;`,
			"ERROR: identifier redeclaration found: (a)",
		},
		{
			`let a = 1; var a = 2;`,
			"ERROR: identifier redeclaration found: (a)",
		},
		{
			`var a = 1; let a = 2;`,
			"ERROR: identifier redeclaration found: (a)",
		},
		{
			`const a = 1; let a = 2;`,
			"ERROR: identifier redeclaration found: (a)",
		},
		{
			`let a = 1; const a = 2;`,
			"ERROR: identifier redeclaration found: (a)",
		},
	}

	for _, tt := range errorTests {
		p := parser.NewParser(tt.Src)
		rootNode := p.Parse()
		evaluator := NewEvaluator()
		evaluator.SetParser(p)
		result := evaluator.Eval(rootNode)
		if result.GetType() != objects.ErrorType {
			t.Errorf("expected %s, got %s", objects.ErrorType, result.GetType())
		}
		if !strings.Contains(result.(*objects.Error).Message, tt.ExpectedErrorMsg) {
			t.Errorf("expected to contain %s, got %s", tt.ExpectedErrorMsg, result.(*objects.Error).Message)
		}
	}
}

func TestEvaluator_Eval_LetReassignmentError(t *testing.T) {
	errorTests := []struct {
		Src              string
		ExpectedErrorMsg string
	}{
		{
			`let a = 1; a = "hello";`,
			"ERROR: can't assign `string` to variable (a) of type `int`",
		},
		{
			`let a = 1.5; a = 2;`,
			"ERROR: can't assign `int` to variable (a) of type `float`",
		},
		{
			`let a = true; a = 1;`,
			"ERROR: can't assign `int` to variable (a) of type `bool`",
		},
		{
			`let a = nil; a = 1;`,
			"ERROR: can't assign `int` to variable (a) of type `nil`",
		},
	}

	for _, tt := range errorTests {
		p := parser.NewParser(tt.Src)
		rootNode := p.Parse()
		evaluator := NewEvaluator()
		evaluator.SetParser(p)
		result := evaluator.Eval(rootNode)
		if result.GetType() != objects.ErrorType {
			t.Errorf("expected %s, got %s", objects.ErrorType, result.GetType())
		}
		if !strings.Contains(result.(*objects.Error).Message, tt.ExpectedErrorMsg) {
			t.Errorf("expected to contain %s, got %s", tt.ExpectedErrorMsg, result.(*objects.Error).Message)
		}
	}
}

// Compound Assignment Tests
func TestEvaluator_CompoundAssignment_Arithmetic(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		// += operator
		{"var a = 10; a += 5; a", 15},
		{"var a = 100; a += 50; a += 25; a", 175},
		// -= operator
		{"var a = 20; a -= 5; a", 15},
		{"var a = 100; a -= 30; a -= 20; a", 50},
		// *= operator
		{"var a = 5; a *= 4; a", 20},
		{"var a = 2; a *= 3; a *= 4; a", 24},
		// /= operator
		{"var a = 20; a /= 4; a", 5},
		{"var a = 100; a /= 2; a /= 5; a", 10},
		// %= operator
		{"var a = 17; a %= 5; a", 2},
		{"var a = 100; a %= 30; a %= 7; a", 3},
	}

	for _, tt := range tests {
		p := parser.NewParser(tt.input)
		rootNode := p.Parse()
		evaluator := NewEvaluator()
		evaluator.SetParser(p)
		result := evaluator.Eval(rootNode)
		AssertInteger(t, result, tt.expected)
	}
}

func TestEvaluator_CompoundAssignment_Bitwise(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		// &= operator (bitwise AND)
		{"var a = 12; a &= 10; a", 8}, // 1100 & 1010 = 1000
		{"var a = 15; a &= 7; a", 7},  // 1111 & 0111 = 0111
		// |= operator (bitwise OR)
		{"var a = 12; a |= 3; a", 15}, // 1100 | 0011 = 1111
		{"var a = 8; a |= 4; a", 12},  // 1000 | 0100 = 1100
		// ^= operator (bitwise XOR)
		{"var a = 12; a ^= 5; a", 9},  // 1100 ^ 0101 = 1001
		{"var a = 15; a ^= 15; a", 0}, // 1111 ^ 1111 = 0000
		// <<= operator (left shift)
		{"var a = 4; a <<= 2; a", 16}, // 100 << 2 = 10000
		{"var a = 1; a <<= 4; a", 16}, // 1 << 4 = 10000
		// >>= operator (right shift)
		{"var a = 16; a >>= 2; a", 4}, // 10000 >> 2 = 100
		{"var a = 32; a >>= 3; a", 4}, // 100000 >> 3 = 100
	}

	for _, tt := range tests {
		p := parser.NewParser(tt.input)
		rootNode := p.Parse()
		evaluator := NewEvaluator()
		evaluator.SetParser(p)
		result := evaluator.Eval(rootNode)
		AssertInteger(t, result, tt.expected)
	}
}

func TestEvaluator_CompoundAssignment_InLoops(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		// += in for loop (i is declared in for loop initializer)
		{`var sum = 0; var i = 0; for(i = 1; i <= 5; i += 1){ sum += i; } sum`, 15},
		// *= in while loop
		{`var product = 1; var i = 1; while(i <= 4){ product *= i; i += 1; } product`, 24},
		// Combined compound assignments
		{`var a = 100; var i = 0; for(i = 0; i < 5; i += 1){ a -= 10; } a`, 50},
	}

	for _, tt := range tests {
		p := parser.NewParser(tt.input)
		rootNode := p.Parse()
		evaluator := NewEvaluator()
		evaluator.SetParser(p)
		result := evaluator.Eval(rootNode)
		AssertInteger(t, result, tt.expected)
	}
}

func TestEvaluator_CompoundAssignment_WithExpressions(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		// Compound assignment with complex expressions
		{"var a = 10; a += 2 * 3; a", 16},
		{"var a = 100; a -= 5 + 5; a", 90},
		{"var a = 5; a *= 2 + 3; a", 25},
		{"var a = 100; a /= 2 * 5; a", 10},
		// Chained compound assignments
		{"var a = 10; var b = 5; a += b; b += a; b", 20},
	}

	for _, tt := range tests {
		p := parser.NewParser(tt.input)
		rootNode := p.Parse()
		evaluator := NewEvaluator()
		evaluator.SetParser(p)
		result := evaluator.Eval(rootNode)
		AssertInteger(t, result, tt.expected)
	}
}
