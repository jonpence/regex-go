/*
Package main implements a regular expression engine.
*/
package main

import (
	"fmt"
)

func main() {

	parser := initParser()
	parser.parse("A*B")

	for i := 1; i <= len(parser.graph.nodes); i++ {
		parser.graph.nodes[i].printNode()
	}

	fmt.Println("ALL POSSIBLE INPUTS:")
	parser.graph.inputs.print()

	automaton := initAutomaton()
	automaton.convertFrom(parser.graph)

	fmt.Println("\nAUTOMATON:")

	for state := range automaton.states {
		automaton.states[state].printState()
	}

	fmt.Printf("WITH START STATE OF %s\n", automaton.start.name)

	testInputs := []string{"A", "B", "AB", "BA", "ABA", "C", "", "AA", "AAA", "AAAA", "AAAAA"}

	for i := 0; i < len(testInputs); i++ {
		fmt.Println(automaton.validate(testInputs[i]))
	}

	fmt.Println("Exiting...")
}
