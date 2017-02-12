package common

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEnvironmentPanicsIfReferenceNotFound(t *testing.T) {
	assert.Panics(t, func() {
		GlobalEnvironment.resolveRef(REF("unset"))
	})
}

func TestEnvironmentReturnsReferencesIfFound(t *testing.T) {
	GlobalEnvironment.createRef("one", I(1))
	result := GlobalEnvironment.resolveRef(REF("one"))
	assert.Equal(t, I(1), result)
}
