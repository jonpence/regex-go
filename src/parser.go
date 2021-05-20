/* parser.go
/*
/* Implements a recursive descent parser for regular expressions.
 */

package main

import (
	"fmt"
)

type Parser struct {
	token Token
	lexer Lexer
	depth int
	nfa   *Automaton
	debug bool
}

/* METHODS */
//  -- initParser() -> Parser
//  -- *Parser.setDebug()
//  -- *Parser.unsetDebug()
//  -- Parser.indent()
//  -- Parser.enter(string)
//  -- Parser.leave(string)
//  -- Parser.require(Token) -> bool
//  -- Parser.accept(Token) -> bool
//  -- *Parser.parse(string) -> bool
//  -- *Parser.parseS() -> int, int, bool
//  -- *Parser.parseA() -> int, int, bool
//  -- *Parser.parseB() -> int, int, bool
//  -- *Parser.parseC() -> int, int, bool
//  -- *Parser.parseD() -> int, int, bool
//  -- *Parser.parseE() -> int, int, bool
//  -- *Parser.parseF() -> int, int, bool
//  -- *Parser.or(int, int, int, int) -> int, int
//  -- *Parser.concatenate(int, int, int, int) -> int, int
//  -- *Parser.kleene(int, int) -> int, int
//  -- *Parser.symbol(string) -> int, int

/* initParser() -> Parser
/*
 */
func initParser() Parser {
	return Parser{DEFAULT, initLexer(), 0, nil, false}
}

/* *Parser.setDebug()
/*
 */
func (p *Parser) setDebug() {
	p.debug = true
}

/* *Parser.unsetDebug()
/*
 */
func (p *Parser) unsetDebug() {
	p.debug = false
}

/* Parser.indent()
/*
 */
func (p Parser) indent() {
	for i := 0; i < p.depth; i += 1 {
		fmt.Print(" ")
	}
}

/* Parser.enter(string)
/*
 */
func (p Parser) enter(nt string) {
	p.indent()
	fmt.Println("ENTERED: " + nt)
}

/* Parser.leave(string)
/*
 */
func (p Parser) leave(nt string) {
	p.indent()
	fmt.Println("LEFT: " + nt)
}

/* Parser.require(Token) -> bool
/*
 */
func (p Parser) require(expected Token) bool {
	if p.token == expected {
		return true
	} else {
		return false
	}
}

/* *Parser.accept(Token) -> bool
/*
 */
func (p *Parser) accept(expected Token) bool {
	if p.token == expected {
		if p.debug {
			p.indent()
			fmt.Printf("ACCEPTED %s\n", tokenStr(expected))
		}
		p.token = p.lexer.nextToken()
		return true
	} else {
		return false
	}
}

/* *Parser.parse(string) -> bool
/*
 */
func (p *Parser) parse(input string) bool {
	// init automaton
	p.nfa = initAutomaton()

	// set up lexer
	p.lexer.storeInput(input)
	p.token = p.lexer.nextToken()

	// do parse
	s_start, s_end, _ := p.parseS()

	// if regex input terminates correctly
	if p.accept(END) {
		// debug statements
		if p.debug {
			fmt.Printf("Valid parse with start state of %d\n", s_start)
		}

		// set terminating state
		p.nfa.getState(intToString(s_end)).setTerminates()

		// set start
		p.nfa.setStart(p.nfa.getState(intToString(s_start)))

		// succesful parse
		return true
	} else {

		// debug statements
		if p.debug {
			fmt.Println("Invalid parse.")
		}

		// unsuccessful parse
		return false
	}
}

/* *Parser.parseS() -> int, int, bool
/*
 */
func (p *Parser) parseS() (int, int, bool) {
	if p.debug {
		p.enter("S")
		p.depth += 1
	}

	var n_start, n_end int

	a_start, a_end, a_ok := p.parseA()

	if !a_ok {
		return 0, 0, false
	}

	b_start, b_end, b_ok := p.parseB()

	if !b_ok {
		return 0, 0, false
	}

	if p.debug {
		p.depth -= 1
		p.leave("S")
	}

	if b_start != 0 {
		n_start, n_end = p.or(a_start, b_start, a_end, b_end)
	} else {
		n_start, n_end = a_start, a_end
	}

	return n_start, n_end, true
}

/* *Parser.parseA() -> int, int, bool
/*
 */
func (p *Parser) parseA() (int, int, bool) {
	if p.debug {
		p.enter("A")
		p.depth += 1
	}

	var n_start, n_end int

	c_start, c_end, c_ok := p.parseC();

	if !c_ok {
		return 0, 0, false
	}

	d_start, d_end, d_ok := p.parseD();

	if !d_ok {
		return 0, 0, false
	}

	if d_start != 0 {
		n_start, n_end = p.concatenate(c_start, d_start, c_end, d_end)
	} else {
		n_start, n_end = c_start, c_end
	}

	if p.debug {
		p.depth -= 1
		p.leave("A")
	}

	return n_start, n_end, true
}

/* *Parser.parseB() -> int, int, bool
/*
 */
func (p *Parser) parseB() (int, int, bool) {
	if p.debug {
		p.enter("B")
		p.depth += 1
	}

	var n_start, n_end int

	switch {
		case p.accept(BAR):

			a_start, a_end, a_ok := p.parseA();

			if !a_ok {
				return 0, 0, false
			}

			b_start, b_end, b_ok := p.parseB();

			if !b_ok {
				return 0, 0, false
			}

			if b_start != 0 {
				n_start, n_end = p.or(a_start, b_start, a_end, b_end)
			} else {
				n_start, n_end = a_start, a_end
			}

		default:
			n_start, n_end = 0, 0
	}

	if p.debug {
		p.depth -= 1
		p.leave("B")
	}

	return n_start, n_end, true;
}

/* *Parser.parseC() -> int, int, bool
/*
 */
func (p *Parser) parseC() (int, int, bool) {
	if p.debug {
		p.enter("C")
		p.depth += 1
	}

	var n_start, n_end int

	e_start, e_end, e_ok := p.parseE();

	if !e_ok {
		return 0, 0, false
	}

	f_start, _, f_ok := p.parseF();

	if !f_ok {
		return 0, 0, false
	}

	if f_start == 0 {
		n_start, n_end = p.kleene(e_start, e_end)
	} else {
		n_start, n_end = e_start, e_end
	}

	if p.debug {
		p.depth -= 1
		p.leave("C")
	}

	return n_start, n_end, true;
}

/* *Parser.parseD() -> int, int, bool
/*
 */
func (p *Parser) parseD() (int, int, bool) {
	if p.debug {
		p.enter("D")
		p.depth += 1
	}

	var n_start, n_end int

	switch {
		case p.accept(DOT):
			c_start, c_end, c_ok := p.parseC();

			if !c_ok {
				return 0, 0, false
			}

			d_start, d_end, d_ok := p.parseD();

			if !d_ok {
				return 0, 0, false
			}

			if d_start != 0 {
				n_start, n_end = p.concatenate(c_start, d_start, c_end, d_end)
			} else {
				n_start, n_end = c_start, c_end
			}

		default:
			n_start, n_end = 0, 0
	}

	if p.debug {
		p.depth -= 1
		p.leave("D")
	}

	return n_start, n_end, true;
}

/* *Parser.parseE() -> int, int, bool
/*
 */
func (p *Parser) parseE() (int, int, bool) {
	if p.debug {
		p.enter("E")
		p.depth += 1
	}

	var n_start, n_end int

	switch {
		case p.require(SYMBOL):
			n_start, n_end = p.symbol(string(p.lexer.char))
			p.accept(SYMBOL)
		case p.accept(LPAREN):
			s_start, s_end, s_ok := p.parseS();

			if !s_ok {
				return 0, 0, false
			}

			if !p.accept(RPAREN) {
				return 0, 0, false
			}

			n_start, n_end = s_start, s_end
		default:
			return 0, 0, false
	}

	if p.debug {
		p.depth -= 1
		p.leave("E")
	}

	return n_start, n_end, true;
}

/* *Parser.parseF() -> int, int, bool
/*
 */
func (p *Parser) parseF() (int, int, bool) {
	if p.debug {
		p.enter("F")
		p.depth += 1
	}

	var n_start, n_end int

	switch {
		case p.accept(STAR):
			if _, _, f_ok := p.parseF(); !f_ok {
				return 0, 0, false
			}
			n_start, n_end = 0, 0
		default:
			n_start, n_end = -1, -1
	}

	if p.debug {
		p.depth -= 1
		p.leave("F")
	}

	return n_start, n_end, true;
}


/* *Parser.or(int, int, int, int) -> int, int
/*
 */
func (p *Parser) or(startA int, startB int, endA int, endB int) (int, int) {
	if p.debug {
		fmt.Println("ADDING OR")
	}

	currentCount := p.nfa.getCount()

	ns0 := initNFAState(intToString(currentCount))
	ns1 := initNFAState(intToString(currentCount + 1))
	ns2 := initNFAState(intToString(currentCount + 2))
	ns3 := initNFAState(intToString(currentCount + 3))

	p.nfa.addEdge(ns0.name, ns1.name, "")
	p.nfa.addEdge(ns2.name, ns3.name, "")
	p.nfa.addEdge(ns1.name, intToString(startA), "")
	p.nfa.addEdge(ns1.name, intToString(startB), "")
	p.nfa.addEdge(intToString(endA), ns2.name, "")
	p.nfa.addEdge(intToString(endB), ns2.name, "")

	return currentCount, currentCount + 3
}

/* *Parser.concatenate(int, int, int, int) -> int, int
/*
 */
func (p *Parser) concatenate(startA int, startB int, endA int, endB int) (int, int) {
	if p.debug {
		fmt.Println("ADDING CONCATENATE")
	}
	p.nfa.addEdge(intToString(endA), intToString(startB), "")
	return startA, endB
}

/* *Parser.kleene(int, int) -> int, int
/*
 */
func (p *Parser) kleene(startA int, endA int) (int, int) {
	if p.debug {
		fmt.Println("ADDING KLEENE")
	}
	currentCount := p.nfa.getCount()

	ns0 := initNFAState(intToString(currentCount))
	ns1 := initNFAState(intToString(currentCount + 1))

	p.nfa.addEdge(ns0.name, ns1.name, "")
	p.nfa.addEdge(ns0.name, intToString(startA), "")
	p.nfa.addEdge(intToString(endA), ns0.name, "")
	p.nfa.addEdge(intToString(endA), ns1.name, "")

	return currentCount , currentCount + 1
}

/* *Parser.symbol(string) -> int, int
/*
 */
func (p *Parser) symbol(symbol string) (int, int) {
	if p.debug {
		fmt.Println("ADDING SYMBOL")
	}
	currentCount := p.nfa.getCount()

	ns0 := initNFAState(intToString(currentCount))
	ns1 := initNFAState(intToString(currentCount + 1))

	p.nfa.addEdge(ns0.name, ns1.name, symbol)

	return currentCount, currentCount + 1
}
