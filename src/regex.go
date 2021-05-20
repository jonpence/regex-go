package main

type Regex struct {
	expression string
	compiled   bool
	dfa        *Automaton
}

func initRegex() Regex {
	return Regex{"", false, initAutomaton()}
}

func newRegex(input string) (Regex, bool) {
	r := initRegex()

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
	p := initParser()

	if !p.parse(r.expression) {
		return false
	}

	for _, state := range p.nfa.states {
		state.printNFAState()
	}

	r.dfa = p.nfa.determinize()
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
