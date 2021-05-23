/* iset.go
/*
/* Implements a set of ints.
 */

package iset

import (
	"fmt"
)

type Set map[int]bool

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
//  -- Set.Print()
//  -- Set.ToString() -> string


/*******************************************************************************/


/* InitSet() -> Set
/*
/* Instantiates a new map[int]bool and returns it cast to a Set.
 */
func InitSet() Set {
	return Set(make(map[int]bool))
}

/* InitSetElements([]int) -> Set
/*
/* Instantiates a new map[int]bool and initializes its elements with */
/* the string slice parameter. Then returns it cast to a Set.
 */
func InitSetElements(elements []int) Set {
	newSet := InitSet()

	for _, element := range elements {
		newSet.Add(element)
	}

	return newSet
}

/* Set.size() -> int
/*
/* Returns the number of elements in the set.
 */
func (s Set) Size() int {
	return len(s)
}

/* InitSetRange(int, int) -> Set
/*
/* Instantiates a new set and fills it with elements ranging from low to high
/* exclusive. Then returns the new set.
 */
func InitSetRange(lower int, upper int) Set {
	newSet := InitSet()

	for ; lower < upper; lower++ {
		newSet.Add(lower)
	}

	return newSet
}

/* Set.IsMember(int) -> bool
/*
/* Returns true if int parameter is a member of the set.
 */
func (s Set) IsMember(num int) bool {
	present, member := s[num]

	return present && member
}

/* *Set.Add(int)
/*
/* Adds string parameter as member of the set.
 */
func (s *Set) Add(num int) {
	(*s)[num] = true
}

/* *Set.Multiadd([]int)
/*
/* Adds all members of []int parameter as members of the set.
 */
func (s *Set) Multiadd(nums []int) {
	for _, num := range nums {
		s.Add(num)
	}
}

/* *Set.Remove(int) -> bool
/*
/* If int parameter is a member of the set, Remove it and return true.
/* Otherwise return false if int parameter is not a member.
 */
func (s *Set) Remove(num int) bool {
	if s.IsMember(num) {
		(*s)[num] = false
		return true
	} else {
		return false
	}
}

/* *Set.Discard(int)
/*
/* Remove int parameter from set and do not report if the int parameter
/* was not detected as an element of the set.
 */
func (s *Set) Discard(num int) {
	if s.IsMember(num) {
		(*s)[num] = false
	}
}

/* *Set.Pop() -> (int, bool)
/*
/* Remove a random element from the set and return it and return true. If the
/* set is empty, return 0 and report false.
 */
func (s *Set) Pop() (int, bool) {
	for element := range *s {
		if s.IsMember(element) {
			s.Remove(element)
			return element, true
		}
	}

	return 0, false
}

/* *Set.Clear()
/*
/* Removes all elements from the set.
 */
func (s *Set) Clear() {
	for element := range *s {
		s.Remove(element)
	}
}

/* Set.Copy() -> Set
/*
/* Returns a Copy of the set.
 */
func (s Set) Copy() Set {
	newSet := InitSet()

	for element := range s {
		if s.IsMember(element) {
			newSet.Add(element)
		}
	}

	return newSet
}

/* Set.Intersection(Set) -> Set
/*
/* Returns the Intersection of the sets.
 */
func (setA Set) Intersection(setB Set) Set {
	var larger, smaller *Set

	newSet := InitSet()

	if setA.Size() > setB.Size() {
		larger = &setA
		smaller = &setB
	} else {
		larger = &setB
		smaller = &setA
	}

	for element := range *larger {
		if larger.IsMember(element) && smaller.IsMember(element) {
			newSet.Add(element)
		}
	}

	return newSet
}

/* Set.Subset(Set) -> bool
/*
/* Returns true if the set is a Subset of the other set.
 */
func (setA Set) Subset(setB Set) bool {
	for element := range setA {
		if setA.IsMember(element) && !setB.IsMember(element) {
			return false
		}
	}

	return true
}

/* Set.Superset(Set) -> bool
/*
/* Returns true if the set is a Superset of the other set.
 */
func (setA Set) Superset(setB Set) bool {
	return setB.Subset(setA)
}

/* Set.ProperSubset(Set) -> bool
/*
/* Returns true if the set is a Proper Subset of the other set.
 */
func (setA Set) ProperSubset(setB Set) bool {
	return setA.Subset(setB) && !setB.Subset(setA)
}

/* Set.ProperSuperset(Set) -> bool
/*
/* Returns true if the set is a Proper Superset of the other set.
 */
func (setA Set) ProperSuperset(setB Set) bool {
	return setA.Superset(setB) && !setB.Superset(setA)
}

/* Set.Equivalent(Set) -> bool
/*
/* Returns true if the set is Equivalent with the other set.
 */
func (setA Set) Equivalent(setB Set) bool {
	return setA.Subset(setB) && setB.Subset(setA)
}

/* Set.Difference(Set) -> Set
/*
/* Returns the Difference between the set and the other set.
 */
func (setA Set) Difference(setB Set) Set {
	newSet := InitSet()
	Intersection := setA.Intersection(setB)

	for element := range setA {
		if setA.IsMember(element) && !Intersection.IsMember(element) {
			newSet.Add(element)
		}
	}

	return newSet
}

/* Set.Union(Set) -> Set
/*
/* Returns the Union of the set with the other set.
 */
func (setA Set) Union(setB Set) Set {
	newSet := InitSet()

	for element := range setA {
		if setA.IsMember(element) {
			newSet.Add(element)
		}
	}

	for element := range setB {
		if setB.IsMember(element) {
			newSet.Add(element)
		}
	}

	return newSet
}

/* Set.SymmetricDifference(Set) -> Set
/*
/* Returns the Difference of the Union of the set and the other set with the */
/* Intersection of the set and the other set.
 */
func (setA Set) SymmetricDifference(setB Set) Set {
	return setA.Union(setB).Difference(setA.Intersection(setB))
}

/* Set.Disjoint(Set) -> bool
/*
/* Returns true if the set is Disjoint with the other set.
 */
func (setA Set) Disjoint(setB Set) bool {
	return setA.Intersection(setB).IsEmpty()
}

/* Set.IsEmpty() -> bool
/*
/* Returns true if the set is empty.
 */
func (s Set) IsEmpty() bool {
	if len(s) == 0 {
		return true
	} else {
		for element := range s {
			if s.IsMember(element) {
				return false
			}
		}

		return true
	}
}

func (s Set) ToSlice() []int {
	slice := []int{}

	for {
		element, ok := s.Pop()

		if !ok {
			break
		}

		slice = append(slice, element)
	}

	return slice
}

/* Set.Print()
/*
/* Prints out the set's elements.
 */
func (s Set) Print() {
	fmt.Println(s.ToString())
}

/* Set.ToString() -> string
/*
/* Returns a string of the set's elements.
 */
func (s Set) ToString() string {
	var buf string
	buf = buf + "{"

	first := true

	for element := range s {
		if s.IsMember(element) {
			if !first {
				buf = buf + fmt.Sprintf(", %d", element)
			} else {
				buf = buf + fmt.Sprintf("%d", element)
				first = false
			}
		}
	}
	buf = buf + "}"

	return buf
}
