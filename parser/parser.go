package parser

import (
	"errors"
	"fmt"
	"github.com/mikeyhu/mekkanism/interfaces"
	"github.com/mikeyhu/mekkanism/common"
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
	pos := s.Pos()
	if tok == scanner.EOF {
		return nil, errors.New("Unexpected EOF")
	}
	fmt.Println("At position", pos, ":", text)
	if text == "(" { //new Expression
		return parseExpression(s)
	} else {
		return nil, errors.New("no Expression found")
	}

}

func parseExpression(s scanner.Scanner) (*common.Expression, error) {
	var tok rune
	if tok != scanner.EOF {
		tok := s.Scan()
		functionName := s.TokenText()
		fmt.Println("FunctionName: ", functionName)
		args := []interfaces.Argument{}
		for tok != scanner.EOF {
			tok = s.Scan()
			token := s.TokenText()
			if token == ")" {
				return &common.Expression{FunctionName: functionName, Arguments: args}, nil
			}
			fmt.Println("Argument: ", token)
			integer, _ := strconv.Atoi(token)
			args = append(args, common.I(integer))
		}
		return nil, errors.New("Expected end of Expression")
	} else {
		return nil, errors.New("Unexpected EOF")
	}
}
