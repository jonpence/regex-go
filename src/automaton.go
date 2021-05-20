package main

import (
	"fmt"
)

// MISC METHODS
// getIds

// STATE METHODS
// initState
// addNeighbor
// setTerminates
// unsetTerminates
// getNeighborList
// getNeighbor
// reachableFrom
// printState

// AUTOMATON METHODS
// initAutomaton
// addState
// addEdge
// setStart
// setCurrent
// reset
// transition
// convertFrom
// validate

/* STATE */

type State struct {
	name           string
	neighbors      map[string][]string
	composites     Set
	terminates     bool
	dfaState       bool
}

func initNFAState(name string) State {
	return State{name, make(map[string][]string), initSet(), false, false}
}

func initDFAState(composites Set) State {
	return State{composites.toString(), make(map[string][]string), composites, false, true}
}

func (s *State) addNeighbor(dest string, input string) {
	if s.dfaState {
		s.neighbors[input][0] = dest
	} else {
		s.neighbors[input] = append(s.neighbors[input], dest)
	}
}

func (s *State) setTerminates() {
	s.terminates = true
}

func (s *State) unsetTerminates() {
	s.terminates = false
}

func (s State) getNeighborList(input string) ([]string, bool) {
	neighborList, present := s.neighbors[input]

	if present {
		return neighborList, true
	} else {
		return nil, false
	}
}

func (s State) getNeighbor(input string) (string, bool) {
	neighborList, present := s.getNeighborList(input)

	if present {
		return neighborList[0], true
	} else {
		return "", false
	}
}

func getIds(nodes []*Node) []int {
	ids := []int{}

	for _, node := range nodes {
		ids = append(ids, node.id)
	}

	return ids
}

func (s State) printNFAState() {
	if s.terminates {
		fmt.Printf("%s*", s.name)
	} else {
		fmt.Printf("%s", s.name)
	}
	if len(s.neighbors) == 0 {
		fmt.Print("\tEMPTY\n")
	} else {
		for symbol := range s.neighbors {
			fmt.Print("\t[")
			if symbol == "" {
				fmt.Print("ε")
			} else {
				fmt.Print(symbol)
			}
			fmt.Print(" -> ")
			for i := 0; i < len(s.neighbors[symbol]) - 1; i++ {
				fmt.Printf("%s, ", s.neighbors[symbol][i])
			}
			fmt.Printf("%s]\n", s.neighbors[symbol][len(s.neighbors[symbol]) - 1])
		}
	}
}

func (s State) printDFAState() {
	if s.terminates {
		fmt.Printf("%s*\n", s.name)
	} else {
		fmt.Printf("%s\n", s.name)
	}
	if len(s.neighbors) == 0 {
		fmt.Print("\tEMPTY\n")
	} else {
		for symbol := range s.neighbors {
			fmt.Printf("\t[%s -> %s]\n", symbol, s.neighbors[symbol][0])
		}
	}
}

/* AUTOMATON */

type Automaton struct {
	states  map[string]*State
	start   *State
	current *State
	inputs  Set
	count   int
}

func initAutomaton() *Automaton {
	return &Automaton{make(map[string]*State), nil, nil, initSet(), 0}
}

func (a *Automaton) addState(s *State) {
	a.states[s.name] = s
	a.count += 1
}

func (a *Automaton) addEdge(src string, dest string, input string) {
	a.states[src].addNeighbor(dest, input)
}

func (a *Automaton) setStart(s *State) {
	a.start = s
}

func (a *Automaton) setCurrent(s *State) {
	a.current = s
}

func (a Automaton) getState(input string) *State {
	return a.states[input]
}

func (a Automaton) getCount() int {
	return a.count
}

func (a *Automaton) reset() {
	a.setCurrent(a.start)
}

func (a *Automaton) transition(char byte) bool {
	dest, present := a.current.getNeighbor(string(char))

	if present {
		a.setCurrent(a.states[dest])
		return true
	}

	return false
}

func (a *Automaton) closureOfOn(state *State, input string) *State {
	newComposites := initSet()
	terminates    := false

	// initialize composites with all reachable states from state on input
	for composite := range state.composites {
		neighbors, _ := a.getState(composite).getNeighborList(input)
		newComposites.multiadd(neighbors)
	}

	// reach all nil-reachable states
	addedNew := true
	for addedNew == true {
		addedNew = false

		for newComposite := range newComposites {
			nullReachable, _ := a.getState(newComposite).getNeighborList("")

			for _, nullReached := range nullReachable {
				if !newComposites.isMember(nullReached) {
					addedNew = true
					newComposites.add(nullReached)
				}
			}
		}
	}

	for newComposite := range newComposites {
		if a.getState(newComposite).terminates {
			terminates = true
		}
	}

	newState := initDFAState(newComposites)

	if terminates {
		newState.setTerminates()
	}

	return &newState
}

func (a *Automaton) determinize() *Automaton {
	dfa   := initAutomaton()
	queue := initQueue()

	initialState := a.closureOfOn(a.start, "")

	dfa.addState(initialState)
	dfa.setStart(initialState)

	queue.enqueue(initialState.name)

	for {
		currentState, ok := queue.dequeue()

		if !ok {
			break
		}

		for input := range a.inputs {
			newState := a.closureOfOn(dfa.getState(currentState), input)

			if newState.name == "{}" {
				continue
			}

			if _, present := a.states[newState.name]; !present {
				dfa.addState(newState)
				queue.enqueue(newState.name)
			}

			dfa.addEdge(currentState, newState.name, input)
		}
	}

	return dfa
}

func (a Automaton) validate(input string) bool {
	a.reset()

	for i := 0; i < len(input); i++ {
		if !a.transition(input[i]) {
			return false
		}
	}

	return a.current.terminates
}
