package prelude

import (
	"fmt"
	"github.com/mikeyhu/glipso/interfaces"
	"github.com/mikeyhu/glipso/parser"
)

// ParsePrelude loads a number of definitions such as functions into global scope
func ParsePrelude(scope interfaces.Scope) {
	code := `
	(do
		(def defmacro (macro [n a e] (def n (macro a e))))
		(defmacro defn [nn aa ee] (def nn (fn aa ee)))

		(defn last [list]
			(if
				(empty (tail list))
				(first list)
				(last (tail list))))

		(defn repeat [item times]
			(if
				(> times 1)
				(lazypair item (repeat item (- times 1)))
				(cons item)))
	)
	`
	exp, err := parser.Parse(code)
	if err != nil {
		panic(fmt.Sprintf("Error parsing prelude, error %v", err))
	}
	_, err = exp.Evaluate(scope)
	if err != nil {
		panic(fmt.Sprintf("Error evaluating prelude, error %v", err))
	}
}
