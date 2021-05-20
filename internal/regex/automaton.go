package regex

import (
	"github.com/jonpence/regex-go/internal/set"
	"github.com/jonpence/regex-go/internal/deque"
)

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

/* AUTOMATON */

type Automaton struct {
	states  map[string]*State
	start   *State
	current *State
	inputs  set.Set
	count   int
}

func initAutomaton() *Automaton {
	return &Automaton{make(map[string]*State), nil, nil, set.InitSet(), 0}
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
	newComposites := set.InitSet()
	terminates    := false

	// initialize composites with all reachable states from state on input
	for composite := range state.composites {
		neighbors, _ := a.getState(composite).getNeighborList(input)
		newComposites.Multiadd(neighbors)
	}

	// reach all nil-reachable states
	addedNew := true
	for addedNew == true {
		addedNew = false

		for newComposite := range newComposites {
			nullReachable, _ := a.getState(newComposite).getNeighborList("")

			for _, nullReached := range nullReachable {
				if !newComposites.IsMember(nullReached) {
					addedNew = true
					newComposites.Add(nullReached)
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

	return newState
}

func (a *Automaton) determinize() *Automaton {
	dfa   := initAutomaton()
	queue := deque.InitQueue()

	initialState := a.closureOfOn(a.start, "")

	dfa.addState(initialState)
	dfa.setStart(initialState)

	queue.Enqueue(initialState.name)

	for {
		currentState, ok := queue.Dequeue()

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
				queue.Enqueue(newState.name)
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
