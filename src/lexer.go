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
)

func tokenStr(tok Token) string {
	switch (tok) {
	    case SYMBOL: return "SYMBOL"
	    case STAR:   return "STAR"
	    case BAR:    return "BAR"
	    case DOT:    return "DOT"
	    case LPAREN: return "LPAREN"
	    case RPAREN: return "RPAREN"
	    case END:    return "END"
	    default:     return "ERR"
	}
}

type Lexer struct {
	input string
	char  byte
	index int
}

func (l *Lexer) storeInput(input string) {
	l.index = 0
	l.input = input
}

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
