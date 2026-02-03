package objects

import (
	"fmt"
)

type GoMixType string

const (
	IntegerType GoMixType = "int"
	FloatType   GoMixType = "float"
	StringType  GoMixType = "string"
	BooleanType GoMixType = "bool"
	NilType     GoMixType = "nil"
	ErrorType   GoMixType = "error"

	FunctionType GoMixType = "func"
)

type GoMixObject interface {
	GetType() GoMixType
	ToString() string
	ToObject() string
}

// types
// Integer: int64
type Integer struct {
	Value int64
}

func (i *Integer) GetType() GoMixType {
	return IntegerType
}

func (i *Integer) ToString() string {
	return fmt.Sprintf("%d", i.Value)
}

func (i *Integer) ToObject() string {
	return fmt.Sprintf("<int(%d)>", i.Value)
}

// Float: float64
type Float struct {
	Value float64
}

func (f *Float) GetType() GoMixType {
	return FloatType
}

func (f *Float) ToString() string {
	return fmt.Sprintf("%f", f.Value)
}

func (f *Float) ToObject() string {
	return fmt.Sprintf("<float(%f)>", f.Value)
}

// String: string
type String struct {
	Value string
}

func (s *String) GetType() GoMixType {
	return StringType
}

func (s *String) ToString() string {
	return s.Value
}

func (s *String) ToObject() string {
	return fmt.Sprintf("<string(%s)>", s.Value)
}

// Boolean: bool
type Boolean struct {
	Value bool
}

func (b *Boolean) GetType() GoMixType {
	return BooleanType
}

func (b *Boolean) ToString() string {
	return fmt.Sprintf("%t", b.Value)
}

func (b *Boolean) ToObject() string {
	return fmt.Sprintf("<bool(%t)>", b.Value)
}

// Null: nil

type Nil struct {
	Value interface{}
}

func (n *Nil) GetType() GoMixType {
	return NilType
}

func (n *Nil) ToString() string {
	return "nil"
}

func (n *Nil) ToObject() string {
	return "<nil()>"
}

// Error: error

type Error struct {
	Message string
}

func (e *Error) GetType() GoMixType {
	return ErrorType
}

func (e *Error) ToString() string {
	return fmt.Sprintf("[ERROR]: %s", e.Message)
}

func (e *Error) ToObject() string {
	return fmt.Sprintf("<error(%s)>", e.Message)
}

// ReturnValue wraps a value returned from a function
type ReturnValue struct {
	Value GoMixObject
}

func (r *ReturnValue) GetType() GoMixType {
	return r.Value.GetType()
}

func (r *ReturnValue) ToString() string {
	return r.Value.ToString()
}

func (r *ReturnValue) ToObject() string {
	return r.Value.ToObject()
}
