package regex

import (
	"math"
	"github.com/jonpence/regex-go/internal/set"
	"github.com/jonpence/regex-go/internal/deque")

type Automaton struct {
	states  map[string]*State
	start   *State
	current *State
	inputs  set.Set
	count   int
}

func initAutomaton() *Automaton {
	return &Automaton{make(map[string]*State), nil, nil, set.InitSet(), 1}
}

func (a *Automaton) addState(s *State) {
	a.states[s.name] = s
	a.count += 1
}

func (a *Automaton) addEdge(src string, dest string, input string) {
	a.states[src].addNeighbor(dest, input)
}

func (a *Automaton) setStart(s *State) {
	a.start = s
}

func (a *Automaton) setCurrent(s *State) {
	a.current = s
}

func (a Automaton) getState(input string) *State {
	return a.states[input]
}

func (a Automaton) getCount() int {
	return a.count
}

func (a *Automaton) reset() {
	a.setCurrent(a.start)
}

func (a *Automaton) transition(char byte) bool {
	dest, present := a.current.getNeighbor(string(char))

	if present {
		a.setCurrent(a.states[dest])
		return true
	}

	return false
}

func max(x int, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func pair(x int, y int) int {
	if x == max(x, y) {
		return x * x + x + y
	} else {
		return y * y + x
	}
}

func unpair(z int) (int, int) {
	floatZ := float64(z)

	floorSqrt := math.Floor(math.Sqrt(floatZ))
	if floatZ - math.Pow(floorSqrt, 2) < floorSqrt {
		return int(floatZ - math.Pow(floorSqrt, 2)), int(floorSqrt)
	} else {
		return int(floorSqrt), int(floatZ - math.Pow(floorSqrt, 2) - floorSqrt)
	}
}

func stoi(input string) int {
	num := 0

	for i := 0; i < len(input); i++ {
		num *= 10
		num += int(input[i] - 48)
	}

	return num
}

func pairN(nums []string, n int) int {
	if len(nums) == 1 {
		return stoi(nums[0])
	}

	z := pair(stoi(nums[0]), stoi(nums[1]))

	for i := 2; i < n; i++ {
		if i >= len(nums) {
			z = pair(z, 0)
		} else {
			z = pair(z, stoi(nums[i]))
		}
	}

	return z;
}

func (a *Automaton) closureOfOn(state *State, input string) *State {
	newComposites := set.InitSet()
	terminates    := false

	// initialize composites with all reachable states from state on input
	for composite := range state.composites {
		neighbors, _ := a.getState(composite).getNeighborList(input)
		newComposites.Multiadd(neighbors)
	}

	// reach all nil-reachable states
	addedNew := true
	for addedNew == true {
		addedNew = false

		for newComposite := range newComposites {
			nullReachable, _ := a.getState(newComposite).getNeighborList("")

			for _, nullReached := range nullReachable {
				if !newComposites.IsMember(nullReached) {
					addedNew = true
					newComposites.Add(nullReached)
				}
			}
		}
	}

	if newComposites.IsEmpty() {
		return nil
	}

	for newComposite := range newComposites {
		if a.getState(newComposite).terminates {
			terminates = true
		}
	}

	newState := initDFAState(newComposites, itos(pairN(newComposites.ToSlice(), len(a.states))))

	if terminates {
		newState.setTerminates()
	}

	return newState
}

func (a *Automaton) determinize() *Automaton {
	dfa   := initAutomaton()
	queue := deque.InitQueue()

	initialState := a.closureOfOn(a.start, "")

	dfa.addState(initialState)
	dfa.setStart(initialState)

	queue.Enqueue(initialState.name)

	for {
		currentState, ok := queue.Dequeue()

		if !ok {
			break
		}

		for input := range a.inputs {
			newState := a.closureOfOn(dfa.getState(currentState), input)

			if newState == nil {
				continue
			}

			if _, present := dfa.states[newState.name]; !present {
				dfa.addState(newState)
				queue.Enqueue(newState.name)
			}

			dfa.addEdge(currentState, newState.name, input)
		}
	}

	return dfa
}

func (a Automaton) validate(input string) bool {
	a.reset()

	for i := 0; i < len(input); i++ {
		if !a.transition(input[i]) {
			return false
		}
	}

	return a.current.terminates
}
