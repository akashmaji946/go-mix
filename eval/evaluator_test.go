package eval

import (
	"fmt"
	"math"
	"testing"

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
	}

	for _, tt := range tests {
		rootNode := parser.NewParser(tt.input).Parse()
		evaluator := NewEvaluator()
		result := evaluator.Eval(rootNode)
		if result.GetType() != IntegerType {
			t.Errorf("expected %s, got %s", IntegerType, result.GetType())
		}
		if result.(*Integer).Value != tt.expected {
			t.Errorf("expected %d, got %d", tt.expected, result.(*Integer).Value)
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
	}

	for _, tt := range tests {
		rootNode := parser.NewParser(tt.input).Parse()
		evaluator := NewEvaluator()
		result := evaluator.Eval(rootNode)
		if result.GetType() != FloatType {
			t.Errorf("expected %s, got %s", FloatType, result.GetType())
		}
		if result.(*Float).Value != tt.expected {
			t.Errorf("expected %f, got %f", tt.expected, result.(*Float).Value)
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
	}

	for _, tt := range tests {
		rootNode := parser.NewParser(tt.input).Parse()
		evaluator := NewEvaluator()
		result := evaluator.Eval(rootNode)
		if result.GetType() != BooleanType {
			t.Errorf("expected %s, got %s", BooleanType, result.GetType())
		}
		if result.(*Boolean).Value != tt.expected {
			t.Errorf("expected %t, got %t", tt.expected, result.(*Boolean).Value)
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
		if result.GetType() != NilType {
			t.Errorf("expected %s, got %s", NilType, result.GetType())
		}

		if val, ok := result.(*Nil); ok {
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
		if result.GetType() != StringType {
			t.Errorf("expected %s, got %s", StringType, result.GetType())
		}
		if result.(*String).Value != tt.expected {
			t.Errorf("expected %s, got %s", tt.expected, result.(*String).Value)
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
		if result.GetType() != IntegerType {
			t.Errorf("expected %s, got %s", IntegerType, result.GetType())
		}
		if result.(*Integer).Value != tt.expected {
			t.Errorf("expected %d, got %d", tt.expected, result.(*Integer).Value)
		}
	}
}
