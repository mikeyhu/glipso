package common

import (
	"errors"
	"github.com/mikeyhu/glipso/interfaces"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEvaluatePlusWith2Arguments(t *testing.T) {
	exp := EXP{Function: REF("+"), Arguments: []interfaces.Type{I(1), I(2)}}
	result, err := exp.Evaluate(GlobalEnvironment)
	assert.NoError(t, err)
	assert.Equal(t, I(3), result)
}

func TestEvaluatePlusWithManyArguments(t *testing.T) {
	exp := EXP{Function: REF("+"), Arguments: []interfaces.Type{I(1), I(2), I(3)}}
	result, err := exp.Evaluate(GlobalEnvironment)
	assert.NoError(t, err)
	assert.Equal(t, I(6), result)
}

func TestEvaluateMinusWith2Arguments(t *testing.T) {
	exp := EXP{Function: REF("-"), Arguments: []interfaces.Type{I(5), I(1)}}
	result, err := exp.Evaluate(GlobalEnvironment)
	assert.NoError(t, err)
	assert.Equal(t, I(4), result)
}

func TestEvaluateNestedFunction(t *testing.T) {
	exp := &EXP{Function: REF("+"), Arguments: []interfaces.Type{
		I(1),
		&EXP{Function: REF("-"), Arguments: []interfaces.Type{I(2), I(1)}}}}
	result, err := exp.Evaluate(GlobalEnvironment)
	assert.NoError(t, err)
	assert.Equal(t, I(2), result)
}

func TestEvaluateFN(t *testing.T) {
	exp := EXP{Function: FN{
		VEC{[]interfaces.Type{REF("a")}},
		&EXP{Function: REF("+"), Arguments: []interfaces.Type{REF("a"), I(1)}}},
		Arguments: []interfaces.Type{I(2)}}
	result, err := exp.Evaluate(GlobalEnvironment)
	assert.NoError(t, err)
	assert.Equal(t, I(3), result)
}

func TestEvaluateFNHasMoreArgumentsThanProvided(t *testing.T) {
	exp := EXP{Function: FN{
		VEC{[]interfaces.Type{REF("a"), REF("b")}},
		&EXP{Function: REF("+"), Arguments: []interfaces.Type{REF("a"), I(1)}}},
		Arguments: []interfaces.Type{I(2)}}

	result, err := exp.Evaluate(GlobalEnvironment)
	assert.Equal(t, NILL, result)
	assert.EqualError(t, err, "too few arguments")
}

func TestEvaluateFNHasLessArgumentsThanProvided(t *testing.T) {
	exp := EXP{Function: FN{
		VEC{[]interfaces.Type{REF("a")}},
		&EXP{Function: REF("+"), Arguments: []interfaces.Type{REF("a"), I(1)}}},
		Arguments: []interfaces.Type{I(2), I(3)}}

	result, err := exp.Evaluate(GlobalEnvironment)
	assert.Equal(t, NILL, result)
	assert.EqualError(t, err, "too many arguments")
}

func TestErrorsWhenEvaluatingUnresolvedREF(t *testing.T) {
	ref := REF("notset")
	result, err := ref.Evaluate(GlobalEnvironment)
	assert.Equal(t, NILL, result)
	assert.Error(t, errors.New("..."), err)
}

func TestErrorsWhenFunctionNotFound(t *testing.T) {
	exp := EXP{Function: REF("not-a-function"), Arguments: []interfaces.Type{REF("a"), I(1)}}

	result, err := exp.Evaluate(GlobalEnvironment)
	assert.Equal(t, NILL, result)
	assert.EqualError(t, err, "evaluate : function 'not-a-function' not found")
}

func TestErrorsWhenEvaluateToValueIsNeitherEvaluatableOrResult(t *testing.T) {
	exp := EXP{}
	result, err := exp.Evaluate(GlobalEnvironment)
	assert.Equal(t, NILL, result)
	assert.EqualError(t, err, "evaluateToValue : value <nil> of type <nil> is neither evaluatable or a result")
}
