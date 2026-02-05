/*
File    : go-mix/scope/scope.go
Author  : Akash Maji
Contact : akashmaji(@iisc.ac.in)
*/
package scope

import "github.com/akashmaji946/go-mix/objects"

// Scope defines a lexical scope boundary for variable lifetime and accessibility.
//
// Scope implements a hierarchical scope chain that enables lexical scoping and closures.
// Each scope maintains its own variable bindings and can access variables from parent scopes.
// This structure supports:
// - Variable shadowing: inner scopes can redefine variables from outer scopes
// - Closures: functions capture their defining scope and can access outer variables
// - Block scoping: each block (function, loop, etc.) can have its own scope
//
// The scope chain is traversed upward (from child to parent) during variable lookup,
// implementing the standard lexical scoping rules found in most programming languages.
type Scope struct {
	// Variables maps variable names to their current values in this scope
	Variables map[string]objects.GoMixObject

	// Consts tracks which variables are declared as constants (immutable)
	Consts map[string]bool

	// LetVars tracks which variables are declared with 'let' (type-safe)
	LetVars map[string]bool

	// LetTypes stores the declared types of 'let' variables for type checking
	LetTypes map[string]objects.GoMixType

	// Parent points to the enclosing scope, forming a scope chain
	// nil indicates this is the global (root) scope
	Parent *Scope
}

// NewScope creates and initializes a new Scope with the specified parent scope.
//
// This constructor initializes all internal maps and establishes the parent-child
// relationship in the scope chain. The parent parameter determines the scope's
// position in the hierarchy:
// - parent == nil: Creates a global (root) scope with no parent
// - parent != nil: Creates a nested scope that can access parent variables
//
// Each new scope starts with empty variable bindings but inherits access to
// all variables in parent scopes through the lookup chain.
//
// Parameters:
//   - parent: The enclosing scope, or nil for a global scope
//
// Returns:
//   - *Scope: A fully initialized scope ready for variable bindings
//
// Example usage:
//
//	globalScope := NewScope(nil)              // Create global scope
//	functionScope := NewScope(globalScope)    // Create function scope
//	blockScope := NewScope(functionScope)     // Create nested block scope
func NewScope(parent *Scope) *Scope {
	return &Scope{
		Variables: make(map[string]objects.GoMixObject),
		Consts:    make(map[string]bool),
		LetVars:   make(map[string]bool),
		LetTypes:  make(map[string]objects.GoMixType),
		Parent:    parent,
	}
}

// LookUp searches for a variable by name in this scope and all parent scopes.
//
// This method implements the core variable resolution algorithm for lexical scoping:
// 1. First checks the current scope's Variables map
// 2. If not found and a parent scope exists, recursively searches the parent
// 3. Continues up the scope chain until the variable is found or the root is reached
//
// This traversal order ensures that:
// - Variables in inner scopes shadow those in outer scopes
// - All variables in the scope chain are accessible
// - The most recent binding is always returned
//
// The method is safe to call even if Variables map is nil (lazy initialization).
//
// Parameters:
//   - varName: The name of the variable to look up
//
// Returns:
//   - objects.GoMixObject: The value bound to the variable (if found)
//   - bool: true if the variable was found in this scope or any parent, false otherwise
//
// Example:
//
//	var x = 10;           // Bound in global scope
//	func foo() {
//	    var y = 20;       // Bound in function scope
//	    return x + y;     // LookUp finds both x (in parent) and y (in current)
//	}
func (s *Scope) LookUp(varName string) (objects.GoMixObject, bool) {
	if s.Variables == nil {
		s.Variables = make(map[string]objects.GoMixObject)
	}
	obj, ok := s.Variables[varName]
	if !ok && s.Parent != nil {
		obj, ok = s.Parent.LookUp(varName)
	}
	return obj, ok
}

// Bind creates a new variable binding in the current scope.
//
// This method adds or updates a variable binding in the current scope only,
// without affecting parent scopes. It performs the following operations:
// 1. Initializes the Variables map if needed (lazy initialization)
// 2. Checks if the variable already exists in the CURRENT scope
// 3. Creates or updates the binding with the provided value
//
// Important behaviors:
// - Only checks the current scope, not parent scopes (use LookUp for that)
// - Returns true if the variable already existed in THIS scope (redeclaration)
// - Does not prevent shadowing variables from parent scopes
// - Used for variable declarations (var, const, let)
//
// Parameters:
//   - varName: The name of the variable to bind
//   - obj: The value to bind to the variable
//
// Returns:
//   - string: The variable name (echoed back)
//   - bool: true if the variable already existed in the current scope, false if new
//
// Example:
//
//	scope.Bind("x", &objects.Integer{Value: 10})  // New binding, returns ("x", false)
//	scope.Bind("x", &objects.Integer{Value: 20})  // Redeclaration, returns ("x", true)
func (s *Scope) Bind(varName string, obj objects.GoMixObject) (string, bool) {
	if s.Variables == nil {
		s.Variables = make(map[string]objects.GoMixObject)
	}
	_, has := s.Variables[varName]
	s.Variables[varName] = obj
	return varName, has
}

// Assign updates an existing variable in the scope where it was originally defined.
//
// This method is crucial for proper closure behavior and variable assignment.
// Unlike Bind (which creates new bindings in the current scope), Assign:
// 1. Searches for the variable in the current scope
// 2. If found, updates it in place and returns this scope
// 3. If not found, recursively searches parent scopes
// 4. Updates the variable in the scope where it was originally defined
//
// This ensures that:
// - Closures can modify variables from their captured scope
// - Assignments affect the original binding, not create new ones
// - Inner scopes can modify outer scope variables
//
// The method is safe to call even if Variables map is nil (lazy initialization).
//
// Parameters:
//   - varName: The name of the variable to assign to
//   - obj: The new value to assign
//
// Returns:
//   - *Scope: The scope where the variable was found and updated (nil if not found)
//   - bool: true if the variable was found and updated, false otherwise
//
// Example:
//
//	var x = 10;           // Bound in outer scope
//	func increment() {
//	    x = x + 1;        // Assign finds and updates x in outer scope
//	}
func (s *Scope) Assign(varName string, obj objects.GoMixObject) (*Scope, bool) {
	if s.Variables == nil {
		s.Variables = make(map[string]objects.GoMixObject)
	}
	if _, ok := s.Variables[varName]; ok {
		s.Variables[varName] = obj
		return s, true
	}
	if s.Parent != nil {
		return s.Parent.Assign(varName, obj)
	}
	return nil, false
}

// IsConstant checks if a variable is declared as a constant in this scope or any parent.
//
// This method traverses the scope chain to determine if a variable was declared
// with the 'const' keyword, which makes it immutable. The search follows the
// same pattern as LookUp:
// 1. Checks the current scope's Consts map
// 2. If not found and a parent exists, recursively checks the parent
// 3. Returns true if found anywhere in the chain
//
// This is used during assignment operations to prevent modification of constants.
// The method is safe to call even if Consts map is nil (lazy initialization).
//
// Parameters:
//   - varName: The name of the variable to check
//
// Returns:
//   - bool: true if the variable is a constant in this scope or any parent, false otherwise
//
// Example:
//
//	const PI = 3.14;
//	PI = 3.15;        // Error: IsConstant("PI") returns true, assignment rejected
func (s *Scope) IsConstant(varName string) bool {
	if s.Consts == nil {
		s.Consts = make(map[string]bool)
	}
	if _, ok := s.Consts[varName]; ok {
		return true
	}
	if s.Parent != nil {
		return s.Parent.IsConstant(varName)
	}
	return false
}

// IsLetVariable checks if a variable is declared with 'let' in this scope or any parent.
//
// This method traverses the scope chain to determine if a variable was declared
// with the 'let' keyword, which enforces type safety. The search follows the
// same pattern as LookUp:
// 1. Checks the current scope's LetVars map
// 2. If not found and a parent exists, recursively checks the parent
// 3. Returns true if found anywhere in the chain
//
// This is used during assignment operations to enforce type checking for 'let' variables.
// The method is safe to call even if LetVars map is nil (lazy initialization).
//
// Parameters:
//   - varName: The name of the variable to check
//
// Returns:
//   - bool: true if the variable is a 'let' variable in this scope or any parent, false otherwise
//
// Example:
//
//	let x = 10;       // Integer type
//	x = 20;           // OK: same type
//	x = "hello";      // Error: IsLetVariable("x") returns true, type mismatch
func (s *Scope) IsLetVariable(varName string) bool {
	if s.LetVars == nil {
		s.LetVars = make(map[string]bool)
	}
	if _, ok := s.LetVars[varName]; ok {
		return true
	}
	if s.Parent != nil {
		return s.Parent.IsLetVariable(varName)
	}
	return false
}

// GetLetType retrieves the declared type of a 'let' variable from this scope or any parent.
//
// This method traverses the scope chain to find the type that was recorded when
// a 'let' variable was declared. The search follows the same pattern as LookUp:
// 1. Checks the current scope's LetTypes map
// 2. If not found and a parent exists, recursively checks the parent
// 3. Returns the type and true if found anywhere in the chain
//
// This is used during assignment operations to verify type compatibility for 'let' variables.
// The method is safe to call even if LetTypes map is nil (lazy initialization).
//
// Parameters:
//   - varName: The name of the variable whose type to retrieve
//
// Returns:
//   - objects.GoMixType: The declared type of the variable (if found)
//   - bool: true if the variable's type was found, false otherwise
//
// Example:
//
//	let x = 10;                    // Records type as IntegerType
//	typ, ok := scope.GetLetType("x")  // Returns (IntegerType, true)
//	x = 20;                        // Type check: new value must be IntegerType
func (s *Scope) GetLetType(varName string) (objects.GoMixType, bool) {
	if s.LetTypes == nil {
		s.LetTypes = make(map[string]objects.GoMixType)
	}
	if typ, ok := s.LetTypes[varName]; ok {
		return typ, true
	}
	if s.Parent != nil {
		return s.Parent.GetLetType(varName)
	}
	return "", false
}

// Copy creates a shallow copy of this scope for closure capture.
//
// This method is essential for implementing closures correctly. When a function
// is defined, it needs to capture the current state of its enclosing scope.
// Copy creates a new scope that:
// 1. Has the same parent as the original (maintains scope chain)
// 2. Contains copies of all variable bindings from the original
// 3. Contains copies of all const, let, and type information
// 4. Is independent of the original (modifications don't affect each other)
//
// The copy is "shallow" in that:
// - The scope structure itself is copied
// - The maps are new instances
// - The map entries (variable bindings) reference the same objects
// - The parent pointer is shared (not copied recursively)
//
// This allows closures to:
// - Capture variables from their defining scope
// - Maintain access even after the original scope is gone
// - Have their own independent variable bindings
//
// Returns:
//   - *Scope: A new scope with copied bindings and the same parent
//
// Example:
//
//	func makeCounter() {
//	    var count = 0;
//	    func increment() {
//	        count = count + 1;  // Accesses captured 'count'
//	        return count;
//	    }
//	    return increment;  // Function captures scope.Copy() of outer scope
//	}
func (s *Scope) Copy() *Scope {
	newScope := &Scope{
		Variables: make(map[string]objects.GoMixObject),
		Consts:    make(map[string]bool),
		LetVars:   make(map[string]bool),
		LetTypes:  make(map[string]objects.GoMixType),
		Parent:    s.Parent,
	}
	// Copy variables
	for k, v := range s.Variables {
		newScope.Variables[k] = v
	}
	// Copy constants
	for k, v := range s.Consts {
		newScope.Consts[k] = v
	}
	// Copy let variables
	for k, v := range s.LetVars {
		newScope.LetVars[k] = v
	}
	// Copy let types
	for k, v := range s.LetTypes {
		newScope.LetTypes[k] = v
	}
	return newScope
}
