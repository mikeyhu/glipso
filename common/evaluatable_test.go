package common

import (
	"errors"
	"github.com/mikeyhu/glipso/interfaces"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Evaluate_AddWith2Arguments(t *testing.T) {
	exp := EXP{Function: REF("+"), Arguments: []interfaces.Type{I(1), I(2)}}
	result, err := exp.Evaluate(GlobalEnvironment)
	assert.NoError(t, err)
	assert.Equal(t, I(3), result)
}

func Test_Evaluate_AddWithManyArguments(t *testing.T) {
	exp := EXP{Function: REF("+"), Arguments: []interfaces.Type{I(1), I(2), I(3)}}
	result, err := exp.Evaluate(GlobalEnvironment)
	assert.NoError(t, err)
	assert.Equal(t, I(6), result)
}

func Test_Evaluate_MinusWith2Arguments(t *testing.T) {
	exp := EXP{Function: REF("-"), Arguments: []interfaces.Type{I(5), I(1)}}
	result, err := exp.Evaluate(GlobalEnvironment)
	assert.NoError(t, err)
	assert.Equal(t, I(4), result)
}

func Test_Evaluate_NestedExpression(t *testing.T) {
	exp := &EXP{Function: REF("+"), Arguments: []interfaces.Type{
		I(1),
		&EXP{Function: REF("-"), Arguments: []interfaces.Type{I(2), I(1)}}}}
	result, err := exp.Evaluate(GlobalEnvironment)
	assert.NoError(t, err)
	assert.Equal(t, I(2), result)
}

func Test_Evaluate_FN(t *testing.T) {
	exp := EXP{Function: FN{
		VEC{[]interfaces.Type{REF("a")}},
		&EXP{Function: REF("+"), Arguments: []interfaces.Type{REF("a"), I(1)}}},
		Arguments: []interfaces.Type{I(2)}}
	result, err := exp.Evaluate(GlobalEnvironment)
	assert.NoError(t, err)
	assert.Equal(t, I(3), result)
}

func Test_Evaluate_FNHasMoreArgumentsThanProvided(t *testing.T) {
	exp := EXP{Function: FN{
		VEC{[]interfaces.Type{REF("a"), REF("b")}},
		&EXP{Function: REF("+"), Arguments: []interfaces.Type{REF("a"), I(1)}}},
		Arguments: []interfaces.Type{I(2)}}

	result, err := exp.Evaluate(GlobalEnvironment)
	assert.Equal(t, NILL, result)
	assert.EqualError(t, err, "too few arguments")
}

func Test_Evaluate_FNHasLessArgumentsThanProvided(t *testing.T) {
	exp := EXP{Function: FN{
		VEC{[]interfaces.Type{REF("a")}},
		&EXP{Function: REF("+"), Arguments: []interfaces.Type{REF("a"), I(1)}}},
		Arguments: []interfaces.Type{I(2), I(3)}}

	result, err := exp.Evaluate(GlobalEnvironment)
	assert.Equal(t, NILL, result)
	assert.EqualError(t, err, "too many arguments")
}

func Test_Evaluate_UnresolvedREF(t *testing.T) {
	ref := REF("notset")
	result, err := ref.Evaluate(GlobalEnvironment)
	assert.Equal(t, NILL, result)
	assert.Error(t, errors.New("..."), err)
}

func Test_Evaluate_FunctionNotFound(t *testing.T) {
	exp := EXP{Function: REF("not-a-function"), Arguments: []interfaces.Type{REF("a"), I(1)}}

	result, err := exp.Evaluate(GlobalEnvironment)
	assert.Equal(t, NILL, result)
	assert.EqualError(t, err, "evaluate : function 'not-a-function' not found")
}

func Test_Evaluate_ValueIsNeitherEvaluatableOrResult(t *testing.T) {
	exp := EXP{}
	result, err := exp.Evaluate(GlobalEnvironment)
	assert.Equal(t, NILL, result)
	assert.EqualError(t, err, "evaluateToValue : value <nil> of type <nil> is neither evaluatable or a result")
}
