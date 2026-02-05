/*
File    : go-mix/eval/evaluator.go
Author  : Akash Maji
Contact : akashmaji(@iisc.ac.in)
*/
package eval

import (
	"fmt"
	"io"
	"os"

	"github.com/akashmaji946/go-mix/function"
	"github.com/akashmaji946/go-mix/objects"
	"github.com/akashmaji946/go-mix/parser"
	"github.com/akashmaji946/go-mix/scope"
)

// Evaluator holds the state for evaluating GoMix AST nodes,
// including parser, scope, builtins, and output writer.
// It serves as the main execution engine for the GoMix interpreter,
// managing the evaluation context and providing access to built-in functions.
type Evaluator struct {
	Par      *parser.Parser              // Parser instance for error reporting with line/column information
	Scp      *scope.Scope                // Current scope for variable bindings and lexical scoping
	Builtins map[string]*objects.Builtin // Map of builtin functions (e.g., print, len, push, pop)
	Writer   io.Writer                   // Output writer for builtin functions (default: os.Stdout)
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
//   - *Evaluator: A fully initialized evaluator ready to execute GoMix code
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
		Builtins: make(map[string]*objects.Builtin),
		Writer:   os.Stdout, // Default to stdout
	}
	for _, builtin := range objects.Builtins {
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
func (e *Evaluator) RegisterFunction(n *parser.FunctionStatementNode) objects.GoMixObject {
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
func (e *Evaluator) InvokeBuiltin(name string, args ...objects.GoMixObject) objects.GoMixObject {

	if builtin, ok := e.Builtins[name]; ok {
		return builtin.Callback(e.Writer, args...)
	}
	return &objects.Nil{}
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
func (e *Evaluator) CreateError(format string, a ...interface{}) *objects.Error {
	msg := fmt.Sprintf(format, a...)
	fullMsg := fmt.Sprintf("[%d:%d] %s", e.Par.Lex.Line, e.Par.Lex.Column, msg)
	return &objects.Error{Message: fullMsg}
}
