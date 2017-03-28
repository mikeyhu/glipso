package parser

import "unicode/utf8"

type token int

var DELIMITERS = []rune{'(', ')', '[', ']'}

func ScanTokens(data []byte, atEOF bool) (advance int, token []byte, err error) {
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
