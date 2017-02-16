package common

import (
	"github.com/mikeyhu/glipso/interfaces"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEvaluatePlusWith2Arguments(t *testing.T) {
	exp := EXP{FunctionName: "+", Arguments: []interfaces.Type{I(1), I(2)}}
	result := exp.Evaluate(GlobalEnvironment)
	assert.Equal(t, I(3), result)
}

func TestEvaluatePlusWithManyArguments(t *testing.T) {
	exp := EXP{FunctionName: "+", Arguments: []interfaces.Type{I(1), I(2), I(3)}}
	result := exp.Evaluate(GlobalEnvironment)
	assert.Equal(t, I(6), result)
}

func TestEvaluateMinusWith2Arguments(t *testing.T) {
	exp := EXP{FunctionName: "-", Arguments: []interfaces.Type{I(5), I(1)}}
	result := exp.Evaluate(GlobalEnvironment)
	assert.Equal(t, I(4), result)
}

func TestEvaluateNestedFunction(t *testing.T) {
	exp := EXP{FunctionName: "+", Arguments: []interfaces.Type{
		I(1),
		EXP{FunctionName: "-", Arguments: []interfaces.Type{I(2), I(1)}}}}
	result := exp.Evaluate(GlobalEnvironment)
	assert.Equal(t, I(2), result)
}

func TestEvaluateFN(t *testing.T) {
	exp := EXPN{Function: FN{
		VEC{[]interfaces.Type{REF("a")}},
		EXP{FunctionName: "+", Arguments: []interfaces.Type{REF("a"), I(1)}}},
		Arguments: []interfaces.Type{I(2)}}
	result := exp.Evaluate(GlobalEnvironment)
	assert.Equal(t, I(3), result)
}

func TestEvaluateFNHasMoreArgumentsThanProvided(t *testing.T) {
	exp := EXPN{Function: FN{
		VEC{[]interfaces.Type{REF("a"), REF("b")}},
		EXP{FunctionName: "+", Arguments: []interfaces.Type{REF("a"), I(1)}}},
		Arguments: []interfaces.Type{I(2)}}

	assert.Panics(t, func() {
		exp.Evaluate(GlobalEnvironment)
	})
}

func TestEvaluateFNHasLessArgumentsThanProvided(t *testing.T) {
	exp := EXPN{Function: FN{
		VEC{[]interfaces.Type{REF("a")}},
		EXP{FunctionName: "+", Arguments: []interfaces.Type{REF("a"), I(1)}}},
		Arguments: []interfaces.Type{I(2), I(3)}}

	assert.Panics(t, func() {
		exp.Evaluate(GlobalEnvironment)
	})
}
