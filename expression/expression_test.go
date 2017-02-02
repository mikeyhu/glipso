package expression

import (
	"github.com/mikeyhu/mekkanism/interfaces"
	"github.com/mikeyhu/mekkanism/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEvaluatePlusWith2Arguments(t *testing.T) {
	exp := &Expression{FunctionName: "+", Arguments: []interfaces.Argument{&types.Argtype{Integer: 1}, &types.Argtype{Integer: 2}}}
	result := exp.Evaluate()
	assert.Equal(t, 3, result.Integer)
}

func TestEvaluatePlusWithManyArguments(t *testing.T) {
	exp := &Expression{FunctionName: "+", Arguments: []interfaces.Argument{&types.Argtype{Integer: 1}, &types.Argtype{Integer: 2}, &types.Argtype{Integer: 3}}}
	result := exp.Evaluate()
	assert.Equal(t, 6, result.Integer)
}

func TestEvaluateMinusWith2Arguments(t *testing.T) {
	exp := &Expression{FunctionName: "-", Arguments: []interfaces.Argument{&types.Argtype{Integer: 5}, &types.Argtype{Integer: 1}}}
	result := exp.Evaluate()
	assert.Equal(t, 4, result.Integer)
}
