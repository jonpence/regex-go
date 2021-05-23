/* Implements a stack using a deque.
 */

package ideque

type Stack Deque

func InitStack() Stack {
	return Stack(InitDeque())
}

func (s Stack) IsEmpty() bool {
	return Deque(s).IsEmpty()
}

func (s *Stack) Push(num int) {
	(*Deque)(s).Prepend(num)
}

func (s *Stack) Pop() (int, bool) {
	return (*Deque)(s).PopTop()
}

func (s Stack) Peek() (int, bool) {
	return (Deque)(s).PeekTop()
}
