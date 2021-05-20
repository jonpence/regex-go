/* lexer.go
 *
 * Implements a lexer which tokenizes regular expressions.
 */
package main

type Token int

const (
	SYMBOL Token = iota + 1
	STAR
	BAR
	DOT
	LPAREN
	RPAREN
	END
	DEFAULT
)

/* METHODS */
//  -- tokenStr(Token) -> string
//  -- initLexer() -> Lexer
//  -- *Lexer.storeInput(string)
//  -- *Lexer.nextToken() -> Token


/*******************************************************************************/

/* tokenStr(Token) -> string
/*
/* Takes a token and returns its corresponding spelling as a string.
 */
func tokenStr(tok Token) string {
	switch (tok) {
	    case SYMBOL:  return "SYMBOL"
	    case STAR:    return "STAR"
	    case BAR:     return "BAR"
	    case DOT:     return "DOT"
	    case LPAREN:  return "LPAREN"
	    case RPAREN:  return "RPAREN"
		case END:     return "END"
		case DEFAULT: return "DEFAULT"
	    default:      return "ERR"
	}
}

type Lexer struct {
	input string
	char  byte
	index int
}

/* initLexer() -> Lexer
/*
/* Returns a new Lexer.
 */
func initLexer() Lexer {
	return Lexer{"", 0,  0}
}

/* *Lexer.storeInput(string)
/*
/* Stores the string parameter into the Lexer's input field.
 */
func (l *Lexer) storeInput(input string) {
	l.index = 0
	l.input = input
}

/* *Lexer.nextToken() -> Token
/*
/* Reads the next character of the input and returns the corresponding token.
 */
func (l *Lexer) nextToken() Token {
	var val Token

	if (l.index == len(l.input)) {
		return END
	}

	l.char = l.input[l.index]

	switch (l.char) {
		case '(': val = LPAREN
		case ')': val = RPAREN
		case '*': val = STAR
		case '|': val = BAR
		case '.': val = DOT
		default:  val = SYMBOL
	}

	l.index += 1

	return val
}
