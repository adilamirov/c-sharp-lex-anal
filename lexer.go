package main

import (
	"strings"
)

const (
	OPERATORS   = "+-=/*%<>^&|[]()."
	SPACES      = " \n\t"
	SCOPE_SIGNS = "{}"
	DELIMITER   = ";,."
	DIGITS      = "1234567890"
	LETTERS     = "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM"
	KEYWORDS    = "abstract as base bool break byte case catch char checked class const continue decimal default delegate do double else enum event explicit extern false finally fixed float for foreach goto if implicit in int interface internal is lock long namespace new null object operator out override params private protected public readonly ref return sbyte sealed short sizeof stackalloc static string struct switch this throw true try typeof uint ulong unchecked unsafe ushort using using static virtual void volatile while"
)

type TokenType int

const (
	PUNCTUATOR TokenType = 0
	OPERATOR   TokenType = 1
	LITERAL    TokenType = 2
	KEYWORD    TokenType = 3
	IDENTIFIER TokenType = 4
)

type Lexer struct {
	text   []byte
	cursor int
	token  string
	tipe   TokenType
}

func (l *Lexer) LoadText(text []byte) {
	if l != nil {
		l.text = text
		l.cursor = 0
	}
}

func (l *Lexer) NextToken() bool {
	// Skip spaces, tabulation, \n's
	for l.cursor < len(l.text) && l.isSpace() {
		l.cursor++
	}

	if l.cursor >= len(l.text) {
		return false
	}

	if l.isDelimiter() {
		l.token = string(l.text[l.cursor])
		l.tipe = PUNCTUATOR
		l.cursor++
		return true
	}

	if l.isScopeSign() {
		l.token = string(l.text[l.cursor])
		l.tipe = PUNCTUATOR
		l.cursor++
		return true
	}

	// Operator
	if l.isOperator() {
		l.tipe = OPERATOR
		// Operators may consist of two chars
		if l.cursor+1 < len(l.text) && strings.IndexByte("|&+-=<>", l.text[l.cursor+1]) != -1 {
			l.token = string(l.text[l.cursor]) + string(l.text[l.cursor+1])
			l.cursor += 2
			return true
		}
		l.token = string(l.text[l.cursor])
		l.cursor++
		return true
	}

	// Numeric literal
	if l.isDigit() {
		// Parse integral part of a number
		for l.token = ""; l.cursor < len(l.text) && l.isDigit(); l.cursor++ {
			l.token += string(l.text[l.cursor])
		}
		// Check if it's float
		if l.cursor < len(l.text) && l.text[l.cursor] == '.' {
			l.token += "."
			l.cursor++
			// Parse fractional part
			for ; l.cursor < len(l.text) && l.isDigit(); l.cursor++ {
				l.token += string(l.text[l.cursor])
			}
			// Numbers may have some modifiers e.g. 3.14f (f is modifier)
			for ; l.cursor < len(l.text) && l.isLetter(); l.cursor++ {
				l.token += string(l.text[l.cursor])
			}
		}
		l.tipe = LITERAL
		return true
	}

	l.token = ""
	// If it's string-literal it starts and ends with `"`
	if l.text[l.cursor] == '"' {
		l.token = "\""
		l.cursor++
		// Note that we need to ignore `\"` combinations
		for l.cursor < len(l.text) && (l.text[l.cursor] != '"' || l.text[l.cursor-1] == '\\') {
			l.token += string(l.text[l.cursor])
			l.cursor++
		}
		l.token += string('"')
		l.cursor++
		l.tipe = LITERAL
		return true
	}
	// Keyword or identifier
	for l.cursor < len(l.text) && !l.isSpace() && !l.isDelimiter() && !l.isScopeSign() && !l.isOperator() {
		l.token += string(l.text[l.cursor])
		if (strings.Index(KEYWORDS, l.token) != -1) {
			l.tipe = KEYWORD
		} else {
			l.tipe = IDENTIFIER
		}
		l.cursor++
	}
	return true
}

func (l *Lexer) GetToken() (token string, err error) {
	return l.token, nil
}

func (l *Lexer) isSpace() bool {
	return strings.IndexByte(SPACES, l.text[l.cursor]) != -1
}

func (l *Lexer) isDelimiter() bool {
	return strings.IndexByte(DELIMITER, l.text[l.cursor]) != -1
}

func (l *Lexer) isScopeSign() bool {
	return strings.IndexByte(SCOPE_SIGNS, l.text[l.cursor]) != -1
}

func (l *Lexer) isOperator() bool {
	return strings.IndexByte(OPERATORS, l.text[l.cursor]) != -1
}

func (l *Lexer) isDigit() bool {
	return strings.IndexByte(DIGITS, l.text[l.cursor]) != -1
}

func (l *Lexer) isLetter() bool {
	return strings.IndexByte(LETTERS, l.text[l.cursor]) != -1
}
