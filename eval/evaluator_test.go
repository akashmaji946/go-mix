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
		{"1 + 2", 3},
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
