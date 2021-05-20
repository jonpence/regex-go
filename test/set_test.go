package main

import (
	"testing"
)

// if condition {
//     t.Errorf("error message")
// }


func TestInitSet(t *testing.T) {
	testSet := initSet()

	if testSet == nil {
		t.Errorf("testSet is nil, expected non-nil value")
	}
}

func TestInitSetElements(t *testing.T) {
	testSlice := []string{"0", "1", "2", "3", "4"}
	testSet   := initSetElements(testSlice)

	if !testSet.isMember("0") {
		t.Errorf("expected testSet to contain 0")
	}

	if !testSet.isMember("1") {
		t.Errorf("expected testSet to contain 1")
	}

	if !testSet.isMember("2") {
		t.Errorf("expected testSet to contain 2")
	}

	if !testSet.isMember("3") {
		t.Errorf("expected testSet to contain 3")
	}

	if !testSet.isMember("4") {
		t.Errorf("expected testSet to contain 4")
	}
}

func TestSize(t *testing.T) {
	testSlice := []string{"0", "1", "2", "3", "4"}
	testSet   := initSetElements(testSlice)

	if testSet.size() != 5 {
		t.Errorf("testSize size is %d, expected 5", testSet.size())
	}
}

func TestInitSetRange(t *testing.T) {
	testSet := initSetRange(0, 5)

	if !testSet.isMember("0") {
		t.Errorf("expected testSet to contain 0")
	}

	if !testSet.isMember("1") {
		t.Errorf("expected testSet to contain 1")
	}

	if !testSet.isMember("2") {
		t.Errorf("expected testSet to contain 2")
	}

	if !testSet.isMember("3") {
		t.Errorf("expected testSet to contain 3")
	}

	if !testSet.isMember("4") {
		t.Errorf("expected testSet to contain 4")
	}

	if testSet.isMember("5") {
		t.Errorf("did not expect testSet to contain 5")
	}
}

func TestIsMember(t *testing.T) {
	testSet := initSetElements([]string{"0"})

	if !testSet.isMember("0") {
		t.Errorf("expected isMember() to report that 0 is a member of testSet")
	}
}

func TestAdd(t *testing.T) {
	testSet := initSet()
	testSet.add("0")

	if !testSet.isMember("0") {
		t.Errorf("expected isMember() to report that 0 is a member of testSet")
	}
}

func TestMultiadd(t *testing.T) {
	testSlice := []string{"0", "1", "2", "3", "4"}
	testSet   := initSet()

	testSet.multiadd(testSlice)

	if !testSet.isMember("0") {
		t.Errorf("expected testSet to contain 0")
	}

	if !testSet.isMember("1") {
		t.Errorf("expected testSet to contain 1")
	}

	if !testSet.isMember("2") {
		t.Errorf("expected testSet to contain 2")
	}

	if !testSet.isMember("3") {
		t.Errorf("expected testSet to contain 3")
	}

	if !testSet.isMember("4") {
		t.Errorf("expected testSet to contain 4")
	}
}

func TestRemove(t *testing.T) {
	testSet := initSet()
	testSet.add("0")

	if !testSet.remove("0") {
		t.Errorf("expected return value of true since 0 is element of testSet")
	}

	if testSet.remove("1") {
		t.Errorf("expected return value of false since 1 is not an element of testSet")
	}
}

func TestDiscard(t *testing.T) {
	testSet := initSet()
	testSet.add("0")

	testSet.discard("0")
	testSet.discard("1")

	if testSet.isMember("0") {
		t.Errorf("expected false as 0 is discarded")
	}

	if testSet.isMember("1") {
		t.Errorf("expected false as 1 is discarded and was never a member of testSet")
	}
}

func TestPop(t *testing.T) {
	testSet := initSetRange(1, 11)

	for testSet.pop() { }

	if !testSet.isEmpty() {
		t.Errorf("expected testSet to be empty")
	}
}

func TestClear(t *testing.T) {
	testSet := initSetRange(0, 101)

	testSet.clear()

	if !testSet.isEmpty() {
		t.Errorf("expected testSet to be empty")
	}
}

func TestCopy(t *testing.T) {
	testSetA := initSetRange(0, 11)
	testSetB := testSetA.copy()

	if !testSetA.equivalent(testSetB) {
		t.Errorf("expected testSetA to be equivalent with testSetB")
	}
}

func TestIntersection(t *testing.T) {
	testSetA := initSetRange(1, 11)
	testSetB := initSetRange(5, 16)
	testSetC := testSetA.intersection(testSetB)
	testSetD := initSetRange(5, 11)

	if !testSetC.equivalent(testSetD) {
		t.Errorf("expected intersection to contain 5 to 10 inclusive")
	}
}

func TestSubset(t *testing.T) {
	testSetA := initSetRange(0, 11)
	testSetB := initSetRange(0, 11)
	testSetC := initSetRange(3, 6)

	if !testSetA.subset(testSetB) {
		t.Errorf("expected testSetA to be subset of testSetB")
	}

	if !testSetC.subset(testSetA) {
		t.Errorf("expected testSetC to be subset of testSetA")
	}
}

func TestSuperset(t *testing.T) {
	testSetA := initSetRange(0, 11)
	testSetB := initSetRange(0, 11)
	testSetC := initSetRange(3, 6)

	if !testSetA.superset(testSetB) {
		t.Errorf("expected testSetA to be superset of testSetB")
	}

	if !testSetA.superset(testSetC) {
		t.Errorf("expected testSetA to be superset of testSetC")
	}
}

func TestProperSubset(t *testing.T) {
	testSetA := initSetRange(0, 11)
	testSetB := initSetRange(0, 11)
	testSetC := initSetRange(3, 6)

	if testSetA.properSubset(testSetB) {
		t.Errorf("did not expect testSetA to be proper subset of testSetB")
	}

	if !testSetC.properSubset(testSetA) {
		t.Errorf("expected testSetC to be subset of testSetA")
	}
}

func TestProperSuperset(t *testing.T) {
	testSetA := initSetRange(0, 11)
	testSetB := initSetRange(0, 11)
	testSetC := initSetRange(3, 6)

	if testSetA.properSuperset(testSetB) {
		t.Errorf("did not expect for testSetA to be proper superset of testSetB")
	}

	if !testSetA.properSuperset(testSetC) {
		t.Errorf("expected testSetA to be proper superset of testSetC")
	}
}

func TestEquivalent(t *testing.T) {
	testSetA := initSetRange(0, 101)
	testSetB := initSetRange(0, 101)
	testSetC := initSetRange(1, 100)
	testSetD := initSetRange(200, 301)

	if !testSetA.equivalent(testSetB) {
		t.Errorf("expected testSetA to be equivalent with testSetB")
	}

	if testSetA.equivalent(testSetC) {
		t.Errorf("did not expect testSetA to be equivalent with testSetC")
	}

	if testSetA.equivalent(testSetD) {
		t.Errorf("did not expect testSetA to be equivalent with testSetD")
	}
}

func TestDifference(t *testing.T) {
	testSetA := initSetRange(1, 21)
	testSetB := initSetRange(1, 11)
	testSetC := initSetRange(10, 21)

	if !testSetC.equivalent(testSetA.difference(testSetB)) {
		t.Errorf("expected testSetC to be equivalent with testSetA - testSetB")
	}
}

func TestUnion(t *testing.T) {
	testSetA := initSetRange(1, 11)
	testSetB := initSetRange(11, 21)
	testSetC := initSetRange(1, 21)

	if !testSetC.equivalent(testSetA.union(testSetB)) {
		t.Errorf("expected testSetC to be equivalent with testSetA U testSetB")
	}
}

func TestSymmetricDifference(t *testing.T) {
	testSetA := initSetRange(1, 16)
	testSetB := initSetRange(10, 21)

	testSlice := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "16", "17", "18", "19", "20"}
	testSetC  := initSetElements(testSlice)

	if !testSetC.equivalent(testSetA.symmetricDifference(testSetB)) {
		t.Errorf("Expected testSetC to be equivalent with symmetric difference of testSetA and testSetB")
	}
}

func TestDisjoint(t *testing.T) {
	testSetA := initSetRange(1, 11)
	testSetB := initSetRange(5, 21)
	testSetC := initSetRange(15, 30)

	if testSetA.disjoint(testSetB) {
		t.Errorf("did not expect testSetA to be disjoint from testSetB")
	}

	if testSetB.disjoint(testSetC) {
		t.Errorf("did not expect testSetB to be disjoint from testSetA")
	}

	if !testSetA.disjoint(testSetC) {
		t.Errorf("expected testSetA to be disjoint from testSetC")
	}
}

func TestPrint(t *testing.T) {
	testSetA := initSetRange(0, 21)
	testSetA.print()
}

func TestToString(t *testing.T) {
	testSetA := initSetRange(1, 6)

	if "{1, 2, 3, 4, 5, 6}" != testSetA.toString() {
		t.Errorf("expeted testSetA's string to be {1, 2, 3, 4, 5, 6}")
	}
}
