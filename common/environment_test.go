package common

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEnvironmentReturnsFalseIfReferenceNotFound(t *testing.T) {
	_, ok := GlobalEnvironment.ResolveRef(REF("unset"))
	assert.False(t, ok)
}

func TestEnvironmentReturnsReferencesIfFound(t *testing.T) {
	GlobalEnvironment.CreateRef(REF("one"), I(1))
	result, _ := GlobalEnvironment.ResolveRef(REF("one"))
	assert.Equal(t, I(1), result)
}
