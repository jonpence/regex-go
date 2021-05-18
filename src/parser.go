// TODO Make sure that the regex string contains only ASCII symbols.

package main

import (
	"fmt"
)

type Parser struct {
	token Token
	lexer Lexer
	depth int
	graph Graph
}

func initParser() Parser {
	return Parser{DEFAULT, initLexer(), 0, initGraph()}
}

func (p Parser) require(expected Token) bool {
	if p.token == expected {
		return true
	} else {
		return false
	}
}

func (p *Parser) accept(expected Token) bool {
	if p.token == expected {
		p.indent()
		fmt.Printf("ACCEPTED %s\n", tokenStr(expected))
		p.token = p.lexer.nextToken()
		return true
	} else {
		return false
	}
}

func (p Parser) indent() {
	for i := 0; i < p.depth; i += 1 {
		fmt.Print(" ")
	}
}

func (p Parser) enter(nt string) {
	p.indent()
	fmt.Println("ENTERED: " + nt)
}

func (p Parser) leave(nt string) {
	p.indent()
	fmt.Println("LEFT: " + nt)
}

func (p *Parser) parse(input string) bool {
	p.lexer.storeInput(input)
	p.token = p.lexer.nextToken()

	var s_start int

	s_start, _, _ = p.parseS()
	if p.accept(END) {
		fmt.Printf("Valid parse with start state of %d\n", s_start)
		p.graph.setStart(s_start)
		p.graph.setTerminal(p.graph.findTerminal())
		return true
	} else {
		fmt.Println("Invalid parse.")
		return false
	}
}

func (p *Parser) parseS() (int, int, bool) {
	p.enter("S")
	p.depth += 1

	var n_start, n_end int

	a_start, a_end, a_ok := p.parseA()

	if !a_ok {
		return 0, 0, false
	}

	b_start, b_end, b_ok := p.parseB()

	if !b_ok {
		return 0, 0, false
	}

	p.depth -= 1
	p.leave("S")

	if b_start != 0 {
		n_start, n_end = p.or(a_start, b_start, a_end, b_end)
	} else {
		n_start, n_end = a_start, a_end
	}

	return n_start, n_end, true
}

func (p *Parser) parseA() (int, int, bool) {
	p.enter("A")
	p.depth += 1

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

	p.depth -= 1
	p.leave("A")

	return n_start, n_end, true
}

func (p *Parser) parseB() (int, int, bool) {
	p.enter("B")
	p.depth += 1

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

	p.depth -= 1
	p.leave("B")

	return n_start, n_end, true;
}

func (p *Parser) parseC() (int, int, bool) {
	p.enter("C")
	p.depth += 1

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

	p.depth -= 1
	p.leave("C")

	return n_start, n_end, true;
}

func (p *Parser) parseD() (int, int, bool) {
	p.enter("D")
	p.depth += 1

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

	p.depth -= 1
	p.leave("D")

	return n_start, n_end, true;
}

func (p *Parser) parseE() (int, int, bool) {
	p.enter("E")
	p.depth += 1

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

	p.depth -= 1
	p.leave("E")

	return n_start, n_end, true;
}

func (p *Parser) parseF() (int, int, bool) {
	p.enter("F")
	p.depth += 1

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

	p.depth -= 1
	p.leave("F")

	return n_start, n_end, true;
}


func (p *Parser) or(startA int, startB int, endA int, endB int) (int, int) {
	fmt.Println("ADDING OR")
	origCount := p.graph.count

	p.graph.addNode()
	p.graph.addNode()
	p.graph.addNode()
	p.graph.addNode()

	p.graph.addEdge(origCount + 1, origCount + 2, "")
	p.graph.addEdge(origCount + 3, origCount + 4, "")
	p.graph.addEdge(origCount + 2, startA, "")
	p.graph.addEdge(origCount + 2, startB, "")
	p.graph.addEdge(endA, origCount + 3, "")
	p.graph.addEdge(endB, origCount + 3, "")

	return origCount + 1, origCount + 4
}

func (p *Parser) concatenate(startA int, startB int, endA int, endB int) (int, int) {
	fmt.Println("ADDING CONCATENATE")
	p.graph.addEdge(endA, startB, "")
	return startA, endB
}

func (p *Parser) kleene(startA int, endA int) (int, int) {
	fmt.Println("ADDING KLEENE")
	origCount := p.graph.count

	p.graph.addNode()
	p.graph.addNode()

	p.graph.addEdge(origCount + 1, origCount + 2, "")
	p.graph.addEdge(origCount + 1, startA, "")
	p.graph.addEdge(endA, origCount + 1, "")
	p.graph.addEdge(endA, origCount + 2, "")

	return origCount + 1, origCount + 2
}

func (p *Parser) symbol(symbol string) (int, int) {
	fmt.Println("ADDING SYMBOL")
	origCount := p.graph.count

	p.graph.addNode()
	p.graph.addNode()

	p.graph.addEdge(origCount + 1, origCount + 2, symbol)

	return origCount + 1, origCount + 2
}
