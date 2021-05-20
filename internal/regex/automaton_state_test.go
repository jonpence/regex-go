/* State tests.
 */

package regex

import (
	"testing"
	"github.com/jonpence/regex-go/internal/set"
)

func TestInitNFAState(t *testing.T) {
	testNFAState := initNFAState("")

	if &testNFAState == nil {
		t.Errorf("expected non-nil value for &testNFAState")
	}

	if testNFAState.dfaState {
		t.Errorf("expected testNFAState.dfaState to be false")
	}
}

func TestInitDFAState(t *testing.T) {
	testDFAState := initDFAState(set.Set{})

	if &testDFAState == nil {
		t.Errorf("expected non-nil value for &testDFAState")
	}

	if !testDFAState.dfaState {
		t.Errorf("expected testDFAState.dfaState to be true")
	}
}

func TestAddNeighbor(t *testing.T) {
	testNFAState := initNFAState("")
	testDFAState := initDFAState(set.Set{})

	testNFAState.addNeighbor("0", "0")
	testNFAState.addNeighbor("1", "0")

	testDFAState.addNeighbor("0", "0")
	testDFAState.addNeighbor("1", "0")

	if len(testNFAState.neighbors["0"]) != 2 {
		t.Errorf("expected testNFAState.neighbors[\"0\"] to be of size 2")
	}

	if len(testDFAState.neighbors["0"]) != 1 {
		t.Errorf("expected testDFAState.neighbors[\"0\"] to be of size 1")
	}
}

func TestSetTerminates(t *testing.T) {
	testState := initNFAState("")

	testState.setTerminates()

	if !testState.terminates {
		t.Errorf("expected testState.terminates to be true")
	}
}

func TestUnsetTerminates(t *testing.T) {
	testState := initNFAState("")
	testState.terminates = true

	testState.unsetTerminates()

	if testState.terminates {
		t.Errorf("expected testState.terminates to be false")
	}
}

func TestGetNeighborList(t *testing.T) {
	testState := initNFAState("")

	_, present := testState.getNeighborList("1")

	if present {
		t.Errorf("expected present to be false")
	}

	testState.addNeighbor("0", "0")
	testState.addNeighbor("1", "0")
	testState.addNeighbor("2", "0")

	testList := []string{"0", "1", "2"}
	list, _ := testState.getNeighborList("0")

	for i, item := range testList {
		if item != list[i] {
			t.Errorf("expected testState.getNeighborList(\"0\") to equal testList")
		}
	}
}

func TestGetNeighbor(t *testing.T) {
	testState := initDFAState(set.Set{})

	_, present := testState.getNeighbor("0")

	if present {
		t.Errorf("expected present to be false")
	}

	testState.addNeighbor("0", "0")

	neighbor, _ :=  testState.getNeighbor("0")

	if neighbor != "0" {
		t.Errorf("expected neighbor to equal \"0\"")
	}
}
