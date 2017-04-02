package parser

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Tokenize_OpeningBracket(t *testing.T) {
	data := []byte("(somestuff)")
	advance, token, err := Tokenize(data, false)
	assert.NoError(t, err)
	assert.Equal(t, 1, advance)
	assert.Equal(t, []byte("("), token)
}

func Test_Tokenize_WordFollowedClosingBracket(t *testing.T) {
	data := []byte("somestuff)")
	advance, token, err := Tokenize(data, false)
	assert.NoError(t, err)
	assert.Equal(t, 9, advance)
	assert.Equal(t, []byte("somestuff"), token)
}

func Test_Tokenize_WordFollowedByAnArgument(t *testing.T) {
	data := []byte("somestuff 1 1)")
	advance, token, err := Tokenize(data, false)
	assert.NoError(t, err)
	assert.Equal(t, 9, advance)
	assert.Equal(t, []byte("somestuff"), token)
}

func Test_Tokenize_ClosingBracket(t *testing.T) {
	data := []byte("))")
	advance, token, err := Tokenize(data, false)
	assert.NoError(t, err)
	assert.Equal(t, 1, advance)
	assert.Equal(t, []byte(")"), token)
}

func Test_Tokenize_OpeningVector(t *testing.T) {
	data := []byte("[]")
	advance, token, err := Tokenize(data, false)
	assert.NoError(t, err)
	assert.Equal(t, 1, advance)
	assert.Equal(t, []byte("["), token)
}

func Test_Tokenize_ClosingVector(t *testing.T) {
	data := []byte("]")
	advance, token, err := Tokenize(data, false)
	assert.NoError(t, err)
	assert.Equal(t, 1, advance)
	assert.Equal(t, []byte("]"), token)
}

func Test_Tokenize_ClosingBracketAtEndOfReader(t *testing.T) {
	data := []byte(")")
	advance, token, err := Tokenize(data, true)
	assert.NoError(t, err)
	assert.Equal(t, 1, advance)
	assert.Equal(t, []byte(")"), token)
}

func Test_Tokenize_OpeningBracketAfterSpace(t *testing.T) {
	data := []byte(" (")
	advance, token, err := Tokenize(data, true)
	assert.NoError(t, err)
	assert.Equal(t, 2, advance)
	assert.Equal(t, []byte("("), token)
}
