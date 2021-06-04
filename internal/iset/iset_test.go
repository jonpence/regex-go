package iset

import (
	"testing"
	"fmt"
)

func TestIsMember(t *testing.T) {
	newSet := InitSetRange(0, 10001)

	if newSet.IsMember(10001) {
		t.Errorf("expected newSet.IsMember(10001) to be false")
	}

	if !newSet.IsMember(235) {
		t.Errorf("expected newSet.IsMember(235) to be true")
	}
}

func TestAdd(t *testing.T) {
	elms := []int{8, 19, 4, 8, 2, 1, 9, 0}

	newSet := InitSetElements(elms)

	fmt.Println(newSet.ToString())
}
