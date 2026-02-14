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
	{Name: "getpid", Callback: getpid},     // Returns the process ID
	{Name: "hostname", Callback: hostname}, // Returns the system hostname
	{Name: "user", Callback: userFunc},     // Returns the current username
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

// getpid returns the process ID of the current process.
//
// Syntax: getpid()
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
