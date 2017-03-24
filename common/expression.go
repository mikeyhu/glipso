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
func (exp *EXP) Evaluate(sco interfaces.Scope) interfaces.Value {
	exp.printStartExpression()
	var result interfaces.Value
	function := exp.Function

	if toREF, ok := function.(REF); ok {
		if fn, ok := sco.ResolveRef(toREF); ok {
			function = fn
		} else {
			panic(fmt.Sprintf("evaluate : function %v not found", toREF))
		}
	}

	if toMacro, ok := function.(interfaces.Expandable); ok {
		result = toMacro.Expand(exp.Arguments).Evaluate(sco)
	} else {
		function := evaluateToResult(function, sco)
		if toFN, ok := function.(interfaces.Function); ok {
			result = toFN.Apply(exp.Arguments, sco)
		}
	}

	exp.printEndExpression(result)
	if result == nil {
		panic(fmt.Sprintf("Evaluate : evaluation down to nil"))
	}
	return result
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
func (bexp *BOUNDEXP) Evaluate(sco interfaces.Scope) interfaces.Value {
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
