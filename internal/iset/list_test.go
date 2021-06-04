package iset

import (
	"testing"
)

func TestSize(t *testing.T) {
	l := List([]int{1, 2, 3, 4, 5})

	if l.size() != 5 {
		t.Errorf("expected l.size() to be 5, is %d", l.size())
	}
}

func TestAt(t *testing.T) {
	l := List([]int{1, 2, 3, 4, 5})

	if _, val := l.at(0); val != 1 {
		t.Errorf("expected l.at(0) to be 1, is %d", val)
	}

	if _, val := l.at(l.size() - 1); val != 5 {
		t.Errorf("expected l.at(l.size() - 1) to be 5, is %d", val)
	}

	if _, val := l.at(2); val != 3 {
		t.Errorf("expected l.at(2) to be 3, is %d", val)
	}
}

func TestInsert(t *testing.T) {
	l := List([]int{1, 2, 3, 4, 5})

	if l.insert(0, -1) {
		t.Errorf("expected l.insert(0, -1) to be false")
	}

	if l.insert(0, 6) {
		t.Errorf("expected l.insert(0, 6) to be false")
	}

	l.insert(0, 3)

	if _, val := l.at(3); val != 0 {
		t.Errorf("expected l.at(3) to be 0, instead is %d with list of %s", val, l.toString())
	}
}

func TestRemove(t *testing.T) {
	l := List([]int{1, 2, 3, 4, 5})

	if l.remove(-1) {
		t.Errorf("expected l.remove(-1) to be false")
	}

	if l.remove(6) {
		t.Errorf("expected l.remove(6) to be false")
	}

	l.remove(3)

	if _, val := l.at(3); val != 5 {
		t.Errorf("expected l.at(3) to be 5, instead is %d with list of %s", val, l.toString())
	}
}

func TestSearch(t *testing.T) {
	l := List([]int{1, 2, 3, 4, 5})

	present, index := l.search(3)

	if !present {
		t.Errorf("expected present to be true for 3")
	}

	if index != 2 {
		t.Errorf("expected index to be 2, instead is %d", index)
	}

	l = List([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20})

	present, index  = l.search(7)

	if !present {
		t.Errorf("expected present to be true for 7")
	}

	if index != 6 {
		t.Errorf("expected index to be 6, instead is %d", index)
	}

	l = List([]int{1, 6, 17, 32, 49, 51, 56, 60, 69, 70})

	present, _ = l.search(7)

	if present {
		t.Errorf("expected present to be false for 7")
	}

	l = initList(0)

	for i := 0; i < 101; i++ {
		l.insert(i, i)
	}

	present, _ = l.search(47)

	if !present {
		t.Errorf("expected present to be true for 42")
	}
}

func TestPrepend(t *testing.T) {
	l := List([]int{1, 2, 3, 4, 5})

	l.prepend(100)

	if _, val := l.at(0); val != 100 {
		t.Errorf("expected l.at(0) to be 100, instead is %d", val)
	}
}

func TestAppend(t *testing.T) {
	l := List([]int{1, 2, 3, 4, 5})

	l.append(100)

	if _, val := l.at(l.size() - 1); val != 100 {
		t.Errorf("expected l.at(l.size() - 1) to be 100, instead is %d", val)
	}
}
