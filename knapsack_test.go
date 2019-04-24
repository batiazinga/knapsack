package knapsack

import (
	"strconv"
	"testing"
)

func TestDynamic(t *testing.T) {
	for i, golden := range goldens {
		// create a test instance
		ti := testInstance{
			golden: golden,
			solve:  Dynamic,
		}
		// run test
		t.Run(
			strconv.Itoa(i),
			ti.test,
		)
	}
}

func TestDynamicLight(t *testing.T) {
	for i, golden := range goldens {
		// create a test instance
		ti := testInstance{
			golden: golden,
			solve:  DynamicLight,
		}
		// run test
		t.Run(
			strconv.Itoa(i),
			ti.test,
		)
	}
}

func TestExhaustive(t *testing.T) {
	for i, golden := range goldens {
		// create a test instance
		ti := testInstance{
			golden: golden,
			solve:  Exhaustive,
		}
		// run test
		t.Run(
			strconv.Itoa(i),
			ti.test,
		)
	}
}

func BenchmarkDynamic(b *testing.B) {
	for i, golden := range goldens {
		// create a test instance
		ti := testInstance{
			golden: golden,
			solve:  Dynamic,
		}
		// run benchmark
		b.Run(
			strconv.Itoa(i),
			ti.bench,
		)
	}
}

func BenchmarkDynamicLight(b *testing.B) {
	for i, golden := range goldens {
		// create a test instance
		ti := testInstance{
			golden: golden,
			solve:  DynamicLight,
		}
		// run benchmark
		b.Run(
			strconv.Itoa(i),
			ti.bench,
		)
	}
}

func BenchmarkExhaustive(b *testing.B) {
	for i, golden := range goldens {
		// create a test instance
		ti := testInstance{
			golden: golden,
			solve:  Exhaustive,
		}
		// run benchmark
		b.Run(
			strconv.Itoa(i),
			ti.bench,
		)
	}
}

type testInstance struct {
	// data
	golden
	// function to test
	solve func(Interface) []bool
}

func (ti testInstance) test(t *testing.T) {
	// build Instance
	instance, err := NewInstance(ti.k, ti.values, ti.weights)
	if err != nil {
		t.Fatal(err)
	}

	// solve
	selection := ti.solve(instance)

	// check results
	if len(ti.solution) != len(selection) {
		t.Fatalf("selection has %v items instead of %v", len(selection), len(ti.solution))
	}
	for j := range ti.solution {
		if ti.solution[j] != selection[j] {
			t.Errorf("error on item %v", j)
		}
	}
}

func (ti testInstance) bench(b *testing.B) {
	// build instance
	instance, err := NewInstance(ti.k, ti.values, ti.weights)
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ti.solve(instance)
	}
}

type golden struct {
	// data
	k       int
	values  []int
	weights []int
	// solution
	solution []bool
}

var goldens = []golden{
	{
		k:        11,
		weights:  []int{4, 5, 8, 3},
		values:   []int{8, 10, 15, 4},
		solution: []bool{false, false, true, true},
	},

	// taken from
	// https://people.sc.fsu.edu/~jburkardt/f77_src/knapsack/knapsack_prb_output.txt
	// items are sorted by density, dense items first (density=value/weight)
	{
		k:        104,
		weights:  []int{25, 35, 45, 5, 25, 3, 2, 2},
		values:   []int{350, 400, 450, 20, 70, 8, 5, 5},
		solution: []bool{true, false, true, true, true, false, true, true},
	},
	// same as above with ten times larger capacity and weights
	// solution is the same
	// Dynamic algorithms should slower (roughly 10 times)
	// Exhaustive search should not be affected
	{
		k:        1040,
		weights:  []int{250, 350, 450, 50, 250, 30, 20, 20},
		values:   []int{350, 400, 450, 20, 70, 8, 5, 5},
		solution: []bool{true, false, true, true, true, false, true, true},
	},
	// same but items are reversed
	// solution is the same
	// Dynamic algorithms should not be affected
	// Exhaustive search should be slower
	{
		k:        104,
		weights:  []int{2, 2, 3, 25, 5, 45, 35, 25},
		values:   []int{5, 5, 8, 70, 20, 450, 400, 350},
		solution: []bool{true, true, false, true, true, true, false, true},
	},

	// another example with dense items first
	{
		k:        102,
		weights:  []int{2, 20, 20, 30, 40, 30, 60, 10},
		values:   []int{15, 100, 90, 60, 40, 15, 10, 1},
		solution: []bool{true, true, true, true, false, true, false, false},
	},
	// another example with dense items last
	{
		k:        102,
		weights:  []int{10, 60, 30, 40, 30, 20, 20, 2},
		values:   []int{1, 10, 15, 40, 60, 90, 100, 15},
		solution: []bool{false, false, true, false, true, true, true, true},
	},
}
