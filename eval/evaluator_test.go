package eval

import (
	"fmt"
	"math"
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
		rootNode := parser.NewParser(tt.input).Parse()
		evaluator := NewEvaluator()
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
		rootNode := parser.NewParser(tt.input).Parse()
		evaluator := NewEvaluator()
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
		rootNode := parser.NewParser(tt.input).Parse()
		evaluator := NewEvaluator()
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
		rootNode := parser.NewParser(tt.input).Parse()
		evaluator := NewEvaluator()
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
		rootNode := parser.NewParser(tt.input).Parse()
		evaluator := NewEvaluator()
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
		rootNode := parser.NewParser(tt.input).Parse()
		evaluator := NewEvaluator()
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
			"[ERROR]: Operator (+) not implemented for (int) and (bool)",
		},
		{
			"1 - true",
			"[ERROR]: Operator (-) not implemented for (int) and (bool)",
		},
		{
			"1 * true",
			"[ERROR]: Operator (*) not implemented for (int) and (bool)",
		},
		{
			"1 / true",
			"[ERROR]: Operator (/) not implemented for (int) and (bool)",
		},
		{
			"1 % true",
			"[ERROR]: Operator (%) not implemented for (int) and (bool)",
		},
		{
			"1 & true",
			"[ERROR]: Operator (&) not implemented for (int) and (bool)",
		},
		{
			"1 | true",
			"[ERROR]: Operator (|) not implemented for (int) and (bool)",
		},
		{
			"1 ^ true",
			"[ERROR]: Operator (^) not implemented for (int) and (bool)",
		},
		{
			"~true",
			"[ERROR]: Operator (~) not implemented for (bool)",
		},
		{
			"1 << true",
			"[ERROR]: Operator (<<) not implemented for (int) and (bool)",
		},
		{
			"1 >> true",
			"[ERROR]: Operator (>>) not implemented for (int) and (bool)",
		},
		{
			"true + true",
			"[ERROR]: Operator (+) not implemented for (bool) and (bool)",
		},
		{
			"true - true",
			"[ERROR]: Operator (-) not implemented for (bool) and (bool)",
		},
		{
			"true * true",
			"[ERROR]: Operator (*) not implemented for (bool) and (bool)",
		},
		{
			"true / true",
			"[ERROR]: Operator (/) not implemented for (bool) and (bool)",
		},
		{
			"true % true",
			"[ERROR]: Operator (%) not implemented for (bool) and (bool)",
		},
		{
			"true & true",
			"[ERROR]: Operator (&) not implemented for (bool) and (bool)",
		},
		{
			"true | true",
			"[ERROR]: Operator (|) not implemented for (bool) and (bool)",
		},
		{
			"true ^ true",
			"[ERROR]: Operator (^) not implemented for (bool) and (bool)",
		},
		{
			"~true",
			"[ERROR]: Operator (~) not implemented for (bool)",
		},
		{
			"true << true",
			"[ERROR]: Operator (<<) not implemented for (bool) and (bool)",
		},
		{
			"true >> true",
			"[ERROR]: Operator (>>) not implemented for (bool) and (bool)",
		},

		{
			`
				if (true) {
					!1
					false
				}
				`,
			"[ERROR]: Operator (!) not implemented for (int)",
		},
	}

	for _, tt := range tests {
		rootNode := parser.NewParser(tt.Src).Parse()
		evaluator := NewEvaluator()
		result := evaluator.Eval(rootNode)
		if result.GetType() != objects.ErrorType {
			t.Errorf("expected %s, got %s", objects.ErrorType, result.GetType())
		}
		if result.(*objects.Error).Message != tt.ExpectedErrorMsg {
			t.Errorf("expected %s, got %s", tt.ExpectedErrorMsg, result.(*objects.Error).Message)
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
		rootNode := parser.NewParser(tt.Src).Parse()
		evaluator := NewEvaluator()
		result := evaluator.Eval(rootNode)
		AssertBoolean(t, result, tt.Expected)

	}

	errorTests := []struct {
		Src              string
		ExpectedErrorMsg string
	}{
		{
			`if(1) { true }`,
			"[ERROR]: Conditional expression must be (bool)",
		},
		{
			`if(2 + 2 * 3) { true }`,
			"[ERROR]: Conditional expression must be (bool)",
		},
	}
	for _, tt := range errorTests {
		rootNode := parser.NewParser(tt.Src).Parse()
		evaluator := NewEvaluator()
		result := evaluator.Eval(rootNode)
		if result.GetType() != objects.ErrorType {
			t.Errorf("expected %s, got %s", objects.ErrorType, result.GetType())
		}
		if result.(*objects.Error).Message != tt.ExpectedErrorMsg {
			t.Errorf("expected %s, got %s", tt.ExpectedErrorMsg, result.(*objects.Error).Message)
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
		rootNode := parser.NewParser(tt.Src).Parse()
		evaluator := NewEvaluator()
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
		rootNode := parser.NewParser(tt.Src).Parse()
		evaluator := NewEvaluator()
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
			"[ERROR]: identifier not found: (b)",
		},
		{
			`var a = 1; var b = 2; var c = a + b + c;`,
			"[ERROR]: identifier not found: (c)",
		},
		{
			`var a = 1; var a = 2; var c = a;`,
			"[ERROR]: identifier redeclaration found: (a)",
		},
	}

	for _, tt := range errorTests {
		rootNode := parser.NewParser(tt.Src).Parse()
		evaluator := NewEvaluator()
		result := evaluator.Eval(rootNode)
		if result.GetType() != objects.ErrorType {
			t.Errorf("expected %s, got %s", objects.ErrorType, result.GetType())
		}
		if result.(*objects.Error).Message != tt.ExpectedErrorMsg {
			t.Errorf("expected %s, got %s", tt.ExpectedErrorMsg, result.(*objects.Error).Message)
		}

	}

}

// function declarartion
