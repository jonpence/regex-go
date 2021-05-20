/* Queue tests.
 */

package deque

import (
	"testing"
)

func TestInitQueue(t *testing.T) {
	testQueue := InitQueue()

	if &testQueue == nil {
		t.Errorf("expected non-nil value for &testQueue")
	}
}

func TestIsEmptyQueue(t *testing.T) {
	testQueue := InitQueue()

	if !testQueue.IsEmpty() {
		t.Errorf("expected testQueue.IsEmpty() to be true, instead is false")
	}

	testQueue.Enqueue("")

	if testQueue.IsEmpty() {
		t.Errorf("expected testQueue.IsEmpty() to be false, instead is true")
	}
}

func TestEnqueue(t *testing.T) {
	testQueue := InitQueue()

	testQueue.Enqueue("")

	if testQueue.top == nil {
		t.Errorf("expected testQueue.top to have non-nil value")
	}
}

func TestDequeue(t *testing.T) {
	testQueue := InitQueue()

	testQueue.Enqueue("")

	_, present := testQueue.Dequeue()

	if !present {
		t.Errorf("expected testQueue.Dequeue to return true, instead returned false")
	}
}

func TestPeekQueue(t *testing.T) {
	testString := "test"
	testQueue  := InitQueue()

	_, present := testQueue.Peek()

	if present {
		t.Errorf("expected testQueue.Peek() to return false, instead returned true")
	}

	testQueue.Enqueue(testString)

	value, _ := testQueue.Peek()

	if value != testString {
		t.Errorf("expected testQueue.Peek() to return %s, instead returned %s", testString, value)
	}

	if testQueue.IsEmpty() {
		t.Errorf("expected testQueue.IsEmpty() to return false, instead returned true")
	}
}
