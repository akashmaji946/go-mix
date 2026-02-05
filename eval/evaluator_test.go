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

// TestEvaluator_Ints verifies integer literal evaluation and arithmetic operations
func TestEvaluator_Ints(t *testing.T) {
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

// TestEvaluator_Floats verifies float literal evaluation and arithmetic operations
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

// TestEvaluator_Bools verifies boolean literal evaluation and comparison operations
func TestEvaluator_Bools(t *testing.T) {
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

// TestEvaluator_Nil verifies nil value evaluation
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

// TestEvaluator_Strings verifies string literal evaluation
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

// TestEvaluator_IntExpr verifies integer expressions including arithmetic, bitwise, and unary operations
func TestEvaluator_IntExpr(t *testing.T) {
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

// TestEvaluator_ExprErr verifies error handling for invalid expression operations
func TestEvaluator_ExprErr(t *testing.T) {
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

// TestEvaluator_Conds verifies if/else conditional statement evaluation
func TestEvaluator_Conds(t *testing.T) {
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

// TestEvaluator_Return verifies return statement evaluation and early exit behavior
func TestEvaluator_Return(t *testing.T) {
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

// TestEvaluator_Decls verifies variable declaration and initialization
func TestEvaluator_Decls(t *testing.T) {
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

// TestEvaluator_DeclErr verifies error handling for invalid variable declarations
func TestEvaluator_DeclErr(t *testing.T) {
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

// TestEvaluator_FuncCall verifies function call evaluation with parameters
func TestEvaluator_FuncCall(t *testing.T) {
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

// TestEvaluator_FuncArgErr verifies error handling for incorrect function argument counts
func TestEvaluator_FuncArgErr(t *testing.T) {
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

// TestEvaluator_Let verifies let keyword variable declaration evaluation
func TestEvaluator_Let(t *testing.T) {
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

// TestEvaluator_ConstDeclErr verifies error handling for const redeclaration
func TestEvaluator_ConstDeclErr(t *testing.T) {
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

// TestEvaluator_ConstReassignErr verifies error handling for const reassignment attempts
func TestEvaluator_ConstReassignErr(t *testing.T) {
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

// TestEvaluator_LetDeclErr verifies error handling for let redeclaration
func TestEvaluator_LetDeclErr(t *testing.T) {
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

// TestEvaluator_LetReassignErr verifies error handling for let type mismatch on reassignment
func TestEvaluator_LetReassignErr(t *testing.T) {
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

// TestEvaluator_CompoundArith verifies arithmetic compound assignment operators (+=, -=, *=, /=, %=)
func TestEvaluator_CompoundArith(t *testing.T) {
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

// TestEvaluator_CompoundBitwise verifies bitwise compound assignment operators (&=, |=, ^=, <<=, >>=)
func TestEvaluator_CompoundBitwise(t *testing.T) {
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

// TestEvaluator_CompoundLoops verifies compound assignments within for and while loops
func TestEvaluator_CompoundLoops(t *testing.T) {
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

// TestEvaluator_CompoundExpr verifies compound assignments with complex expressions
func TestEvaluator_CompoundExpr(t *testing.T) {
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

// TestEvaluator_RangeSimple verifies simple range expression evaluation
func TestEvaluator_RangeSimple(t *testing.T) {
	tests := []struct {
		input    string
		expStart int64
		expEnd   int64
	}{
		{"2...5", 2, 5},
		{"1...10", 1, 10},
		{"0...100", 0, 100},
		{"-5...5", -5, 5},
	}

	for _, tt := range tests {
		p := parser.NewParser(tt.input)
		rootNode := p.Parse()
		evaluator := NewEvaluator()
		evaluator.SetParser(p)
		result := evaluator.Eval(rootNode)

		if result.GetType() != objects.RangeType {
			t.Errorf("expected %s, got %s", objects.RangeType, result.GetType())
		}

		rangeObj := result.(*objects.Range)
		if rangeObj.Start != tt.expStart {
			t.Errorf("expected start %d, got %d", tt.expStart, rangeObj.Start)
		}
		if rangeObj.End != tt.expEnd {
			t.Errorf("expected end %d, got %d", tt.expEnd, rangeObj.End)
		}
	}
}

// TestEvaluator_RangeVar verifies range assignment to variables
func TestEvaluator_RangeVar(t *testing.T) {
	src := `var x = 2...9; x`
	p := parser.NewParser(src)
	rootNode := p.Parse()
	evaluator := NewEvaluator()
	evaluator.SetParser(p)
	result := evaluator.Eval(rootNode)

	if result.GetType() != objects.RangeType {
		t.Errorf("expected %s, got %s", objects.RangeType, result.GetType())
	}

	rangeObj := result.(*objects.Range)
	if rangeObj.Start != 2 || rangeObj.End != 9 {
		t.Errorf("expected range(2,9), got range(%d,%d)", rangeObj.Start, rangeObj.End)
	}
}

// TestEvaluator_RangeBuiltin verifies range() builtin function
func TestEvaluator_RangeBuiltin(t *testing.T) {
	tests := []struct {
		input    string
		expStart int64
		expEnd   int64
	}{
		{"range(2, 9)", 2, 9},
		{"range(1, 5)", 1, 5},
		{"range(10, 20)", 10, 20},
	}

	for _, tt := range tests {
		p := parser.NewParser(tt.input)
		rootNode := p.Parse()
		evaluator := NewEvaluator()
		evaluator.SetParser(p)
		result := evaluator.Eval(rootNode)

		if result.GetType() != objects.RangeType {
			t.Errorf("expected %s, got %s", objects.RangeType, result.GetType())
		}

		rangeObj := result.(*objects.Range)
		if rangeObj.Start != tt.expStart {
			t.Errorf("expected start %d, got %d", tt.expStart, rangeObj.Start)
		}
		if rangeObj.End != tt.expEnd {
			t.Errorf("expected end %d, got %d", tt.expEnd, rangeObj.End)
		}
	}
}

// TestEvaluator_ForeachRange verifies foreach loop with range
func TestEvaluator_ForeachRange(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{`var sum = 0; foreach i in 1...5 { sum += i; } sum`, 15}, // 1+2+3+4+5
		{`var sum = 0; foreach i in 1...3 { sum += i; } sum`, 6},  // 1+2+3
		{`var count = 0; foreach i in 1...10 { count += 1; } count`, 10},
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

// TestEvaluator_ForeachRangeVar verifies foreach loop with range variable
func TestEvaluator_ForeachRangeVar(t *testing.T) {
	src := `var r = 1...5; var sum = 0; foreach i in r { sum += i; } sum`
	p := parser.NewParser(src)
	rootNode := p.Parse()
	evaluator := NewEvaluator()
	evaluator.SetParser(p)
	result := evaluator.Eval(rootNode)
	AssertInteger(t, result, 15) // 1+2+3+4+5
}

// TestEvaluator_ForeachArray verifies foreach loop with arrays
func TestEvaluator_ForeachArray(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{`var sum = 0; foreach i in [1, 2, 3] { sum += i; } sum`, 6},
		{`var sum = 0; foreach i in [10, 20, 30] { sum += i; } sum`, 60},
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

// TestEvaluator_ForeachNested verifies nested foreach loops
func TestEvaluator_ForeachNested(t *testing.T) {
	src := `var sum = 0; foreach i in 1...2 { foreach j in 1...2 { sum += i * 10 + j; } } sum`
	p := parser.NewParser(src)
	rootNode := p.Parse()
	evaluator := NewEvaluator()
	evaluator.SetParser(p)
	result := evaluator.Eval(rootNode)
	AssertInteger(t, result, 66) // 11+12+21+22 = 66 (actually: 11+12+21+22 = 66)
}

// TestEvaluator_RangeBuiltinError verifies error handling for range() builtin
func TestEvaluator_RangeBuiltinError(t *testing.T) {
	errorTests := []struct {
		Src              string
		ExpectedErrorMsg string
	}{
		{
			`range(1)`,
			"ERROR: wrong number of arguments",
		},
		{
			`range(1, 2, 3)`,
			"ERROR: wrong number of arguments",
		},
		{
			`range("a", 5)`,
			"ERROR: first argument to `range` must be an integer",
		},
		{
			`range(1, "b")`,
			"ERROR: second argument to `range` must be an integer",
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

// TestEvaluator_ForeachError verifies error handling for foreach loops
func TestEvaluator_ForeachError(t *testing.T) {
	errorTests := []struct {
		Src              string
		ExpectedErrorMsg string
	}{
		{
			`foreach i in 123 { }`,
			"ERROR: foreach requires a `range` or `array`, got 'int'",
		},
		{
			`foreach i in "hello" { }`,
			"ERROR: foreach requires a `range` or `array`, got 'string'",
		},
		{
			`foreach i in true { }`,
			"ERROR: foreach requires a `range` or `array`, got 'bool'",
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

// TestEvaluator_MapKeys verifies keys_map() builtin function
func TestEvaluator_MapKeys(t *testing.T) {
	tests := []struct {
		input        string
		expectedLen  int
		expectedKeys []string
	}{
		{`var m = map{"a": 1, "b": 2, "c": 3}; keys_map(m)`, 3, []string{"a", "b", "c"}},
		{`var m = map{"x": 10}; keys_map(m)`, 1, []string{"x"}},
		{`var m = map{}; keys_map(m)`, 0, []string{}},
	}

	for _, tt := range tests {
		p := parser.NewParser(tt.input)
		rootNode := p.Parse()
		evaluator := NewEvaluator()
		evaluator.SetParser(p)
		result := evaluator.Eval(rootNode)

		if result.GetType() != objects.ArrayType {
			t.Errorf("expected array, got %s", result.GetType())
			continue
		}

		arr := result.(*objects.Array)
		if len(arr.Elements) != tt.expectedLen {
			t.Errorf("expected %d keys, got %d", tt.expectedLen, len(arr.Elements))
		}
	}
}

// TestEvaluator_MapInsert verifies insert_map() builtin function
func TestEvaluator_MapInsert(t *testing.T) {
	src := `var m = map{"a": 1}; insert_map(m, "b", 2); m`
	p := parser.NewParser(src)
	rootNode := p.Parse()
	evaluator := NewEvaluator()
	evaluator.SetParser(p)
	result := evaluator.Eval(rootNode)

	if result.GetType() != objects.MapType {
		t.Errorf("expected map, got %s", result.GetType())
	}

	mapObj := result.(*objects.Map)
	if len(mapObj.Keys) != 2 {
		t.Errorf("expected 2 keys, got %d", len(mapObj.Keys))
	}
}

// TestEvaluator_MapRemove verifies remove_map() builtin function
func TestEvaluator_MapRemove(t *testing.T) {
	src := `var m = map{"a": 1, "b": 2}; remove_map(m, "a"); m`
	p := parser.NewParser(src)
	rootNode := p.Parse()
	evaluator := NewEvaluator()
	evaluator.SetParser(p)
	result := evaluator.Eval(rootNode)

	if result.GetType() != objects.MapType {
		t.Errorf("expected map, got %s", result.GetType())
	}

	mapObj := result.(*objects.Map)
	if len(mapObj.Keys) != 1 {
		t.Errorf("expected 1 key, got %d", len(mapObj.Keys))
	}
}

// TestEvaluator_MapContain verifies contain_map() builtin function
func TestEvaluator_MapContain(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{`var m = map{"a": 1, "b": 2}; contain_map(m, "a")`, true},
		{`var m = map{"a": 1, "b": 2}; contain_map(m, "c")`, false},
		{`var m = map{}; contain_map(m, "x")`, false},
	}

	for _, tt := range tests {
		p := parser.NewParser(tt.input)
		rootNode := p.Parse()
		evaluator := NewEvaluator()
		evaluator.SetParser(p)
		result := evaluator.Eval(rootNode)
		AssertBoolean(t, result, tt.expected)
	}
}

// TestEvaluator_EnumerateMap verifies enumerate_map() builtin function
func TestEvaluator_EnumerateMap(t *testing.T) {
	tests := []struct {
		input       string
		expectedLen int
	}{
		{`var m = map{"a": 1, "b": 2, "c": 3}; enumerate_map(m)`, 3},
		{`var m = map{"x": 10}; enumerate_map(m)`, 1},
		{`var m = map{}; enumerate_map(m)`, 0},
	}

	for _, tt := range tests {
		p := parser.NewParser(tt.input)
		rootNode := p.Parse()
		evaluator := NewEvaluator()
		evaluator.SetParser(p)
		result := evaluator.Eval(rootNode)

		if result.GetType() != objects.ArrayType {
			t.Errorf("expected array, got %s", result.GetType())
			continue
		}

		arr := result.(*objects.Array)
		if len(arr.Elements) != tt.expectedLen {
			t.Errorf("expected %d pairs, got %d", tt.expectedLen, len(arr.Elements))
		}

		// Verify each element is an array of 2 elements
		for i, elem := range arr.Elements {
			if elem.GetType() != objects.ArrayType {
				t.Errorf("pair %d: expected array, got %s", i, elem.GetType())
				continue
			}
			pair := elem.(*objects.Array)
			if len(pair.Elements) != 2 {
				t.Errorf("pair %d: expected 2 elements, got %d", i, len(pair.Elements))
			}
		}
	}
}

// TestEvaluator_SetInsert verifies insert_set() builtin function
func TestEvaluator_SetInsert(t *testing.T) {
	src := `var s = set{1, 2}; insert_set(s, 3); s`
	p := parser.NewParser(src)
	rootNode := p.Parse()
	evaluator := NewEvaluator()
	evaluator.SetParser(p)
	result := evaluator.Eval(rootNode)

	if result.GetType() != objects.SetType {
		t.Errorf("expected set, got %s", result.GetType())
	}

	setObj := result.(*objects.Set)
	if len(setObj.Values) != 3 {
		t.Errorf("expected 3 values, got %d", len(setObj.Values))
	}
}

// TestEvaluator_SetRemove verifies remove_set() builtin function
func TestEvaluator_SetRemove(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{`var s = set{1, 2, 3}; remove_set(s, 2)`, true},
		{`var s = set{1, 2, 3}; remove_set(s, 5)`, false},
		{`var s = set{}; remove_set(s, 1)`, false},
	}

	for _, tt := range tests {
		p := parser.NewParser(tt.input)
		rootNode := p.Parse()
		evaluator := NewEvaluator()
		evaluator.SetParser(p)
		result := evaluator.Eval(rootNode)
		AssertBoolean(t, result, tt.expected)
	}
}

// TestEvaluator_SetContains verifies contains_set() builtin function
func TestEvaluator_SetContains(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{`var s = set{1, 2, 3}; contains_set(s, 2)`, true},
		{`var s = set{1, 2, 3}; contains_set(s, 5)`, false},
		{`var s = set{}; contains_set(s, 1)`, false},
	}

	for _, tt := range tests {
		p := parser.NewParser(tt.input)
		rootNode := p.Parse()
		evaluator := NewEvaluator()
		evaluator.SetParser(p)
		result := evaluator.Eval(rootNode)
		AssertBoolean(t, result, tt.expected)
	}
}

// TestEvaluator_SetValues verifies values_set() builtin function
func TestEvaluator_SetValues(t *testing.T) {
	tests := []struct {
		input       string
		expectedLen int
	}{
		{`var s = set{1, 2, 3, 4, 5}; values_set(s)`, 5},
		{`var s = set{42}; values_set(s)`, 1},
		{`var s = set{}; values_set(s)`, 0},
	}

	for _, tt := range tests {
		p := parser.NewParser(tt.input)
		rootNode := p.Parse()
		evaluator := NewEvaluator()
		evaluator.SetParser(p)
		result := evaluator.Eval(rootNode)

		if result.GetType() != objects.ArrayType {
			t.Errorf("expected array, got %s", result.GetType())
			continue
		}

		arr := result.(*objects.Array)
		if len(arr.Elements) != tt.expectedLen {
			t.Errorf("expected %d values, got %d", tt.expectedLen, len(arr.Elements))
		}
	}
}

// TestEvaluator_LengthMapSet verifies length() function works with maps and sets
func TestEvaluator_LengthMapSet(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		// Map length tests
		{`var m = map{"a": 1, "b": 2, "c": 3}; length(m)`, 3},
		{`var m = map{}; length(m)`, 0},
		{`var m = map{"x": 10}; length(m)`, 1},

		// Set length tests
		{`var s = set{1, 2, 3, 4, 5}; length(s)`, 5},
		{`var s = set{}; length(s)`, 0},
		{`var s = set{"a"}; length(s)`, 1},

		// Existing types still work
		{`length("hello")`, 5},
		{`length([1, 2, 3])`, 3},
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

// TestEvaluator_SizeFunction verifies size() function as an alias for length()
func TestEvaluator_SizeFunction(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		// size() is an alias for length()
		{`var m = map{"a": 1, "b": 2}; size(m)`, 2},
		{`var s = set{1, 2, 3}; size(s)`, 3},
		{`size("test")`, 4},
		{`size([1, 2, 3, 4])`, 4},
		{`var m = map{}; size(m)`, 0},
		{`var s = set{}; size(s)`, 0},
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

// TestEvaluator_MapSetErrors verifies error handling for map and set functions
func TestEvaluator_MapSetErrors(t *testing.T) {
	errorTests := []struct {
		Src              string
		ExpectedErrorMsg string
	}{
		// Map function errors
		{`keys_map(123)`, "ERROR: argument to `keys_map` must be a map"},
		{`insert_map(123, "key", "value")`, "ERROR: first argument to `insert_map` must be a map"},
		{`remove_map("not a map", "key")`, "ERROR: first argument to `remove_map` must be a map"},
		{`contain_map(true, "key")`, "ERROR: first argument to `contain_map` must be a map"},
		{`enumerate_map([1, 2, 3])`, "ERROR: argument to `enumerate_map` must be a map"},
		// Set function errors
		{`insert_set(123, 5)`, "ERROR: first argument to `insert_set` must be a set"},
		{`remove_set("not a set", 5)`, "ERROR: first argument to `remove_set` must be a set"},
		{`contains_set(true, 5)`, "ERROR: first argument to `contains_set` must be a set"},
		{`values_set([1, 2, 3])`, "ERROR: argument to `values_set` must be a set"},
		// Wrong number of arguments
		{`keys_map()`, "ERROR: wrong number of arguments"},
		{`insert_map(map{}, "key")`, "ERROR: wrong number of arguments"},
		{`remove_map(map{})`, "ERROR: wrong number of arguments"},
		{`contain_map(map{})`, "ERROR: wrong number of arguments"},
		{`enumerate_map()`, "ERROR: wrong number of arguments"},
		{`insert_set(set{})`, "ERROR: wrong number of arguments"},
		{`remove_set(set{})`, "ERROR: wrong number of arguments"},
		{`contains_set(set{})`, "ERROR: wrong number of arguments"},
		{`values_set()`, "ERROR: wrong number of arguments"},
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
