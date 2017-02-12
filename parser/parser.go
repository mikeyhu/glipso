package parser

import (
	"errors"
	"github.com/mikeyhu/mekkanism/common"
	"github.com/mikeyhu/mekkanism/interfaces"
	"strconv"
	"strings"
	"text/scanner"
)

func Parse(input string) (*common.Expression, error) {
	var s scanner.Scanner
	s.Filename = "input"
	s.Init(strings.NewReader(input))

	tok := s.Scan()
	text := s.TokenText()
	if tok == scanner.EOF {
		return nil, errors.New("Unexpected EOF")
	}
	if text == "(" {
		_, exp, err := parseExpression(s)
		return exp, err
	} else {
		return nil, errors.New("no Expression found")
	}
}

func parseExpression(s scanner.Scanner) (scanner.Scanner, *common.Expression, error) {
	var tok rune
	if tok != scanner.EOF {
		tok := s.Scan()
		functionName := s.TokenText()
		args := []interfaces.Argument{}
		for tok != scanner.EOF {
			tok = s.Scan()
			token := s.TokenText()
			if token == ")" {
				return s, &common.Expression{FunctionName: functionName, Arguments: args}, nil
			} else if token == "(" {
				ms, arg, err := parseExpression(s)
				if err != nil {
					return ms, arg, err
				}
				s = ms
				args = append(args, *arg)
			} else {

				if integer, err := strconv.Atoi(token); err == nil {
					args = append(args, common.I(integer))
				} else {
					args = append(args, common.REF(token))
				}
			}
		}
		return s, nil, errors.New("Expected end of Expression")
	} else {
		return s, nil, errors.New("Unexpected EOF")
	}
}
