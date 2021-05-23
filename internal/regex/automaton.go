package regex

import (
	"github.com/jonpence/regex-go/internal/iset"
	"github.com/jonpence/regex-go/internal/set"
	"github.com/jonpence/regex-go/internal/ideque")

type Automaton struct {
	states  map[int]*State
	start   *State
	current *State
	inputs  set.Set
	count   int
}

func initAutomaton() *Automaton {
	return &Automaton{make(map[int]*State), nil, nil, set.InitSet(), 1}
}

func (a *Automaton) addState(s *State) {
	a.states[s.id] = s
	a.count += 1
}

func (a *Automaton) addEdge(src int, dest int, input string) {
	a.states[src].addNeighbor(dest, input)
}

func (a *Automaton) setStart(s *State) {
	a.start = s
}

func (a *Automaton) setCurrent(s *State) {
	a.current = s
}

func (a Automaton) getState(num int) *State {
	return a.states[num]
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
	newComposites := iset.InitSet()
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

	if newComposites.IsEmpty() {
		return nil
	}

	for newComposite := range newComposites {
		if a.getState(newComposite).terminates {
			terminates = true
		}
	}

	newState := initDFAState(newComposites, a.count)
	a.count += 1

	if terminates {
		newState.setTerminates()
	}

	return newState
}

func (a *Automaton) determinize() *Automaton {
	dfa   := initAutomaton()
	queue := ideque.InitQueue()

	initialState := a.closureOfOn(a.start, "")

	dfa.addState(initialState)
	dfa.setStart(initialState)

	queue.Enqueue(initialState.id)

	for {
		present := false
		currentState, ok := queue.Dequeue()

		if !ok {
			break
		}

		for input := range a.inputs {
			newState := a.closureOfOn(dfa.getState(currentState), input)

			if newState == nil {
				continue
			}

			// TODO Perhaps find a way to optimize this
			for state := range dfa.states {
				if dfa.states[state].composites.Equivalent(newState.composites) {
					present = true
					newState = dfa.states[state]
					break
				}
			}

			if !present {
				dfa.addState(newState)
				queue.Enqueue(newState.id)
			}

			dfa.addEdge(currentState, newState.id, input)
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
