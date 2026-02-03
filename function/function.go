package function

import (
	"fmt"

	"github.com/akashmaji946/go-mix/objects"
	"github.com/akashmaji946/go-mix/parser"
	"github.com/akashmaji946/go-mix/scope"
)

// Function represents a function object
type Function struct {
	Name   string
	Params []*parser.IdentifierExpressionNode
	Body   *parser.BlockStatementNode
	Scp    *scope.Scope
}

func (f *Function) GetType() objects.GoMixType {
	return "func"
}

func (f *Function) ToString() string {
	return fmt.Sprintf("func(%s)", f.Name)
}

func (f *Function) ToObject() string {
	args := ""
	for i, param := range f.Params {
		if i > 0 {
			args += ", "
		}
		args += param.Name
	}
	return fmt.Sprintf("<func[%s(%s)]>", f.Name, args)
}
