package expression

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"github.com/mikeyhu/mekkanism/types"
)

func TestEvaluatePlusWith2Arguments(t *testing.T) {
	exp := &Expression{FunctionName: "+", Arguments: []types.Argtype{types.Argtype{Integer:1}, types.Argtype{Integer:2}}}
	result := exp.Evaluate()
	assert.Equal(t, 3, result.Integer)
}

func TestEvaluateMinusWith2Arguments(t *testing.T) {
	exp := &Expression{FunctionName: "-", Arguments: []types.Argtype{types.Argtype{Integer:5}, types.Argtype{Integer:1}}}
	result := exp.Evaluate()
	assert.Equal(t, 4, result.Integer)
}
