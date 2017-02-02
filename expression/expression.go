package expression

import "fmt"
import "github.com/mikeyhu/mekkanism/types"
import "github.com/mikeyhu/mekkanism/inbuilt"

type Expression struct {
	FunctionName  string
	File          string
	StartPosition string
	Arguments     []types.Argtype
}

func (exp *Expression) Evaluate() types.Argtype {
	if exp.FunctionName == "+" {
		return inbuilt.PlusAll(exp.Arguments)
	} else if exp.FunctionName == "-" {
		return inbuilt.MinusAll(exp.Arguments)
	} else {
		panic(fmt.Sprintf("Panic - Cannot resolve FunctionName '%s'", exp.FunctionName))
	}
}
