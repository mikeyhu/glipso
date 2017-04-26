package main

import (
	"flag"
	"fmt"
	"github.com/mikeyhu/glipso/common"
	"github.com/mikeyhu/glipso/parser"
	"github.com/mikeyhu/glipso/prelude"
	"os"
)

func main() {
	env := common.GlobalEnvironment

	prelude.ParsePrelude(env)

	debug := flag.Bool("debug", false, "Enable debug output")
	flag.Parse()

	common.DEBUG = *debug
	args := flag.Args()

	var exp *common.EXP
	if len(args) > 0 {
		file, _ := os.Open(args[0])
		exp, _ = parser.ParseFile(file)
	} else {
		exp, _ = parser.ParseFile(os.Stdin)
	}
	output, err := exp.Evaluate(env)
	fmt.Println(output)
	if err != nil {
		panic(err)
	}
	common.GlobalEnvironment.DisplayDiagnostics()
}
