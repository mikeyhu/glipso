package main

import (
	"fmt"
	"github.com/mikeyhu/glipso/common"
	"github.com/mikeyhu/glipso/parser"
	"github.com/mikeyhu/glipso/prelude"
	"os"
)

func main() {
	env := common.GlobalEnvironment
	prelude.ParsePrelude(env)
	args := os.Args[1:]
	var exp *common.EXP
	if len(args) > 0 {
		file, _ := os.Open(args[0])
		exp, _ = parser.ParseFile(file)
	} else {
		exp, _ = parser.ParseFile(os.Stdin)
	}
	fmt.Println(exp.Evaluate(env))
}
