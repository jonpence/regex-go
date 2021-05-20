/* Implements a stack using a deque.
 */

package deque

type Stack Deque

func InitStack() Stack {
	return Stack(InitDeque())
}

func (s Stack) IsEmpty() bool {
	return Deque(s).IsEmpty()
}

func (s *Stack) Push(str string) {
	(*Deque)(s).Prepend(str)
}

func (s *Stack) Pop() (string, bool) {
	return (*Deque)(s).PopTop()
}

func (s Stack) Peek() (string, bool) {
	return (Deque)(s).PeekTop()
}
