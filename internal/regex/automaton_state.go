package regex

import (
	"fmt"
	"github.com/jonpence/regex-go/internal/iset"
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
	id             int
	neighbors      map[string][]int
	composites     iset.Set
	terminates     bool
	dfaState       bool
}

func initNFAState(id int) *State {
	return &State{id, make(map[string][]int), iset.InitSet(), false, false}
}

func initDFAState(composites iset.Set, id int) *State {
	return &State{id, make(map[string][]int), composites, false, true}
}

func (s *State) addNeighbor(dest int, input string) {
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

func (s State) getNeighborList(input string) ([]int, bool) {
	neighborList, present := s.neighbors[input]

	if present {
		return neighborList, true
	} else {
		return nil, false
	}
}

func (s State) getNeighbor(input string) (int, bool) {
	neighborList, present := s.getNeighborList(input)

	if present {
		return neighborList[0], true
	} else {
	return 0, false
	}
}

func (s State) printNFAState() {
	if s.terminates {
		fmt.Printf("%d*", s.id)
	} else {
		fmt.Printf("%d", s.id)
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
				fmt.Printf("%d, ", s.neighbors[symbol][i])
			}
			fmt.Printf("%d]\n", s.neighbors[symbol][len(s.neighbors[symbol]) - 1])
		}
	}
}

func (s State) printDFAState() {
	if s.terminates {
		fmt.Printf("%d*\n", s.id)
	} else {
		fmt.Printf("%d\n", s.id)
	}
	if len(s.neighbors) == 0 {
		fmt.Print("\tEMPTY\n")
	} else {
		for symbol := range s.neighbors {
			fmt.Printf("\t[%s -> %d]\n", symbol, s.neighbors[symbol][0])
		}
	}
}
