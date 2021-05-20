/* Implements a deque, queue, and stack.
 */

package deque

type Deque struct {
	top    *DequeNode
	bottom *DequeNode
}

func InitDeque() Deque {
	return Deque{nil, nil}
}

func (d Deque) Top() *DequeNode {
	return d.top
}

func (d Deque) Bottom() *DequeNode {
	return d.bottom
}

func (d Deque) IsEmpty() bool {
	return d.top == nil && d.bottom == nil
}

func (d *Deque) Append(s string) {
	node := initNode(s)

	if d.top == nil {
		d.top    = node
		d.bottom = node
	} else {
		node.setPrevious(d.bottom)
		d.bottom.setNext(node)
		d.bottom = node
	}
}

func (d *Deque) Prepend(s string) {
	node := initNode(s)

	if d.top == nil {
		d.top    = node
		d.bottom = node
	} else {
		node.setNext(d.top)
		d.top.setPrevious(node)
		d.top = node
	}
}

func (d *Deque) PopTop() (string, bool) {
	if d.IsEmpty() {
		return "", false
	}

	topString := d.top.value

	if d.top.next != nil {
		d.top = d.top.next
		d.top.setPrevious(nil)
	} else {
		d.top = nil
		d.bottom = nil
	}

	return topString, true
}

func (d *Deque) PopBottom() (string, bool) {
	if d.IsEmpty() {
		return "", false
	}

	bottomString := d.bottom.value

	if d.bottom.previous != nil {
		d.bottom = d.bottom.previous
		d.bottom.setNext(nil)
	} else {
		d.top = nil
		d.bottom = nil
	}

	return bottomString, true
}

func (d Deque) PeekTop() (string, bool) {
	if d.IsEmpty() {
		return "", false
	}

	return d.top.value, true
}

func (d Deque) PeekBottom() (string, bool) {
	if d.IsEmpty() {
		return "", false
	}

	return d.bottom.value, true
}

