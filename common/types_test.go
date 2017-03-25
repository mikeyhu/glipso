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
