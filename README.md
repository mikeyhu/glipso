# Glipso

A basic Lisp written using Go to help me learn both Go and Lisp.

## Current Features

Glipso has *very few* features. So far it supports the following functions:

```
(= arg...)              return true if all arguments are equal, otherwise false
(+ arg...)              sum all arguments
(- arg...)              minus all arguments from the first argument
(apply func list)       apply list of items as arguments to func
(cons arg list?)        add arg to beginning of list. If list is not provided then creates a new list
(def var exp)           set a variable in the global environment
(defn name [args] exp)  performs 'def' and 'fn' functions together
(do exp...)             run the expressions in order
(filter fn list)        filter out items in a list by applying fn to them and dropping false responses
(first list)            get first element in list
(fn [args] exp)         creates a function that accepts n arguments are an expression
(if test exp1 exp2)     if test is 'true' evaluate exp1, otherwise evaluate exp2
(map fn list)           generate a new list by applying fn to each element in a list
(range start end)       creates a lazily evaluated list from start to end (inclusive)
(tail list)             get tail of the list
```

### Types

Glipso internally supports the following types:
```
I   integer
B   boolean
P   pair/list
EXP expression
REF reference
VEC Vector
S   String
```

### Example Code
```
	(do
		(def add1 (fn [a] (+ 1 a)))
		(add1 5)
	)
```

### Building
```bash
#run all tests
go test ./...

#build
go build

#test, build & run all acceptance tests
./precommit.sh
```

### Running
```bash
./glipso examples/summing-range.glipso

#or

echo "(+ 1 2 3)" | ./glipso
```

## Roadmap

In no particular order:

* make `map` and `filter` functions work with lazy lists
* add function to create a lazy pair
* make list functions work on lazy, non-lazy and vector lists
* add initial set of functions written in Lisp, i.e. some kind of Prelude
* support some kind of HashMap datatype
* support for more datatypes, i.e. Decimal
* support functions such as `<`, `>` & `%`
* investigate the need for a `nil` datatype (or maybe just empty list?)
* better parser support so names like `some-function` can be used
* implement some goroutine support to push expressions onto other threads and receive notifications when complete
* macro implementation to reduce some special case code

