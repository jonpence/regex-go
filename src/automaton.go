package main

import (
	"fmt"
)


/* STATE */

type State struct {
	composites     []*Node
	name           string
	transitions    map[string]string
	terminates     bool
}

func initState(composites []*Node, compositeNames Set) State {
	return State{composites, compositeNames.toString(), make(map[string]string), false}
}

func (s *State) addTransition(dest string, input string) {
	s.transitions[input] = dest
}

func (s *State) setTerminates() {
	s.terminates = true
}

func (s *State) unsetTerminates() {
	s.terminates = false
}

func getIds(nodes []*Node) []int {
	ids := []int{}

	for _, node := range nodes {
		ids = append(ids, node.id)
	}

	return ids
}

func (s *State) reachableFrom(input string) ([]*Node, Set) {
	composites := []*Node{}
	compositeNames := initSet()

	// initialize composites with all reachable states from s on input
	for _, node := range s.composites {
		composites = append(composites, node.transitionOn(input)...)
	}

	for _, node := range composites {
		compositeNames.add(intToString(node.id))
	}

	// reach all nil-reachable states
	addedNew := true
	for addedNew == true {
		addedNew = false

		for _, node := range composites {
			temp := node.transitionOn("")

			for _, tempNode := range temp {
				tempNodeName := intToString(tempNode.id)

				if !compositeNames.isMember(tempNodeName) {
					addedNew = true
					compositeNames.add(tempNodeName)
				}
			}

			composites = append(composites, node.transitionOn("")...)
		}
	}

	// add rest of names to set, set terminates
	for _, node := range composites {
		compositeNames.add(intToString(node.id))
	}

	return composites, compositeNames
}

func (s State) transition(input string) (string, bool) {
	dest, present := s.transitions[input]

	if present {
		return dest, true
	} else {
		return "", false
	}
}

func (s State) printState() {
	if s.terminates {
		fmt.Printf("%s*\n", s.name)
	} else {
		fmt.Printf("%s\n", s.name)
	}
	if len(s.transitions) == 0 {
		fmt.Print("\tEMPTY\n")
	} else {
		for i := range s.transitions {
			fmt.Printf("\t[%s -> %s]\n", i, s.transitions[i])
		}
	}
}

/* AUTOMATON */

type Automaton struct {
	states  map[string]*State
	start   *State
	current *State
	inputs Set
}

func initAutomaton() Automaton {
	return Automaton{make(map[string]*State), nil, nil, initSet()}
}

func (a *Automaton) addState(s *State) {
	a.states[s.name] = s
}

func (a *Automaton) addEdge(src string, dest string, input string) {
	a.states[src].addTransition(dest, input)
}

func (a *Automaton) setStart(s *State) {
	a.start = s
}

func (a *Automaton) setCurrent(s *State) {
	a.current = s
}

func (a *Automaton) reset() {
	a.setCurrent(a.start)
}

func (a *Automaton) transition(char byte) bool {
	dest, present := a.current.transition(string(char))

	if present {
		a.setCurrent(a.states[dest])
		return true
	}

	return false
}

// TODO Move this to graph, or refactor code to avoid this hack
func (g Graph) containsTerminus(nodes []*Node) bool {
	for _, node := range nodes {
		if node == g.terminal {
			return true
		}
	}

	return false
}

func (a *Automaton) convertFrom(g Graph) {
	a.inputs = g.inputs
	queue := initQueue()

	startState   := initState([]*Node{g.start}, initSetRange(g.start.id, g.start.id + 1))
	initialStates, initialNames := startState.reachableFrom("")
	initialStates = append(initialStates, g.start)
	initialNames.add(intToString(g.start.id))
	initialState := initState(initialStates, initialNames)

	if g.containsTerminus(initialStates) {
		initialState.setTerminates()
	}

	a.addState(&initialState)
	a.setStart(&initialState)
	queue.enqueue(initialState.name)

	for {
		currentStateName, ok := queue.dequeue()

		if !ok {
			break
		}

		for input := range a.inputs {
			newState := initState(a.states[currentStateName].reachableFrom(input))

			if newState.name == "{}" {
				continue
			}

			if _, present := a.states[newState.name]; !present {
				a.addState(&newState)
				queue.enqueue(newState.name)
			}

			if g.containsTerminus(newState.composites) {
				newState.setTerminates()
			}

			a.addEdge(currentStateName, newState.name, input)
		}
	}

	a.setCurrent(a.start)
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
