package scope

import "github.com/akashmaji946/go-mix/objects"

// Scope defines a boundary of variables' lifetime and accessability
type Scope struct {
	// The variables bound to this Scope instance
	Variables map[string]objects.GoMixObject
	// The constants bound to this Scope instance
	Consts map[string]bool
	// The parent Scope if we are inside a function
	// if this is nil, this is the global Scope instance.
	Parent *Scope
}

// Creates a new Scope with the given parent.
func NewScope(parent *Scope) *Scope {
	return &Scope{
		Variables: make(map[string]objects.GoMixObject),
		Consts:    make(map[string]bool),
		Parent:    parent,
	}
}

// LookUp: Looks up the object bound to the varName
// The lookup should explore the
// parent(s) Scope as well ans should return a tuple (obj, true)
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

// Bind: Binds the object
// Returns true if the variable already exists in the CURRENT scope (not parent scopes)
func (s *Scope) Bind(varName string, obj objects.GoMixObject) (string, bool) {
	if s.Variables == nil {
		s.Variables = make(map[string]objects.GoMixObject)
	}
	_, has := s.Variables[varName]
	s.Variables[varName] = obj
	return varName, has
}

// Assign: Assigns to an existing variable (does not create new binding)
// Returns the scope where the variable was found, and whether it was found
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

// IsConstant: Checks if a variable is constant in this scope or any parent
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

// Copy: Creates a shallow copy of this scope
// This is used for closures to capture the scope at function definition time
func (s *Scope) Copy() *Scope {
	newScope := &Scope{
		Variables: make(map[string]objects.GoMixObject),
		Consts:    make(map[string]bool),
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
	return newScope
}
