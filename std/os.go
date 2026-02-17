/*
File    : go-mix/std/os.go
Author  : Akash Maji
Contact : akashmaji(@iisc.ac.in)
*/

// Package std - os.go
// This file defines the operating system interaction builtin functions for the Go-Mix language.
// It provides functions for environment variables and command execution.
package std

import (
	"io"
	"os"
	"os/exec"
	"os/user"
	"runtime"
	"time"
)

var osMethods = []*Builtin{
	{Name: "getenv", Callback: getenv},     // Gets an environment variable
	{Name: "setenv", Callback: setenv},     // Sets an environment variable
	{Name: "unsetenv", Callback: unsetenv}, // Unsets an environment variable
	{Name: "exec", Callback: execCmd},      // Executes a shell command
	{Name: "exit", Callback: exitFunc},     // Terminates the program
	{Name: "args", Callback: argsFunc},     // Returns command-line arguments
	{Name: "sleep", Callback: sleepFunc},   // Pauses execution for N milliseconds

	{Name: "getcwd", Callback: getcwd},     // Returns the current working directory
	{Name: "getpid", Callback: getpid},     // Returns the process ID
	{Name: "hostname", Callback: hostname}, // Returns the system hostname
	{Name: "user", Callback: userFunc},     // Returns the current username
	{Name: "platform", Callback: platform}, // Returns the OS name
	{Name: "arch", Callback: arch},         // Returns the system architecture

	// assertions
	{Name: "assert", Callback: assert},            // Asserts that a condition is true
	{Name: "assert_equal", Callback: assertEqual}, // Asserts that two values are equal
	{Name: "assert_true", Callback: assertTrue},   // Asserts that a condition is true
	{Name: "assert_false", Callback: assertFalse}, // Asserts that a condition is false

}

// init registers the OS methods as global builtins and as a package for import.
func init() {
	// Register as global builtins (for backward compatibility)
	Builtins = append(Builtins, osMethods...)

	// Register as a package (for import functionality)
	osPackage := &Package{
		Name:      "os",
		Functions: make(map[string]*Builtin),
	}
	for _, method := range osMethods {
		osPackage.Functions[method.Name] = method
	}
	RegisterPackage(osPackage)
}

// getenv retrieves the value of the environment variable named by the key.
// It returns the value, which will be empty if the variable is not present.
//
// Syntax: getenv(key)
//
// Example:
//
//	var path = getenv("PATH");
func getenv(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 1 {
		return createError("ERROR: getenv expects 1 argument (key)")
	}
	if args[0].GetType() != StringType {
		return createError("ERROR: argument to `getenv` must be a string, got '%s'", args[0].GetType())
	}

	key := args[0].ToString()
	val := os.Getenv(key)
	return &String{Value: val}
}

// setenv sets the value of the environment variable named by the key.
//
// Syntax: setenv(key, value)
func setenv(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 2 {
		return createError("ERROR: setenv expects 2 arguments (key, value)")
	}
	key := args[0].ToString()
	val := args[1].ToString()
	err := os.Setenv(key, val)
	if err != nil {
		return createError("ERROR: setenv failed: %v", err)
	}
	return &Nil{}
}

// unsetenv unsets a single environment variable.
//
// Syntax: unsetenv(key)
func unsetenv(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 1 {
		return createError("ERROR: unsetenv expects 1 argument (key)")
	}
	key := args[0].ToString()
	err := os.Unsetenv(key)
	if err != nil {
		return createError("ERROR: unsetenv failed: %v", err)
	}
	return &Nil{}
}

// execCmd executes the named program with the given arguments.
// It returns the combined standard output and standard error as a string.
//
// Syntax: exec(command, arg1, arg2, ...)
//
// Example:
//
//	var output = exec("ls", "-la");
func execCmd(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) == 0 {
		return createError("ERROR: exec expects at least 1 argument (command)")
	}

	cmdName := args[0].ToString()
	cmdArgs := make([]string, len(args)-1)
	for i, arg := range args[1:] {
		cmdArgs[i] = arg.ToString()
	}

	// Create the command
	cmd := exec.Command(cmdName, cmdArgs...)

	// Execute and capture combined output (stdout + stderr)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return createError("ERROR: command execution failed: %v\nOutput: %s", err, string(output))
	}

	return &String{Value: string(output)}
}

// exitFunc terminates the current process with an optional exit code.
//
// Syntax: exit([code])
// Default code is 0.
func exitFunc(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	code := 0
	if len(args) > 0 {
		if args[0].GetType() == IntegerType {
			code = int(args[0].(*Integer).Value)
		} else {
			return createError("ERROR: exit code must be an integer")
		}
	}
	os.Exit(code)
	return &Nil{}
}

// argsFunc returns an array of strings containing the command-line arguments.
//
// Syntax: args()
func argsFunc(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	elements := make([]GoMixObject, len(os.Args))
	for i, arg := range os.Args {
		elements[i] = &String{Value: arg}
	}
	return &Array{Elements: elements}
}

// sleepFunc pauses the current process for the specified number of milliseconds.
//
// Syntax: sleep(milliseconds)
func sleepFunc(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 1 {
		return createError("ERROR: sleep expects 1 argument (milliseconds)")
	}
	if args[0].GetType() != IntegerType {
		return createError("ERROR: argument to `sleep` must be an integer, got '%s'", args[0].GetType())
	}
	ms := args[0].(*Integer).Value
	time.Sleep(time.Duration(ms) * time.Millisecond)
	return &Nil{}
}

// getcwd returns the current working directory.
//
// Syntax: getcwd()
// Example:
//
//	var cwd = getcwd(); // say "/home/user/projects"
//
// This would set cwd to the current working directory of the process.
func getcwd(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 0 {
		return createError("ERROR: getcwd expects 0 arguments")
	}
	dir, err := os.Getwd()
	if err != nil {
		return createError("ERROR: could not get current working directory: %v", err)
	}
	return &String{Value: dir}
}

// getpid returns the process ID of the current process.
//
// Syntax: getpid()
// Example:
//
//	var pid = getpid(); // say 12345
//
// This would set pid to the process ID of the current process.
func getpid(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 0 {
		return createError("ERROR: getpid expects 0 arguments")
	}
	return &Integer{Value: int64(os.Getpid())}
}

// hostname returns the system hostname.
//
// Syntax: hostname()
func hostname(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 0 {
		return createError("ERROR: hostname expects 0 arguments")
	}
	name, err := os.Hostname()
	if err != nil {
		return createError("ERROR: could not get hostname: %v", err)
	}
	return &String{Value: name}
}

// userFunc returns the current username.
//
// Syntax: user()
func userFunc(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 0 {
		return createError("ERROR: user expects 0 arguments")
	}
	u, err := user.Current()
	if err != nil {
		return createError("ERROR: could not get current user: %v", err)
	}
	return &String{Value: u.Username}
}

// platform returns the operating system name.
//
// Syntax: platform()
func platform(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 0 {
		return createError("ERROR: platform expects 0 arguments")
	}
	return &String{Value: runtime.GOOS}
}

// arch returns the running program's architecture.
//
// Syntax: arch()
func arch(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 0 {
		return createError("ERROR: arch expects 0 arguments")
	}
	return &String{Value: runtime.GOARCH}
}

// assert checks if a condition is true and raises an error if it is not.
//
// Syntax: assert(condition, message)
//
// Example:
//
//	assert(1 == 1, "This will not raise an error"); // Passes
//	assert(1 == 2, "This will raise an error"); // Raises an error with the message "This will raise an error"
func assert(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 2 {
		return createError("ERROR: assert expects 2 arguments (condition, message)")
	}
	condition := toBool(rt, writer, args[0])
	if condition.GetType() != BooleanType {
		return createError("ERROR: condition argument to `assert` must be a boolean")
	}
	if !condition.(*Boolean).Value {
		message := args[1].ToString()
		return createError("Assertion failed: %s", message)
	}
	return &Nil{}
}

// assertEqual checks if two values are equal and raises an error if they are not.
//
// Syntax: assert_equal(value1, value2, message)
//
// Example:
//
//	assert_equal(1, 1, "Values are not equal"); // Passes
//	assert_equal(1, 2, "Values are not equal"); // Raises an error with the message "Values are not equal"
func assertEqual(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 3 {
		return createError("ERROR: assert_equal expects 3 arguments (value1, value2, message)")
	}
	value1 := args[0]
	value2 := args[1]
	message := args[2].ToString()

	if !isEqual(value1, value2) {
		printf(rt, writer, &String{Value: "[FAIL] Assertion failed: %s\n"}, &String{Value: message})
		exitFunc(rt, writer, &Integer{Value: 1})
	}
	printf(rt, writer, &String{Value: "[PASS] Assertion passed: %s\n"}, &String{Value: message})
	return &Nil{}
}

func isEqual(a, b GoMixObject) bool {
	if a.GetType() != b.GetType() {
		return false
	}
	switch a.GetType() {
	case IntegerType:
		return a.(*Integer).Value == b.(*Integer).Value
	case FloatType:
		return a.(*Float).Value == b.(*Float).Value
	case StringType:
		return a.(*String).Value == b.(*String).Value
	case BooleanType:
		return a.(*Boolean).Value == b.(*Boolean).Value
	case ArrayType:
		arrA := a.(*Array).Elements
		arrB := b.(*Array).Elements
		if len(arrA) != len(arrB) {
			return false
		}
		for i := range arrA {
			if !isEqual(arrA[i], arrB[i]) {
				return false
			}
		}
		return true
	case MapType:
		mapA := a.(*Map).Pairs
		mapB := b.(*Map).Pairs
		if len(mapA) != len(mapB) {
			return false
		}
		for key, valA := range mapA {
			valB, exists := mapB[key]
			if !exists || !isEqual(valA, valB) {
				return false
			}
		}
		return true
	case ListType:
		listA := a.(*List).Elements
		listB := b.(*List).Elements
		if len(listA) != len(listB) {
			return false
		}
		for i := range listA {
			if !isEqual(listA[i], listB[i]) {
				return false
			}
		}
		return true
	case TupleType:
		tupleA := a.(*Tuple).Elements
		tupleB := b.(*Tuple).Elements
		if len(tupleA) != len(tupleB) {
			return false
		}
		for i := range tupleA {
			if !isEqual(tupleA[i], tupleB[i]) {
				return false
			}
		}
		return true
	case CharType:
		return a.(*Char).Value == b.(*Char).Value
	case NilType:
		if b.GetType() != NilType {
			return false
		}
		return true
	case FunctionType:
		if b.GetType() != FunctionType {
			return false
		}
		fnA, okA := a.(FunctionInterface)
		fnB, okB := b.(FunctionInterface)
		if !okA || !okB {
			return false
		}
		return fnA.GetName() == fnB.GetName()
	default:
		return a == b // Fallback to pointer equality for unsupported types
	}
}

// assertTrue checks if a condition is true and raises an error if it is not.
//
// Syntax: assert_true(condition, message)
//
// Example:
//
//	assert_true(1 == 1, "Condition is not true"); // Passes
//	assert_true(1 == 2, "Condition is not true"); // Raises an error with the message "Condition is not true"
func assertTrue(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 2 {
		return createError("ERROR: assert_true expects 2 arguments (condition, message)")
	}
	condition := toBool(rt, writer, args[0])
	message := args[1].ToString()
	if condition.GetType() != BooleanType {
		return createError("ERROR: condition argument to `assert_true` must be a boolean")
	}
	if !condition.(*Boolean).Value {
		printf(rt, writer, &String{Value: "[FAIL] Assertion failed: %s\n"}, &String{Value: message})
		exitFunc(rt, writer, &Integer{Value: 1})
	}
	printf(rt, writer, &String{Value: "[PASS] Assertion passed: %s\n"}, &String{Value: message})
	return &Nil{}
}

// assertFalse checks if a condition is false and raises an error if it is not.
//
// Syntax: assert_false(condition, message)
//
// Example:
//
//	assert_false(1 == 2, "Condition is not false"); // Passes
//	assert_false(1 == 1, "Condition is not false"); // Raises an error with the message "Condition is not false"
func assertFalse(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 2 {
		return createError("ERROR: assert_false expects 2 arguments (condition, message)")
	}
	condition := toBool(rt, writer, args[0])
	message := args[1].ToString()
	if condition.GetType() != BooleanType {
		return createError("ERROR: condition argument to `assert_false` must be a boolean")
	}
	if condition.(*Boolean).Value {

		printf(rt, writer, &String{Value: "[FAIL] Assertion failed: %s\n"}, &String{Value: message})
		exitFunc(rt, writer, &Integer{Value: 1})
	}
	printf(rt, writer, &String{Value: "[PASS] Assertion passed: %s\n"}, &String{Value: message})
	return &Nil{}
}
