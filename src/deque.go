/* Implements a deque, queue, and stack.
 */

package main

/* NODE */

type DequeNode struct {
	value    string
	previous *DequeNode
	next     *DequeNode
}

/* DEQUE */

type Deque struct {
	top    *DequeNode
	bottom *DequeNode
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

func (d Deque) isEmpty() bool {
	return d.top == nil && d.bottom == nil
}

func (d *Deque) append(s string) {
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

func (d *Deque) prepend(s string) {
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

func (d *Deque) popTop() (string, bool) {
	if d.isEmpty() {
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

func (d *Deque) popBottom() (string, bool) {
	if d.isEmpty() {
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

func (d Deque) peekTop() (string, bool) {
	if d.isEmpty() {
		return "", false
	}

	return d.top.value, true
}

func (d Deque) peekBottom() (string, bool) {
	if d.isEmpty() {
		return "", false
	}

	return d.bottom.value, true
}

/* QUEUE */

type Queue Deque

func (q Queue) isEmpty() bool {
	return Deque(q).isEmpty()
}

func (q *Queue) enqueue(s string) {
	(*Deque)(q).append(s)
}

func (q *Queue) dequeue() (string, bool) {
	return (*Deque)(q).popTop()
}

func (q Queue) peek() (string, bool) {
	return Deque(q).peekTop()
}

/* STACK */

type Stack Deque

func (s Stack) isEmpty() bool {
	return Deque(s).isEmpty()
}

func (s *Stack) push(str string) {
	(*Deque)(s).prepend(str)
}

func (s *Stack) pop() (string, bool) {
	return (*Deque)(s).popTop()
}

func (s Stack) peek() (string, bool) {
	return (Deque)(s).peekTop()
}
