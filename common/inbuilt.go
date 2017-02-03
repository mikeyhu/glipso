package common
import (
	"github.com/mikeyhu/mekkanism/interfaces"
)

func PlusAll(arguments []interfaces.Argument) interfaces.Argument {
	all := I(0)
	for _, v := range arguments {
		all += v.(I)
	}
	return all
}

func MinusAll(arguments []interfaces.Argument) interfaces.Argument {
	var all I
	head := true
	for _, v := range arguments {
		if head {
			all = v.(I)
			head = false
		} else {
			all -= v.(I)
		}
	}
	return all
}
