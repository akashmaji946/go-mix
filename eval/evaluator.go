/*
File    : go-mix/eval/evaluator.go
Author  : Akash Maji
Contact : akashmaji(@iisc.ac.in)
*/
package eval

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/akashmaji946/go-mix/function"
	"github.com/akashmaji946/go-mix/parser"
	"github.com/akashmaji946/go-mix/scope"
	"github.com/akashmaji946/go-mix/std"
)

// Evaluator holds the state for evaluating Go-Mix AST nodes,
// including parser, scope, builtins, and output writer.
// It serves as the main execution engine for the Go-Mix interpreter,
// managing the evaluation context and providing access to built-in functions.
type Evaluator struct {
	Par      *parser.Parser              // Parser instance for error reporting with line/column information
	Scp      *scope.Scope                // Current scope for variable bindings and lexical scoping
	Builtins map[string]*std.Builtin     // Map of builtin functions (e.g., print, len, push, pop)
	Types    map[string]*std.GoMixStruct // Map of user-defined struct types (name to struct definition)
	Writer   io.Writer                   // Output writer for builtin functions (default: os.Stdout)
	Reader   *bufio.Reader               // Input reader for builtin functions (default: os.Stdin)
}

// NewEvaluator creates and initializes a new Evaluator instance with default configuration.
//
// This constructor performs the following initialization:
// - Creates a new root scope with no parent (global scope)
// - Initializes an empty builtin functions map
// - Sets the output writer to os.Stdout for default console output
// - Registers all available builtin functions from the objects package
//
// Returns:
//   - *Evaluator: A fully initialized evaluator ready to execute Go-Mix code
//
// Example usage:
//
//	ev := NewEvaluator()
//	ev.SetParser(parser)
//	result := ev.Eval(astNode)
func NewEvaluator() *Evaluator {
	ev := &Evaluator{
		Par:      nil,
		Scp:      scope.NewScope(nil),
		Builtins: make(map[string]*std.Builtin),
		Types:    make(map[string]*std.GoMixStruct),
		Writer:   os.Stdout, // Default to stdout
		Reader:   bufio.NewReader(os.Stdin),
	}
	for _, builtin := range std.Builtins {
		ev.Builtins[builtin.Name] = builtin
	}
	return ev
}

// SetWriter configures the output destination for the evaluator's builtin functions.
//
// This method allows redirecting output from builtin functions (like print, println)
// to any io.Writer implementation. This is particularly useful for:
// - Testing: capturing output to verify program behavior
// - Logging: redirecting output to log files
// - Custom output handling: sending output to network streams, buffers, etc.
//
// Parameters:
//   - w: An io.Writer implementation that will receive output from builtin functions
//
// Example usage:
//
//	var buf bytes.Buffer
//	ev.SetWriter(&buf)  // Redirect output to buffer for testing
func (e *Evaluator) SetWriter(w io.Writer) {
	e.Writer = w
}

// SetReader configures the input source for the evaluator's builtin functions.
func (e *Evaluator) SetReader(r io.Reader) {
	e.Reader = bufio.NewReader(r)
}

// GetInputReader returns the buffered input reader.
// This implements the std.Runtime interface.
func (e *Evaluator) GetInputReader() *bufio.Reader {
	return e.Reader
}

// SetParser assigns a parser instance to the evaluator for enhanced error reporting.
//
// The parser reference is used by CreateError() to include source code position
// information (line and column numbers) in error messages, making debugging easier.
// This should be called before evaluating any code to ensure accurate error locations.
//
// Parameters:
//   - p: A pointer to the Parser instance that parsed the AST being evaluated
//
// Example usage:
//
//	parser := parser.NewParser(lexer)
//	ev.SetParser(parser)
//	result := ev.Eval(parser.Parse())
func (e *Evaluator) SetParser(p *parser.Parser) {
	e.Par = p
}

// RegisterFunction creates and registers a user-defined function in the current scope.
//
// This method processes function declarations by:
// 1. Creating a Function object with the function's name, parameters, body, and captured scope
// 2. Checking for redeclaration conflicts in the current scope
// 3. Binding the function to its name in the current scope for later invocation
//
// The function captures the current scope (closure), allowing it to access variables
// from its defining scope even when called from different contexts.
//
// Parameters:
//   - n: A FunctionStatementNode containing the function's AST representation
//
// Returns:
//   - objects.GoMixObject: The created Function object on success, or an Error object
//     if the function name is already declared in the current scope
//
// Example:
//
//	func add(a, b) { return a + b; }  // Creates and registers 'add' function
func (e *Evaluator) RegisterFunction(n *parser.FunctionStatementNode) std.GoMixObject {
	function := &function.Function{
		Name:   n.FuncName.Name,
		Params: n.FuncParams,
		Body:   &n.FuncBody,
		Scp:    e.Scp, // Reference the current scope directly, not a copy
	}
	// redeclared?
	name, has := e.Scp.Bind(n.FuncName.Name, function)
	if has && name != "" {
		return e.CreateError("ERROR: function redeclaration found: (%s)", n.FuncName.Name)
	}
	e.Scp.Bind(n.FuncName.Name, function)
	return function
}

// IsBuiltin checks if a given identifier name corresponds to a registered builtin function.
//
// This method is used during function call evaluation to determine whether to
// invoke a builtin function or look up a user-defined function in the scope chain.
// Builtin functions are registered during evaluator initialization and include
// functions like print, println, len, push, pop, shift, etc.
//
// Parameters:
//   - name: The identifier name to check (e.g., "print", "len", "push")
//
// Returns:
//   - bool: true if the name matches a registered builtin function, false otherwise
//
// Example usage:
//
//	if e.IsBuiltin("print") {
//	    // Handle builtin function call
//	}
func (e *Evaluator) IsBuiltin(name string) bool {
	_, ok := e.Builtins[name]
	return ok
}

// InvokeBuiltin executes a builtin function by name with the provided arguments.
//
// This method looks up the builtin function in the Builtins map and invokes its
// callback function with the evaluator's writer and the provided arguments.
// Builtin functions handle their own argument validation and type checking.
//
// Parameters:
//   - name: The name of the builtin function to invoke (e.g., "print", "len")
//   - args: Variable number of GoMixObject arguments to pass to the builtin function
//
// Returns:
//   - objects.GoMixObject: The result returned by the builtin function's callback,
//     or a Nil object if the builtin function is not found in the registry
//
// Example usage:
//
//	result := e.InvokeBuiltin("len", arrayObject)  // Returns Integer with array length
//	e.InvokeBuiltin("print", stringObject)         // Prints to writer, returns Nil
func (e *Evaluator) InvokeBuiltin(name string, args ...std.GoMixObject) std.GoMixObject {

	if builtin, ok := e.Builtins[name]; ok {
		return builtin.Callback(e, e.Writer, args...)
	}
	return &std.Nil{}
}

// CallFunction executes a Go-Mix function object with the provided arguments.
// This implements the std.Runtime interface.
func (e *Evaluator) CallFunction(fn std.GoMixObject, args ...std.GoMixObject) std.GoMixObject {
	if fn.GetType() != std.FunctionType {
		return e.CreateError("ERROR: object is not a function")
	}
	functionObject := fn.(*function.Function)

	if len(args) != len(functionObject.Params) {
		return e.CreateError("ERROR: wrong number of arguments: expected %d, got %d", len(functionObject.Params), len(args))
	}

	callSiteScope := scope.NewScope(functionObject.Scp)
	for i, param := range functionObject.Params {
		callSiteScope.Bind(param.Name, args[i])
	}

	oldScope := e.Scp
	e.Scp = callSiteScope
	result := e.Eval(functionObject.Body)
	e.Scp = oldScope

	return UnwrapReturnValue(result)
}

// CreateError creates a new Error object with a formatted message including source position.
//
// This method constructs detailed error messages that include:
// - Line number: The line in the source code where the error occurred
// - Column number: The column position in that line
// - Error message: A formatted description of the error
//
// The parser must be set via SetParser() before calling this method to ensure
// accurate position information. The format string and arguments follow the
// same conventions as fmt.Sprintf().
//
// Parameters:
//   - format: A format string following fmt.Sprintf conventions
//   - a: Variable arguments to be formatted into the error message
//
// Returns:
//   - *objects.Error: An Error object with the formatted message including position info
//
// Example usage:
//
//	return e.CreateError("ERROR: identifier not found: (%s)", varName)
//	// Output: "[10:5] ERROR: identifier not found: (myVar)"
func (e *Evaluator) CreateError(format string, a ...interface{}) *std.Error {
	msg := fmt.Sprintf(format, a...)
	fullMsg := fmt.Sprintf("[%d:%d] %s", e.Par.Lex.Line, e.Par.Lex.Column, msg)
	return &std.Error{Message: fullMsg}
}

// NamedParameter represents a parameter passed to a function or method call.
//
// It encapsulates both the parameter name (from the function definition) and
// the evaluated value passed as an argument. This structure is primarily used
// during method invocation on objects to bind argument values to parameter names
// in the method's execution scope.
//
// Fields:
//   - Name: The name of the parameter as defined in the function signature.
//     Used for binding the value in the function's scope.
//   - Value: The evaluated runtime object passed as an argument.
type NamedParameter struct {
	Name  string          // The name of the parameter (e.g., "a", "b")
	Value std.GoMixObject //  The value of the parameter, which can be any GoMixObject (e.g., Integer, String, Array)
}

// IndexOfDot finds the index of the first period (.) character in a string.
//
// This helper function is used by the evaluator to detect method calls in
// identifier names (e.g., "obj.method"). It scans the string from left to right.
//
// Parameters:
//   - s: The string to search
//
// Returns:
//   - int: The index of the first dot, or -1 if no dot is found
func IndexOfDot(s string) int {
	for i, c := range s {
		if c == '.' {
			return i
		}
	}
	return -1
}
