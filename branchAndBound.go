package knapsack

// Exhaustive performs an (almost) exhaustive search of all solutions.
// It is guaranted to find the optimal solution.
func Exhaustive(data Interface) []bool {
	// backtrack
	// init best solution so far: zero item solution
	bestSoFar := &solution{
		selection: make([]bool, data.Len()),
	}
	// start with no item selected
	current := &solution{
		selection: make([]bool, data.Len()),
	}
	for i := 0; i != data.Len(); i++ {
		branchBound(data, i, current, bestSoFar)
	}

	// return the solution
	return bestSoFar.selection
}

// branchBound adds item 'index' to the bag.
// If a solution is found it compares it to the current best solution.
// If a solution is not found yet it recurses until it finds a solution.
func branchBound(data Interface, item int, current, bestSoFar *solution) {
	if item >= data.Len() {
		// a solution has been found
		// save it if it is better than the best solution so far
		if current.value > bestSoFar.value {
			// copy current solution into best solution so far
			for i, selected := range current.selection {
				bestSoFar.selection[i] = selected
			}
			bestSoFar.value = current.value
		}
		return
	}

	// put selected item in the bag
	current.selectItem(data, item)
	// does this exceed the capacity?
	// is this worth going on?
	if current.weight > data.K() || current.value+maxRemainingValue(data, item) <= bestSoFar.value {
		// with this item, we exceed the capacity
		// or there is no chance to find a better solution
		// backtrack: unselect item and return
		current.unselectItem(data, item)
		return
	}
	// with this new selected item there is a change to find a better solution

	// recursive: choose next item to select
	for i := item; i < data.Len(); i++ {
		branchBound(data, i+1, current, bestSoFar)
	}

	// backtrack: unselect item and return
	current.unselectItem(data, item)
}

// maxRemainingValue computes the value which can be added to the bag
// if all items strictly after item are selected.
func maxRemainingValue(data Interface, item int) int {
	remain := 0
	for i := item + 1; i < data.Len(); i++ {
		remain += data.Value(i)
	}
	return remain
}

type solution struct {
	selection []bool
	// value and weight of the selection
	value, weight int
}

func (s *solution) selectItem(data Interface, i int) {
	s.selection[i] = true
	s.value += data.Value(i)
	s.weight += data.Weight(i)
}
func (s *solution) unselectItem(data Interface, i int) {
	s.selection[i] = false
	s.value -= data.Value(i)
	s.weight -= data.Weight(i)
}
