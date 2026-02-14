/*
File    : go-mix/std/math.go
Author  : Akash Maji
Contact : akashmaji(@iisc.ac.in)
*/

// Package std - math.go
// This file defines the math builtin functions available in the Go-Mix language.
//	It includes functions for absolute value, min/max, rounding, square root, power, and trigonometric functions.

package std

import (
	"io"
	"math"
	"math/rand"
	"time"
)

var mathMethods = []*Builtin{
	{Name: "abs", Callback: abs},          // Returns the absolute value of a number
	{Name: "fabs", Callback: fabs},        // Returns the absolute value of a floating-point number (alias for abs)
	{Name: "min", Callback: min},          // Returns the smaller of two numbers
	{Name: "max", Callback: max},          // Returns the larger of two numbers
	{Name: "floor", Callback: floor},      // Returns the largest integer less than or equal to a number
	{Name: "ceil", Callback: ceil},        // Returns the smallest integer greater than or equal to a number
	{Name: "round", Callback: round},      // Returns the nearest integer to a number
	{Name: "sqrt", Callback: sqrt},        // Returns the square root of a number
	{Name: "pow", Callback: pow},          // Returns the result of raising a number to a power
	{Name: "sin", Callback: sin},          // Returns the sine of the radian argument
	{Name: "cos", Callback: cos},          // Returns the cosine of the radian argument
	{Name: "tan", Callback: tan},          // Returns the tangent of the radian argument
	{Name: "asin", Callback: asin},        // Returns the arcsine, in radians
	{Name: "acos", Callback: acos},        // Returns the arccosine, in radians
	{Name: "atan", Callback: atan},        // Returns the arctangent, in radians
	{Name: "atan2", Callback: atan2},      // Returns the arctangent of y/x, in radians
	{Name: "log", Callback: log},          // Returns the natural logarithm
	{Name: "log10", Callback: log10},      // Returns the decimal logarithm
	{Name: "exp", Callback: exp},          // Returns e**x
	{Name: "rand", Callback: randFunc},    // Returns a random float [0.0, 1.0)
	{Name: "rand_int", Callback: randInt}, // Returns a random integer in range
}

// init registers the math methods as global builtins by appending them to the Builtins slice.
// It also registers the math package for import functionality.
func init() {
	// Register as global builtins (for backward compatibility)
	Builtins = append(Builtins, mathMethods...)

	// Register as a package (for import functionality)
	mathPackage := &Package{
		Name:      "math",
		Functions: make(map[string]*Builtin),
	}
	for _, method := range mathMethods {
		mathPackage.Functions[method.Name] = method
	}
	RegisterPackage(mathPackage)

	rand.Seed(time.Now().UnixNano())
}

// abs returns the absolute value of an integer.
//
// Syntax: abs(integer)
//
// Usage:
//
//	Returns the non-negative value of the given integer.
//	If the input is negative, it returns the positive equivalent.
//
// Example:
//
//	abs(-5); // Returns 5
//	abs(10); // Returns 10
func abs(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 1 {
		return createError("ERROR: wrong number of arguments. got=%d, want=1", len(args))
	}
	if args[0].GetType() != IntegerType {
		return createError("ERROR: argument to `abs` must be an integer, got '%s'", args[0].GetType())
	}

	value := args[0].(*Integer).Value
	if value < 0 {
		value = -value
	}
	return &Integer{Value: value}
}

// fabs returns the absolute value of a floating-point number.
//
// Syntax: fabs(float)
//
// Usage:
//
//	Returns the non-negative value of the given floating-point number.
//	This is the floating-point equivalent of the abs() function.
//
// Example:
//
//	fabs(-3.14); // Returns 3.14
//	fabs(2.5);   // Returns 2.5
func fabs(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 1 {
		return createError("ERROR: wrong number of arguments. got=%d, want=1", len(args))
	}
	if args[0].GetType() != FloatType {
		return createError("ERROR: argument to `fabs` must be a float, got '%s'", args[0].GetType())
	}

	value := args[0].(*Float).Value
	if value < 0 {
		value = -value
	}
	return &Float{Value: value}
}

// min returns the smaller of two numbers.
//
// Syntax: min(num1, num2)
//
// Usage:
//
//	Compares two numbers (integers or floats) and returns the smaller one.
//	If types are mixed, the result is promoted to float if necessary.
//
// Example:
//
//	min(10, 20);   // Returns 10
//	min(5.5, 2.1); // Returns 2.1
func min(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 2 {
		return createError("ERROR: wrong number of arguments. got=%d, want=2", len(args))
	}
	if (args[0].GetType() != IntegerType && args[0].GetType() != FloatType) ||
		(args[1].GetType() != IntegerType && args[1].GetType() != FloatType) {
		return createError("ERROR: arguments to `min` must be integers or floats, got '%s' and '%s'", args[0].GetType(), args[1].GetType())
	}

	var a, b float64
	if args[0].GetType() == IntegerType {
		a = float64(args[0].(*Integer).Value)
	} else {
		a = args[0].(*Float).Value
	}
	if args[1].GetType() == IntegerType {
		b = float64(args[1].(*Integer).Value)
	} else {
		b = args[1].(*Float).Value
	}

	if a < b {
		if args[0].GetType() == IntegerType {
			return &Integer{Value: int64(a)}
		}
		return &Float{Value: a}
	} else {
		if args[1].GetType() == IntegerType {
			return &Integer{Value: int64(b)}
		}
		return &Float{Value: b}
	}
}

// max returns the larger of two numbers.
//
// Syntax: max(num1, num2)
//
// Usage:
//
//	Compares two numbers (integers or floats) and returns the larger one.
//	If types are mixed, the result is promoted to float if necessary.
//
// Example:
//
//	max(10, 20);   // Returns 20
//	max(5.5, 2.1); // Returns 5.5
func max(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 2 {
		return createError("ERROR: wrong number of arguments. got=%d, want=2", len(args))
	}
	if (args[0].GetType() != IntegerType && args[0].GetType() != FloatType) ||
		(args[1].GetType() != IntegerType && args[1].GetType() != FloatType) {
		return createError("ERROR: arguments to `max` must be integers or floats, got '%s' and '%s'", args[0].GetType(), args[1].GetType())
	}

	var a, b float64
	if args[0].GetType() == IntegerType {
		a = float64(args[0].(*Integer).Value)
	} else {
		a = args[0].(*Float).Value
	}
	if args[1].GetType() == IntegerType {
		b = float64(args[1].(*Integer).Value)
	} else {
		b = args[1].(*Float).Value
	}

	if a > b {
		if args[0].GetType() == IntegerType {
			return &Integer{Value: int64(a)}
		}
		return &Float{Value: a}
	} else {
		if args[1].GetType() == IntegerType {
			return &Integer{Value: int64(b)}
		}
		return &Float{Value: b}
	}
}

// floor returns the largest integer less than or equal to a number.
//
// Syntax: floor(float)
//
// Usage:
//
//	Returns the greatest integer value that is less than or equal to the input float.
//
// Example:
//
//	floor(3.9);  // Returns 3
//	floor(-3.1); // Returns -4
func floor(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 1 {
		return createError("ERROR: wrong number of arguments. got=%d, want=1", len(args))
	}
	if args[0].GetType() != FloatType {
		return createError("ERROR: argument to `floor` must be a float, got '%s'", args[0].GetType())
	}

	value := args[0].(*Float).Value
	floorValue := int64(value)
	if value < 0 && value != float64(floorValue) {
		floorValue--
	}
	return &Integer{Value: floorValue}
}

// ceil returns the smallest integer greater than or equal to a number.
//
// Syntax: ceil(float)
//
// Usage:
//
//	Returns the smallest integer value that is greater than or equal to the input float.
//
// Example:
//
//	ceil(3.1);  // Returns 4
//	ceil(-3.9); // Returns -3
func ceil(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 1 {
		return createError("ERROR: wrong number of arguments. got=%d, want=1", len(args))
	}
	if args[0].GetType() != FloatType {
		return createError("ERROR: argument to `ceil` must be a float, got '%s'", args[0].GetType())
	}

	value := args[0].(*Float).Value
	ceilValue := int64(value)
	if value > 0 && value != float64(ceilValue) {
		ceilValue++
	}
	return &Integer{Value: ceilValue}
}

// round returns the nearest value to a number with specified precision.
//
// Syntax: round(float, [precision])
//
// Usage:
//
//	Rounds the given float to the nearest value.
//	The optional precision argument specifies the number of decimal places (defaults to 0).
//	Uses standard rounding rules (0.5 rounds up).
//
// Example:
//
//	round(3.5);       // Returns 4.0
//	round(3.14159, 2); // Returns 3.14
func round(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) == 0 || len(args) > 2 {
		return createError("ERROR: wrong number of arguments. got=%d, want=1 or 2", len(args))
	}
	if args[0].GetType() != FloatType {
		return createError("ERROR: first argument to `round` must be a float, got '%s'", args[0].GetType())
	}

	value := args[0].(*Float).Value
	precision := 0
	if len(args) == 2 {
		if args[1].GetType() != IntegerType {
			return createError("ERROR: second argument to `round` must be an integer, got '%s'", args[1].GetType())
		}
		precision = int(args[1].(*Integer).Value)
	}

	factor := math.Pow(10, float64(precision))
	return &Float{Value: math.Round(value*factor) / factor}
}

// sqrt returns the square root of a number.
//
// Syntax: sqrt(number)
//
// Usage:
//
//	Returns the square root of the given non-negative number (integer or float).
//	Returns an error if the input is negative.
//
// Example:
//
//	sqrt(16);   // Returns 4.0
//	sqrt(2.25); // Returns 1.5
func sqrt(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 1 {
		return createError("ERROR: wrong number of arguments. got=%d, want=1", len(args))
	}
	if args[0].GetType() != FloatType {
		// convert integer to float if needed
		if args[0].GetType() == IntegerType {
			args[0] = &Float{Value: float64(args[0].(*Integer).Value)}
		} else {
			return createError("argument to `sqrt` must be a `number`, got '%s'", args[0].GetType())
		}
	}

	value := args[0].(*Float).Value
	if value < 0 {
		return createError("ERROR: cannot compute square root of a negative number")
	}
	powVal := pow(rt, writer, &Float{Value: value}, &Float{Value: 0.5})
	return powVal
}

// pow returns the result of raising a number to a power.
//
// Syntax: pow(base, exponent)
//
// Usage:
//
//	Returns the base raised to the power of the exponent (base^exponent).
//	Supports both integer and floating-point arguments.
//
// Example:
//
//	pow(2, 3);   // Returns 8.0
//	pow(9, 0.5); // Returns 3.0
func pow(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 2 {
		return createError("ERROR: wrong number of arguments. got=%d, want=2", len(args))
	}
	if (args[0].GetType() != IntegerType && args[0].GetType() != FloatType) ||
		(args[1].GetType() != IntegerType && args[1].GetType() != FloatType) {
		return createError("ERROR: arguments to `pow` must be integers or floats, got '%s' and '%s'", args[0].GetType(), args[1].GetType())
	}

	var base, exponent float64
	if args[0].GetType() == IntegerType {
		base = float64(args[0].(*Integer).Value)
	} else {
		base = args[0].(*Float).Value
	}
	if args[1].GetType() == IntegerType {
		exponent = float64(args[1].(*Integer).Value)
	} else {
		exponent = args[1].(*Float).Value
	}

	return &Float{Value: math.Pow(base, exponent)}
}

// sin returns the sine of the radian argument.
//
// Syntax: sin(radians)
//
// Usage:
//
//	Returns the trigonometric sine of the angle specified in radians.
//
// Example:
//
//	sin(0); // Returns 0.0
func sin(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 1 {
		return createError("ERROR: wrong number of arguments. got=%d, want=1", len(args))
	}
	if args[0].GetType() != FloatType {
		if args[0].GetType() == IntegerType {
			args[0] = &Float{Value: float64(args[0].(*Integer).Value)}
		} else {
			return createError("argument to `sin` must be a `number`, got '%s'", args[0].GetType())
		}
	}
	return &Float{Value: math.Sin(args[0].(*Float).Value)}
}

// cos returns the cosine of the radian argument.
//
// Syntax: cos(radians)
//
// Usage:
//
//	Returns the trigonometric cosine of the angle specified in radians.
//
// Example:
//
//	cos(0); // Returns 1.0
func cos(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 1 {
		return createError("ERROR: wrong number of arguments. got=%d, want=1", len(args))
	}
	if args[0].GetType() != FloatType {
		if args[0].GetType() == IntegerType {
			args[0] = &Float{Value: float64(args[0].(*Integer).Value)}
		} else {
			return createError("argument to `cos` must be a `number`, got '%s'", args[0].GetType())
		}
	}
	return &Float{Value: math.Cos(args[0].(*Float).Value)}
}

// tan returns the tangent of the radian argument.
//
// Syntax: tan(radians)
//
// Usage:
//
//	Returns the trigonometric tangent of the angle specified in radians.
//
// Example:
//
//	tan(0); // Returns 0.0
func tan(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 1 {
		return createError("ERROR: wrong number of arguments. got=%d, want=1", len(args))
	}
	if args[0].GetType() != FloatType {
		if args[0].GetType() == IntegerType {
			args[0] = &Float{Value: float64(args[0].(*Integer).Value)}
		} else {
			return createError("argument to `tan` must be a `number`, got '%s'", args[0].GetType())
		}
	}
	return &Float{Value: math.Tan(args[0].(*Float).Value)}
}

// asin returns the arcsine, in radians.
//
// Syntax: asin(value)
//
// Usage:
//
//	Returns the inverse sine of the given value in radians.
//
// Example:
//
//	asin(0); // Returns 0.0
func asin(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 1 {
		return createError("ERROR: wrong number of arguments. got=%d, want=1", len(args))
	}
	if args[0].GetType() != FloatType {
		if args[0].GetType() == IntegerType {
			args[0] = &Float{Value: float64(args[0].(*Integer).Value)}
		} else {
			return createError("argument to `asin` must be a `number`, got '%s'", args[0].GetType())
		}
	}
	return &Float{Value: math.Asin(args[0].(*Float).Value)}
}

// acos returns the arccosine, in radians.
//
// Syntax: acos(value)
//
// Usage:
//
//	Returns the inverse cosine of the given value in radians.
//
// Example:
//
//	acos(1); // Returns 0.0
func acos(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 1 {
		return createError("ERROR: wrong number of arguments. got=%d, want=1", len(args))
	}
	if args[0].GetType() != FloatType {
		if args[0].GetType() == IntegerType {
			args[0] = &Float{Value: float64(args[0].(*Integer).Value)}
		} else {
			return createError("argument to `acos` must be a `number`, got '%s'", args[0].GetType())
		}
	}
	return &Float{Value: math.Acos(args[0].(*Float).Value)}
}

// atan returns the arctangent, in radians.
//
// Syntax: atan(value)
//
// Usage:
//
//	Returns the inverse tangent of the given value in radians.
//
// Example:
//
//	atan(0); // Returns 0.0
func atan(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 1 {
		return createError("ERROR: wrong number of arguments. got=%d, want=1", len(args))
	}
	if args[0].GetType() != FloatType {
		if args[0].GetType() == IntegerType {
			args[0] = &Float{Value: float64(args[0].(*Integer).Value)}
		} else {
			return createError("argument to `atan` must be a `number`, got '%s'", args[0].GetType())
		}
	}
	return &Float{Value: math.Atan(args[0].(*Float).Value)}
}

// atan2 returns the arctangent of y/x, in radians.
//
// Syntax
//
//	atan2(y, x)
//
// Usage:
//
//	Returns the angle in radians between the positive x-axis and the point (x, y).
//	This is useful for determining the angle of a vector from the origin to (x, y).
//
// Example:
//
//	atan2(1, 0); // Returns π/2 radians (90 degrees)
func atan2(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 2 {
		return createError("ERROR: wrong number of arguments. got=%d, want=2", len(args))
	}
	// Check y
	if args[0].GetType() != FloatType {
		if args[0].GetType() == IntegerType {
			args[0] = &Float{Value: float64(args[0].(*Integer).Value)}
		} else {
			return createError("first argument to `atan2` must be a `number`, got '%s'", args[0].GetType())
		}
	}
	// Check x
	if args[1].GetType() != FloatType {
		if args[1].GetType() == IntegerType {
			args[1] = &Float{Value: float64(args[1].(*Integer).Value)}
		} else {
			return createError("second argument to `atan2` must be a `number`, got '%s'", args[1].GetType())
		}
	}
	return &Float{Value: math.Atan2(args[0].(*Float).Value, args[1].(*Float).Value)}
}

// log returns the natural logarithm.
//
// Syntax: log(value)
//
// Usage:
//
//	Returns the natural logarithm of the given value.
//	The input must be a positive number (integer or float).
//
// Example:
//
//	log(1); // Returns 0.0
func log(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 1 {
		return createError("ERROR: wrong number of arguments. got=%d, want=1", len(args))
	}
	if args[0].GetType() != FloatType {
		if args[0].GetType() == IntegerType {
			args[0] = &Float{Value: float64(args[0].(*Integer).Value)}
		} else {
			return createError("argument to `log` must be a `number`, got '%s'", args[0].GetType())
		}
	}
	return &Float{Value: math.Log(args[0].(*Float).Value)}
}

// log10 returns the decimal logarithm.
// Syntax: log10(value)
//
// Usage:
//
//	Returns the base-10 logarithm of the given value.
//	The input must be a positive number (integer or float).
//
// Example:
//
//	log10(100); // Returns 2.0
func log10(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 1 {
		return createError("ERROR: wrong number of arguments. got=%d, want=1", len(args))
	}
	if args[0].GetType() != FloatType {
		if args[0].GetType() == IntegerType {
			args[0] = &Float{Value: float64(args[0].(*Integer).Value)}
		} else {
			return createError("argument to `log10` must be a `number`, got '%s'", args[0].GetType())
		}
	}
	return &Float{Value: math.Log10(args[0].(*Float).Value)}
}

// exp returns e**x.
// Syntax: exp(value)
//
// Usage:
//
//	Returns the exponential of the given value (e raised to the power of the input).
//	The input can be an integer or a float.
//
// Example:
//
//	exp(1); // Returns e (approximately 2.71828)
func exp(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 1 {
		return createError("ERROR: wrong number of arguments. got=%d, want=1", len(args))
	}
	if args[0].GetType() != FloatType {
		if args[0].GetType() == IntegerType {
			args[0] = &Float{Value: float64(args[0].(*Integer).Value)}
		} else {
			return createError("argument to `exp` must be a `number`, got '%s'", args[0].GetType())
		}
	}
	return &Float{Value: math.Exp(args[0].(*Float).Value)}
}

// randFunc returns a random floating-point number in [0.0, 1.0).
// Syntax: rand()
func randFunc(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 0 {
		return createError("ERROR: rand expects 0 arguments")
	}
	return &Float{Value: rand.Float64()}
}

// randInt returns a random integer between min and max (inclusive).
// Syntax: rand_int(min, max)
func randInt(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 2 {
		return createError("ERROR: rand_int expects 2 arguments (min, max)")
	}
	if args[0].GetType() != IntegerType || args[1].GetType() != IntegerType {
		return createError("ERROR: arguments to `rand_int` must be integers")
	}

	min := args[0].(*Integer).Value
	max := args[1].(*Integer).Value

	if min > max {
		return createError("ERROR: min cannot be greater than max in `rand_int`")
	}

	return &Integer{Value: min + rand.Int63n(max-min+1)}
}
