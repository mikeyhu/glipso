package parser

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldParseOpeningBracket(t *testing.T) {
	data := []byte("(somestuff)")
	advance, token, err := ScanTokens(data, false)
	assert.NoError(t, err)
	assert.Equal(t, 1, advance)
	assert.Equal(t, []byte("("), token)
}

func TestShouldParseAWordFollowedClosingBracket(t *testing.T) {
	data := []byte("somestuff)")
	advance, token, err := ScanTokens(data, false)
	assert.NoError(t, err)
	assert.Equal(t, 9, advance)
	assert.Equal(t, []byte("somestuff"), token)
}

func TestShouldParseAWordFollowedByAnArgument(t *testing.T) {
	data := []byte("somestuff 1 1)")
	advance, token, err := ScanTokens(data, false)
	assert.NoError(t, err)
	assert.Equal(t, 9, advance)
	assert.Equal(t, []byte("somestuff"), token)
}

func TestShouldParseAnClosingBracket(t *testing.T) {
	data := []byte("))")
	advance, token, err := ScanTokens(data, false)
	assert.NoError(t, err)
	assert.Equal(t, 1, advance)
	assert.Equal(t, []byte(")"), token)
}

func TestShouldParseAnOpeningVector(t *testing.T) {
	data := []byte("[]")
	advance, token, err := ScanTokens(data, false)
	assert.NoError(t, err)
	assert.Equal(t, 1, advance)
	assert.Equal(t, []byte("["), token)
}

func TestShouldParseAnClosingVector(t *testing.T) {
	data := []byte("]")
	advance, token, err := ScanTokens(data, false)
	assert.NoError(t, err)
	assert.Equal(t, 1, advance)
	assert.Equal(t, []byte("]"), token)
}

func TestShouldParseAnClosingBracketAtEndOfReader(t *testing.T) {
	data := []byte(")")
	advance, token, err := ScanTokens(data, true)
	assert.NoError(t, err)
	assert.Equal(t, 1, advance)
	assert.Equal(t, []byte(")"), token)
}

func TestShouldParseAnOpeningBracketAfterSpace(t *testing.T) {
	data := []byte(" (")
	advance, token, err := ScanTokens(data, true)
	assert.NoError(t, err)
	assert.Equal(t, 2, advance)
	assert.Equal(t, []byte("("), token)
}
