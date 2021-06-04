/* list.go
/*
/* Implements a list of integers.
 */

package iset

import (
	"fmt"
)

type List []int

func initList(size int) List {
	return List(make([]int, size))
}

func (l List)size() int {
	return len(l)
}

func (l List)at(index int) (bool, int) {
	if index < 0 || index >= l.size() {
		return false, 0
	}

	return true, l[index]
}

func (l *List)insert(val int, index int) bool {
	if index < 0 || index > l.size() {
		return false
	} else if index == l.size() {
		*l = append(*l, val)
	} else {
		*l = append((*l)[:index + 1], (*l)[index:]...)
		(*l)[index] = val
	}

	return true
}

func (l *List)prepend(val int) {
	l.insert(val, 0)
}

func (l *List)append(val int) {
	l.insert(val, l.size())
}

func (l *List)remove(index int) bool {
	if index < 0 || index > l.size() {
		return false
	}

	*l = append((*l)[:index - 1], (*l)[index:]...)

	return true
}

// binary search
func (l List)search(val int) (bool, int) {
	offset := (l.size() + 1) % 2
	low    := 0
	high   := l.size()

	i := (l.size() / 2) + offset

	for i > low && i < high {
		if l[i] == val {
			return true, i
		} else if val < l[i] {
			high = i
			i    = ((i + low) / 2) + offset
		} else {
			low = i
			i   = ((i + high) / 2) + offset
		}
	}

	if i >= low && i < l.size() {
		return l[i] == val, i
	} else {
		return false, 0
	}
}

func (l List)toString() string {
	buf := fmt.Sprintf("{%d", l[0])

	for _, elmt := range l[1:] {
		buf += fmt.Sprintf(", %d", elmt)
	}

	return buf + "}"
}
