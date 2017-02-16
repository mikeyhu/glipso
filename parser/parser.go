/*
The parser package takes a string (either in memony or in a file) representation of some code and parses it into a tree of Expressions
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
	var tok rune
	if tok != scanner.EOF {
		tok := s.Scan()
		functionName := s.TokenText()
		args := []interfaces.Type{}
		for tok != scanner.EOF {
			tok = s.Scan()
			token := s.TokenText()
			if token == ")" {
				return s, &common.EXP{FunctionName: functionName, Arguments: args}, nil
			} else if token == "(" {
				ms, arg, err := parseExpression(s)
				if err != nil {
					return ms, arg, err
				}
				s = ms
				args = append(args, *arg)
			} else if token == "[" {
				ms, vec, err := parseVector(s)
				if err != nil {
					return ms, nil, err
				}
				s = ms
				args = append(args, *vec)
			} else {

				if integer, err := strconv.Atoi(token); err == nil {
					args = append(args, common.I(integer))
				} else {
					args = append(args, common.REF(token))
				}
			}
		}
		return s, nil, errors.New("Expected end of Expression")
	}
	return s, nil, errors.New("Unexpected EOF while parsing EXP")
}

func parseVector(s scanner.Scanner) (scanner.Scanner, *common.VEC, error) {
	vec := []interfaces.Type{}
	var tok rune
	tok = s.Scan()
	for tok != scanner.EOF {
		token := s.TokenText()
		if token == "]" {
			return s, &common.VEC{vec}, nil
		} else if token == "(" {
			ms, item, err := parseExpression(s)
			if err != nil {
				return ms, nil, err
			}
			s = ms
			vec = append(vec, *item)
		} else {

			if integer, err := strconv.Atoi(token); err == nil {
				vec = append(vec, common.I(integer))
			} else {
				vec = append(vec, common.REF(token))
			}
		}
		tok = s.Scan()
	}
	return s, nil, errors.New("Unexpected EOF while parsing VEC")
}
