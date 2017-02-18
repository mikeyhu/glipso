package prelude

import (
	"github.com/mikeyhu/glipso/common"
	"github.com/mikeyhu/glipso/parser"
	"github.com/stretchr/testify/assert"
	"testing"
)

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
	result := exp.Evaluate(common.GlobalEnvironment)
	assert.Equal(t, common.I(105), result)
}
