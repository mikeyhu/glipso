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

func (exp EXP) Evaluate(sco interfaces.Scope) interfaces.Argument {
	var result interfaces.Argument
	if fi, ok := inbuilt[exp.FunctionName]; ok {
		if fi.evaluateArgs {
			exp.evaluateArguments(sco)
		}
		result = fi.function(exp.Arguments, sco)
	} else {
		panic(fmt.Sprintf("Panic - Cannot resolve FunctionName '%s'", exp.FunctionName))
	}
	exp.printExpression(result)
	if exp, ok := result.(EXP); ok {
		result = exp.Evaluate(sco.NewChildScope())
	}
	return result
}

func (exp EXP) evaluateArguments(sco interfaces.Scope) {
	for p, arg := range exp.Arguments {
		if e, ok := arg.(interfaces.Evaluatable); ok {
			exp.Arguments[p] = e.Evaluate(sco.NewChildScope())
		}
	}
}

func (exp EXP) printExpression(result interfaces.Argument) {
	if DEBUG {
		fmt.Printf("%v = %v\n", exp, result)
	}
}

type EXPN struct {
	Function  FN
	Arguments []interfaces.Argument
}

func (e EXPN) IsArg() {

}
func (e EXPN) Evaluate(sco interfaces.Scope) interfaces.Argument {
	env := sco.NewChildScope()
	if len(e.Function.Arguments.Vector) != len(e.Arguments) {
		panic("Invalid number of arguments")
	}
	for i, v := range e.Function.Arguments.Vector {
		env.CreateRef(v.(REF), e.Arguments[i])
	}
	return e.Function.Expression.Evaluate(env)
}
