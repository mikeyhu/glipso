package common

import (
	"github.com/mikeyhu/glipso/interfaces"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Pair_IterateToTail(t *testing.T) {
	pair := P{I(1), &P{I(2), ENDED}}

	assert.Equal(t, I(1), pair.Head())
	assert.True(t, pair.HasTail())

	next, _ := pair.Iterate(GlobalEnvironment)
	assert.Equal(t, I(2), next.Head())
	assert.False(t, next.HasTail())
}

func Test_Pair_TailNotIterable(t *testing.T) {
	pair := LAZYP{I(1), &EXP{REF("+"), []interfaces.Type{I(2)}}}

	assert.Equal(t, I(1), pair.Head())
	assert.True(t, pair.HasTail())

	next, err := pair.Iterate(GlobalEnvironment)
	assert.Equal(t, ENDED, next)
	assert.EqualError(t, err, "lazypair : iterable expected got 2")
}
