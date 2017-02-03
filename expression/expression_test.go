package expression

import (
	"github.com/mikeyhu/mekkanism/interfaces"
	"github.com/mikeyhu/mekkanism/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEvaluatePlusWith2Arguments(t *testing.T) {
	exp := &Expression{FunctionName: "+", Arguments: []interfaces.Argument{types.I(1), types.I(2)}}
	result := exp.Evaluate()
	assert.Equal(t, types.I(3), result)
}

func TestEvaluatePlusWithManyArguments(t *testing.T) {
	exp := &Expression{FunctionName: "+", Arguments: []interfaces.Argument{types.I(1), types.I(2), types.I(3)}}
	result := exp.Evaluate()
	assert.Equal(t, types.I(6), result)
}

func TestEvaluateMinusWith2Arguments(t *testing.T) {
	exp := &Expression{FunctionName: "-", Arguments: []interfaces.Argument{types.I(5), types.I(1)}}
	result := exp.Evaluate()
	assert.Equal(t, types.I(4), result)
}
