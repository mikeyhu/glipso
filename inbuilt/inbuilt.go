package inbuilt

import (
	"github.com/mikeyhu/mekkanism/interfaces"
	"github.com/mikeyhu/mekkanism/types"
)

func PlusAll(arguments []interfaces.Argument) interfaces.Argument {
	all := types.I(0)
	for _, v := range arguments {
		all += v.(types.I)
	}
	return all
}

func MinusAll(arguments []interfaces.Argument) interfaces.Argument {
	var all types.I
	head := true
	for _, v := range arguments {
		if head {
			all = v.(types.I)
			head = false
		} else {
			all -= v.(types.I)
		}
	}
	return all
}
