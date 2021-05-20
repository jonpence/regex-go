/* Implements nodes for use in the deque, queue, and stack data structures.
 */

package deque

type DequeNode struct {
	value    string
	previous *DequeNode
	next     *DequeNode
}

func initNode(s string) *DequeNode {
	return &DequeNode{s, nil, nil}
}

func (n *DequeNode) setValue(value string) {
	n.value = value
}

func (n *DequeNode) setPrevious(previous *DequeNode) {
	n.previous = previous
}

func (n *DequeNode) setNext(next *DequeNode) {
	n.next = next
}
