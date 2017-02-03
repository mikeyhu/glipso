package parser

import (
	"github.com/mikeyhu/mekkanism/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParserReturnsExpressionFunctionName(t *testing.T) {
	result, err := Parse("(+ 1 2)")
	assert.NoError(t, err)
	assert.Equal(t, result.FunctionName, "+")
}

func TestParserReturnsExpressionArguments(t *testing.T) {
	result, err := Parse("(+ 1 3)")
	assert.NoError(t, err)
	assert.Equal(t, result.Arguments[0].(common.I).Int(), 1, "1st element")
	assert.Equal(t, result.Arguments[1].(common.I).Int(), 3, "2nd element")
	assert.Equal(t, len(result.Arguments), 2, "array length")
}
