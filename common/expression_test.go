package common

import (
	"github.com/mikeyhu/mekkanism/interfaces"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEvaluatePlusWith2Arguments(t *testing.T) {
	exp := Expression{FunctionName: "+", Arguments: []interfaces.Argument{I(1), I(2)}}
	result := exp.Evaluate()
	assert.Equal(t, I(3), result)
}

func TestEvaluatePlusWithManyArguments(t *testing.T) {
	exp := Expression{FunctionName: "+", Arguments: []interfaces.Argument{I(1), I(2), I(3)}}
	result := exp.Evaluate()
	assert.Equal(t, I(6), result)
}

func TestEvaluateMinusWith2Arguments(t *testing.T) {
	exp := Expression{FunctionName: "-", Arguments: []interfaces.Argument{I(5), I(1)}}
	result := exp.Evaluate()
	assert.Equal(t, I(4), result)
}

func TestEvaluateNestedFunction(t *testing.T) {
	exp := Expression{FunctionName: "+", Arguments: []interfaces.Argument{I(1),
		Expression{FunctionName: "-", Arguments: []interfaces.Argument{I(2), I(1)}}}}
	result := exp.Evaluate()
	assert.Equal(t, I(2), result)
}
