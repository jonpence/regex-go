package set

import (
	"testing"
)

func TestInitSet(t *testing.T) {
	testSet := InitSet()

	if testSet == nil {
		t.Errorf("testSet is nil, expected non-nil value")
	}
}

func TestInitSetElements(t *testing.T) {
	testSlice := []string{"0", "1", "2", "3", "4"}
	testSet   := InitSetElements(testSlice)

	if !testSet.IsMember("0") {
		t.Errorf("expected testSet to contain 0")
	}

	if !testSet.IsMember("1") {
		t.Errorf("expected testSet to contain 1")
	}

	if !testSet.IsMember("2") {
		t.Errorf("expected testSet to contain 2")
	}

	if !testSet.IsMember("3") {
		t.Errorf("expected testSet to contain 3")
	}

	if !testSet.IsMember("4") {
		t.Errorf("expected testSet to contain 4")
	}
}

func TestSize(t *testing.T) {
	testSlice := []string{"0", "1", "2", "3", "4"}
	testSet   := InitSetElements(testSlice)

	if testSet.Size() != 5 {
		t.Errorf("testSize size is %d, expected 5", testSet.Size())
	}
}

func TestInitSetRange(t *testing.T) {
	testSet := InitSetRange(0, 5)

	if !testSet.IsMember("0") {
		t.Errorf("expected testSet to contain 0")
	}

	if !testSet.IsMember("1") {
		t.Errorf("expected testSet to contain 1")
	}

	if !testSet.IsMember("2") {
		t.Errorf("expected testSet to contain 2")
	}

	if !testSet.IsMember("3") {
		t.Errorf("expected testSet to contain 3")
	}

	if !testSet.IsMember("4") {
		t.Errorf("expected testSet to contain 4")
	}

	if testSet.IsMember("5") {
		t.Errorf("did not expect testSet to contain 5")
	}
}

func TestIsMember(t *testing.T) {
	testSet := InitSetElements([]string{"0"})

	if !testSet.IsMember("0") {
		t.Errorf("expected IsMember() to report that 0 is a member of testSet")
	}
}

func TestAdd(t *testing.T) {
	testSet := InitSet()
	testSet.Add("0")

	if !testSet.IsMember("0") {
		t.Errorf("expected IsMember() to report that 0 is a member of testSet")
	}
}

func TestMultiadd(t *testing.T) {
	testSlice := []string{"0", "1", "2", "3", "4"}
	testSet   := InitSet()

	testSet.Multiadd(testSlice)

	if !testSet.IsMember("0") {
		t.Errorf("expected testSet to contain 0")
	}

	if !testSet.IsMember("1") {
		t.Errorf("expected testSet to contain 1")
	}

	if !testSet.IsMember("2") {
		t.Errorf("expected testSet to contain 2")
	}

	if !testSet.IsMember("3") {
		t.Errorf("expected testSet to contain 3")
	}

	if !testSet.IsMember("4") {
		t.Errorf("expected testSet to contain 4")
	}
}

func TestRemove(t *testing.T) {
	testSet := InitSet()
	testSet.Add("0")

	if !testSet.Remove("0") {
		t.Errorf("expected return value of true since 0 is element of testSet")
	}

	if testSet.Remove("1") {
		t.Errorf("expected return value of false since 1 is not an element of testSet")
	}
}

func TestDiscard(t *testing.T) {
	testSet := InitSet()
	testSet.Add("0")

	testSet.Discard("0")
	testSet.Discard("1")

	if testSet.IsMember("0") {
		t.Errorf("expected false as 0 is Discarded")
	}

	if testSet.IsMember("1") {
		t.Errorf("expected false as 1 is Discarded and was never a member of testSet")
	}
}

func TestPop(t *testing.T) {
	testSet := InitSetRange(1, 11)

	for _, ok := testSet.Pop(); ok; _, ok = testSet.Pop() { }

	if !testSet.IsEmpty() {
		t.Errorf("expected testSet to be empty")
	}
}

func TestClear(t *testing.T) {
	testSet := InitSetRange(0, 101)

	testSet.Clear()

	if !testSet.IsEmpty() {
		t.Errorf("expected testSet to be empty")
	}
}

func TestCopy(t *testing.T) {
	testSetA := InitSetRange(0, 11)
	testSetB := testSetA.Copy()

	if !testSetA.Equivalent(testSetB) {
		t.Errorf("expected testSetA to be Equivalent with testSetB")
	}
}

func TestIntersection(t *testing.T) {
	testSetA := InitSetRange(1, 11)
	testSetB := InitSetRange(5, 16)
	testSetC := testSetA.Intersection(testSetB)
	testSetD := InitSetRange(5, 11)

	if !testSetC.Equivalent(testSetD) {
		t.Errorf("expected intersection to contain 5 to 10 inclusive")
	}
}

func TestSubset(t *testing.T) {
	testSetA := InitSetRange(0, 11)
	testSetB := InitSetRange(0, 11)
	testSetC := InitSetRange(3, 6)

	if !testSetA.Subset(testSetB) {
		t.Errorf("expected testSetA to be Subset of testSetB")
	}

	if !testSetC.Subset(testSetA) {
		t.Errorf("expected testSetC to be subset of testSetA")
	}
}

func TestSuperset(t *testing.T) {
	testSetA := InitSetRange(0, 11)
	testSetB := InitSetRange(0, 11)
	testSetC := InitSetRange(3, 6)

	if !testSetA.Superset(testSetB) {
		t.Errorf("expected testSetA to be superset of testSetB")
	}

	if !testSetA.Superset(testSetC) {
		t.Errorf("expected testSetA to be superset of testSetC")
	}
}

func TestProperSubset(t *testing.T) {
	testSetA := InitSetRange(0, 11)
	testSetB := InitSetRange(0, 11)
	testSetC := InitSetRange(3, 6)

	if testSetA.ProperSubset(testSetB) {
		t.Errorf("did not expect testSetA to be proper subset of testSetB")
	}

	if !testSetC.ProperSubset(testSetA) {
		t.Errorf("expected testSetC to be subset of testSetA")
	}
}

func TestProperSuperset(t *testing.T) {
	testSetA := InitSetRange(0, 11)
	testSetB := InitSetRange(0, 11)
	testSetC := InitSetRange(3, 6)

	if testSetA.ProperSuperset(testSetB) {
		t.Errorf("did not expect for testSetA to be proper superset of testSetB")
	}

	if !testSetA.ProperSuperset(testSetC) {
		t.Errorf("expected testSetA to be proper superset of testSetC")
	}
}

func TestEquivalent(t *testing.T) {
	testSetA := InitSetRange(0, 101)
	testSetB := InitSetRange(0, 101)
	testSetC := InitSetRange(1, 100)
	testSetD := InitSetRange(200, 301)

	if !testSetA.Equivalent(testSetB) {
		t.Errorf("expected testSetA to be equivalent with testSetB")
	}

	if testSetA.Equivalent(testSetC) {
		t.Errorf("did not expect testSetA to be equivalent with testSetC")
	}

	if testSetA.Equivalent(testSetD) {
		t.Errorf("did not expect testSetA to be equivalent with testSetD")
	}
}

func TestDifference(t *testing.T) {
	testSetA := InitSetRange(1, 21)
	testSetB := InitSetRange(1, 10)
	testSetC := InitSetRange(10, 21)

	if !testSetC.Equivalent(testSetA.Difference(testSetB)) {
		t.Errorf("expected testSetC to be equivalent with testSetA - testSetB, instead:\n\ttestSetC is \n\t\t%s\n\ttestSetA - testSetB is\n\t\t%s", testSetC.ToString(), testSetA.Difference(testSetB).ToString())
	}
}

func TestUnion(t *testing.T) {
	testSetA := InitSetRange(1, 11)
	testSetB := InitSetRange(11, 21)
	testSetC := InitSetRange(1, 21)

	if !testSetC.Equivalent(testSetA.Union(testSetB)) {
		t.Errorf("expected testSetC to be Equivalent with testSetA U testSetB")
	}
}

func TestSymmetricDifference(t *testing.T) {
	testSetA := InitSetRange(1, 16)
	testSetB := InitSetRange(10, 21)

	testSlice := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "16", "17", "18", "19", "20"}
	testSetC  := InitSetElements(testSlice)

	if !testSetC.Equivalent(testSetA.SymmetricDifference(testSetB)) {
		t.Errorf("Expected testSetC to be Equivalent with Symmetric Difference of testSetA and testSetB")
	}
}

func TestDisjoint(t *testing.T) {
	testSetA := InitSetRange(1, 11)
	testSetB := InitSetRange(5, 21)
	testSetC := InitSetRange(15, 30)

	if testSetA.Disjoint(testSetB) {
		t.Errorf("did not expect testSetA to be Disjoint from testSetB")
	}

	if testSetB.Disjoint(testSetC) {
		t.Errorf("did not expect testSetB to be Disjoint from testSetA")
	}

	if !testSetA.Disjoint(testSetC) {
		t.Errorf("expected testSetA to be Disjoint from testSetC")
	}
}
