# Glipso

A basic Lisp written using Go to help me learn both Go and Lisp.

## Current Features

Glipso has *very few* features. So far it supports the following functions:

```
(= arg...)              return true if all arguments are equal, otherwise false
(+ arg...)              sum all arguments
(- arg...)              minus all arguments from the first argument
(apply func list)       apply list of items as arguments to func
(def var exp)           set a variable in the global environment
(if test exp1 exp2)     if test is 'true' evaluate exp1, otherwise evaluate exp2
(cons arg list?)        add arg to beginning of list. If list is not provided then creates a new list
(first list)            get first element in list
(tail list)             get tail of the list
(do exp...)             run the expressions in order
```

## Types

Glipso internally supports the following types:
```
I   integer
B   boolean
P   pair/list
EXP expression
REF reference
```

## Building
```bash
#run all tests
go test ./...

#build
go build

#test, build & run all acceptance tests
./precommit.sh
```

## Running
```bash
./glipso acceptance/summing-numbers.glipso

#or

echo "(+ 1 2 3)" | ./glipso
```