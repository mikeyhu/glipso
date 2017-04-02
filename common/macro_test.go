package common

import (
	"github.com/mikeyhu/glipso/interfaces"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Macro_Expand(t *testing.T) {
	macro := MAC{
		VEC{[]interfaces.Type{REF("a")}},
		&EXP{REF("+"), []interfaces.Type{REF("a"), I(1)}},
	}

	result := macro.Expand([]interfaces.Type{I(10)})

	assert.Equal(t, &EXP{REF("+"), []interfaces.Type{I(10), I(1)}}, result)
}

func Test_Macro_NestedExpansion(t *testing.T) {
	macro := MAC{
		VEC{[]interfaces.Type{REF("a")}},
		&EXP{REF("+"), []interfaces.Type{
			&EXP{REF("+"), []interfaces.Type{REF("a"), I(1)}}},
		},
	}

	result := macro.Expand([]interfaces.Type{I(10)})

	assert.Equal(t, &EXP{REF("+"), []interfaces.Type{I(10), I(1)}}, result.(*EXP).Arguments[0])
}

func Test_Macro_FoundAndExpanded(t *testing.T) {
	GlobalEnvironment.CreateRef(REF("adder"), MAC{
		VEC{[]interfaces.Type{REF("a")}},
		&EXP{REF("+"), []interfaces.Type{REF("a"), I(1)}},
	})

	expression := &EXP{REF("adder"), []interfaces.Type{I(10)}}

	result, _ := expression.Evaluate(GlobalEnvironment)
	assert.Equal(t, I(11), result)
}
