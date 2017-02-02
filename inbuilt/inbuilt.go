package inbuilt

import "github.com/mikeyhu/mekkanism/types"

func PlusAll (arguments []types.Argtype) types.Argtype {
	all := 0
	for _, v := range arguments {
		all += v.Integer
	}
	return types.Argtype{Integer:all}
}

func MinusAll (arguments []types.Argtype) types.Argtype {
	var all int
	head := true
	for _,v := range arguments {
		if head {
			all = v.Integer
			head = false
		} else {
			all -= v.Integer
		}
	}
	return types.Argtype{Integer:all}
}