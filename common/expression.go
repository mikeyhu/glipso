package common

import (
	"fmt"
	"github.com/mikeyhu/glipso/interfaces"
)

// DEBUG enable to display debug information for function evaluation
var DEBUG = false

// EXP is a an Expression that has a Function and Arguments and can be evaluated against a scope
type EXP struct {
	Function  interfaces.Type
	Arguments []interfaces.Type
}

// IsType for EXP
func (exp *EXP) IsType() {}

// String representation of EXP
func (exp *EXP) String() string {
	return fmt.Sprintf("(%v %v)", exp.Function, exp.Arguments)
}

// Evaluate evaluates the Function provided with the Arguments and Scope
func (exp *EXP) Evaluate(sco interfaces.Scope) interfaces.Type {
	exp.printStartExpression()
	var result interfaces.Type
	function := exp.Function

	if toREF, ok := function.(REF); ok {
		if fn, ok := sco.ResolveRef(toREF); ok {
			function = fn
		} else if fi, ok := inbuilt[toREF.String()]; ok {
			result = exp.evaluateInbuilt(fi, sco.NewChildScope())
		}
	}

	if toMacro, ok := function.(interfaces.Expandable); ok {
		result = toMacro.Expand(exp.Arguments)
	} else {
		if toEval, ok := function.(*EXP); ok {
			function = toEval.Evaluate(sco)
		}
		if toFN, ok := function.(FN); ok {
			result = exp.evaluateFN(toFN, sco)
		}
	}

	exp.printEndExpression(result)
	if exp, ok := result.(*EXP); ok {
		result = exp.Evaluate(sco)
	}
	return result
}

func (exp *EXP) evaluateFN(fn FN, env interfaces.Scope) interfaces.Type {
	if len(fn.Arguments.Vector) != len(exp.Arguments) {
		panic("Invalid number of arguments")
	}
	fnenv := env.NewChildScope()
	for i, v := range fn.Arguments.Vector {
		if ev, ok := exp.Arguments[i].(interfaces.Evaluatable); ok {
			fnenv.CreateRef(v.(REF), ev.Evaluate(env))
		} else {
			fnenv.CreateRef(v.(REF), exp.Arguments[i])
		}
	}
	return fn.Expression.Evaluate(fnenv)
}

func (exp *EXP) evaluateInbuilt(fi FunctionInfo, env interfaces.Scope) interfaces.Type {
	evaluatedArgs := make([]interfaces.Type, len(exp.Arguments))
	if fi.evaluateArgs {
		for p, arg := range exp.Arguments {
			if r, ok := arg.(REF); ok {
				arg = r.Evaluate(env)
			}
			if e, ok := arg.(interfaces.Evaluatable); ok {
				arg = e.Evaluate(env)
			}
			evaluatedArgs[p] = arg
		}
	} else {
		copy(evaluatedArgs, exp.Arguments)
	}
	return fi.function(evaluatedArgs, env)
}

func (exp *EXP) printStartExpression() {
	if DEBUG {
		fmt.Printf("%v = ?\n", exp)
	}
}

func (exp *EXP) printEndExpression(result interfaces.Type) {
	if DEBUG {
		fmt.Printf("%v = %v\n", exp, result)
	}
}

// BOUNDEXP provides a way for a Expression to be bound to a particular scope for later evaluation
type BOUNDEXP struct {
	Evaluatable interfaces.Evaluatable
	Scope       interfaces.Scope
}

// Evaluate on a Bound Expression replaces the provided scope with the bound scope
func (bexp *BOUNDEXP) Evaluate(sco interfaces.Scope) interfaces.Type {
	return bexp.Evaluatable.Evaluate(bexp.Scope)
}

// String representation of a BEXP
func (bexp *BOUNDEXP) String() string {
	return fmt.Sprintf("BEXP(%v %v)", bexp.Evaluatable, bexp.Scope)
}

// BindEvaluation Creates a new BOUNDEXP with the Expression and Scope provided
func BindEvaluation(ev interfaces.Evaluatable, sco interfaces.Scope) interfaces.Evaluatable {
	return &BOUNDEXP{ev, sco}
}
