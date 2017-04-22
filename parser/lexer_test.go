package parser

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_tokenize_OpeningBracket(t *testing.T) {
	data := []byte("(somestuff)")
	advance, token, err := tokenize(data, false)
	assert.NoError(t, err)
	assert.Equal(t, 1, advance)
	assert.Equal(t, []byte("("), token)
}

func Test_tokenize_WordFollowedClosingBracket(t *testing.T) {
	data := []byte("somestuff)")
	advance, token, err := tokenize(data, false)
	assert.NoError(t, err)
	assert.Equal(t, 9, advance)
	assert.Equal(t, []byte("somestuff"), token)
}

func Test_tokenize_WordFollowedByAnArgument(t *testing.T) {
	data := []byte("somestuff 1 1)")
	advance, token, err := tokenize(data, false)
	assert.NoError(t, err)
	assert.Equal(t, 9, advance)
	assert.Equal(t, []byte("somestuff"), token)
}

func Test_tokenize_ClosingBracket(t *testing.T) {
	data := []byte("))")
	advance, token, err := tokenize(data, false)
	assert.NoError(t, err)
	assert.Equal(t, 1, advance)
	assert.Equal(t, []byte(")"), token)
}

func Test_tokenize_OpeningVector(t *testing.T) {
	data := []byte("[]")
	advance, token, err := tokenize(data, false)
	assert.NoError(t, err)
	assert.Equal(t, 1, advance)
	assert.Equal(t, []byte("["), token)
}

func Test_tokenize_ClosingVector(t *testing.T) {
	data := []byte("]")
	advance, token, err := tokenize(data, false)
	assert.NoError(t, err)
	assert.Equal(t, 1, advance)
	assert.Equal(t, []byte("]"), token)
}

func Test_tokenize_ClosingBracketAtEndOfReader(t *testing.T) {
	data := []byte(")")
	advance, token, err := tokenize(data, true)
	assert.NoError(t, err)
	assert.Equal(t, 1, advance)
	assert.Equal(t, []byte(")"), token)
}

func Test_tokenize_OpeningBracketAfterSpace(t *testing.T) {
	data := []byte(" (")
	advance, token, err := tokenize(data, true)
	assert.NoError(t, err)
	assert.Equal(t, 2, advance)
	assert.Equal(t, []byte("("), token)
}

func Test_tokenize_QuotedStringWithSpaces(t *testing.T) {
	data := []byte("\"hello world\" ")
	advance, token, err := tokenize(data, true)
	assert.NoError(t, err)
	assert.Equal(t, 13, advance)
	assert.Equal(t, []byte("\"hello world\""), token)
}

func Test_tokenize_Error_UnclosedString(t *testing.T) {
	data := []byte("\"hello world ")
	_, _, err := tokenize(data, true)
	assert.EqualError(t, err, "string not closed")
}

func Test_tokenize_QuotedStringContainingEscapedQuote(t *testing.T) {
	data := []byte(`"quote \" here" `)
	advance, token, err := tokenize(data, true)
	assert.NoError(t, err)
	assert.Equal(t, 15, advance)
	assert.Equal(t, []byte(`"quote \" here"`), token)
}
