package parser

import (
	"errors"
	"unicode/utf8"
)

type token int

var DELIMITERS = []rune{'(', ')', '[', ']'}

func Tokenize(data []byte, atEOF bool) (advance int, token []byte, err error) {
	start := 0
	for width := 0; start < len(data); start += width {
		var r rune
		r, width = utf8.DecodeRune(data[start:])
		if !isSpace(r) {
			break
		}
	}
	char, width := utf8.DecodeRune(data[start:])
	if isDelimiter(char) {
		return start + width, data[start : start+width], nil
	}
	if isStringDelimiter(char) {
		for width, i := 0, start+1; i < len(data); i += width {
			var r rune
			r, width = utf8.DecodeRune(data[i:])
			if isEscape(r) {
				i += width
				_, width = utf8.DecodeRune(data[i:])
			} else if isStringDelimiter(r) {
				return i + width, data[start : i+width], nil
			}
		}
		return len(data), nil, errors.New("string not closed")
	}
	for width, i := 0, start; i < len(data); i += width {
		var r rune
		r, width = utf8.DecodeRune(data[i:])
		if isSpaceOrDelimiter(r) {
			return i, data[start:i], nil
		}
	}
	if atEOF && len(data) > start {
		return len(data), data[start:], nil
	}
	return start, nil, nil
}

func isStringDelimiter(r rune) bool {
	return r == '"'
}

func isEscape(r rune) bool {
	return r == '\\'
}

func isSpace(r rune) bool {
	if r <= '\u00FF' {
		switch r {
		case ' ', '\t', '\n', '\v', '\f', '\r':
			return true
		case '\u0085', '\u00A0':
			return true
		}
		return false
	}
	if '\u2000' <= r && r <= '\u200a' {
		return true
	}
	switch r {
	case '\u1680', '\u2028', '\u2029', '\u202f', '\u205f', '\u3000':
		return true
	}
	return false
}

func isSpaceOrDelimiter(r rune) bool {
	return isSpace(r) || isDelimiter(r)
}

func isDelimiter(r rune) bool {
	for _, d := range DELIMITERS {
		if d == r {
			return true
		}
	}
	return false
}
