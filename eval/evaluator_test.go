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

	"github.com/akashmaji946/go-mix/parser"
	"github.com/akashmaji946/go-mix/std"
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
		if result.GetType() != std.IntegerType {
			t.Errorf("expected %s, got %s", std.IntegerType, result.GetType())
		}
		if result.(*std.Integer).Value != tt.expected {
			t.Errorf("expected %d, got %d", tt.expected, result.(*std.Integer).Value)
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
		if result.GetType() != std.FloatType {
			t.Errorf("expected %s, got %s", std.FloatType, result.GetType())
		}
		if result.(*std.Float).Value != tt.expected {
			t.Errorf("expected %f, got %f", tt.expected, result.(*std.Float).Value)
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
		if result.GetType() != std.BooleanType {
			t.Errorf("expected %s, got %s", std.BooleanType, result.GetType())
		}
		if result.(*std.Boolean).Value != tt.expected {
			t.Errorf("expected %t, got %t", tt.expected, result.(*std.Boolean).Value)
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
		if result.GetType() != std.NilType {
			t.Errorf("expected %s, got %s", std.NilType, result.GetType())
		}

		if val, ok := result.(*std.Nil); ok {
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
		if result.GetType() != std.StringType {
			t.Errorf("expected %s, got %s", std.StringType, result.GetType())
		}
		if result.(*std.String).Value != tt.expected {
			t.Errorf("expected %s, got %s", tt.expected, result.(*std.String).Value)
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
		if result.GetType() != std.IntegerType {
			t.Errorf("expected %s, got %s", std.IntegerType, result.GetType())
		}
		if result.(*std.Integer).Value != tt.expected {
			t.Errorf("expected %d, got %d", tt.expected, result.(*std.Integer).Value)
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
		if result.GetType() != std.ErrorType {
			t.Errorf("expected %s, got %s", std.ErrorType, result.GetType())
		}
		if !strings.Contains(result.(*std.Error).Message, tt.ExpectedErrorMsg) {
			t.Errorf("expected to contain %s, got %s", tt.ExpectedErrorMsg, result.(*std.Error).Message)
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
		if result.GetType() != std.ErrorType {
			t.Errorf("expected %s, got %s", std.ErrorType, result.GetType())
		}
		if !strings.Contains(result.(*std.Error).Message, tt.ExpectedErrorMsg) {
			t.Errorf("expected to contain %s, got %s", tt.ExpectedErrorMsg, result.(*std.Error).Message)
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
		if result.GetType() != std.ErrorType {
			t.Errorf("expected %s, got %s", std.ErrorType, result.GetType())
		}
		if !strings.Contains(result.(*std.Error).Message, tt.ExpectedErrorMsg) {
			t.Errorf("expected to contain %s, got %s", tt.ExpectedErrorMsg, result.(*std.Error).Message)
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
		if result.GetType() != std.ErrorType {
			t.Errorf("expected %s, got %s", std.ErrorType, result.GetType())
		}
		if !strings.Contains(result.(*std.Error).Message, tt.ExpectedErrorMsg) {
			t.Errorf("expected to contain %s, got %s", tt.ExpectedErrorMsg, result.(*std.Error).Message)
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
		if result.GetType() != std.ErrorType {
			t.Errorf("expected %s, got %s", std.ErrorType, result.GetType())
		}
		if !strings.Contains(result.(*std.Error).Message, tt.ExpectedErrorMsg) {
			t.Errorf("expected to contain %s, got %s", tt.ExpectedErrorMsg, result.(*std.Error).Message)
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
		if result.GetType() != std.ErrorType {
			t.Errorf("expected %s, got %s", std.ErrorType, result.GetType())
		}
		if !strings.Contains(result.(*std.Error).Message, tt.ExpectedErrorMsg) {
			t.Errorf("expected to contain %s, got %s", tt.ExpectedErrorMsg, result.(*std.Error).Message)
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
		if result.GetType() != std.ErrorType {
			t.Errorf("expected %s, got %s", std.ErrorType, result.GetType())
		}
		if !strings.Contains(result.(*std.Error).Message, tt.ExpectedErrorMsg) {
			t.Errorf("expected to contain %s, got %s", tt.ExpectedErrorMsg, result.(*std.Error).Message)
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

		if result.GetType() != std.RangeType {
			t.Errorf("expected %s, got %s", std.RangeType, result.GetType())
		}

		rangeObj := result.(*std.Range)
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

	if result.GetType() != std.RangeType {
		t.Errorf("expected %s, got %s", std.RangeType, result.GetType())
	}

	rangeObj := result.(*std.Range)
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

		if result.GetType() != std.RangeType {
			t.Errorf("expected %s, got %s", std.RangeType, result.GetType())
		}

		rangeObj := result.(*std.Range)
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
			"ERROR: foreach requires an `iterable`, got `int`",
		},
		{
			`foreach i in "hello" { }`,
			"ERROR: foreach requires an `iterable`, got `string`",
		},
		{
			`foreach i in true { }`,
			"ERROR: foreach requires an `iterable`, got `bool`",
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

		if result.GetType() != std.ArrayType {
			t.Errorf("expected array, got %s", result.GetType())
			continue
		}

		arr := result.(*std.Array)
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

	if result.GetType() != std.MapType {
		t.Errorf("expected map, got %s", result.GetType())
	}

	mapObj := result.(*std.Map)
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

	if result.GetType() != std.MapType {
		t.Errorf("expected map, got %s", result.GetType())
	}

	mapObj := result.(*std.Map)
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

		if result.GetType() != std.ArrayType {
			t.Errorf("expected array, got %s", result.GetType())
			continue
		}

		arr := result.(*std.Array)
		if len(arr.Elements) != tt.expectedLen {
			t.Errorf("expected %d pairs, got %d", tt.expectedLen, len(arr.Elements))
		}

		// Verify each element is an array of 2 elements
		for i, elem := range arr.Elements {
			if elem.GetType() != std.ArrayType {
				t.Errorf("pair %d: expected array, got %s", i, elem.GetType())
				continue
			}
			pair := elem.(*std.Array)
			if len(pair.Elements) != 2 {
				t.Errorf("pair %d: expected 2 elements, got %d", i, len(pair.Elements))
			}
		}
	}
}

// TestEvaluator_SetInsert verifies insert_set() builtin function
func TestEvaluator_SetInsert(t *testing.T) {
	src := `import sets; var s = set{1, 2}; sets.insert_set(s, 3); s`
	p := parser.NewParser(src)
	rootNode := p.Parse()
	evaluator := NewEvaluator()
	evaluator.SetParser(p)
	result := evaluator.Eval(rootNode)

	if result.GetType() != std.SetType {
		t.Errorf("expected set, got %s", result.GetType())
	}

	setObj := result.(*std.Set)
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

		if result.GetType() != std.ArrayType {
			t.Errorf("expected array, got %s", result.GetType())
			continue
		}

		arr := result.(*std.Array)
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

// TestEvaluator_ListInsert verifies insert_list function with various indices
func TestEvaluator_ListInsert(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{`var l = list(1, 2, 4); insert_list(l, 2, 3); l`, "list(1, 2, 3, 4)"},
		{`var l = list(2, 3, 4); insert_list(l, 0, 1); l`, "list(1, 2, 3, 4)"},
		{`var l = list(1, 2, 3); insert_list(l, 3, 4); l`, "list(1, 2, 3, 4)"},
		{`var l = list(1, 2, 3); insert_list(l, -1, 4); l`, "list(1, 2, 3, 4)"},
		{`var l = list(); insert_list(l, 0, 1); l`, "list(1)"},
	}

	for _, tt := range tests {
		p := parser.NewParser(tt.input)
		rootNode := p.Parse()
		evaluator := NewEvaluator()
		evaluator.SetParser(p)
		result := evaluator.Eval(rootNode)
		if result.ToString() != tt.expected {
			t.Errorf("expected %s, got %s", tt.expected, result.ToString())
		}
	}
}

// TestEvaluator_ListRemove verifies remove_list function with various indices
func TestEvaluator_ListRemove(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{`var l = list(1, 2, 3, 4); remove_list(l, 2)`, 3},
		{`var l = list(1, 2, 3, 4); remove_list(l, 0)`, 1},
		{`var l = list(1, 2, 3, 4); remove_list(l, -1)`, 4},
		{`var l = list(1, 2, 3, 4); remove_list(l, -2)`, 3},
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

// TestEvaluator_ListContains verifies contains_list function
func TestEvaluator_ListContains(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{`var l = list(1, 2, 3, 4); contains_list(l, 3)`, true},
		{`var l = list(1, 2, 3, 4); contains_list(l, 5)`, false},
		{`var l = list("a", "b", "c"); contains_list(l, "b")`, true},
		{`var l = list("a", "b", "c"); contains_list(l, "d")`, false},
		{`var l = list(1, 2, 3); contains_list(l, "2")`, false},
		{`var l = list(); contains_list(l, 1)`, false},
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

// TestEvaluator_TupleContains verifies contains_tuple function
func TestEvaluator_TupleContains(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{`var t = tuple(1, 2, 3, 4); contains_tuple(t, 3)`, true},
		{`var t = tuple(1, 2, 3, 4); contains_tuple(t, 5)`, false},
		{`var t = tuple("a", "b", "c"); contains_tuple(t, "b")`, true},
		{`var t = tuple("a", "b", "c"); contains_tuple(t, "d")`, false},
		{`var t = tuple(1, 2, 3); contains_tuple(t, "2")`, false},
		{`var t = tuple(); contains_tuple(t, 1)`, false},
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

// TestEvaluator_ListInsertErrors verifies error handling for insert_list
func TestEvaluator_ListInsertErrors(t *testing.T) {
	errorTests := []struct {
		Src              string
		ExpectedErrorMsg string
	}{
		{`insert_list(123, 0, 1)`, "ERROR: first argument to `insert_list` must be a list"},
		{`insert_list(list(), "0", 1)`, "ERROR: second argument to `insert_list` must be an integer"},
		{`insert_list(list(), 10, 1)`, "ERROR: list index out of bounds"},
		{`insert_list(list(), -10, 1)`, "ERROR: list index out of bounds"},
		{`insert_list(list())`, "ERROR: wrong number of arguments"},
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

// TestEvaluator_ListRemoveErrors verifies error handling for remove_list
func TestEvaluator_ListRemoveErrors(t *testing.T) {
	errorTests := []struct {
		Src              string
		ExpectedErrorMsg string
	}{
		{`remove_list(123, 0)`, "ERROR: first argument to `remove_list` must be a list"},
		{`remove_list(list(), "0")`, "ERROR: second argument to `remove_list` must be an integer"},
		{`remove_list(list(), 0)`, "ERROR: cannot remove from empty list"},
		{`remove_list(list(1, 2), 10)`, "ERROR: list index out of bounds"},
		{`remove_list(list(1, 2), -10)`, "ERROR: list index out of bounds"},
		{`remove_list(list())`, "ERROR: wrong number of arguments"},
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

// TestEvaluator_ListTupleContainsErrors verifies error handling for contains functions
func TestEvaluator_ListTupleContainsErrors(t *testing.T) {
	errorTests := []struct {
		Src              string
		ExpectedErrorMsg string
	}{
		{`contains_list(123, 1)`, "ERROR: first argument to `contains_list` must be a list"},
		{`contains_list(list())`, "ERROR: wrong number of arguments"},
		{`contains_tuple(123, 1)`, "ERROR: first argument to `contains_tuple` must be a tuple"},
		{`contains_tuple(tuple())`, "ERROR: wrong number of arguments"},
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

// TestEvaluator_ListCreation verifies list() constructor function
func TestEvaluator_ListCreation(t *testing.T) {
	tests := []struct {
		input       string
		expectedLen int
	}{
		{`list()`, 0},
		{`list(1, 2, 3)`, 3},
		{`list(1, "hello", true, 3.14)`, 4},
		{`list(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)`, 10},
	}

	for _, tt := range tests {
		p := parser.NewParser(tt.input)
		rootNode := p.Parse()
		evaluator := NewEvaluator()
		evaluator.SetParser(p)
		result := evaluator.Eval(rootNode)

		if result.GetType() != std.ListType {
			t.Errorf("expected list, got %s", result.GetType())
			continue
		}

		listObj := result.(*std.List)
		if len(listObj.Elements) != tt.expectedLen {
			t.Errorf("expected %d elements, got %d", tt.expectedLen, len(listObj.Elements))
		}
	}
}

// TestEvaluator_TupleCreation verifies tuple() constructor function
func TestEvaluator_TupleCreation(t *testing.T) {
	tests := []struct {
		input       string
		expectedLen int
	}{
		{`tuple()`, 0},
		{`tuple(1, 2, 3)`, 3},
		{`tuple("Alice", 25, true, 5.8)`, 4},
		{`tuple(1, 2, 3, 4, 5)`, 5},
	}

	for _, tt := range tests {
		p := parser.NewParser(tt.input)
		rootNode := p.Parse()
		evaluator := NewEvaluator()
		evaluator.SetParser(p)
		result := evaluator.Eval(rootNode)

		if result.GetType() != std.TupleType {
			t.Errorf("expected tuple, got %s", result.GetType())
			continue
		}

		tupleObj := result.(*std.Tuple)
		if len(tupleObj.Elements) != tt.expectedLen {
			t.Errorf("expected %d elements, got %d", tt.expectedLen, len(tupleObj.Elements))
		}
	}
}

// TestEvaluator_ListIndexing verifies list indexing operations
func TestEvaluator_ListIndexing(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{`var l = list(10, 20, 30, 40, 50); l[0]`, 10},
		{`var l = list(10, 20, 30, 40, 50); l[2]`, 30},
		{`var l = list(10, 20, 30, 40, 50); l[4]`, 50},
		{`var l = list(10, 20, 30, 40, 50); l[-1]`, 50},
		{`var l = list(10, 20, 30, 40, 50); l[-2]`, 40},
		{`var l = list(10, 20, 30, 40, 50); l[-5]`, 10},
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

// TestEvaluator_TupleIndexing verifies tuple indexing operations
func TestEvaluator_TupleIndexing(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{`var t = tuple(10, 20, 30); t[0]`, 10},
		{`var t = tuple(10, 20, 30); t[1]`, 20},
		{`var t = tuple(10, 20, 30); t[2]`, 30},
		{`var t = tuple(10, 20, 30); t[-1]`, 30},
		{`var t = tuple(10, 20, 30); t[-2]`, 20},
		{`var t = tuple(10, 20, 30); t[-3]`, 10},
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

// TestEvaluator_ListSlicing verifies list slicing operations
func TestEvaluator_ListSlicing(t *testing.T) {
	tests := []struct {
		input       string
		expectedLen int
	}{
		{`var l = list(0, 10, 20, 30, 40, 50); l[1:4]`, 3},
		{`var l = list(0, 10, 20, 30, 40, 50); l[:3]`, 3},
		{`var l = list(0, 10, 20, 30, 40, 50); l[3:]`, 3},
		{`var l = list(0, 10, 20, 30, 40, 50); l[:]`, 6},
		{`var l = list(0, 10, 20, 30, 40, 50); l[1:-1]`, 4},
	}

	for _, tt := range tests {
		p := parser.NewParser(tt.input)
		rootNode := p.Parse()
		evaluator := NewEvaluator()
		evaluator.SetParser(p)
		result := evaluator.Eval(rootNode)

		if result.GetType() != std.ArrayType {
			t.Errorf("expected array from slice, got %s", result.GetType())
			continue
		}

		arr := result.(*std.Array)
		if len(arr.Elements) != tt.expectedLen {
			t.Errorf("expected %d elements, got %d", tt.expectedLen, len(arr.Elements))
		}
	}
}

// TestEvaluator_TupleSlicing verifies tuple slicing operations
func TestEvaluator_TupleSlicing(t *testing.T) {
	tests := []struct {
		input       string
		expectedLen int
	}{
		{`var t = tuple(1, 2, 3, 4, 5); t[1:3]`, 2},
		{`var t = tuple(1, 2, 3, 4, 5); t[:2]`, 2},
		{`var t = tuple(1, 2, 3, 4, 5); t[2:]`, 3},
		{`var t = tuple(1, 2, 3, 4, 5); t[:]`, 5},
		{`var t = tuple(1, 2, 3, 4, 5); t[1:-1]`, 3},
	}

	for _, tt := range tests {
		p := parser.NewParser(tt.input)
		rootNode := p.Parse()
		evaluator := NewEvaluator()
		evaluator.SetParser(p)
		result := evaluator.Eval(rootNode)

		if result.GetType() != std.ArrayType {
			t.Errorf("expected array from slice, got %s", result.GetType())
			continue
		}

		arr := result.(*std.Array)
		if len(arr.Elements) != tt.expectedLen {
			t.Errorf("expected %d elements, got %d", tt.expectedLen, len(arr.Elements))
		}
	}
}

// TestEvaluator_ListPushPop verifies list push and pop operations
func TestEvaluator_ListPushPop(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{`var l = list(1, 2, 3); pushback_list(l, 4); l[3]`, 4},
		{`var l = list(1, 2, 3); pushfront_list(l, 0); l[0]`, 0},
		{`var l = list(1, 2, 3); popback_list(l)`, 3},
		{`var l = list(1, 2, 3); popfront_list(l)`, 1},
		{`var l = list(1, 2, 3); pushback_list(l, 4); size_list(l)`, 4},
		{`var l = list(1, 2, 3); popback_list(l); size_list(l)`, 2},
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

// TestEvaluator_ListPeek verifies list peek operations
func TestEvaluator_ListPeek(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{`var l = list(10, 20, 30); peekfront_list(l)`, 10},
		{`var l = list(10, 20, 30); peekback_list(l)`, 30},
		{`var l = list(5); peekfront_list(l)`, 5},
		{`var l = list(5); peekback_list(l)`, 5},
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

// TestEvaluator_TuplePeek verifies tuple peek operations
func TestEvaluator_TuplePeek(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{`var t = tuple(10, 20, 30); peekfront_tuple(t)`, 10},
		{`var t = tuple(10, 20, 30); peekback_tuple(t)`, 30},
		{`var t = tuple(42); peekfront_tuple(t)`, 42},
		{`var t = tuple(42); peekback_tuple(t)`, 42},
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

// TestEvaluator_ListSize verifies list size operations
func TestEvaluator_ListSize(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{`size_list(list())`, 0},
		{`size_list(list(1, 2, 3))`, 3},
		{`size_list(list(1, 2, 3, 4, 5))`, 5},
		{`var l = list(1, 2); pushback_list(l, 3); size_list(l)`, 3},
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

// TestEvaluator_TupleSize verifies tuple size operations
func TestEvaluator_TupleSize(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{`size_tuple(tuple())`, 0},
		{`size_tuple(tuple(1, 2, 3))`, 3},
		{`size_tuple(tuple("a", "b", "c", "d"))`, 4},
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

// TestEvaluator_ListLength verifies length() function with lists
func TestEvaluator_ListLength(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{`length(list())`, 0},
		{`length(list(1, 2, 3))`, 3},
		{`length(list(1, 2, 3, 4, 5))`, 5},
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

// TestEvaluator_TupleLength verifies length() function with tuples
func TestEvaluator_TupleLength(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{`length(tuple())`, 0},
		{`length(tuple(1, 2, 3))`, 3},
		{`length(tuple("a", "b"))`, 2},
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

// TestEvaluator_ForeachList verifies foreach loops with lists
func TestEvaluator_ForeachList(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{`var sum = 0; foreach i in list(1, 2, 3) { sum += i; } sum`, 6},
		{`var sum = 0; foreach i in list(10, 20, 30) { sum += i; } sum`, 60},
		{`var count = 0; foreach i in list(1, 2, 3, 4, 5) { count += 1; } count`, 5},
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

// TestEvaluator_ForeachTuple verifies foreach loops with tuples
func TestEvaluator_ForeachTuple(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{`var sum = 0; foreach i in tuple(1, 2, 3) { sum += i; } sum`, 6},
		{`var sum = 0; foreach i in tuple(5, 10, 15) { sum += i; } sum`, 30},
		{`var count = 0; foreach i in tuple(1, 2, 3) { count += 1; } count`, 3},
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

// TestEvaluator_ListNested verifies nested list operations
func TestEvaluator_ListNested(t *testing.T) {
	src := `var matrix = list(list(1, 2), list(3, 4)); matrix[0][1]`
	p := parser.NewParser(src)
	rootNode := p.Parse()
	evaluator := NewEvaluator()
	evaluator.SetParser(p)
	result := evaluator.Eval(rootNode)
	AssertInteger(t, result, 2)
}

// TestEvaluator_TupleNested verifies nested tuple operations
func TestEvaluator_TupleNested(t *testing.T) {
	src := `var nested = tuple(tuple(1, 2), tuple(3, 4)); nested[1][0]`
	p := parser.NewParser(src)
	rootNode := p.Parse()
	evaluator := NewEvaluator()
	evaluator.SetParser(p)
	result := evaluator.Eval(rootNode)
	AssertInteger(t, result, 3)
}

// TestEvaluator_ListMixed verifies lists with mixed types
func TestEvaluator_ListMixed(t *testing.T) {
	src := `var l = list(1, "hello", true); length(l)`
	p := parser.NewParser(src)
	rootNode := p.Parse()
	evaluator := NewEvaluator()
	evaluator.SetParser(p)
	result := evaluator.Eval(rootNode)
	AssertInteger(t, result, 3)
}

// TestEvaluator_TupleMixed verifies tuples with mixed types
func TestEvaluator_TupleMixed2(t *testing.T) {
	src := `var t = tuple("Alice", 25, true); length(t)`
	p := parser.NewParser(src)
	rootNode := p.Parse()
	evaluator := NewEvaluator()
	evaluator.SetParser(p)
	result := evaluator.Eval(rootNode)
	AssertInteger(t, result, 3)
}

// TestEvaluator_Structs verifies struct creation and field access
func TestEvaluator_Structs(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{`struct Point { };`, ""},
		{`struct Person { func init(){} }`, ""},
		{`struct Rectangle { func init(x, y) {  } }`, ""},
		{`struct Circle { func init(radius) { var area = 1.0; } }`, ""},
		{`
		struct Point {

			func init(x, y) {

			}
			func move(dx, dy) {

			}
		}
		`, ""},
	}

	for _, tt := range tests {
		p := parser.NewParser(tt.input)
		rootNode := p.Parse()
		evaluator := NewEvaluator()
		evaluator.SetParser(p)
		result := evaluator.Eval(rootNode)
		if result.GetType() == std.ErrorType {
			t.Errorf("unexpected error: %s", result.ToString())
			continue
		}
		if result.GetType() != std.StructType {
			t.Errorf("expected struct, got %s", result.GetType())
		}

	}
}

// parseOrDie parses source code or panics if parsing fails
func parseOrDie(src string) *parser.RootNode {
	p := parser.NewParser(src)
	rootNode := p.Parse()
	if p.HasErrors() {
		panic("parse failed")
	}
	return rootNode
}

// assertFoundType asserts that a struct type was registered in the evaluator
func assertFoundType(t *testing.T, evaluator *Evaluator, typeName string) *std.GoMixStruct {
	tp, exists := evaluator.Types[typeName]
	if !exists {
		t.Fatalf("expected type '%s' to be registered", typeName)
	}
	return tp
}

// assertFoundInScope asserts that a variable exists in the current scope
func assertFoundInScope(t *testing.T, evaluator *Evaluator, varName string, expectedType std.GoMixType) std.GoMixObject {
	value, exists := evaluator.Scp.LookUp(varName)
	if !exists {
		t.Fatalf("expected variable '%s' to exist in scope", varName)
	}
	if value.GetType() != expectedType {
		t.Errorf("expected type '%s', got '%s'", expectedType, value.GetType())
	}
	return value
}

// TestEvaluator_Eval_EvaluateInstanceCreation verifies struct instance creation with and without constructors
func TestEvaluator_Eval_EvaluateInstanceCreation(t *testing.T) {
	tests := []struct {
		Src        string
		AssertFunc func(*Evaluator)
	}{
		{
			Src: `	
						struct A { 
						}
						var a = new A()
            `,
			AssertFunc: func(evaluator *Evaluator) {
				tp := assertFoundType(t, evaluator, "A")
				s := assertFoundInScope(t, evaluator, "a", std.ObjectType)
				p, ok := s.(*std.GoMixObjectInstance)
				if !ok {
					t.Fatalf("expected GoMixObjectInstance, got %T", s)
				}
				if p.Struct != tp {
					t.Errorf("instance struct mismatch")
				}
				if len(p.InstanceFields) != 0 {
					t.Errorf("expected 0 fields, got %d", len(p.InstanceFields))
				}
			},
		},
		{
			Src: `	
						struct A { 
						}
						var a = new A()
						var b = new A()
            `,
			AssertFunc: func(evaluator *Evaluator) {
				tp := assertFoundType(t, evaluator, "A")
				sa := assertFoundInScope(t, evaluator, "a", std.ObjectType)
				sb := assertFoundInScope(t, evaluator, "b", std.ObjectType)
				pa, oka := sa.(*std.GoMixObjectInstance)
				pb, okb := sb.(*std.GoMixObjectInstance)
				if !oka || !okb {
					t.Fatalf("expected GoMixObjectInstance, got %T and %T", sa, sb)
				}
				if pa.Struct != tp || pb.Struct != tp {
					t.Errorf("instance struct mismatch")
				}
				if len(pa.InstanceFields) != 0 || len(pb.InstanceFields) != 0 {
					t.Errorf("expected 0 fields, got %d and %d", len(pa.InstanceFields), len(pb.InstanceFields))
				}
			},
		},
		{
			Src: `	
						struct Point { 
							func init(x, y) {
								var px = x
								var py = y
							}
						}
						var p = new Point(10, 20)
            `,
			AssertFunc: func(evaluator *Evaluator) {
				tp := assertFoundType(t, evaluator, "Point")
				p := assertFoundInScope(t, evaluator, "p", std.ObjectType)
				inst, ok := p.(*std.GoMixObjectInstance)
				if !ok {
					t.Fatalf("expected GoMixObjectInstance, got %T", p)
				}
				if inst.Struct != tp {
					t.Errorf("instance struct mismatch")
				}
			},
		},
		{
			Src: `	
						struct Rectangle { 
							func init(width, height) {
								var w = width
								var h = height
							}
						}
						var r1 = new Rectangle(5, 10)
						var r2 = new Rectangle(15, 20)
            `,
			AssertFunc: func(evaluator *Evaluator) {
				tp := assertFoundType(t, evaluator, "Rectangle")
				r1 := assertFoundInScope(t, evaluator, "r1", std.ObjectType)
				r2 := assertFoundInScope(t, evaluator, "r2", std.ObjectType)
				ir1, ok1 := r1.(*std.GoMixObjectInstance)
				ir2, ok2 := r2.(*std.GoMixObjectInstance)
				if !ok1 || !ok2 {
					t.Fatalf("expected GoMixObjectInstance, got %T and %T", r1, r2)
				}
				if ir1.Struct != tp || ir2.Struct != tp {
					t.Errorf("instance struct mismatch")
				}
				// Instances should be different
				if ir1 == ir2 {
					t.Errorf("instances should be different")
				}
			},
		},
	}

	for i, test := range tests {
		evaluator := NewEvaluator()
		rootNode := parseOrDie(test.Src)
		result := evaluator.Eval(rootNode)
		if result.GetType() == std.ErrorType {
			t.Errorf("test %d: unexpected error: %s", i, result.ToString())
			continue
		}
		test.AssertFunc(evaluator)
	}
}

// TestEvaluator_StructMethodCall verifies calling methods on struct instances
func TestEvaluator_StructMethodCall(t *testing.T) {
	src := `
	struct Calc {
		func add(a, b) { return a + b; }
	}
	var c = new Calc();
	c.add(10, 20);
	`
	p := parser.NewParser(src)
	root := p.Parse()
	e := NewEvaluator()
	e.SetParser(p)
	result := e.Eval(root)
	AssertInteger(t, result, 30)
}

// TestEvaluator_StructThis verifies 'this' keyword usage in methods
func TestEvaluator_StructThis(t *testing.T) {
	src := `
	struct Counter {
		func init(start) { this.val = start; }
		func inc() { this.val += 1; }
		func get() { return this.val; }
	}
	var c = new Counter(10);
	c.inc();
	c.get();
	`
	p := parser.NewParser(src)
	root := p.Parse()
	e := NewEvaluator()
	e.SetParser(p)
	result := e.Eval(root)
	AssertInteger(t, result, 11)
}

// TestEvaluator_StructStaticFields verifies static field access and modification
func TestEvaluator_StructStaticFields(t *testing.T) {
	src := `
	struct Config {
		var count = 0;
	}
	Config.count = 10;
	var c1 = new Config();
	var c2 = new Config();
	// Access via class and instances (fallback)
	Config.count + c1.count + c2.count
	`
	p := parser.NewParser(src)
	root := p.Parse()
	e := NewEvaluator()
	e.SetParser(p)
	result := e.Eval(root)
	// 10 + 10 + 10 = 30
	AssertInteger(t, result, 30)
}

// TestEvaluator_StructConstField verifies error when assigning to const field
func TestEvaluator_StructConstField(t *testing.T) {
	src := `
	struct Math {
		const PI = 3.14;
	}
	Math.PI = 3.14159;
	`
	p := parser.NewParser(src)
	root := p.Parse()
	e := NewEvaluator()
	e.SetParser(p)
	result := e.Eval(root)
	AssertError(t, result, "ERROR: can't assign to constant field (PI) in struct (Math)")
}

// TestEvaluator_StructLetFieldType verifies type checking for let fields
func TestEvaluator_StructLetFieldType(t *testing.T) {
	src := `
	struct User {
		let name = "Anonymous";
	}
	User.name = 12345;
	`
	p := parser.NewParser(src)
	root := p.Parse()
	e := NewEvaluator()
	e.SetParser(p)
	result := e.Eval(root)
	AssertError(t, result, "ERROR: can't assign `int` to field (name) of type `string` in struct (User)")
}

// TestEvaluator_StructSelfVsThis verifies distinction between self (static) and this (instance)
func TestEvaluator_StructSelfVsThis(t *testing.T) {
	src := `
	struct Counter {
		var total = 0; // static
		func init(start) {
			this.local = start; // instance
		}
		func increment() {
			self.total += 1;
			this.local += 1;
		}
		func get() {
			return self.total + this.local;
		}
	}
	var c1 = new Counter(10);
	var c2 = new Counter(20);
	c1.increment(); // total=1, c1.local=11
	c2.increment(); // total=2, c2.local=21
	c1.get() + c2.get() // (2+11) + (2+21) = 13 + 23 = 36
	`
	p := parser.NewParser(src)
	root := p.Parse()
	e := NewEvaluator()
	e.SetParser(p)
	result := e.Eval(root)
	AssertInteger(t, result, 36)
}

// TestEvaluator_StructComplexLogic verifies complex logic within struct methods
func TestEvaluator_StructComplexLogic(t *testing.T) {
	src := `
	struct Processor {
		const LIMIT = 5;
		var processed = 0;
		
		func process(items) {
			foreach item in items {
				if (self.processed >= self.LIMIT) {
					break;
				}
				if (item < 0) {
					continue;
				}
				self.processed += 1;
			}
			return self.processed;
		}
	}
	var p = new Processor();
	// Should process 1, 2, 3, 4, 5. Skip -1. Stop at 6.
	// Items: 1, -1, 2, 3, 4, 5, 6
	p.process([1, -1, 2, 3, 4, 5, 6])
	`
	p := parser.NewParser(src)
	root := p.Parse()
	e := NewEvaluator()
	e.SetParser(p)
	result := e.Eval(root)
	AssertInteger(t, result, 5)
}

// TestEvaluator_StructFibonacci verifies recursion within a struct method
func TestEvaluator_StructFibonacci(t *testing.T) {
	src := `
    struct Math {
        func fib(n) {
            if (n <= 1) { return n; }
            return this.fib(n-1) + this.fib(n-2);
        }
    }
    var m = new Math();
    m.fib(10)
    `
	p := parser.NewParser(src)
	root := p.Parse()
	e := NewEvaluator()
	e.SetParser(p)
	result := e.Eval(root)
	AssertInteger(t, result, 55)
}

// TestEvaluator_StructLinkedList verifies linked list creation with structs
func TestEvaluator_StructLinkedList(t *testing.T) {
	src := `
    struct Node {
        var next = nil;
        func init(val) { this.val = val; }
    }
    var head = new Node(1);
    head.next = new Node(2);
    head.next.next = new Node(3);
    head.val + head.next.val + head.next.next.val
    `
	p := parser.NewParser(src)
	root := p.Parse()
	e := NewEvaluator()
	e.SetParser(p)
	result := e.Eval(root)
	AssertInteger(t, result, 6)
}

// TestEvaluator_StructStaticCounter verifies static fields for instance counting
func TestEvaluator_StructStaticCounter(t *testing.T) {
	src := `
    struct Item {
        var count = 0;
        func init() { Item.count += 1; }
    }
    new Item();
    new Item();
    new Item();
    Item.count
    `
	p := parser.NewParser(src)
	root := p.Parse()
	e := NewEvaluator()
	e.SetParser(p)
	result := e.Eval(root)
	AssertInteger(t, result, 3)
}

// TestEvaluator_StructArrayManagement verifies array manipulation within struct
func TestEvaluator_StructArrayManagement(t *testing.T) {
	src := `
    struct Stack {
        var data = list();
        func push(x) { pushback_list(this.data, x); }
        func pop() { return popback_list(this.data); }
    }
    var s = new Stack();
    s.push(10);
    s.push(20);
    s.pop() + s.pop()
    `
	p := parser.NewParser(src)
	root := p.Parse()
	e := NewEvaluator()
	e.SetParser(p)
	result := e.Eval(root)
	AssertInteger(t, result, 30)
}

// TestEvaluator_StructNestedAccess verifies accessing nested struct fields
func TestEvaluator_StructNestedAccess(t *testing.T) {
	src := `
    struct Inner { var val = 10; }
    struct Outer { 
        var inner = nil;
        func init() { this.inner = new Inner(); }
    }
    var o = new Outer();
    o.inner.val
    `
	p := parser.NewParser(src)
	root := p.Parse()
	e := NewEvaluator()
	e.SetParser(p)
	result := e.Eval(root)
	AssertInteger(t, result, 10)
}

// TestEvaluator_StructConstProtection verifies error when modifying const static field
func TestEvaluator_StructConstProtection(t *testing.T) {
	src := `struct S { const VAL = 10; } S.VAL = 20;`
	p := parser.NewParser(src)
	root := p.Parse()
	e := NewEvaluator()
	e.SetParser(p)
	result := e.Eval(root)
	AssertError(t, result, "ERROR: can't assign to constant field (VAL) in struct (S)")
}

// TestEvaluator_StructLetTypeProtection verifies error when modifying let static field with wrong type
func TestEvaluator_StructLetTypeProtection(t *testing.T) {
	src := `struct S { let VAL = 10; } S.VAL = "hello";`
	p := parser.NewParser(src)
	root := p.Parse()
	e := NewEvaluator()
	e.SetParser(p)
	result := e.Eval(root)
	AssertError(t, result, "ERROR: can't assign `string` to field (VAL) of type `int` in struct (S)")
}

// TestEvaluator_StructStaticPersistence verifies static fields persist across instances
func TestEvaluator_StructStaticPersistence(t *testing.T) {
	src := `struct S { var count = 0; func inc() { S.count += 1; } } var s1 = new S(); var s2 = new S(); s1.inc(); s2.inc(); S.count`
	p := parser.NewParser(src)
	root := p.Parse()
	e := NewEvaluator()
	e.SetParser(p)
	result := e.Eval(root)
	AssertInteger(t, result, 2)
}

// TestEvaluator_StructSelfAccess verifies self keyword accesses static fields
func TestEvaluator_StructSelfAccess(t *testing.T) {
	src := `struct S { var x = 10; func getX() { return self.x; } } var s = new S(); s.getX()`
	p := parser.NewParser(src)
	root := p.Parse()
	e := NewEvaluator()
	e.SetParser(p)
	result := e.Eval(root)
	AssertInteger(t, result, 10)
}

// TestEvaluator_StructThisInstance verifies this keyword accesses instance fields
func TestEvaluator_StructThisInstance(t *testing.T) {
	src := `struct S { func init(v) { this.v = v; } func get() { return this.v; } } var s = new S(5); s.get()`
	p := parser.NewParser(src)
	root := p.Parse()
	e := NewEvaluator()
	e.SetParser(p)
	result := e.Eval(root)
	AssertInteger(t, result, 5)
}

// TestEvaluator_StructShadowing verifies instance fields shadow static fields
func TestEvaluator_StructShadowing(t *testing.T) {
	src := `struct S { var x = 10; func init(v) { this.x = v; } func getStatic() { return self.x; } func getInstance() { return this.x; } } var s = new S(20); s.getStatic() + s.getInstance()`
	p := parser.NewParser(src)
	root := p.Parse()
	e := NewEvaluator()
	e.SetParser(p)
	result := e.Eval(root)
	AssertInteger(t, result, 30)
}

// TestEvaluator_StructMethodChaining verifies method chaining by returning this
func TestEvaluator_StructMethodChaining(t *testing.T) {
	src := `struct S { var val = 100; func add(v) { self.val += v; return this; } } var s = new S(); var d = s.add(5).add(10).val; d;`
	p := parser.NewParser(src)
	root := p.Parse()
	e := NewEvaluator()
	e.SetParser(p)
	result := e.Eval(root)
	AssertInteger(t, result, 115)
}

// TestEvaluator_StructArrayStatic verifies static array field shared across instances
func TestEvaluator_StructArrayStatic(t *testing.T) {
	src := `struct S { var list = []; func add(v) { self.list = push_array(self.list, v); } } var s1 = new S(); var s2 = new S(); s1.add(1); s2.add(2); length(S.list)`
	p := parser.NewParser(src)
	root := p.Parse()
	e := NewEvaluator()
	e.SetParser(p)
	result := e.Eval(root)
	AssertInteger(t, result, 2)
}

// TestEvaluator_StructComplexLoop verifies complex loop logic in struct method
func TestEvaluator_StructComplexLoop(t *testing.T) {
	src := `struct Math { func sum(n) { var s = 0; for(var i=1; i<=n; i+=1) { s += i; } return s; } } var m = new Math(); m.sum(5)`
	p := parser.NewParser(src)
	root := p.Parse()
	e := NewEvaluator()
	e.SetParser(p)
	result := e.Eval(root)
	AssertInteger(t, result, 15)
}

// TestEvaluator_StructFibonacciRecursive verifies recursive method call in struct
func TestEvaluator_StructFibonacciRecursive(t *testing.T) {
	src := `struct Fib { func calc(n) { if (n <= 1) { return n; } return this.calc(n-1) + this.calc(n-2); } } var f = new Fib(); f.calc(6)`
	p := parser.NewParser(src)
	root := p.Parse()
	e := NewEvaluator()
	e.SetParser(p)
	result := e.Eval(root)
	AssertInteger(t, result, 8)
}

// TestEvaluator_StringFunctions verifies evaluation of string builtin functions
func TestEvaluator_StringFunctions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{`upper("mix")`, "MIX"},
		{`reverse("abc")`, "cba"},
		{`contains("hello", "ell")`, true},
		{`ord('A')`, int64(65)},
		{`substring("hello", 1, 2)`, "el"},
	}

	for _, tt := range tests {
		p := parser.NewParser(tt.input)
		rootNode := p.Parse()
		evaluator := NewEvaluator()
		evaluator.SetParser(p)
		result := evaluator.Eval(rootNode)

		switch exp := tt.expected.(type) {
		case string:
			AssertString(t, result, exp)
		case bool:
			AssertBoolean(t, result, exp)
		case int64:
			AssertInteger(t, result, exp)
		}
	}
}

// TestEvaluator_AdditionalScenarios verifies various complex evaluation scenarios
func TestEvaluator_AdditionalScenarios(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		// 1. Mixed Arithmetic (Int + Float promotion)
		{`1 + 2.5`, 3.5},
		// 2. String Concatenation with non-string types
		{`"Value: " + 10`, "Value: 10"},
		// 3. Short-circuiting AND (Right side would panic if evaluated)
		{"false && (1 / 0 == 0)", false},
		// 4. Short-circuiting OR
		{"true || (1 / 0 == 0)", true},
		// 5. Functional Programming: Closures
		{`var adder = func(x) { return func(y) { return x + y; }; }; var add5 = adder(5); add5(10)`, int64(15)},
		// 6. Data Structures: Array of functions
		{`var arr = [func(x){return x*x;}, func(x){return x*x*x;}]; var f = arr[0]; var g = arr[1]; f(2) + g(2)`, int64(12)},
		// 7. Data Structures: Map of functions
		{`var m = map{"f": func(x){return x+1;}}; var f = m["f"]; f(10)`, int64(11)},
		// 8. Iteration: Foreach over enumerate_map
		{`var m = map{"a": 10, "b": 20}; var s = 0; foreach p in enumerate_map(m) { s += p[1]; } s`, int64(30)},
		// 9. OOP: Struct method returning a new instance
		{`struct Factory { func create() { return new Item(); } } struct Item { var id = 100; } var f = new Factory(); f.create().id`, int64(100)},
		// 10. Integration: Combining multiple builtins and operators
		{`length(split("a,b,c", ",")) * 10 + ord('0')`, int64(78)},
	}

	for _, tt := range tests {
		p := parser.NewParser(tt.input)
		rootNode := p.Parse()
		evaluator := NewEvaluator()
		evaluator.SetParser(p)
		result := evaluator.Eval(rootNode)

		switch exp := tt.expected.(type) {
		case int64:
			AssertInteger(t, result, exp)
		case float64:
			AssertFloat(t, result, exp)
		case string:
			AssertString(t, result, exp)
		case bool:
			AssertBoolean(t, result, exp)
		}
	}
}

// TestEvaluator_TimeFunctions verifies evaluation of time builtin functions
func TestEvaluator_TimeFunctions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{`typeof(now())`, "int"},
		{`typeof(now_ms())`, "int"},
		{`typeof(utc_now())`, "int"},
		{`typeof(timezone())`, "string"},
		{`format_time(parse_time("2023-10-27", "2006-01-02"), "2006-01-02")`, "2023-10-27"},
	}

	for _, tt := range tests {
		p := parser.NewParser(tt.input)
		rootNode := p.Parse()
		evaluator := NewEvaluator()
		evaluator.SetParser(p)
		result := evaluator.Eval(rootNode)

		switch exp := tt.expected.(type) {
		case string:
			AssertString(t, result, exp)
		case int64:
			AssertInteger(t, result, exp)
		}
	}
}

// TestEvaluator_FormatFunctions verifies evaluation of type conversion builtin functions
func TestEvaluator_FormatFunctions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{`to_int("42")`, int64(42)},
		{`to_float("3.14")`, 3.14},
		{`to_bool(1)`, true},
		{`to_string(123)`, "123"},
		{`to_int('A')`, int64(65)},
	}

	for _, tt := range tests {
		p := parser.NewParser(tt.input)
		rootNode := p.Parse()
		evaluator := NewEvaluator()
		evaluator.SetParser(p)
		result := evaluator.Eval(rootNode)

		switch exp := tt.expected.(type) {
		case int64:
			AssertInteger(t, result, exp)
		case float64:
			AssertFloat(t, result, exp)
		case bool:
			AssertBoolean(t, result, exp)
		case string:
			AssertString(t, result, exp)
		}
	}
}

// TestEvaluator_JSONFunctions verifies evaluation of JSON decoding builtin functions
func TestEvaluator_JSONFunctions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{`var m = json_string_to_map("{\"name\": \"John\", \"age\": 25}"); m["name"]`, "John"},
		{`var m = json_string_to_map("{\"name\": \"John\", \"age\": 25}"); m["age"]`, int64(25)},
		{`var a = json_string_to_map("[1, 2, 3]"); a[1]`, int64(2)},
	}

	for _, tt := range tests {
		p := parser.NewParser(tt.input)
		rootNode := p.Parse()
		evaluator := NewEvaluator()
		evaluator.SetParser(p)
		result := evaluator.Eval(rootNode)

		switch exp := tt.expected.(type) {
		case string:
			AssertString(t, result, exp)
		case int64:
			AssertInteger(t, result, exp)
		}
	}
}

// TestEvaluator_SortFunction verifies evaluation of the sort builtin function
func TestEvaluator_SortFunction(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{`var a = [3, 1, 4, 2]; sort(a); a`, "[1, 2, 3, 4]"},
		{`var a = ["c", "a", "b"]; sort(a); a`, "[a, b, c]"},
		{`var a = [1, 2, 3]; sort(a, true); a`, "[3, 2, 1]"},
	}

	for _, tt := range tests {
		p := parser.NewParser(tt.input)
		rootNode := p.Parse()
		evaluator := NewEvaluator()
		evaluator.SetParser(p)
		result := evaluator.Eval(rootNode)
		if result.ToString() != tt.expected {
			t.Errorf("input %s: expected %s, got %s", tt.input, tt.expected, result.ToString())
		}
	}
}

// TestEvaluator_SortedArray verifies evaluation of the sorted_array builtin function
func TestEvaluator_SortedArray(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{`var a = [3, 1, 4, 2]; var b = sorted_array(a); b`, "[1, 2, 3, 4]"},
		{`var a = [3, 1, 4, 2]; sorted_array(a); a`, "[3, 1, 4, 2]"}, // verify original is not modified
		{`var a = [1, 2, 3]; sorted_array(a, true)`, "[3, 2, 1]"},
		{`var a = [1, 2, 3]; sorted(a, true)`, "[3, 2, 1]"},
	}

	for _, tt := range tests {
		p := parser.NewParser(tt.input)
		rootNode := p.Parse()
		evaluator := NewEvaluator()
		evaluator.SetParser(p)
		result := evaluator.Eval(rootNode)
		if result.ToString() != tt.expected {
			t.Errorf("input %s: expected %s, got %s", tt.input, tt.expected, result.ToString())
		}
	}
}

// TestEvaluator_CloneArray verifies evaluation of the clone_array builtin function
func TestEvaluator_CloneArray(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{`var a = [1, 2, 3]; var b = clone_array(a); b`, "[1, 2, 3]"},
		{`var a = [1, 2, 3]; var b = clone_array(a); pop_array(b); a`, "[1, 2, 3]"}, // verify original is not modified
	}

	for _, tt := range tests {
		p := parser.NewParser(tt.input)
		rootNode := p.Parse()
		evaluator := NewEvaluator()
		evaluator.SetParser(p)
		result := evaluator.Eval(rootNode)
		if result.ToString() != tt.expected {
			t.Errorf("input %s: expected %s, got %s", tt.input, tt.expected, result.ToString())
		}
	}
}

// TestEvaluator_ReferenceEquality verifies the is_same_ref builtin function
func TestEvaluator_ReferenceEquality(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{`var a = [1]; var b = a; is_same_ref(a, b)`, true},
		{`var a = [1]; var b = [1]; is_same_ref(a, b)`, false},
		{`var a = map{"x": 1}; var b = a; is_same_ref(a, b)`, true},
		{`is_same_ref(nil, nil)`, true},
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

// TestEvaluator_CustomSort verifies the csort builtin function with a custom comparator
func TestEvaluator_CustomSort(t *testing.T) {
	input := `
		var a = ["apple", "go", "banana", "c"];
		var cmp = func(x, y) { return length(x) < length(y); };
		var b = csorted(a, cmp);
		csort(a, cmp);
		tostring(a) + " " + tostring(b)
	`
	p := parser.NewParser(input)
	rootNode := p.Parse()
	evaluator := NewEvaluator()
	evaluator.SetParser(p)
	result := evaluator.Eval(rootNode)
	expected := "[c, go, apple, banana] [c, go, apple, banana]"
	if result.ToString() != expected {
		t.Errorf("expected %s, got %s", expected, result.ToString())
	}
}

// TestEvaluator_MapArray verifies the map_array builtin function
func TestEvaluator_MapArray(t *testing.T) {
	input := `
		var a = [1, 2, 3];
		var b = map_array(a, func(x) { return x * x; });
		tostring(b)
	`
	p := parser.NewParser(input)
	rootNode := p.Parse()
	evaluator := NewEvaluator()
	evaluator.SetParser(p)
	result := evaluator.Eval(rootNode)
	expected := "[1, 4, 9]"
	if result.ToString() != expected {
		t.Errorf("expected %s, got %s", expected, result.ToString())
	}
}

// TestEvaluator_FilterArray verifies the filter_array builtin function
func TestEvaluator_FilterArray(t *testing.T) {
	input := `
		var a = [1, 2, 3, 4, 5, 6];
		var b = filter_array(a, func(x) { return x % 2 == 0; });
		tostring(b)
	`
	p := parser.NewParser(input)
	rootNode := p.Parse()
	evaluator := NewEvaluator()
	evaluator.SetParser(p)
	result := evaluator.Eval(rootNode)
	expected := "[2, 4, 6]"
	if result.ToString() != expected {
		t.Errorf("expected %s, got %s", expected, result.ToString())
	}
}

// TestEvaluator_ReduceArray verifies the reduce_array builtin function
func TestEvaluator_ReduceArray(t *testing.T) {
	input := `
		var a = [1, 2, 3, 4];
		var sum = reduce_array(a, func(acc, x) { return acc + x; }, 0);
		sum
	`
	p := parser.NewParser(input)
	rootNode := p.Parse()
	evaluator := NewEvaluator()
	evaluator.SetParser(p)
	result := evaluator.Eval(rootNode)
	AssertInteger(t, result, 10)
}

// TestEvaluator_FindArray verifies the find_array builtin function
func TestEvaluator_FindArray(t *testing.T) {
	input := `
		var a = [1, 2, 3, 4];
		var found = find(a, func(x) { return x > 2; });
		found
	`
	p := parser.NewParser(input)
	rootNode := p.Parse()
	evaluator := NewEvaluator()
	evaluator.SetParser(p)
	result := evaluator.Eval(rootNode)
	AssertInteger(t, result, 3)
}

// TestEvaluator_JSONEncode verifies the json_encode builtin function
func TestEvaluator_JSONEncode(t *testing.T) {
	input := `
		var m = map{"a": 1, "b": [true, nil]};
		json_encode(m)
	`
	p := parser.NewParser(input)
	rootNode := p.Parse()
	evaluator := NewEvaluator()
	evaluator.SetParser(p)
	result := evaluator.Eval(rootNode)
	// JSON map keys are sorted alphabetically by Go's json.Marshal
	expected := `{"a":1,"b":[true,null]}`
	if result.ToString() != expected {
		t.Errorf("expected %s, got %s", expected, result.ToString())
	}
}

// TestEvaluator_ImportStatement verifies the import statement evaluation
func TestEvaluator_ImportStatement(t *testing.T) {
	input := `import math;`
	p := parser.NewParser(input)
	rootNode := p.Parse()
	evaluator := NewEvaluator()
	evaluator.SetParser(p)
	result := evaluator.Eval(rootNode)

	// Import should return the package object
	if result.GetType() != std.PackageType {
		t.Errorf("expected PackageType, got %s", result.GetType())
	}
	if pkg, ok := result.(*std.Package); ok {
		if pkg.Name != "math" {
			t.Errorf("expected package name 'math', got '%s'", pkg.Name)
		}
	} else {
		t.Errorf("expected *Package, got %T", result)
	}
}

// TestEvaluator_PackageFunctionCall verifies package function calls
func TestEvaluator_PackageFunctionCall(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
		checkInt bool
	}{
		{`import math; math.abs(-5)`, int64(5), true},
		{`import math; math.abs(10)`, int64(10), true},
		{`import math; math.min(5, 3)`, int64(3), true},
		{`import math; math.max(5, 3)`, int64(5), true},
		{`import math; math.floor(3.7)`, int64(3), true},
		{`import math; math.ceil(3.2)`, int64(4), true},
	}

	for _, tt := range tests {
		p := parser.NewParser(tt.input)
		rootNode := p.Parse()
		evaluator := NewEvaluator()
		evaluator.SetParser(p)
		result := evaluator.Eval(rootNode)

		if tt.checkInt {
			if result.GetType() != std.IntegerType {
				t.Errorf("for input '%s': expected IntegerType, got %s", tt.input, result.GetType())
			}
			intResult := result.(*std.Integer).Value
			if intResult != tt.expected.(int64) {
				t.Errorf("for input '%s': expected %d, got %d", tt.input, tt.expected.(int64), intResult)
			}
		}
	}
}

// TestEvaluator_PackageFunctionCallWithFloat verifies package function calls with float results
func TestEvaluator_PackageFunctionCallWithFloat(t *testing.T) {
	tests := []struct {
		input    string
		expected float64
		delta    float64
	}{
		{`import math; math.sqrt(16)`, 4.0, 1e-9},
		{`import math; math.pow(2, 3)`, 8.0, 1e-9},
		{`import math; math.sqrt(2)`, 1.41421356, 1e-5},
	}

	for _, tt := range tests {
		p := parser.NewParser(tt.input)
		rootNode := p.Parse()
		evaluator := NewEvaluator()
		evaluator.SetParser(p)
		result := evaluator.Eval(rootNode)

		if result.GetType() != std.FloatType {
			t.Errorf("for input '%s': expected FloatType, got %s", tt.input, result.GetType())
		}
		floatResult := result.(*std.Float).Value
		if floatResult < tt.expected-tt.delta || floatResult > tt.expected+tt.delta {
			t.Errorf("for input '%s': expected %f (±%f), got %f", tt.input, tt.expected, tt.delta, floatResult)
		}
	}
}

// TestEvaluator_MultiplePackageImports verifies multiple package imports
func TestEvaluator_MultiplePackageImports(t *testing.T) {
	input := `
		import math;
		var x = math.abs(-10);
		var y = math.pow(2, 3);
		x + y
	`
	p := parser.NewParser(input)
	rootNode := p.Parse()
	evaluator := NewEvaluator()
	evaluator.SetParser(p)
	result := evaluator.Eval(rootNode)

	if result.GetType() != std.FloatType {
		t.Errorf("expected FloatType, got %s", result.GetType())
	}
	// 10 + 8.0 = 18.0
	expected := 18.0
	actual := result.(*std.Float).Value
	if actual < expected-1e-9 || actual > expected+1e-9 {
		t.Errorf("expected %f, got %f", expected, actual)
	}
}

// TestEvaluator_PackageFunctionInLoop verifies package functions work in loops
func TestEvaluator_PackageFunctionInLoop(t *testing.T) {
	input := `
		import math;
		var sum = 0;
		for (var i = 1; i <= 3; i = i + 1) {
			sum = sum + math.pow(i, 2);
		}
		sum
	`
	p := parser.NewParser(input)
	rootNode := p.Parse()
	evaluator := NewEvaluator()
	evaluator.SetParser(p)
	result := evaluator.Eval(rootNode)

	// math.pow returns float, so sum becomes float
	if result.GetType() != std.FloatType {
		t.Errorf("expected FloatType, got %s", result.GetType())
	}
	// 1 + 4 + 9 = 14
	expected := 14.0
	actual := result.(*std.Float).Value
	if actual < expected-1e-9 || actual > expected+1e-9 {
		t.Errorf("expected %f, got %f", expected, actual)
	}
}

// TestEvaluator_PackageFunctionInConditional verifies package functions work in conditionals
func TestEvaluator_PackageFunctionInConditional(t *testing.T) {
	input := `
		import math;
		var x = -5;
		if (math.abs(x) > 3) {
			"big"
		} else {
			"small"
		}
	`
	p := parser.NewParser(input)
	rootNode := p.Parse()
	evaluator := NewEvaluator()
	evaluator.SetParser(p)
	result := evaluator.Eval(rootNode)

	if result.GetType() != std.StringType {
		t.Errorf("expected StringType, got %s", result.GetType())
	}
	if result.(*std.String).Value != "big" {
		t.Errorf("expected 'big', got '%s'", result.(*std.String).Value)
	}
}

// TestEvaluator_PackageFunctionInFunctionCall verifies package functions work in function calls
func TestEvaluator_PackageFunctionInFunctionCall(t *testing.T) {
	input := `
		import math;
		func square(x) {
			return math.pow(x, 2);
		}
		square(5)
	`
	p := parser.NewParser(input)
	rootNode := p.Parse()
	evaluator := NewEvaluator()
	evaluator.SetParser(p)
	result := evaluator.Eval(rootNode)

	// math.pow returns float
	if result.GetType() != std.FloatType {
		t.Errorf("expected FloatType, got %s", result.GetType())
	}
	// 5^2 = 25
	expected := 25.0
	actual := result.(*std.Float).Value
	if actual < expected-1e-9 || actual > expected+1e-9 {
		t.Errorf("expected %f, got %f", expected, actual)
	}
}

// TestEvaluator_PackageNotFound verifies error handling for non-existent packages
func TestEvaluator_PackageNotFound(t *testing.T) {
	input := `import nonexistent;`
	p := parser.NewParser(input)
	rootNode := p.Parse()
	evaluator := NewEvaluator()
	evaluator.SetParser(p)
	result := evaluator.Eval(rootNode)

	if result.GetType() != std.ErrorType {
		t.Errorf("expected ErrorType, got %s", result.GetType())
	}
	if !strings.Contains(result.(*std.Error).Message, "package") {
		t.Errorf("expected error message to contain 'package', got '%s'", result.(*std.Error).Message)
	}
}

// TestEvaluator_PackageFunctionNotFound verifies error handling for non-existent package functions
func TestEvaluator_PackageFunctionNotFound(t *testing.T) {
	input := `import math; math.nonexistent(5);`
	p := parser.NewParser(input)
	rootNode := p.Parse()
	evaluator := NewEvaluator()
	evaluator.SetParser(p)
	result := evaluator.Eval(rootNode)

	if result.GetType() != std.ErrorType {
		t.Errorf("expected ErrorType, got %s", result.GetType())
	}
	if !strings.Contains(result.(*std.Error).Message, "function") {
		t.Errorf("expected error message to contain 'function', got '%s'", result.(*std.Error).Message)
	}
}

// ==================== PACKAGE IMPORT TESTS ====================

// TestPackageImport_Strings tests importing and using the strings package
func TestPackageImport_Strings(t *testing.T) {
	input := `
		import strings;
		strings.upper("hello")
	`
	p := parser.NewParser(input)
	rootNode := p.Parse()
	evaluator := NewEvaluator()
	evaluator.SetParser(p)
	result := evaluator.Eval(rootNode)

	if result.GetType() != std.StringType {
		t.Errorf("expected StringType, got %s", result.GetType())
	}
	if result.(*std.String).Value != "HELLO" {
		t.Errorf("expected 'HELLO', got '%s'", result.(*std.String).Value)
	}
}

// TestPackageImport_StringsMultiple tests multiple string functions
func TestPackageImport_StringsMultiple(t *testing.T) {
	input := `
		import strings;
		var s = "hello world";
		var upper = strings.upper(s);
		var lower = strings.lower("HELLO");
		upper + " " + lower
	`
	p := parser.NewParser(input)
	rootNode := p.Parse()
	evaluator := NewEvaluator()
	evaluator.SetParser(p)
	result := evaluator.Eval(rootNode)

	if result.GetType() != std.StringType {
		t.Errorf("expected StringType, got %s", result.GetType())
	}
	expected := "HELLO WORLD hello"
	if result.(*std.String).Value != expected {
		t.Errorf("expected '%s', got '%s'", expected, result.(*std.String).Value)
	}
}

// TestPackageImport_Time tests importing and using the time package
func TestPackageImport_Time(t *testing.T) {
	input := `
		import time;
		var ts = time.now();
		ts > 0
	`
	p := parser.NewParser(input)
	rootNode := p.Parse()
	evaluator := NewEvaluator()
	evaluator.SetParser(p)
	result := evaluator.Eval(rootNode)

	if result.GetType() != std.BooleanType {
		t.Errorf("expected BooleanType, got %s", result.GetType())
	}
	if !result.(*std.Boolean).Value {
		t.Errorf("expected true, got false")
	}
}

// TestPackageImport_Arrays tests importing and using the arrays package
func TestPackageImport_Arrays(t *testing.T) {
	input := `
		import arrays;
		var a = [1, 2, 3];
		arrays.pop_array(a)
	`
	p := parser.NewParser(input)
	rootNode := p.Parse()
	evaluator := NewEvaluator()
	evaluator.SetParser(p)
	result := evaluator.Eval(rootNode)

	if result.GetType() == std.ErrorType {
		t.Errorf("got error: %s", result.(*std.Error).Message)
		return
	}

	if result.GetType() != std.IntegerType {
		t.Errorf("expected IntegerType, got %s", result.GetType())
	}
	if result.(*std.Integer).Value != 3 {
		t.Errorf("expected 3, got %d", result.(*std.Integer).Value)
	}
}

// TestPackageImport_Os tests importing and using the os package
func TestPackageImport_Os(t *testing.T) {
	input := `
		import os;
		var args = os.args();
		typeof(args)
	`
	p := parser.NewParser(input)
	rootNode := p.Parse()
	evaluator := NewEvaluator()
	evaluator.SetParser(p)
	result := evaluator.Eval(rootNode)

	if result.GetType() != std.StringType {
		t.Errorf("expected StringType, got %s", result.GetType())
	}
	if result.(*std.String).Value != "array" {
		t.Errorf("expected 'array', got '%s'", result.(*std.String).Value)
	}
}

// TestPackageImport_Io tests importing and using the io package
func TestPackageImport_Io(t *testing.T) {
	input := `
		import io;
		var formatted = io.sprintf("Number: %d", 42);
		formatted
	`
	p := parser.NewParser(input)
	rootNode := p.Parse()
	evaluator := NewEvaluator()
	evaluator.SetParser(p)
	result := evaluator.Eval(rootNode)

	if result.GetType() != std.StringType {
		t.Errorf("expected StringType, got %s", result.GetType())
	}
	expected := "Number: 42"
	if result.(*std.String).Value != expected {
		t.Errorf("expected '%s', got '%s'", expected, result.(*std.String).Value)
	}
}

// TestPackageImport_List tests importing and using the list package
func TestPackageImport_List(t *testing.T) {
	input := `
		import list;
		var l = list.list(1, 2, 3);
		list.size_list(l)
	`
	p := parser.NewParser(input)
	rootNode := p.Parse()
	evaluator := NewEvaluator()
	evaluator.SetParser(p)
	result := evaluator.Eval(rootNode)

	if result.GetType() != std.IntegerType {
		t.Errorf("expected IntegerType, got %s", result.GetType())
	}
	if result.(*std.Integer).Value != 3 {
		t.Errorf("expected 3, got %d", result.(*std.Integer).Value)
	}
}

// TestPackageImport_Map tests importing and using the map package
func TestPackageImport_Map(t *testing.T) {
	input := `
		import maps;
		var m = map{"name": "John", "age": 30};
		var keys = maps.keys_map(m);
		1
	`
	p := parser.NewParser(input)
	rootNode := p.Parse()
	evaluator := NewEvaluator()
	evaluator.SetParser(p)
	result := evaluator.Eval(rootNode)

	if result.GetType() != std.IntegerType {
		t.Errorf("expected IntegerType, got %s", result.GetType())
	}
	if result.(*std.Integer).Value != 1 {
		t.Errorf("expected 1, got %d", result.(*std.Integer).Value)
	}
}

// TestPackageImport_Set tests importing and using the set package
func TestPackageImport_Set(t *testing.T) {
	input := `
		import sets;
		var s = set{1, 2, 3}
		var values = sets.values_set(s);
		typeof(values)
	`
	p := parser.NewParser(input)
	rootNode := p.Parse()
	evaluator := NewEvaluator()
	evaluator.SetParser(p)
	result := evaluator.Eval(rootNode)

	if result.GetType() == std.ErrorType {
		t.Errorf("got error: %s", result.(*std.Error).Message)
		return
	}

	if result.GetType() != std.StringType {
		t.Errorf("expected StringType, got %s", result.GetType())
	}
	if result.(*std.String).Value != "array" {
		t.Errorf("expected 'array', got '%s'", result.(*std.String).Value)
	}
}

// TestPackageImport_Format tests importing and using the format package
func TestPackageImport_Format(t *testing.T) {
	input := `
		import format;
		var num = format.to_int("42");
		num + 8
	`
	p := parser.NewParser(input)
	rootNode := p.Parse()
	evaluator := NewEvaluator()
	evaluator.SetParser(p)
	result := evaluator.Eval(rootNode)

	if result.GetType() != std.IntegerType {
		t.Errorf("expected IntegerType, got %s", result.GetType())
	}
	if result.(*std.Integer).Value != 50 {
		t.Errorf("expected 50, got %d", result.(*std.Integer).Value)
	}
}

// TestPackageImport_File tests importing and using the file package
func TestPackageImport_File(t *testing.T) {
	input := `
		import file;
		var exists = file.file_exists("./go-mix");
		typeof(exists)
	`
	p := parser.NewParser(input)
	rootNode := p.Parse()
	evaluator := NewEvaluator()
	evaluator.SetParser(p)
	result := evaluator.Eval(rootNode)

	if result.GetType() != std.StringType {
		t.Errorf("expected StringType, got %s", result.GetType())
	}
	if result.(*std.String).Value != "bool" {
		t.Errorf("expected 'bool', got '%s'", result.(*std.String).Value)
	}
}

// TestPackageImport_Tuple tests importing and using the tuple package
func TestPackageImport_Tuple(t *testing.T) {
	input := `
		import tuple;
		var t = tuple.tuple(1, "hello", true);
		tuple.size_tuple(t)
	`
	p := parser.NewParser(input)
	rootNode := p.Parse()
	evaluator := NewEvaluator()
	evaluator.SetParser(p)
	result := evaluator.Eval(rootNode)

	if result.GetType() != std.IntegerType {
		t.Errorf("expected IntegerType, got %s", result.GetType())
	}
	if result.(*std.Integer).Value != 3 {
		t.Errorf("expected 3, got %d", result.(*std.Integer).Value)
	}
}

// TestPackageImport_Common tests importing and using the common package
func TestPackageImport_Common(t *testing.T) {
	input := `
		import common;
		var arr = [3, 1, 4, 1, 5, 9];
		var sorted = common.sorted(arr);
		typeof(sorted)
	`
	p := parser.NewParser(input)
	rootNode := p.Parse()
	evaluator := NewEvaluator()
	evaluator.SetParser(p)
	result := evaluator.Eval(rootNode)

	if result.GetType() != std.StringType {
		t.Errorf("expected StringType, got %s", result.GetType())
	}
	if result.(*std.String).Value != "array" {
		t.Errorf("expected 'array', got '%s'", result.(*std.String).Value)
	}
}

// TestPackageImport_MultiplePackages tests importing multiple packages in one program
func TestPackageImport_MultiplePackages(t *testing.T) {
	input := `
		import math;
		import strings;
		var x = math.abs(-5);
		var s = strings.upper("hi");
		x + 2
	`
	p := parser.NewParser(input)
	rootNode := p.Parse()
	evaluator := NewEvaluator()
	evaluator.SetParser(p)
	result := evaluator.Eval(rootNode)

	if result.GetType() != std.IntegerType {
		t.Errorf("expected IntegerType, got %s", result.GetType())
	}
	// 5 + 2 = 7
	if result.(*std.Integer).Value != 7 {
		t.Errorf("expected 7, got %d", result.(*std.Integer).Value)
	}
}

// TestPackageImport_PackageInFunction tests using imported packages inside functions
func TestPackageImport_PackageInFunction(t *testing.T) {
	input := `
		import strings;
		func process(str) {
			return strings.upper(str) + strings.lower("WORLD");
		}
		process("hello ")
	`
	p := parser.NewParser(input)
	rootNode := p.Parse()
	evaluator := NewEvaluator()
	evaluator.SetParser(p)
	result := evaluator.Eval(rootNode)

	if result.GetType() != std.StringType {
		t.Errorf("expected StringType, got %s", result.GetType())
	}
	expected := "HELLO world"
	if result.(*std.String).Value != expected {
		t.Errorf("expected '%s', got '%s'", expected, result.(*std.String).Value)
	}
}

func TestListAndTupleSearchFunctions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		// List Search Tests
		{`var l = list(1, 2, 3, 4); find_list(l, func(x) { return x == 3; });`, 3},
		{`var l = list(1, 2, 3, 4); find_list(l, func(x) { return x == 99; });`, nil},
		{`var l = list(1, 2, 3); some_list(l, func(x) { return x > 2; });`, true},
		{`var l = list(1, 2, 3); some_list(l, func(x) { return x > 5; });`, false},
		{`var l = list(1, 2, 3); every_list(l, func(x) { return x > 0; });`, true},
		{`var l = list(1, 2, 3); every_list(l, func(x) { return x > 1; });`, false},

		// Tuple Search Tests
		{`var t = tuple(1, 2, 3, 4); find_tuple(t, func(x) { return x == 3; });`, 3},
		{`var t = tuple(1, 2, 3, 4); find_tuple(t, func(x) { return x == 99; });`, nil},
		{`var t = tuple(1, 2, 3); some_tuple(t, func(x) { return x > 2; });`, true},
		{`var t = tuple(1, 2, 3); some_tuple(t, func(x) { return x > 5; });`, false},
		{`var t = tuple(1, 2, 3); every_tuple(t, func(x) { return x > 0; });`, true},
		{`var t = tuple(1, 2, 3); every_tuple(t, func(x) { return x > 1; });`, false},
	}

	for _, tt := range tests {
		p := parser.NewParser(tt.input)
		rootNode := p.Parse()
		evaluator := NewEvaluator()
		evaluator.SetParser(p)
		result := evaluator.Eval(rootNode)

		switch exp := tt.expected.(type) {
		case int:
			AssertInteger(t, result, int64(exp))
		case bool:
			AssertBoolean(t, result, exp)
		case nil:
			if result.GetType() != std.NilType {
				t.Errorf("expected NilType, got %s", result.GetType())
			}
		}
	}
}
func TestTypeConversionFunctions(t *testing.T) {
	input := `
	var arr = [1, 2, 3];
	var l = to_list(arr);
	var t = to_tuple(l);
	var arr2 = to_array(t);
	return [typeof(l), typeof(t), typeof(arr2)];
	`
	root := parser.NewParser(input).Parse()
	evaluator := NewEvaluator()
	evaluator.SetParser(parser.NewParser(input))
	evaluated := evaluator.Eval(root)

	if evaluated.GetType() != std.ArrayType {
		t.Fatalf("expected array, got %s", evaluated.GetType())
	}

	arr, ok := evaluated.(*std.Array)
	if !ok {
		t.Fatalf("expected array, got %T", evaluated)
	}

	if len(arr.Elements) != 3 {
		t.Fatalf("expected 3 elements, got %d", len(arr.Elements))
	}

	if arr.Elements[0].ToString() != "list" {
		t.Errorf("expected list, got %s", arr.Elements[0].ToString())
	}
	if arr.Elements[1].ToString() != "tuple" {
		t.Errorf("expected tuple, got %s", arr.Elements[1].ToString())
	}
	if arr.Elements[2].ToString() != "array" {
		t.Errorf("expected array, got %s", arr.Elements[2].ToString())
	}
}
