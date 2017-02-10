package common

import (
	"fmt"
	"github.com/mikeyhu/mekkanism/interfaces"
)

type Expression struct {
	FunctionName  string
	File          string
	StartPosition string
	Arguments     []interfaces.Argument
}

func (exp Expression) IsArg() {}

func (exp Expression) String() string {
	return fmt.Sprintf("(%s %v)", exp.FunctionName, exp.Arguments)
}

func (exp Expression) Evaluate() interfaces.Argument {
	exp.evaluateArguments()
	var result interfaces.Argument
	if f, ok := inbuilt[exp.FunctionName]; ok {
		result = f(exp.Arguments)
	} else {
		panic(fmt.Sprintf("Panic - Cannot resolve FunctionName '%s'", exp.FunctionName))
	}
	exp.printExpression(result)
	return result
}

func (exp Expression) evaluateArguments() {
	for p, arg := range exp.Arguments {
		switch t := arg.(type) {
		case Expression:
			exp.Arguments[p] = t.Evaluate().(I)
		}
	}
}

func (exp Expression) printExpression(result interfaces.Argument) {
	fmt.Printf("%v = %v\n", exp, result)
}
