package common

import (
	"fmt"
	"github.com/mikeyhu/glipso/interfaces"
)

var DEBUG = false

type EXP struct {
	FunctionName  string
	File          string
	StartPosition string
	Arguments     []interfaces.Argument
}

func (exp EXP) IsArg() {}

func (exp EXP) String() string {
	return fmt.Sprintf("(%s %v)", exp.FunctionName, exp.Arguments)
}

func (exp EXP) Evaluate() interfaces.Argument {
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
	if exp, ok := result.(EXP); ok {
		result = exp.Evaluate()
	}
	return result
}

func (exp EXP) evaluateArguments() {
	for p, arg := range exp.Arguments {
		if e, ok := arg.(interfaces.Evaluatable); ok {
			exp.Arguments[p] = e.Evaluate()
		}
	}
}

func (exp EXP) printExpression(result interfaces.Argument) {
	if DEBUG {
		fmt.Printf("%v = %v\n", exp, result)
	}
}
