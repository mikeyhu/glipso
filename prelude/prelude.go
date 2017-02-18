package prelude

import (
	"fmt"
	"github.com/mikeyhu/glipso/interfaces"
	"github.com/mikeyhu/glipso/parser"
)

func ParsePrelude(scope interfaces.Scope) {
	prelude := `
	(do
		(def defn (fn [n a e] (def n (fn a e))))
	)
	`
	exp, err := parser.Parse(prelude)
	if err != nil {
		panic(fmt.Sprintf("Error parsing prelude, error %v", err))
	}
	exp.Evaluate(scope)
}
