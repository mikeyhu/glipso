package common

import (
	"github.com/mikeyhu/glipso/interfaces"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_FI_Apply_ReturnsError(t *testing.T) {
	//given
	fi := FI{name: "my-function"}
	//when
	result, err := fi.Apply([]interfaces.Type{}, GlobalEnvironment.NewChildScope())
	//then
	assert.Equal(t, NILL, result)
	assert.EqualError(t, err, "FI : my-function had neither an evaluator or lazy evaluator")
}

func Test_FI_Apply_ValidatesArgumentLengths(t *testing.T) {
	//given
	fi := FI{name: "filter", evaluator: filter, argumentCount: 2}
	//when
	result, err := fi.Apply([]interfaces.Type{}, GlobalEnvironment.NewChildScope())
	//then
	assert.Equal(t, NILL, result)
	assert.EqualError(t, err, "filter : invalid number of arguments [0 of 2]")
}
