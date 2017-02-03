package expression

import (
	"fmt"
	"github.com/mikeyhu/mekkanism/inbuilt"
	"github.com/mikeyhu/mekkanism/interfaces"
)

type Expression struct {
	FunctionName  string
	File          string
	StartPosition string
	Arguments     []interfaces.Argument
}

func (exp Expression) IsArg() {}

func (exp Expression) Evaluate() interfaces.Argument {
	if exp.FunctionName == "+" {
		return inbuilt.PlusAll(exp.Arguments)
	} else if exp.FunctionName == "-" {
		return inbuilt.MinusAll(exp.Arguments)
	} else {
		panic(fmt.Sprintf("Panic - Cannot resolve FunctionName '%s'", exp.FunctionName))
	}
}
