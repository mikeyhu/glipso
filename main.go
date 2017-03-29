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
	debug := flag.Bool("debug", false, "Enable debug output")
	flag.Parse()

	common.DEBUG = *debug
	args := flag.Args()

	env := common.GlobalEnvironment
	prelude.ParsePrelude(env)
	var exp *common.EXP
	if len(args) > 0 {
		file, _ := os.Open(args[0])
		exp, _ = parser.ParseFile(file)
	} else {
		exp, _ = parser.ParseFile(os.Stdin)
	}
	output, _ := exp.Evaluate(env)
	fmt.Println(output)
	common.GlobalEnvironment.DisplayDiagnostics()
}
