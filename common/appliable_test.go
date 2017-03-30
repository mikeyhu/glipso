package common

import (
	"github.com/mikeyhu/glipso/interfaces"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFIApplyReturnsError(t *testing.T) {
	//given
	fi := FI{name: "my-function"}
	//when
	result, err := fi.Apply([]interfaces.Type{}, GlobalEnvironment.NewChildScope())
	//then
	assert.Equal(t, NILL, result)
	assert.EqualError(t, err, "FI : my-function had neither an evaluator or lazy evaluator")
}
