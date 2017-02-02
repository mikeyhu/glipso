package inbuilt

import (
	"github.com/mikeyhu/mekkanism/interfaces"
	"github.com/mikeyhu/mekkanism/types"
)

func PlusAll(arguments []interfaces.Argument) types.Argtype {
	all := 0
	for _, v := range arguments {
		all += v.GetInteger()
	}
	return types.Argtype{Integer: all}
}

func MinusAll(arguments []interfaces.Argument) types.Argtype {
	var all int
	head := true
	for _, v := range arguments {
		if head {
			all = v.GetInteger()
			head = false
		} else {
			all -= v.GetInteger()
		}
	}
	return types.Argtype{Integer: all}
}
