package fsm

import (
	"fmt"
)



// A graph holds its size and an adjacency matrix. The cells of the matrix hold runes,
// which correspond to acceptable chars for transitioning.
type FiniteStateMachine struct {
	size      int
	matrix    [][][]rune
	terminals []int
}

// Initializes a FiniteStateMachine whose cells are all empty.
func InitFSM(size int) FiniteStateMachine {
	fsm := FiniteStateMachine{}
	fsm.size = size
	fsm.matrix = make([][][]rune, fsm.size)

	for i := 0; i < fsm.size; i++ {
		fsm.matrix[i] = make([][]rune, fsm.size)
		for j := 0; j < fsm.size; j++ {
			fsm.matrix[i][j] = []rune{}
		}
	}

	return fsm
}

// Prints a FiniteStateMachine's adjacency matrix.
func (fsm FiniteStateMachine) PrintFSM() {
	for _, col := range fsm.matrix {
		for _, cell := range col {
			fmt.Printf("%c ", cell)
		}
		fmt.Println()
	}
}

// Gets the number of states within the FSM.
func (fsm *FiniteStateMachine) GetSize() int {
	return fsm.size
}

// Sets the slice of transitional characters of some cell in the adjacency matrix.
func (fsm *FiniteStateMachine) SetTransition(x int, y int, newTransitions []rune) {
	fsm.matrix[x][y] = newTransitions
}

// Returns the transitional characters of a cell in the adjacency matrix.
func (fsm FiniteStateMachine) GetTransition(x int, y int) []rune {
	return fsm.matrix[x][y]
}

// Sets the terminal states of the FiniteStateMachine.
func (fsm *FiniteStateMachine) SetTerminals(newTerminals []int) {
	fsm.terminals = newTerminals
}

// Returns the terminals of the FiniteStateMachine.
func (fsm FiniteStateMachine) GetTerminals() []int {
	return fsm.terminals
}

// Given a byte and a slice of bytes, this function returns whether or not the value of
// the byte exists within the slice.
func member(char rune, chars []rune) bool {
	for _, c := range chars {
		if char == c {
			return true
		}
	}

	return false
}

// Attempts a transition. If it succeeds, then return true and the id of the next state.
// Otherwise return false and the id of the state parameter.
func (fsm FiniteStateMachine) transition(state int, char rune) (bool, int) {
	for nextState, chars := range fsm.matrix[state] {
		if member(char, chars) {
			return true, nextState
		}
	}

	return false, state
}

// Determines whether or not a given state id is a valid terminal state id.
func (fsm FiniteStateMachine) isTerminal(state int) bool {
	for _, terminalState := range fsm.terminals {
		if state == terminalState {
			return true
		}
	}

	return false
}

// Determines whether or not a given word is within the language defined by the FSM.
func (fsm FiniteStateMachine) ValidateWord(word string) bool {
	state := 0
	valid := true

	for _, char := range word {
		valid, state = fsm.transition(state, char)
		if !valid {
			return false
		}
	}

	return fsm.isTerminal(state)
}
