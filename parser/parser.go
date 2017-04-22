/*
Package parser takes a string (either in memony or in a file) representation of some code and parses it into a tree of Expressions
*/
package parser

import (
	"bufio"
	"errors"
	"github.com/mikeyhu/glipso/common"
	"github.com/mikeyhu/glipso/interfaces"
	"os"
	"strconv"
	"strings"
)

//Parse parses a string containing some code and returns an EXP that represents it
func Parse(input string) (*common.EXP, error) {
	s := bufio.NewScanner(strings.NewReader(input))
	s.Split(tokenize)
	return root(s)
}

//ParseFile parses code from the provided file and returns an EXP that represents it
func ParseFile(inputFile *os.File) (*common.EXP, error) {
	s := bufio.NewScanner(inputFile)
	s.Split(tokenize)
	return root(s)
}

func root(s *bufio.Scanner) (*common.EXP, error) {
	more := s.Scan()
	text := s.Text()
	if !more {
		return nil, errors.New("Unexpected EOF")
	}
	if text == "(" {
		_, exp, err := parseExpression(s)
		return exp, err
	}
	return nil, errors.New("no EXP found")
}

//
func parseExpression(s *bufio.Scanner) (*bufio.Scanner, *common.EXP, error) {
	args := []interfaces.Type{}
	more := true
	var err error
	for more {
		more = s.Scan()
		token := s.Text()
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

func parseVector(s *bufio.Scanner) (*bufio.Scanner, *common.VEC, error) {
	vec := []interfaces.Type{}
	more := true
	var err error
	for more {
		more = s.Scan()
		token := s.Text()
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

func addElementToArray(s *bufio.Scanner, list []interfaces.Type, token string) (*bufio.Scanner, []interfaces.Type, error) {
	var err error
	if token == "(" {
		var exp *common.EXP
		s, exp, err = parseExpression(s)
		if err != nil {
			return s, nil, err
		}
		return s, append(list, exp), nil
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
		t, err := parseTokenToType(token)
		if err != nil {
			return s, nil, err
		}
		list = append(list, t)
	}
	return s, list, nil
}

func parseTokenToType(token string) (interfaces.Type, error) {
	if token[0] == '"' {
		str, err := strconv.Unquote(token)
		if err != nil {
			return nil, err
		}
		return common.S(str), nil
	}
	if token[0] == ':' {
		return common.SYM(token), nil
	}
	if integer, err := strconv.Atoi(token); err == nil {
		return common.I(integer), nil
	}
	if float, err := strconv.ParseFloat(token, 64); err == nil {
		return common.F(float), nil
	}
	if b, err := strconv.ParseBool(token); err == nil {
		return common.B(b), nil
	}
	return common.REF(token), nil

}
