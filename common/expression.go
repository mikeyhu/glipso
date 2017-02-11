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
	var result interfaces.Argument
	if fi, ok := inbuilt[exp.FunctionName]; ok {
		if fi.evaluateArgs {
			exp.evaluateArguments()
		}
		result = fi.function(exp.Arguments)
	} else {
		panic(fmt.Sprintf("Panic - Cannot resolve FunctionName '%s'", exp.FunctionName))
	}
	exp.printExpression(result)
	if exp, ok := result.(Expression); ok {
		result = exp.Evaluate()
	}
	return result
}

func (exp Expression) evaluateArguments() {
	for p, arg := range exp.Arguments {
		switch t := arg.(type) {
		case Expression:
			exp.Arguments[p] = t.Evaluate()
		}
	}
}

func (exp Expression) printExpression(result interfaces.Argument) {
	fmt.Printf("%v = %v\n", exp, result)
}
