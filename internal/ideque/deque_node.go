/* Implements nodes for use in the deque, queue, and stack data structures.
 */

package ideque

type DequeNode struct {
	value    int
	previous *DequeNode
	next     *DequeNode
}

func initNode(num int) *DequeNode {
	return &DequeNode{num, nil, nil}
}

func (n *DequeNode) setValue(value int) {
	n.value = value
}

func (n *DequeNode) setPrevious(previous *DequeNode) {
	n.previous = previous
}

func (n *DequeNode) setNext(next *DequeNode) {
	n.next = next
}
