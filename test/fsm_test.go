package fsm

import (
	"testing"
	"reflect"
	fsm "../src/fsm"
)

// Gonna use this FSM to test throughout the test functions
var test_fsm fsm.FiniteStateMachine

// Set up test_fsm and test fsm.InitFSM
func TestInitFSM(t *testing.T) {
	test_fsm = fsm.InitFSM(5)

	if test_fsm.GetSize() != 5 {
		t.Errorf("test_fsm.size = %d, wanted 5", test_fsm.GetSize())
	}
}

func TestSetTransition(t *testing.T) {
	test_fsm.SetTransition(0,0,[]rune{'0', '1', '2'})

	if !reflect.DeepEqual(test_fsm.GetTransition(0,0), []rune{'0', '1', '2'}) {
		t.Errorf("test_fsm[0][0] = %c, wanted [0 1 2]", test_fsm.GetTransition(0,0))
	}
}

func TestSetTerminals(t *testing.T) {
	test_fsm.SetTerminals([]int{0, 1, 2})

	if !reflect.DeepEqual(test_fsm.GetTerminals(), []int{0, 1, 2}) {
		t.Errorf("test_fsm.terminals = %d, wanted [0 1 2]", test_fsm.GetTerminals())
	}
}

func TestValidateWord(t *testing.T) {
	// Overwrite 0 -> 0 transition defined in TestSetTransition
	test_fsm.SetTransition(0,0,[]rune{})

	// Overwrite terminals defined in TestSetTerminals
	test_fsm.SetTerminals([]int{2, 4})

	// Set a pathway to validate "c(a)*t"
	test_fsm.SetTransition(0,1,[]rune{'c'})
	test_fsm.SetTransition(1,1,[]rune{'a'})
	test_fsm.SetTransition(1,2,[]rune{'t'})

	// Set a pathway to validate "d(o)*g"
	test_fsm.SetTransition(0,3,[]rune{'d'})
	test_fsm.SetTransition(3,3,[]rune{'o'})
	test_fsm.SetTransition(3,4,[]rune{'g'})

	// Test various words

    if !test_fsm.ValidateWord("ct") {
		t.Errorf("Rejected \"ct\", should've accepted")
	}

	if !test_fsm.ValidateWord("cat") {
		t.Errorf("Rejected \"cat\", should've accepted")
	}

	if !test_fsm.ValidateWord("caaat") {
		t.Errorf("Rejected \"caaat\", should've accepted")
	}

	if test_fsm.ValidateWord("ctt") {
		t.Errorf("Accepted \"ctt\", should've rejected")
	}

	if test_fsm.ValidateWord("catt") {
		t.Errorf("Accepted \"catt\", should've rejected")
	}

	if test_fsm.ValidateWord("caaatt") {
		t.Errorf("Accepted \"caaatt\", should've rejected")
	}

	if !test_fsm.ValidateWord("dg") {
		t.Errorf("Rejected \"dg\", should've accepted")
	}

	if !test_fsm.ValidateWord("dog") {
		t.Errorf("Rejected \"dog\", should've accepted")
	}

	if !test_fsm.ValidateWord("dooog") {
		t.Errorf("Rejected \"dooog\", should've accepted")
	}

	if test_fsm.ValidateWord("dgg") {
		t.Errorf("Accepted \"dgg\", should've rejected")
	}

	if test_fsm.ValidateWord("dogg") {
		t.Errorf("Accepted \"dogg\", should've rejected")
	}

	if test_fsm.ValidateWord("dooogg") {
		t.Errorf("Accepted \"dooogg\", should've rejected")
	}

	if test_fsm.ValidateWord("elephant") {
		t.Errorf("Accepted \"elephant\", should've rejected")
	}

	if test_fsm.ValidateWord("hamster") {
		t.Errorf("Accepted \"hamster\", should've rejected")
	}

	if test_fsm.ValidateWord("cdaotg") {
		t.Errorf("Accepted \"cdaotg\", should've rejected")
	}
}
