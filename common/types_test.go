package common

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPanicsWhenEvaluatingUnresolvedREF(t *testing.T) {
	ref := REF("notset")
	assert.Panics(t, func() {
		ref.Evaluate(GlobalEnvironment)
	})
}

func TestPairCanIterateToTail(t *testing.T) {
	pair := P{I(1), &P{I(2), nil}}

	assert.Equal(t, I(1), pair.Head())
	assert.True(t, pair.HasTail())
	assert.Equal(t, I(2), pair.Iterate(GlobalEnvironment).Head())
	assert.False(t, pair.Iterate(GlobalEnvironment).HasTail())
}
