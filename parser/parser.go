/*
Package parser takes a string (either in memony or in a file) representation of some code and parses it into a tree of Expressions
*/
package parser

import (
	"errors"
	"github.com/mikeyhu/glipso/common"
	"github.com/mikeyhu/glipso/interfaces"
	"os"
	"strconv"
	"strings"
	"text/scanner"
)

//Parse parses a string containing some code and returns an EXP that represents it
func Parse(input string) (*common.EXP, error) {
	var s scanner.Scanner
	s.Filename = "input"
	s.Init(strings.NewReader(input))
	return parseRoot(s)
}

//ParseFile parses code from the provided file and returns an EXP that represents it
func ParseFile(inputFile *os.File) (*common.EXP, error) {
	var s scanner.Scanner
	s.Init(inputFile)
	return parseRoot(s)
}

func parseRoot(s scanner.Scanner) (*common.EXP, error) {
	tok := s.Scan()
	text := s.TokenText()
	if tok == scanner.EOF {
		return nil, errors.New("Unexpected EOF")
	}
	if text == "(" {
		_, exp, err := parseExpression(s)
		return exp, err
	}
	return nil, errors.New("no EXP found")
}

func parseExpression(s scanner.Scanner) (scanner.Scanner, *common.EXP, error) {
	args := []interfaces.Type{}
	var tok rune
	var err error
	for tok != scanner.EOF {
		tok = s.Scan()
		token := s.TokenText()
		if token == ")" {
			head := args[0]
			tail := args[1:]
			return s, &common.EXP{Function: head, Arguments: tail}, nil
		}
		s, args, err = addElementToArray(s, args, token)
		if err != nil {
			return s, nil, err
		}
	}
	return s, nil, errors.New("Unexpected EOF while parsing EXP")
}

func parseVector(s scanner.Scanner) (scanner.Scanner, *common.VEC, error) {
	vec := []interfaces.Type{}
	var tok rune
	var err error
	for tok != scanner.EOF {
		tok = s.Scan()
		token := s.TokenText()
		if token == "]" {
			return s, &common.VEC{vec}, nil
		}
		s, vec, err = addElementToArray(s, vec, token)
		if err != nil {
			return s, nil, err
		}
	}
	return s, nil, errors.New("Unexpected EOF while parsing VEC")
}

func addElementToArray(s scanner.Scanner, list []interfaces.Type, token string) (scanner.Scanner, []interfaces.Type, error) {
	var err error
	if token == "(" {
		var exp *common.EXP
		s, exp, err = parseExpression(s)
		if err != nil {
			return s, nil, err
		}
		return s, append(list, *exp), nil
	}
	if token == "[" {
		var vec *common.VEC
		s, vec, err = parseVector(s)
		if err != nil {
			return s, nil, err
		}
		return s, append(list, *vec), nil
	}
	if len(token) > 0 {
		if token[0] == '"' {
			list = append(list, common.S(token[1:len(token)-1]))
		} else if integer, err := strconv.Atoi(token); err == nil {
			list = append(list, common.I(integer))
		} else {
			list = append(list, common.REF(token))
		}
	}
	return s, list, nil
}
