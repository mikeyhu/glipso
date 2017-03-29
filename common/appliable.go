package common

import (
	"fmt"
	"github.com/mikeyhu/glipso/interfaces"
)

// FN acts as storage for a reusable Appliable by storing a set of arguments to a function and the function expression itself
type FN struct {
	Arguments  VEC
	Expression interfaces.Evaluatable
}

// IsType for FN
func (f FN) IsType()  {}
func (f FN) IsValue() {}

// String output for FN
func (f FN) String() string {
	return fmt.Sprintf("FN(%v, %v)", f.Arguments, f.Expression)
}

func (f FN) Apply(arguments []interfaces.Type, env interfaces.Scope) (interfaces.Value, error) {
	if len(f.Arguments.Vector) != len(arguments) {
		panic("Invalid number of arguments")
	}
	fnenv := env.NewChildScope()
	for i, v := range f.Arguments.Vector {
		eval, _ := evaluateToValue(arguments[i], env)
		fnenv.CreateRef(v.(REF), eval)
	}
	return f.Expression.Evaluate(fnenv)
}

// FI provides information about a built in function
type FI struct {
	name          string
	evaluator     evaluator
	lazyEvaluator lazyEvaluator
}

// IsType for FI
func (fi FI) IsType()  {}
func (fi FI) IsValue() {}

// String for FI
func (fi FI) String() string {
	return fmt.Sprintf("FI(%s)", fi.name)
}
func (fi FI) Apply(arguments []interfaces.Type, sco interfaces.Scope) (interfaces.Value, error) {
	if fi.evaluator != nil {
		evaluatedArgs := make([]interfaces.Value, len(arguments))
		for p, arg := range arguments {
			evaluatedArgs[p], _ = evaluateToValue(arg, sco)
		}
		return fi.evaluator(evaluatedArgs, sco)
	} else if fi.lazyEvaluator != nil {
		unevaluatedArgs := make([]interfaces.Type, len(arguments))
		copy(unevaluatedArgs, arguments)
		return fi.lazyEvaluator(unevaluatedArgs, sco)
	}
	panic(fmt.Sprintf("FI : %v had neither an evaluator or lazy evaluator", fi.name))
}
