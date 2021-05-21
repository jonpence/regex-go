package regex

type Regex struct {
	expression string
	compiled   bool
	dfa        *Automaton
	parser     *Parser
}

func initRegex(parser *Parser) Regex {
	return Regex{"", false, initAutomaton(), parser}
}

func newRegex(input string, parser *Parser) (Regex, bool) {
	r := initRegex(parser)

	r.setExpression(input)

	if r.compile() {
		return r, true
	} else {
		return Regex{}, false
	}
}

func (r *Regex) setExpression(expression string) {
	r.expression = expression
	r.compiled = false
}

func (r *Regex) setAutomaton(dfa *Automaton) {
	r.dfa = dfa
}

func (r *Regex) compile() bool {
	if !r.parser.parse(r.expression) {
		return false
	}

	r.dfa = r.parser.nfa.determinize()
	r.compiled = true

	return true
}

func (r Regex) validate(input string) bool {
	if r.compiled {
		return r.dfa.validate(input)
	} else {
		return false
	}
}
