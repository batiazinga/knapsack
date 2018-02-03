package knapsack

import "github.com/batiazinga/bitarray"

// Dynamic is a dynamic programming solution.
// It is fast but costs a lot of memory: O(n * capacity) bytes.
//
// Note: to save memory, the matrix of changes instead of values is built
// (i.e. four/eight times cheaper).
func Dynamic(data Interface) []bool {
	// initialize the change array
	// items are rows (item 0 is row 1)
	// weights are columns (capacity is column k)
	a := newBoolArray2D(data.Len()+1, data.K()+1)

	// solve with dynamic algorithm
	return dynamic(data, a)
}

// DynamicLightSlow is a dynamic programming solution.
// It is similar to Dynamic
// but it is slower and uses asymptotically eight times less memory.
func DynamicLightSlow(data Interface) []bool {
	// initialize the change array
	// items are rows (item 0 is row 1)
	// weights are columns (capacity is column k)
	a := bitarray.New2D(data.Len()+1, data.K()+1)

	// solve with dynamic algorithm
	return dynamic(data, a)
}

func dynamic(data Interface, a boolGetSetter2D) []bool {
	// build the change array
	buildChangeArray(a, data)

	// build solution from change array
	return solveFromChangeArray(a, data)
}

// build the change array: n+1 * k+1 array of booleans.
// It MUST start from a boolSetter2D full of false.
func buildChangeArray(a boolSetter2D, data Interface) {
	currentValues := make([]int, data.K()+1)
	nextValues := make([]int, data.K()+1)

	// fill it
	//
	// first row: 0 values and no change by definition
	//
	// next rows:
	for i := 0; i < data.Len(); i++ {
		wi := data.Weight(i) // weight of i-th item
		vi := data.Value(i)  // value of the i-th item

		for w := 0; w != data.K()+1; w++ {
			if wi <= w {
				// item i can potentially be selected
				// next value is max(currentValue[j], currentValue[j-wi]+vi)
				if currentValues[w] < currentValues[w-wi]+vi {
					// item i is selected
					// update value...
					nextValues[w] = currentValues[w-wi] + vi
					// ... and record the change
					a.Set(i+1, w, true)
				} else {
					// item i is not selected
					// value does not change...
					nextValues[w] = currentValues[w]
					// ... and no change to record
				}
			} else {
				// item i cannot be selected
				// value does not change...
				nextValues[w] = currentValues[w]
				//... and no change to record
			}
		}

		// prepare for next iteration
		currentValues = nextValues
		nextValues = make([]int, data.K()+1) // reinit nextValues
	}
}

// solveFromChangeArray reads the change array
// to build the solution.
func solveFromChangeArray(a boolGetter2D, data Interface) []bool {
	// init solution: list of unselected items
	selection := make([]bool, data.Len())

	// start reading array at position (last row, last column)
	row := data.Len()
	column := data.K()

	// read
	for row != 0 && column != 0 {
		// move to the last selected item
		// for the given weight
		for row != 0 && !a.Get(row, column) {
			row--
		}
		if row == 0 {
			// no item to select
			continue
		}

		// item is selected
		selection[row-1] = true // item is row-1!!
		// lower weight...
		column -= data.Weight(row - 1)
		// ... and move to previous item
		row--
	}

	return selection
}

// boolGetter2D is a getter on a 2-dimensional array of booleans.
// It returns a boolean.
type boolGetter2D interface {
	Get(i, j int) bool
}

// boolSetter2D is a setter on a 2-dimensional array of booleans.
// It sets a boolean.
type boolSetter2D interface {
	Set(i, j int, b bool)
}

// boolGetSetter2D is both a getter and setter
// on a 2-dimensional array of booleans.
type boolGetSetter2D interface {
	boolGetter2D
	boolSetter2D
}

// boolArray2D is 2-dimensional boolean array.
// It implements both the boolGetter2D and boolSetter2D interfaces.
type boolArray2D struct {
	nColumns int
	content  []bool
}

// Get returns the value at position (i,j).
// Array must have been properly initialized.
func (m boolArray2D) Get(i, j int) bool {
	return m.content[i*m.nColumns+j]
}

// Set the value at position (i,j).
// Array must have been properly initialized.
func (m boolArray2D) Set(i, j int, b bool) {
	m.content[i*m.nColumns+j] = b
}

// newBoolArray2D returns a i*j boolArray2D.
func newBoolArray2D(i, j int) boolArray2D {
	return boolArray2D{
		nColumns: j,
		content:  make([]bool, i*j),
	}
}
