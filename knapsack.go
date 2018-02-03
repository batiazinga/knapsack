package knapsack

import "fmt"

// Interface is a collection of items for the 0-1 Knapsack problem.
// Values and weights must be strictly positive integers.
type Interface interface {
	// Len returns the number of items.
	Len() int
	// K returns the capacity of the knapsack.
	K() int
	// Value returns the value of the i-th item.
	Value(i int) int
	// Weight returns the weight of the i-th item.
	Weight(i int) int
}

// Instance of the 0-1 Knapsack problem.
// It implements Interface.
type Instance struct {
	// capacity
	k int
	// values ane weights (order matters)
	values, weights []int
}

func (inst Instance) Len() int         { return len(inst.values) }
func (inst Instance) K() int           { return inst.k }
func (inst Instance) Value(i int) int  { return inst.values[i] }
func (inst Instance) Weight(i int) int { return inst.weights[i] }

// NewInstance returns an instance of the Knapsack problem or an error if data are invalid.
// It checks that data are valid: Capacity, values and weights must be strictly positive and
// number of items, values and weights must be consistent.
func NewInstance(capacity int, values, weights []int) (inst Instance, err error) {
	// ckeck number of items
	if len(values) != len(weights) {
		err = fmt.Errorf("inconsistent number of items: %v values and %v weights", len(values), len(weights))
		return
	}
	// check capacity
	if capacity <= 0 {
		err = fmt.Errorf("invalid (%v) capacity", capacity)
		return
	}
	// check values
	for i, v := range values {
		if v <= 0 {
			err = fmt.Errorf("item %v has invalid value %v", i, v)
			return
		}
	}
	// check weights
	for i, w := range weights {
		if w <= 0 {
			err = fmt.Errorf("item %v has invalid weight %v", i, w)
			return
		}
	}

	return Instance{
		k:       capacity,
		values:  values,
		weights: weights,
	}, nil
}
