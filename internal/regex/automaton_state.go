package regex

import (
	"fmt"
	"github.com/jonpence/regex-go/internal/set"
)

// STATE METHODS
// initState
// addNeighbor
// setTerminates
// unsetTerminates
// getNeighborList
// getNeighbor
// reachableFrom
// printState

type State struct {
	name           string
	neighbors      map[string][]string
	composites     set.Set
	terminates     bool
	dfaState       bool
}

func initNFAState(name string) *State {
	return &State{name, make(map[string][]string), set.InitSet(), false, false}
}

func initDFAState(composites set.Set) *State {
	return &State{composites.ToString(), make(map[string][]string), composites, false, true}
}

func (s *State) addNeighbor(dest string, input string) {
	if !s.dfaState || len(s.neighbors[input]) == 0 {
		s.neighbors[input] = append(s.neighbors[input], dest)
	} else {
		s.neighbors[input][0] = dest
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
