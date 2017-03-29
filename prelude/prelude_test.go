package prelude

import (
	"github.com/mikeyhu/glipso/common"
	"github.com/mikeyhu/glipso/parser"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TODO : Refactor defn to macro
func TestPreludeDefnDoesNotShareScope(t *testing.T) {
	ParsePrelude(common.GlobalEnvironment)
	code := `
	(do
		(defn add1 [a] (+ a 1))
		(defn add2 [a] (+ a 2))
		(add2 (add1 (add2 100)))
	)
	`
	exp, err := parser.Parse(code)
	assert.NoError(t, err)
	result, _ := exp.Evaluate(common.GlobalEnvironment)
	assert.Equal(t, common.I(105), result)
}

func TestPreludeDefines_defmacro(t *testing.T) {
	ParsePrelude(common.GlobalEnvironment)
	result, ok := common.GlobalEnvironment.ResolveRef(common.REF("defmacro"))
	assert.True(t, ok)
	_, mok := result.(common.MAC)
	assert.True(t, mok)
}

func TestLastReturnsLastNumberInList(t *testing.T) {
	ParsePrelude(common.GlobalEnvironment)
	code := `
	(last (range 1 5))
	`
	exp, err := parser.Parse(code)
	assert.NoError(t, err)
	result, _ := exp.Evaluate(common.GlobalEnvironment)
	assert.Equal(t, common.I(5), result)
}

func TestRepeatReturnsTheItem(t *testing.T) {
	ParsePrelude(common.GlobalEnvironment)
	code := `
	(first (repeat "s" 5))
	`
	exp, err := parser.Parse(code)
	assert.NoError(t, err)
	result, _ := exp.Evaluate(common.GlobalEnvironment)
	assert.Equal(t, common.S("s"), result)
}

func TestRepeatReturnsTheItemNTimes(t *testing.T) {
	ParsePrelude(common.GlobalEnvironment)
	code := `
	(apply + (repeat 10 5))
	`
	exp, err := parser.Parse(code)
	assert.NoError(t, err)
	result, _ := exp.Evaluate(common.GlobalEnvironment)
	assert.Equal(t, common.I(50), result)
}
