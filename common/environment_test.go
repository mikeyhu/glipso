package common

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEnvironmentPanicsIfReferenceNotFound(t *testing.T) {
	assert.Panics(t, func() {
		GlobalEnvironment.ResolveRef(REF("unset"))
	})
}

func TestEnvironmentReturnsReferencesIfFound(t *testing.T) {
	GlobalEnvironment.CreateRef(REF("one"), I(1))
	result := GlobalEnvironment.ResolveRef(REF("one"))
	assert.Equal(t, I(1), result)
}
