package eval

import "fmt"

type GoMixType string

const (
	IntegerType GoMixType = "int"
	FloatType   GoMixType = "float"
	StringType  GoMixType = "string"
	BooleanType GoMixType = "bool"
	NilType     GoMixType = "nil"
)

type GoMixObject interface {
	GetType() GoMixType
	ToString() string
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
