/* Implements a Set data type using a map[string]bool hashmap.
 */

package main

import (
	"fmt"
)

/* DATA */
type Set map[string]bool

/* METHODS */
//  -- initSet() -> Set
//  -- Set.isMember(string) -> bool
//  -- *Set.add(string)
//  -- *Set.remove(string) -> bool
//  -- *Set.discard(string)
//  -- *Set.pop() -> string
//  -- *Set.clear()
//  -- Set.copy() -> Set
//  -- Set.intersection(Set) -> Set
//  -- Set.subset(Set) -> bool
//  -- Set.superset(Set) -> bool
//  -- Set.properSubset(Set) -> bool
//  -- Set.properSuperset(Set) -> bool
//  -- Set.equivalent(Set) -> bool
//  -- Set.difference(Set) -> Set
//  -- Set.union(Set) -> Set
//  -- Set.symmetricDifference(Set) -> Set
//  -- Set.disjoint(Set) -> bool
//  -- Set.print()
//  -- Set.toString() -> string


/*******************************************************************************/


/* initSet() -> Set
/*
/* Instantiates a new map[string]bool and returns it cast to a Set.
 */
func initSet() Set {
	return Set(make(map[string]bool))
}

/* Set.isMember(string) -> bool
/*
/* Returns true if string parameter is a member of the set.
 */
func (s Set) isMember(str string) bool {
	present, member := s[str]

	return present && member
}

/* *Set.add(string)
/*
/* Adds string parameter as member of the set.
 */
func (s *Set) add(str string) {
	(*s)[str] = true
}

/* *Set.remove(string) -> bool
/*
/* If string parameter is a member of the set, remove it and return true.
/* Otherwise return false if string parameter is not a member.
 */
func (s *Set) remove(str string) bool {
	if s.isMember(str) {
		(*s)[str] = false
		return true
	} else {
		return false
	}
}

/* *Set.discard(string)
/*
/* Remove string parameter from set and do not report if the string parameter
/* was not detected as an element of the set.
 */
func (s *Set) discard(str string) {
	if s.isMember(str) {
		(*s)[str] = false
	}
}

/* *Set.pop() -> (string, bool)
/*
/* Remove a random element from the set and return it and return true. If the */
/* set is empty, return the empty string and report false.
 */
func (s *Set) pop() (string, bool) {
	for element := range *s {
		if s.isMember(element) {
			s.remove(element)
			return element, true
		}
	}

	return "", false
}

/* *Set.clear()
/*
/* Removes all elements from the set.
 */
func (s *Set) clear() {
	for element := range *s {
		s.remove(element)
	}
}

/* Set.copy() -> Set
/*
/* Returns a copy of the set.
 */
func (s Set) copy() Set {
	newSet := initSet()

	for element := range s {
		if s.isMember(element) {
			newSet.add(element)
		}
	}

	return newSet
}

/* Set.intersection(Set) -> Set
/*
/* Returns the intersection of the sets.
 */
func (setA Set) intersection(setB Set) Set {
	var larger, smaller *Set

	newSet := initSet()

	if len(setA) > len(setB) {
		larger = &setA
		smaller = &setB
	} else {
		larger = &setB
		smaller = &setA
	}

	for element := range *larger {
		if larger.isMember(element) && smaller.isMember(element) {
			newSet.add(element)
		}
	}

	return newSet
}

/* Set.subset(Set) -> bool
/*
/* Returns true if the set is a subset of the other set.
 */
func (setA Set) subset(setB Set) bool {
	for element := range setA {
		if setA.isMember(element) && !setB.isMember(element) {
			return false
		}
	}

	return true
}

/* Set.superset(Set) -> bool
/*
/* Returns true if the set is a superset of the other set.
 */
func (setA Set) superset(setB Set) bool {
	return setB.subset(setA)
}

/* Set.properSubset(Set) -> bool
/*
/* Returns true if the set is a proper subset of the other set.
 */
func (setA Set) properSubset(setB Set) bool {
	return setA.subset(setB) && !setB.subset(setA)
}

/* Set.properSuperset(Set) -> bool
/*
/* Returns true if the set is a proper superset of the other set.
 */
func (setA Set) properSuperset(setB Set) bool {
	return setA.superset(setB) && !setB.superset(setA)
}

/* Set.equivalent(Set) -> bool
/*
/* Returns true if the set is equivalent with the other set.
 */
func (setA Set) equivalent(setB Set) bool {
	return setA.subset(setB) && setB.subset(setA)
}

/* Set.difference(Set) -> Set
/*
/* Returns the difference between the set and the other set.
 */
func (setA Set) difference(setB Set) Set {
	newSet := initSet()
	intersection := setA.intersection(setB)

	for element := range setA {
		if setA.isMember(element) && !intersection.isMember(element) {
			newSet.add(element)
		}
	}

	return newSet
}

/* Set.union(Set) -> Set
/*
/* Returns the union of the set with the other set.
 */
func (setA Set) union(setB Set) Set {
	newSet := initSet()

	for element := range setA {
		if setA.isMember(element) {
			newSet.add(element)
		}
	}

	for element := range setB {
		if setB.isMember(element) {
			newSet.add(element)
		}
	}

	return newSet
}

/* Set.symmetricDifference(Set) -> Set
/*
/* Returns the difference of the union of the set and the other set with the */
/* intersection of the set and the other set.
 */
func (setA Set) symmetricDifference(setB Set) Set {
	return setA.union(setB).difference(setA.intersection(setB))
}

/* Set.disjoint(Set) -> bool
/*
/* Returns true if the set is disjoint with the other set.
 */
func (setA Set) disjoint(setB Set) bool {
	return setA.intersection(setB).isEmpty()
}

/* Set.isEmpty() -> bool
/*
/* Returns true if the set is empty.
 */
func (s Set) isEmpty() bool {
	if len(s) == 0 {
		return true
	} else {
		for element := range s {
			if s.isMember(element) {
				return false
			}
		}

		return true
	}
}

/* Set.print()
/*
/* Prints out the set's elements.
 */
func (s Set) print() {
	fmt.Println(s.toString())
}

/* Set.toString() -> string
/*
/* Returns a string of the set's elements.
 */
func (s Set) toString() string {
	var buf string
	buf = buf + "{"

	first := true

	for element := range s {
		if s[element] {
			if !first {
				buf = buf + fmt.Sprintf(", %s", element)
			} else {
				buf = buf + element
				first = false
			}
		}
	}
	buf = buf + "}"

	return buf
}
