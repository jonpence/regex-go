/* Implements a queue using a deque.
 */

package ideque

type Queue Deque

func InitQueue() Queue {
	return Queue(InitDeque())
}

func (q Queue) IsEmpty() bool {
	return Deque(q).IsEmpty()
}

func (q *Queue) Enqueue(num int) {
	(*Deque)(q).Append(num)
}

func (q *Queue) Dequeue() (int, bool) {
	return (*Deque)(q).PopTop()
}

func (q Queue) Peek() (int, bool) {
	return Deque(q).PeekTop()
}
