/* Implements a queue using a deque.
 */

package deque

type Queue Deque

func InitQueue() Queue {
	return Queue(InitDeque())
}

func (q Queue) IsEmpty() bool {
	return Deque(q).IsEmpty()
}

func (q *Queue) Enqueue(s string) {
	(*Deque)(q).Append(s)
}

func (q *Queue) Dequeue() (string, bool) {
	return (*Deque)(q).PopTop()
}

func (q Queue) Peek() (string, bool) {
	return Deque(q).PeekTop()
}
