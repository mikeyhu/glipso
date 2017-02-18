package common

import (
	"fmt"
	"github.com/mikeyhu/glipso/interfaces"
)

// DEBUG enable to display debug information for function evaluation
var DEBUG = false

type EXP struct {
	Function  interfaces.Type
	Arguments []interfaces.Type
}

func (exp EXP) IsType() {}

func (exp EXP) String() string {
	return fmt.Sprintf("(%v %v)", exp.Function, exp.Arguments)
}

func (exp EXP) Evaluate(sco interfaces.Scope) interfaces.Type {
	env := sco.NewChildScope()
	var result interfaces.Type
	function := exp.Function

	if toEval, ok := function.(EXP); ok {
		function = toEval.Evaluate(env)
	}

	if toREF, ok := function.(REF); ok {
		if fn, ok := env.ResolveRef(toREF); ok {
			function = fn
		} else if fi, ok := inbuilt[toREF.String()]; ok {
			result = exp.evaluateInbuilt(fi, env)
		}
	}

	if toFN, ok := function.(FN); ok {
		result = exp.evaluateFN(toFN, env)
	}

	exp.printExpression(result)
	if exp, ok := result.(EXP); ok {
		result = exp.Evaluate(env)
	}
	return result
}

func (exp EXP) evaluateFN(fn FN, env interfaces.Scope) interfaces.Type {
	if len(fn.Arguments.Vector) != len(exp.Arguments) {
		panic("Invalid number of arguments")
	}
	for i, v := range fn.Arguments.Vector {
		env.CreateRef(v.(REF), exp.Arguments[i])
	}
	return fn.Expression.Evaluate(env)
}

func (exp EXP) evaluateInbuilt(fi FunctionInfo, env interfaces.Scope) interfaces.Type {
	evaluatedArgs := make([]interfaces.Type, len(exp.Arguments))
	if fi.evaluateArgs {
		for p, arg := range exp.Arguments {
			if r, ok := arg.(REF); ok {
				arg = r.Evaluate(env)
			}
			if e, ok := arg.(interfaces.Evaluatable); ok {
				arg = e.Evaluate(env.NewChildScope())
			}
			evaluatedArgs[p] = arg
		}
	} else {
		copy(evaluatedArgs, exp.Arguments)
	}
	return fi.function(evaluatedArgs, env)
}

func (exp EXP) printExpression(result interfaces.Type) {
	if DEBUG {
		fmt.Printf("%v = %v\n", exp, result)
	}
}
