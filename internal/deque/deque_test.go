/* Deque tests.
 */

package deque

import (
	"testing"
)

func TestInitDeque(t *testing.T) {
	testDeque := InitDeque()

	if &testDeque == nil {
		t.Errorf("expected non-nil value for &testDeque, instead is %p", &testDeque)
	}
}

func TestTop(t *testing.T) {
	testDeque     := InitDeque()
	testDequeNode := initNode("")

	testDeque.top = testDequeNode

	if testDeque.Top() != testDequeNode {
		t.Errorf("expected testDeque.Top() to equal testDequeNode, instead equals %p", testDeque.Top())
	}
}

func TestBottom(t *testing.T) {
	testDeque     := InitDeque()
	testDequeNode := initNode("")

	testDeque.bottom = testDequeNode

	if testDeque.Bottom() != testDequeNode {
		t.Errorf("expected testDeque.Bottom() to equal testDequeNode, instead equals %p", testDeque.Bottom())
	}
}

func TestIsEmptyDeque(t *testing.T) {
	testDequeA    := InitDeque()
	testDequeB    := InitDeque()
	testDequeNode := initNode("")

	testDequeA.top = testDequeNode

	if testDequeA.IsEmpty() {
		t.Errorf("expected testDequeA.IsEmpty() to be false, instead is true")
	}

	if !testDequeB.IsEmpty() {
		t.Errorf("expected testDequeB.IsEmpty() to be true, instead is false")
	}
}

func TestAppend(t *testing.T) {
	testStringA := "testA"
	testStringB := "testB"
	testDeque   := InitDeque()

	testDeque.Append(testStringA)

	if testDeque.top.value != testStringA {
		t.Errorf("expected testDeque.top.value to equal %s, instead is %s", testStringA, testDeque.top.value)
	}

	testDeque.Append(testStringB)

	if testDeque.bottom.value != testStringB {
		t.Errorf("expected testDeque.bottom.value to equal %s, instead is %s", testStringB, testDeque.bottom.value)
	}
}

func TestPrepend(t *testing.T) {
	testStringA := "testA"
	testStringB := "testB"
	testDeque   := InitDeque()

	testDeque.Prepend(testStringA)

	if testDeque.top.value != testStringA {
		t.Errorf("expected testDeque.top.value to equal %s, instead is %s", testStringA, testDeque.top.value)
	}

	testDeque.Prepend(testStringB)

	if testDeque.top.value != testStringB {
		t.Errorf("expected testDeque.top.value to equal %s, instead is %s", testStringB, testDeque.top.value)
	}
}

func TestPopTop(t *testing.T) {
	testString  := "test"
	testDeque   := InitDeque()

	_, present := testDeque.PopTop()

	if present {
		t.Errorf("expected present to be false, instead is true")
	}

	testDeque.Prepend(testString)

	_, present = testDeque.PopTop()

	if !present {
		t.Errorf("expected present to be true, instead is false")
	}

	_, present = testDeque.PopTop()

	if present {
		t.Errorf("expected present to be false, instead is true")
	}
}

func TestPopBottom(t *testing.T) {
	testString  := "test"
	testDeque   := InitDeque()

	_, present := testDeque.PopBottom()

	if present {
		t.Errorf("expected present to be false, instead is true")
	}

	testDeque.Append(testString)

	_, present = testDeque.PopBottom()

	if !present {
		t.Errorf("expected present to be true, instead is false")
	}

	_, present = testDeque.PopBottom()

	if present {
		t.Errorf("expected present to be false, instead is true")
	}
}

func TestPeekTop(t *testing.T) {
	testTop    := "testTop"
	testBottom := "testBottom"
	testDeque  := InitDeque()

	_, present := testDeque.PeekTop()

	if present {
		t.Errorf("expected present to be false, instead is true")
	}

	testDeque.Prepend(testBottom)
	testDeque.Prepend(testTop)

	value, _ := testDeque.PeekTop()

	if value != testTop {
		t.Errorf("expected testDeque.PeekTop() to return %s, instead returned %s", testTop, value)
	}
}

func TestPeekBottom(t *testing.T) {
	testTop    := "testTop"
	testBottom := "testBottom"
	testDeque  := InitDeque()

	_, present := testDeque.PeekBottom()

	if present {
		t.Errorf("expected present to be false, instead is true")
	}

	testDeque.Prepend(testBottom)
	testDeque.Prepend(testTop)

	value, _ := testDeque.PeekBottom()

	if value != testBottom {
		t.Errorf("expected testDeque.PeekBottom() to return %s, instead returned %s", testBottom, value)
	}
}
