/* Stack tests.
 */

package deque

import (
	"testing"
)

func TestInitStack(t *testing.T) {
	testStack := InitStack()

	if &testStack == nil {
		t.Errorf("expected non-nil value for &testStack")
	}
}

func TestIsEmptyStack(t *testing.T) {
	testStack := InitStack()
	testNode  := initNode("")

	if !testStack.IsEmpty() {
		t.Errorf("expected true for testStack.isEmpty(), instead is false")
	}

	testStack.top = testNode

	if testStack.IsEmpty() {
		t.Errorf("expected false for testStack.isEmpty(), instead is true")
	}
}

func TestPush(t *testing.T) {
	testString := "test"
	testStack  := InitStack()

	testStack.Push(testString)

	if testStack.top.value != testString {
		t.Errorf("expected testStack.top.value to equal %s, instead is %s", testString, testStack.top.value)
	}
}

func TestPop(t *testing.T) {
	testString := "test"
	testStack  := InitStack()

	_, present := testStack.Pop()

	if present {
		t.Errorf("expected testStack.Pop() to return (_, false), instead returned (_, true)")
	}

	testStack.Push(testString)

	value, _ := testStack.Pop()

	if value != testString {
		t.Errorf("expected testStack.Pop() to return (%s, _), instead returned (%s, _)", testString, value)
	}
}

func TestPeekStack(t *testing.T) {
	testString := "test"
	testStack  := InitStack()

	_, present := testStack.Peek()

	if present {
		t.Errorf("expected testStack.Peek() to return (_, false), instead returned (_, true)")
	}

	testStack.Push(testString)

	value, _ := testStack.Peek()

	if value != testString {
		t.Errorf("expected testStack.Peek() to return (%s, _), instead returned (%s, _)", testString, value)
	}

	if testStack.IsEmpty() {
		t.Errorf("expected non-empty stack, instead is empty")
	}
}
