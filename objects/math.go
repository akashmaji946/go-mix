package objects

import (
	"io"
	"math"
)

var mathMethods = []*Builtin{
	{
		Name:     "abs", // Returns the absolute value of a number
		Callback: abs,
	},
	{
		Name:     "fabs", // Returns the absolute value of a floating-point number (alias for abs)
		Callback: fabs,
	},
	{
		Name:     "min", // Returns the smaller of two numbers
		Callback: min,
	},
	{
		Name:     "max", // Returns the larger of two numbers
		Callback: max,
	},
	{
		Name:     "floor", // Returns the largest integer less than or equal to a number
		Callback: floor,
	},
	{
		Name:     "ceil", // Returns the smallest integer greater than or equal to a number
		Callback: ceil,
	},
	{
		Name:     "round", // Returns the nearest integer to a number
		Callback: round,
	},
	{
		Name:     "sqrt", // Returns the square root of a number
		Callback: sqrt,
	},
	{
		Name:     "pow", // Returns the result of raising a number to a power
		Callback: pow,
	},
	{
		Name:     "sin", // Returns the sine of the radian argument
		Callback: sin,
	},
	{
		Name:     "cos", // Returns the cosine of the radian argument
		Callback: cos,
	},
	{
		Name:     "tan", // Returns the tangent of the radian argument
		Callback: tan,
	},
	{
		Name:     "asin", // Returns the arcsine, in radians
		Callback: asin,
	},
	{
		Name:     "acos", // Returns the arccosine, in radians
		Callback: acos,
	},
	{
		Name:     "atan", // Returns the arctangent, in radians
		Callback: atan,
	},
	{
		Name:     "atan2", // Returns the arctangent of y/x, in radians
		Callback: atan2,
	},
	{
		Name:     "log", // Returns the natural logarithm
		Callback: log,
	},
	{
		Name:     "log10", // Returns the decimal logarithm
		Callback: log10,
	},
	{
		Name:     "exp", // Returns e**x
		Callback: exp,
	},
}

// init registers the math methods as global builtins by appending them to the Builtins slice.
func init() {
	Builtins = append(Builtins, mathMethods...)
}

// abs returns the absolute value of a number.
// It takes one argument: the number to evaluate.
// If the argument is not an integer, it returns an error object.
func abs(writer io.Writer, args ...GoMixObject) GoMixObject {
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
// It takes one argument: the number to evaluate.
// If the argument is not a float, it returns an error object.
func fabs(writer io.Writer, args ...GoMixObject) GoMixObject {
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
// It takes two arguments: the numbers to compare.
// If the arguments are not integers|floats, it returns an error object.
func min(writer io.Writer, args ...GoMixObject) GoMixObject {
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
// It takes two arguments: the numbers to compare.
// If the arguments are not integers|floats, it returns an error object.
func max(writer io.Writer, args ...GoMixObject) GoMixObject {
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
// It takes one argument: the number to evaluate.
// If the argument is not a float, it returns an error object.
func floor(writer io.Writer, args ...GoMixObject) GoMixObject {
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
// It takes one argument: the number to evaluate.
// If the argument is not a float, it returns an error object.
func ceil(writer io.Writer, args ...GoMixObject) GoMixObject {
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

// round returns the nearest integer to a number.
// It takes two argument: the number to evaluate, and an optional precision for rounding (default is 0).
// If the argument is not a float, it returns an error object.
func round(writer io.Writer, args ...GoMixObject) GoMixObject {
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
// It takes one argument: the number to evaluate.
// If the argument is not a float or is negative, it returns an error object.
func sqrt(writer io.Writer, args ...GoMixObject) GoMixObject {
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
	powVal := pow(writer, &Float{Value: value}, &Float{Value: 0.5})
	return powVal
}

// pow returns the result of raising a number to a power.
// It takes two arguments: the base and the exponent.
// If the arguments are not floats, it returns an error object.
func pow(writer io.Writer, args ...GoMixObject) GoMixObject {
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
func sin(writer io.Writer, args ...GoMixObject) GoMixObject {
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
func cos(writer io.Writer, args ...GoMixObject) GoMixObject {
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
func tan(writer io.Writer, args ...GoMixObject) GoMixObject {
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
func asin(writer io.Writer, args ...GoMixObject) GoMixObject {
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
func acos(writer io.Writer, args ...GoMixObject) GoMixObject {
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
func atan(writer io.Writer, args ...GoMixObject) GoMixObject {
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
func atan2(writer io.Writer, args ...GoMixObject) GoMixObject {
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
func log(writer io.Writer, args ...GoMixObject) GoMixObject {
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
func log10(writer io.Writer, args ...GoMixObject) GoMixObject {
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
func exp(writer io.Writer, args ...GoMixObject) GoMixObject {
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
