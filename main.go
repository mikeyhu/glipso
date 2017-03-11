package main

import (
	"flag"
	"fmt"
	"github.com/mikeyhu/glipso/common"
	"github.com/mikeyhu/glipso/parser"
	"github.com/mikeyhu/glipso/prelude"
	"os"
	"github.com/pkg/profile"
)

func main() {
	debug := flag.Bool("debug", false, "Enable debug output")
	prof := flag.Bool("profile", false, "Enable profiling")
	flag.Parse()
	if *prof {
		defer profile.Start(profile.CPUProfile, profile.ProfilePath(".")).Stop()
	}

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
	fmt.Println(exp.Evaluate(env))

	common.GlobalEnvironment.DisplayDiagnostics()
}
