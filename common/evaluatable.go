package common

import (
	"fmt"
	"github.com/mikeyhu/glipso/interfaces"
)

// DEBUG enable to display debug information for function evaluation
var DEBUG = false

func evaluateToValue(value interfaces.Type, sco interfaces.Scope) (interfaces.Value, error) {
	switch v := value.(type) {
	case interfaces.Evaluatable:
		return v.Evaluate(sco)
	case interfaces.Value:
		return v, nil
	default:
		return NILL, fmt.Errorf("evaluateToValue : value %v of type %v is neither evaluatable or a result", value, v)
	}
}

// EXP is a an Expression that has a Appliable and Arguments and can be evaluated against a scope
type EXP struct {
	Function  interfaces.Type
	Arguments []interfaces.Type
}

// IsType for EXP
func (exp *EXP) IsType() {}

// String representation of EXP
func (exp *EXP) String() string {
	argAsS := fmt.Sprintf("%v", exp.Arguments)
	return fmt.Sprintf("(%v %v)", exp.Function, argAsS[1:len(argAsS)-1])
}

// Evaluate evaluates the Appliable provided with the Arguments and Scope
func (exp *EXP) Evaluate(sco interfaces.Scope) (interfaces.Value, error) {
	exp.printStartExpression()
	var result interfaces.Value
	var err error
	function := exp.Function

	if toREF, ok := function.(REF); ok {
		if fn, ok := sco.ResolveRef(toREF); ok {
			function = fn
		} else {
			return exp.returnAndPrint(NILL, fmt.Errorf("evaluate : function '%v' not found", toREF))
		}
	}

	if toMacro, ok := function.(interfaces.Expandable); ok {
		result, err = toMacro.Expand(exp.Arguments).Evaluate(sco)
		if err != nil {
			return exp.returnAndPrint(NILL, err)
		}
	} else {
		function, err := evaluateToValue(function, sco)
		if err != nil {
			return exp.returnAndPrint(NILL, err)
		}
		if toFN, ok := function.(interfaces.Appliable); ok {
			result, err = toFN.Apply(exp.Arguments, sco)
			if err != nil {
				return exp.returnAndPrint(NILL, err)
			}
		}
	}
	return exp.returnAndPrint(result, nil)
}

func (exp *EXP) printStartExpression() {
	if DEBUG {
		fmt.Printf("%v = ?\n", exp)
	}
}

func (exp *EXP) returnAndPrint(result interfaces.Value, err error) (interfaces.Value, error) {
	if err != nil {
		fmt.Printf("%v\n\tcaused ERROR : %v\n", exp, err)
	}
	if DEBUG {
		fmt.Printf("%v = %v\n", exp, result)
	}
	return result, err
}

// REF (Reference)
// symbol for something in scope, variable or function
type REF string

// IsType for REF
func (r REF) IsType() {}

// String representation of a REF
func (r REF) String() string {
	return string(r)
}

// Evaluate resolves a REF to something in scope
func (r REF) Evaluate(sco interfaces.Scope) (interfaces.Value, error) {
	if DEBUG {
		env := sco.(*Environment)
		fmt.Printf("%v being looked up in scope %v:\n", r, env.id)
		env.DisplayEnvironment()
	}
	if resolved, ok := sco.ResolveRef(r); ok {
		return resolved, nil
	}
	return NILL, fmt.Errorf("unable to resolve REF('%v')", r)
}

// BOUNDEXP provides a way for a Expression to be bound to a particular scope for later evaluation
type BOUNDEXP struct {
	Evaluatable interfaces.Evaluatable
	Scope       interfaces.Scope
}

// Evaluate on a Bound Expression replaces the provided scope with the bound scope
func (bexp *BOUNDEXP) Evaluate(sco interfaces.Scope) (interfaces.Value, error) {
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
