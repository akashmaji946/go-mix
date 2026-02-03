package scope

import "github.com/akashmaji946/go-mix/objects"

// Scope defines a boundary of variables' lifetime and accessability
type Scope struct {
	// The variables bound to this Scope instance
	Variables map[string]objects.GoMixObject
	// The parent Scope if we are inside a function
	// if this is nil, this is the global Scope instance.
	Parent *Scope
}

// Creates a new Scope with the given parent.
func NewScope(parent *Scope) *Scope {
	return &Scope{
		Variables: make(map[string]objects.GoMixObject),
		Parent:    parent,
	}
}

// LookUp: Looks up the object bound to the varName
// The lookup should explore the
// parent(s) Scope as well ans should return a tuple (obj, true)
func (s *Scope) LookUp(varName string) (objects.GoMixObject, bool) {
	obj, ok := s.Variables[varName]
	if !ok && s.Parent != nil {
		obj, ok = s.Parent.LookUp(varName)
	}
	return obj, ok
}

// Bind: Binds the object
func (s *Scope) Bind(varName string, obj objects.GoMixObject) bool {
	_, has := s.LookUp(varName)
	s.Variables[varName] = obj
	return has
}
