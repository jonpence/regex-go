/* Automaton tests.
 */

package regex

import (
	"fmt"
	"testing"
)

func TestInitAutomaton(t *testing.T) {
	testAutomaton := initAutomaton()

	if testAutomaton == nil {
		t.Errorf("expected non-nil value for testAutomaton")
	}
}

func TestAddState(t *testing.T) {
	testAutomaton := initAutomaton()
	testState     := initNFAState("test")

	testAutomaton.addState(testState)

	_, present := testAutomaton.states[testState.name]

	if !present {
		t.Errorf("expected present to be true")
	}
}

func TestAddEdge(t *testing.T) {
	testAutomaton := initAutomaton()
	testStateA    := initNFAState("A")
	testStateB    := initNFAState("B")

	testAutomaton.addState(testStateA)
	testAutomaton.addState(testStateB)

	testAutomaton.addEdge(testStateA.name, testStateB.name, "")

	if testAutomaton.states[testStateA.name].neighbors[""][0] != "B" {
		t.Errorf("expected edge from A to B")
	}
}

func TestSetStart(t *testing.T) {
	testAutomaton := initAutomaton()
	testState     := initNFAState("")

	testAutomaton.addState(testState)
	testAutomaton.setStart(testState)

	if testAutomaton.start != testState {
		t.Errorf("expected testAutomaton.start to equal testState")
	}
}

func TestSetCurrent(t *testing.T) {
	testAutomaton := initAutomaton()
	testState     := initNFAState("")

	testAutomaton.addState(testState)
	testAutomaton.setCurrent(testState)

	if testAutomaton.current != testState {
		t.Errorf("expected testAutomaton.current to equal testState")
	}
}

func TestGetState(t *testing.T) {
	testAutomaton := initAutomaton()
	testState     := initNFAState("")

	testAutomaton.addState(testState)

	if testAutomaton.getState("") != testState {
		t.Errorf("expected testAutomaton.getState(\"\") to equal testState")
	}
}

func TestGetCount(t *testing.T) {
	testAutomaton := initAutomaton()

	testAutomaton.addState(initNFAState(""))
	testAutomaton.addState(initNFAState(""))
	testAutomaton.addState(initNFAState(""))
	testAutomaton.addState(initNFAState(""))

	if testAutomaton.getCount() != 4 {
		t.Errorf("expected testAutomaton.getCount() to equal 4")
	}
}

func TestReset(t *testing.T) {
	testAutomaton := initAutomaton()
	testStateA    := initNFAState("A")
	testStateB    := initNFAState("B")

	testAutomaton.setStart(testStateA)
	testAutomaton.setCurrent(testStateB)

	testAutomaton.reset()

	if testAutomaton.current != testStateA {
		t.Errorf("expected testAutomaton.current to equal testStateA")
	}
}

func TestTransition(t *testing.T) {
	testAutomaton := initAutomaton()

	testStateA := initNFAState("A")
	testStateB := initNFAState("B")

	testAutomaton.addState(testStateA)
	testAutomaton.addState(testStateB)

	testAutomaton.addEdge(testStateA.name, testStateB.name, "A")

	testAutomaton.setCurrent(testStateA)

	testAutomaton.transition('A')

	if testAutomaton.current != testStateB {
		t.Errorf("expected testAutomaton.current to equal testStateB")
	}
}

func TestClosureOfOn(t *testing.T) {
	testParser := initParser()
	testParser.parse("A|B")

	fmt.Println("TestClosureOfOn: ", testParser.nfa.closureOfOn(testParser.nfa.start, "").name)
}

func TestDeterminize(t *testing.T) {
	testParser := initParser()
	testParser.parse("A|B")

	dfa := testParser.nfa.determinize()

	fmt.Println("TestDeterminize:")
	for state := range dfa.states {
		dfa.getState(state).printDFAState()
	}
}

func TestValidateAutomaton(t *testing.T) {
	testParser := initParser()
	testParser.parse("A|B")

	dfa := testParser.nfa.determinize()

	if !dfa.validate("A") {
		t.Errorf("expected A to be valid")
	}

	if !dfa.validate("B") {
		t.Errorf("expected B to be valid")
	}

	if dfa.validate("C") {
		t.Errorf("expected C to be invalid")
	}

	if dfa.validate("") {
		t.Errorf("expected empty string to be invalid")
	}
}
