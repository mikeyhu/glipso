# Glipso

A basic Lisp written using Go to help me learn both Go and Lisp.

## Current Features

Glipso has *very few* features. So far it supports the following functions:

```
(= arg...)                  return true if all arguments are equal, otherwise false
(+ arg...)                  sum all arguments
(- arg...)                  minus all arguments from the first argument
(apply func list)           apply list of items as arguments to func
(assoc hash key val ...)    creates a new hash map that combines the original whash map with provided new key value pairs
(cons arg list?)            add arg to beginning of list. If list is not provided then creates a new list
(def var exp)               set a variable in the global environment
(defn name [args] exp)      performs 'def' and 'fn' functions together
(defmacro name [args] exp)  performs 'def' and 'macro' functions together
(do exp...)                 run the expressions in order
(empty list)                returns true if a list is empty
(filter fn list)            filter out items in a list by applying fn to them and dropping false responses
(first list)                get first element in list
(fn [args] exp)             creates a function that accepts n arguments are an expression
(hash-map key val ...)      creates a hashmap with the provided key value pairs
(if test exp1 exp2)         if test is 'true' evaluate exp1, otherwise evaluate exp2
(last list)                 returns the last value in list
(lazypair a b)              returns a pair with head 'a' that will evaluate 'b' lazily to generate a tail
(let [arg pairs] exp)       creates a new scope for exp in which arg pairs have been evaluated and put into scope
(macro [args] exp)          creates a macro that will replace args in the exp with arguments provided for evaluation
(map fn list)               generate a new list by applying fn to each element in a list
(panic message)             exit with a message
(range start end)           creates a lazily evaluated list from start to end (inclusive)
(repeat item times)         returns a list consisting of times number of items 
(tail list)                 get tail of the list
(take num list)             returns a lazily evaluated list that is the first 'num' elements in 'list'
```

### Example Code : A lazy list of primes
```lisp
(do
    (defn notdivbyany [num listofdivs]
        (empty
            (filter
                (fn [z] (= 0 z))
                (map (fn [head] (% num head)) listofdivs)
            )
        )
    )

    (defn getprimes [num listofprimes]
        (if
            (notdivbyany num listofprimes)
            (lazypair num (getprimes (+ num 1) (cons num listofprimes)))
            (getprimes (+ num 1) listofprimes)
        )
    )

    (apply print (take 100 (getprimes 3 (cons 2))))
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

#benchmark acceptance tests
go test -bench=.
```

### Running
```bash
./glipso examples/summing-range.glipso

#or

echo "(+ 1 2 3)" | ./glipso
```

### Types

Glipso internally supports the following types:
```
B       boolean
EXP     expression
I       integer
F       float
LAZYP   lazily evaluated pair
MAC     Macro
P       pair/list
REF     reference
S       String
VEC     Vector
```

### Interfaces

The codebase uses a number of different interfaces:
```
Type        root type
Appliable   functions (either inbuilt or defined) that can be applied
Comparable  Comparable Values
Equalable   Values with Equality
Evaluatable s-expressions that have an Appliable, arguments and Scope and will return a Value
Expandable  Macro expandable to Evaluatable
Iterable    Values that have Head and Tails
Numeric     Numeric Values
Scope       provided to Appliables alongside arguments to lookup and store values
Value       result of an Evaluatable
```

## Roadmap

In no particular order:

* implicit expression around parsed expressions to cope with multiple sequential expressions
* support some kind of HashMap datatype along with :symbols
* make `map` and `filter` functions work with lazy lists
* make list functions work on lazy, non-lazy and vector lists
* support for more datatypes, i.e. Decimal
* implement some goroutine support to push expressions onto other threads and receive notifications when complete
* improve macro implementation with better substitution options
