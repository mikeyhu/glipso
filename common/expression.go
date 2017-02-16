package common

import (
	"fmt"
	"github.com/mikeyhu/glipso/interfaces"
)

// DEBUG enable to display debug information for function evaluation
var DEBUG = false

type EXP struct {
	FunctionName string
	Arguments    []interfaces.Type
}

func (exp EXP) IsType() {}

func (exp EXP) String() string {
	return fmt.Sprintf("(%s %v)", exp.FunctionName, exp.Arguments)
}

func (exp EXP) Evaluate(sco interfaces.Scope) interfaces.Type {
	var result interfaces.Type
	if fi, ok := inbuilt[exp.FunctionName]; ok {
		if fi.evaluateArgs {
			exp.evaluateArguments(sco)
		}
		result = fi.function(exp.Arguments, sco)
	} else {
		function := sco.ResolveRef(REF(exp.FunctionName))
		if function, ok := function.(FN); ok {
			expn := EXPN{function, exp.Arguments}
			result = expn.Evaluate(sco)
		} else {
			panic(fmt.Sprintf("Panic - Cannot resolve FunctionName '%s'", exp.FunctionName))
		}
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

func (exp EXP) printExpression(result interfaces.Type) {
	if DEBUG {
		fmt.Printf("%v = %v\n", exp, result)
	}
}

type EXPN struct {
	Function  FN
	Arguments []interfaces.Type
}

func (e EXPN) IsArg() {

}
func (e EXPN) Evaluate(sco interfaces.Scope) interfaces.Type {
	env := sco.NewChildScope()
	if len(e.Function.Arguments.Vector) != len(e.Arguments) {
		panic("Invalid number of arguments")
	}
	for i, v := range e.Function.Arguments.Vector {
		env.CreateRef(v.(REF), e.Arguments[i])
	}
	return e.Function.Expression.Evaluate(env)
}
