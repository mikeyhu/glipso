package common

import "github.com/mikeyhu/glipso/interfaces"

type EXPBuilder struct {
	function  interfaces.Type
	arguments []interfaces.Type
}

func EXPBuild(function interfaces.Type) EXPBuilder {
	return EXPBuilder{function: function}
}

func (e EXPBuilder) withArgs(args ...interfaces.Type) EXPBuilder {
	e.arguments = args
	return e
}

func (e EXPBuilder) build() *EXP {
	return &EXP{e.function, e.arguments}
}

type FNBuilder struct {
	arguments  []interfaces.Type
	expression EXPBuilder
}

func FNBuild() FNBuilder {
	return FNBuilder{}
}

func (f FNBuilder) withArgs(args ...interfaces.Type) FNBuilder {
	f.arguments = args
	return f
}

func (f FNBuilder) withEXPBuilder(expression EXPBuilder) FNBuilder {
	f.expression = expression
	return f
}

func (f FNBuilder) build() FN {
	return FN{
		VEC{f.arguments},
		f.expression.build(),
	}
}
