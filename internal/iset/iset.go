/* iset.go
/*
/* Implements a set of ints.
 */

package iset

import (
	"fmt"
)

type Set List

/* METHODS */
//  -- InitSet() -> Set
//  -- InitSetElements([]int) -> Set
//  -- Set.size() -> int
//  -- itos(int) -> int
//  -- InitSetRange(int, int) -> Set
//  -- Set.IsMember(int) -> bool
//  -- *Set.Add(int)
//  -- *Set.Multiadd([]int)
//  -- *Set.Remove(int) -> bool
//  -- *Set.Discard(int)
//  -- *Set.Pop() -> int
//  -- *Set.Clear()
//  -- Set.Copy() -> Set
//  -- Set.Intersection(Set) -> Set
//  -- Set.Subset(Set) -> bool
//  -- Set.Superset(Set) -> bool
//  -- Set.ProperSubset(Set) -> bool
//  -- Set.ProperSuperset(Set) -> bool
//  -- Set.Equivalent(Set) -> bool
//  -- Set.Difference(Set) -> Set
//  -- Set.Union(Set) -> Set
//  -- Set.SymmetricDifference(Set) -> Set
//  -- Set.Disjoint(Set) -> bool
//  -- Set.ToString() -> string


/*******************************************************************************/


/* InitSet() -> Set
 */
func InitSet() Set {
	return Set(initList(0))
}

/* InitSetElements([]int) -> Set
 */
func InitSetElements(elements []int) Set {
	newSet := InitSet()

	for _, elmt := range elements {
		newSet.Add(elmt)
		fmt.Println(newSet.ToString())
	}

	return newSet
}

/* InitSetRange(int, int) -> Set
 */
func InitSetRange(lower int, upper int) Set {
	newSet := InitSet()

	for ; lower < upper ; lower++ {
		newSet.Add(lower)
	}

	return newSet
}

/* Set.size() -> int
 */
func (s Set) Size() int {
	return List(s).size()
}

/* Set.IsMember(int) -> bool
 */
func (s Set) IsMember(num int) bool {
	present, _ := List(s).search(num)

	return present
}

/* *Set.Add(int)
 */
func (s *Set) Add(num int) {
	for i := 0; i < s.Size(); i++ {
		if num == (*s)[i] {
			return
		} else if num < (*s)[i] {
			(*List)(s).insert(num, i)
			return
		}
	}

	(*List)(s).append(num)
}

/* *Set.Multiadd([]int)
 */
func (s *Set) Multiadd(nums []int) {
	for _, num := range nums {
		s.Add(num)
	}
}

/* *Set.Remove(int) -> bool
 */
func (s *Set) Remove(num int) bool {
	if !s.IsMember(num) {
		return false
	}

	present, index := List(*s).search(num)

	if !present {
		return false
	}

	(*List)(s).remove(index)

	return true
}

/* *Set.Discard(int)
 */
func (s *Set) Discard(num int) {
	s.Remove(num)
}

/* *Set.Pop() -> (int, bool)
 */
func (s *Set) Pop() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	}

	elmt := (*s)[0]

	s.Discard(elmt)

	return elmt, true
}

/* *Set.Clear()
 */
func (s *Set) Clear() {
	for !s.IsEmpty() {
		s.Pop()
	}
}

/* Set.Copy() -> Set
 */
func (s Set) Copy() Set {
	elements := []int{}
	copy(elements, s)

	return InitSetElements(elements)
}

/* Set.Intersection(Set) -> Set
 */
func (setA Set) Intersection(setB Set) Set {
	newSet := InitSet()

	for _, elmt := range setA {
		if setB.IsMember(elmt) {
			newSet.Add(elmt)
		}
	}

	return newSet
}

/* Set.Subset(Set) -> bool
 */
func (setA Set) Subset(setB Set) bool {
	for _, elmt := range setA {
		if !setB.IsMember(elmt) {
			return false
		}
	}

	return true
}

/* Set.Superset(Set) -> bool
 */
func (setA Set) Superset(setB Set) bool {
	return setB.Subset(setA)
}

/* Set.ProperSubset(Set) -> bool
 */
func (setA Set) ProperSubset(setB Set) bool {
	return setA.Subset(setB) && !setB.Subset(setA)
}

/* Set.ProperSuperset(Set) -> bool
 */
func (setA Set) ProperSuperset(setB Set) bool {
	return setA.Superset(setB) && !setB.Superset(setA)
}

/* Set.Equivalent(Set) -> bool
 */
func (setA Set) Equivalent(setB Set) bool {
	return setA.Subset(setB) && setB.Subset(setA)
}

/* Set.Difference(Set) -> Set
 */
func (setA Set) Difference(setB Set) Set {
	intersection := setA.Intersection(setB)
	newSet       := InitSet()

	for _, elmt := range setA {
		if !intersection.IsMember(elmt) {
			newSet.Add(elmt)
		}
	}

	return newSet
}

/* Set.Union(Set) -> Set
 */
func (setA Set) Union(setB Set) Set {
	newSet := InitSet()

	for _, elmt := range setA {
		newSet.Add(elmt)
	}

	for _, elmt := range setB {
		newSet.Add(elmt)
	}

	return newSet
}

/* Set.SymmetricDifference(Set) -> Set
 */
func (setA Set) SymmetricDifference(setB Set) Set {
	return setA.Union(setB).Difference(setA.Intersection(setB))
}

/* Set.Disjoint(Set) -> bool
 */
func (setA Set) Disjoint(setB Set) bool {
	return setA.Intersection(setB).IsEmpty()
}

/* Set.IsEmpty() -> bool
 */
func (s Set) IsEmpty() bool {
	return len(s) == 0
}

/* Set.ToString() -> string
 */
func (s Set) ToString() string {
	return List(s).toString()
}
