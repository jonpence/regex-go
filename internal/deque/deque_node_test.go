/* DequeNode tests.
 */

package deque

import (
	"testing"
)

func TestinitNode(t *testing.T) {
	testNode := initNode("")

	if testNode == nil {
		t.Errorf("expected non-nil value for testNode")
	}
}

func TestSetValue(t *testing.T) {
	testNode := initNode("")

	testNode.setValue("test")

	if testNode.value != "test" {
		t.Errorf("expected testNode.value() to equal \"test\"")
	}
}

func TestSetPrevious(t *testing.T) {
	testNodeA := initNode("testNodeA")
	testNodeB := initNode("testNodeB")

	testNodeB.setPrevious(testNodeA)

	if testNodeB.previous != testNodeA {
		t.Errorf("expected value of testNodeB.previous to be testNodeA")
	}
}

func TestSetNext(t *testing.T) {
	testNodeA := initNode("testNodeA")
	testNodeB := initNode("testNodeB")

	testNodeA.setNext(testNodeB)

	if testNodeA.next != testNodeB {
		t.Errorf("expected value of testNodeA.next to be testNodeB")
	}
}
