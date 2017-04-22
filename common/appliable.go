package common

import (
	"errors"
	"fmt"
	"github.com/mikeyhu/glipso/interfaces"
)

// FN acts as storage for a reusable Appliable by storing a set of arguments to a function and the function expression itself
type FN struct {
	Arguments  VEC
	Expression interfaces.Evaluatable
}

// IsType for FN
func (f FN) IsType() {}

// IsValue for FN
func (f FN) IsValue() {}

// String output for FN
func (f FN) String() string {
	return fmt.Sprintf("FN(%v, %v)", f.Arguments, f.Expression)
}

// Apply for FN : validates the number of args and then applies the FN to the arguments
func (f FN) Apply(arguments []interfaces.Type, env interfaces.Scope) (interfaces.Value, error) {
	if len(f.Arguments.Vector) < len(arguments) {
		return NILL, errors.New("too many arguments")
	} else if len(f.Arguments.Vector) > len(arguments) {
		return NILL, errors.New("too few arguments")
	}
	fnenv := env.NewChildScope()
	for i, v := range f.Arguments.Vector {
		eval, err := evaluateToValue(arguments[i], env)
		if err != nil {
			return NILL, err
		}
		fnenv.CreateRef(v.(REF), eval)
	}
	return f.Expression.Evaluate(fnenv)
}

// FI provides information about a built in function
type FI struct {
	name          string
	evaluator     evaluator
	lazyEvaluator lazyEvaluator
	argumentCount int
}

// IsType for FI
func (fi FI) IsType() {}

// IsValue for FI
func (fi FI) IsValue() {}

// String for FI
func (fi FI) String() string {
	return fmt.Sprintf("FI(%s)", fi.name)
}

// Apply for FI : validates the number of args and then applies the FI to the arguments
func (fi FI) Apply(arguments []interfaces.Type, sco interfaces.Scope) (interfaces.Value, error) {
	if fi.argumentCount > 0 && len(arguments) != fi.argumentCount {
		return NILL, fmt.Errorf("%v : invalid number of arguments [%d of %d]", fi.name, len(arguments), fi.argumentCount)
	}
	if fi.evaluator != nil {
		evaluatedArgs := make([]interfaces.Value, len(arguments))
		for p, arg := range arguments {
			var err error
			evaluatedArgs[p], err = evaluateToValue(arg, sco)
			if err != nil {
				return NILL, err
			}
		}
		return fi.evaluator(evaluatedArgs, sco)
	} else if fi.lazyEvaluator != nil {
		unevaluatedArgs := make([]interfaces.Type, len(arguments))
		copy(unevaluatedArgs, arguments)
		return fi.lazyEvaluator(unevaluatedArgs, sco)
	}
	return NILL, fmt.Errorf("FI : %v had neither an evaluator or lazy evaluator", fi.name)
}
