package main

import "testing"

/*----- Part 1: Table-Driven Tests & Math Operations -----*/

// 1. Factorial
func TestFactorial(t *testing.T) {
	tests := []struct {
		name    string
		input   int
		want    int
		wantErr bool
	}{
		// Add your test cases here
		{name: "factorial of 0", input: 0, want: 1, wantErr: false},
		// Add at least 5 more test cases
		{name: "factorial of 1", input: 1, want: 1, wantErr: false},
		{name: "factorial of 5", input: 5, want: 120, wantErr: false},
		{name: "factorial of 7", input: 7, want: 5040, wantErr: false},
		{name: "factorial of 10", input: 10, want: 3628800, wantErr: false},
		{name: "negative number -3", input: -3, want: 0, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Factorial(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Factorial() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Factorial() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 2. IsPrime
func TestIsPrime(t *testing.T) {
	tests := []struct {
		name    string
		input   int
		want    bool
		wantErr bool
	}{
		{name: "prime number 2", input: 2, want: true, wantErr: false},
		{name: "prime number 3", input: 3, want: true, wantErr: false},
		{name: "prime number 17", input: 17, want: true, wantErr: false},
		{name: "composite number 4", input: 4, want: false, wantErr: false},
		{name: "composite number 15", input: 15, want: false, wantErr: false},
		{name: "composite number 25", input: 25, want: false, wantErr: false},
		{name: "number 1", input: 1, want: false, wantErr: true},
		{name: "negative number -5", input: -5, want: false, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IsPrime(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsPrime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IsPrime() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 3. Power
func TestPower(t *testing.T) {
	tests := []struct {
		name     string
		base     int
		exponent int
		want     int
		wantErr  bool
	}{
		{name: "2^3", base: 2, exponent: 3, want: 8, wantErr: false},
		{name: "5^2", base: 5, exponent: 2, want: 25, wantErr: false},
		{name: "3^4", base: 3, exponent: 4, want: 81, wantErr: false},
		{name: "any^0", base: 10, exponent: 0, want: 1, wantErr: false},
		{name: "0^5", base: 0, exponent: 5, want: 0, wantErr: false},
		{name: "negative exponent", base: 2, exponent: -3, want: 0, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Power(tt.base, tt.exponent)
			if (err != nil) != tt.wantErr {
				t.Errorf("Power() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Power() = %v, want %v", got, tt.want)
			}
		})
	}
}

/*----- Part 2: Function Factory & Closures -----*/

// 1. TestMakeCounter
func TestMakeCounter(t *testing.T) {
	tests := []struct {
		name        string
		start       int
		increments  int
		wantResults []int
	}{
		{name: "counter from 0, 3 calls", start: 0, increments: 3, wantResults: []int{1, 2, 3}},
		{name: "counter from 10, 2 calls", start: 10, increments: 2, wantResults: []int{11, 12}},
		{name: "counter from -5, 4 calls", start: -5, increments: 4, wantResults: []int{-4, -3, -2, -1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			counter := MakeCounter(tt.start)

			for i := 0; i < tt.increments; i++ {
				got := counter()
				if got != tt.wantResults[i] {
					t.Errorf("MakeCounter() call %d = %v, want %v", i+1, got, tt.wantResults[i])
				}
			}
		})
	}

	// Test independence of counters
	t.Run("independent counters", func(t *testing.T) {
		counter1 := MakeCounter(0)
		counter2 := MakeCounter(100)

		// Call counter1 twice
		if got := counter1(); got != 1 {
			t.Errorf("counter1 first call = %v, want 1", got)
		}
		if got := counter1(); got != 2 {
			t.Errorf("counter1 second call = %v, want 2", got)
		}

		// Call counter2 once
		if got := counter2(); got != 101 {
			t.Errorf("counter2 first call = %v, want 101", got)
		}

		// Counter1 should still be at 3 on next call
		if got := counter1(); got != 3 {
			t.Errorf("counter1 third call = %v, want 3", got)
		}
	})
}

// 2. TestMakeMultiplier
func TestMakeMultiplier(t *testing.T) {
	tests := []struct {
		name   string
		factor int
		inputs []int
		wants  []int
	}{
		{name: "doubler", factor: 2, inputs: []int{5, 10, 0, -3}, wants: []int{10, 20, 0, -6}},
		{name: "tripler", factor: 3, inputs: []int{4, 7, 1, -2}, wants: []int{12, 21, 3, -6}},
		{name: "zero multiplier", factor: 0, inputs: []int{100, 50, -25}, wants: []int{0, 0, 0}},
		{name: "negative multiplier", factor: -5, inputs: []int{2, 3, -4}, wants: []int{-10, -15, 20}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			multiplier := MakeMultiplier(tt.factor)

			for i, input := range tt.inputs {
				got := multiplier(input)
				if got != tt.wants[i] {
					t.Errorf("MakeMultiplier(%d)(%d) = %v, want %v", tt.factor, input, got, tt.wants[i])
				}
			}
		})
	}
}

// 3. TestMakeAccumulator
func TestMakeAccumulator(t *testing.T) {
	tests := []struct {
		name       string
		initial    int
		operations []struct {
			op  string // "add" or "subtract"
			val int
		}
		wantAfterEach []int
	}{
		{
			name:    "basic accumulation",
			initial: 100,
			operations: []struct {
				op  string
				val int
			}{
				{op: "add", val: 50},
				{op: "subtract", val: 30},
				{op: "add", val: 20},
				{op: "subtract", val: 10},
			},
			wantAfterEach: []int{150, 120, 140, 130},
		},
		{
			name:    "negative operations",
			initial: 0,
			operations: []struct {
				op  string
				val int
			}{
				{op: "add", val: -10},
				{op: "subtract", val: -5}, // Subtracting negative = adding
				{op: "add", val: 20},
			},
			wantAfterEach: []int{-10, -5, 15},
		},
		{
			name:    "only adds",
			initial: 50,
			operations: []struct {
				op  string
				val int
			}{
				{op: "add", val: 10},
				{op: "add", val: 20},
				{op: "add", val: 30},
			},
			wantAfterEach: []int{60, 80, 110},
		},
		{
			name:    "only subtracts",
			initial: 100,
			operations: []struct {
				op  string
				val int
			}{
				{op: "subtract", val: 25},
				{op: "subtract", val: 15},
				{op: "subtract", val: 10},
			},
			wantAfterEach: []int{75, 60, 50},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			add, subtract, get := MakeAccumulator(tt.initial)

			// Check initial value
			if got := get(); got != tt.initial {
				t.Errorf("initial get() = %v, want %v", got, tt.initial)
			}

			// Perform operations and check after each
			for i, op := range tt.operations {
				switch op.op {
				case "add":
					add(op.val)
				case "subtract":
					subtract(op.val)
				default:
					t.Fatalf("unknown operation: %s", op.op)
				}

				if got := get(); got != tt.wantAfterEach[i] {
					t.Errorf("after operation %d (%s %d), get() = %v, want %v",
						i+1, op.op, op.val, got, tt.wantAfterEach[i])
				}
			}
		})
	}

	// Test that functions share state properly
	t.Run("shared state test", func(t *testing.T) {
		add, subtract, get := MakeAccumulator(0)

		add(100)
		if got := get(); got != 100 {
			t.Errorf("after add(100), get() = %v, want 100", got)
		}

		subtract(30)
		if got := get(); got != 70 {
			t.Errorf("after subtract(30), get() = %v, want 70", got)
		}

		add(50)
		if got := get(); got != 120 {
			t.Errorf("after add(50), get() = %v, want 120", got)
		}

		// Create another accumulator - should be independent
		add2, _, get2 := MakeAccumulator(200)
		add2(50)
		if got := get2(); got != 250 {
			t.Errorf("second accumulator after add(50), get() = %v, want 250", got)
		}

		// First accumulator should still be at 120
		if got := get(); got != 120 {
			t.Errorf("first accumulator after second accumulator operations, get() = %v, want 120", got)
		}
	})
}

/*----- Part 3: Higher-Order Functions -----*/

// 1. Apply
func TestApply(t *testing.T) {
	tests := []struct {
		name      string
		nums      []int
		operation func(int) int
		want      []int
	}{
		{
			name:      "squaring numbers",
			nums:      []int{1, 2, 3, 4, 5},
			operation: func(x int) int { return x * x },
			want:      []int{1, 4, 9, 16, 25},
		},
		{
			name:      "doubling numbers",
			nums:      []int{0, 1, 2, 3, 4},
			operation: func(x int) int { return x * 2 },
			want:      []int{0, 2, 4, 6, 8},
		},
		{
			name:      "negating numbers",
			nums:      []int{5, -5, 0, 10, -10},
			operation: func(x int) int { return -x },
			want:      []int{-5, 5, 0, -10, 10},
		},
		{
			name:      "adding 10 to each",
			nums:      []int{1, 2, 3},
			operation: func(x int) int { return x + 10 },
			want:      []int{11, 12, 13},
		},
		{
			name:      "empty slice",
			nums:      []int{},
			operation: func(x int) int { return x * 2 },
			want:      []int{},
		},
		{
			name: "absolute value",
			nums: []int{-3, -2, -1, 0, 1, 2, 3},
			operation: func(x int) int {
				if x < 0 {
					return -x
				}
				return x
			},
			want: []int{3, 2, 1, 0, 1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Apply(tt.nums, tt.operation)

			// Check length
			if len(got) != len(tt.want) {
				t.Errorf("Apply() length = %v, want %v", len(got), len(tt.want))
				return
			}

			// Check each element
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("Apply()[%d] = %v, want %v", i, got[i], tt.want[i])
				}
			}

			// Verify original slice is not modified
			if len(tt.nums) > 0 {
				originalCopy := make([]int, len(tt.nums))
				copy(originalCopy, tt.nums)
				Apply(tt.nums, tt.operation)
				for i := range tt.nums {
					if tt.nums[i] != originalCopy[i] {
						t.Errorf("Apply() modified original slice at index %d", i)
					}
				}
			}
		})
	}
}

// 2. Filter
func TestFilter(t *testing.T) {
	tests := []struct {
		name      string
		nums      []int
		predicate func(int) bool
		want      []int
	}{
		{
			name:      "even numbers",
			nums:      []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			predicate: func(x int) bool { return x%2 == 0 },
			want:      []int{2, 4, 6, 8, 10},
		},
		{
			name:      "positive numbers",
			nums:      []int{-5, -3, -1, 0, 1, 3, 5},
			predicate: func(x int) bool { return x > 0 },
			want:      []int{1, 3, 5},
		},
		{
			name:      "numbers greater than 10",
			nums:      []int{5, 10, 15, 20, 25, 30},
			predicate: func(x int) bool { return x > 10 },
			want:      []int{15, 20, 25, 30},
		},
		{
			name:      "negative numbers",
			nums:      []int{-10, -5, 0, 5, 10},
			predicate: func(x int) bool { return x < 0 },
			want:      []int{-10, -5},
		},
		{
			name:      "divisible by 3",
			nums:      []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			predicate: func(x int) bool { return x%3 == 0 },
			want:      []int{3, 6, 9},
		},
		{
			name:      "empty slice",
			nums:      []int{},
			predicate: func(x int) bool { return x > 0 },
			want:      []int{},
		},
		{
			name:      "no matches",
			nums:      []int{1, 2, 3, 4, 5},
			predicate: func(x int) bool { return x > 10 },
			want:      []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Filter(tt.nums, tt.predicate)

			// Check length
			if len(got) != len(tt.want) {
				t.Errorf("Filter() length = %v, want %v", len(got), len(tt.want))
				return
			}

			// Check each element
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("Filter()[%d] = %v, want %v", i, got[i], tt.want[i])
				}
			}
		})
	}
}

// 3. Reduce
func TestReduce(t *testing.T) {
	tests := []struct {
		name      string
		nums      []int
		initial   int
		operation func(accumulator, current int) int
		want      int
	}{
		{
			name:      "sum of numbers",
			nums:      []int{1, 2, 3, 4, 5},
			initial:   0,
			operation: func(acc, curr int) int { return acc + curr },
			want:      15,
		},
		{
			name:      "product of numbers",
			nums:      []int{1, 2, 3, 4},
			initial:   1,
			operation: func(acc, curr int) int { return acc * curr },
			want:      24,
		},
		{
			name:    "find maximum",
			nums:    []int{5, 2, 9, 1, 7},
			initial: -1000000, // Very small initial value
			operation: func(acc, curr int) int {
				if curr > acc {
					return curr
				}
				return acc
			},
			want: 9,
		},
		{
			name:    "find minimum",
			nums:    []int{5, 2, 9, 1, 7},
			initial: 1000000, // Very large initial value
			operation: func(acc, curr int) int {
				if curr < acc {
					return curr
				}
				return acc
			},
			want: 1,
		},
		{
			name:    "concatenate as string length",
			nums:    []int{1, 23, 456},
			initial: 0,
			operation: func(acc, curr int) int {
				// Count digits (simplified)
				count := 0
				n := curr
				if n == 0 {
					return acc + 1
				}
				if n < 0 {
					n = -n
				}
				for n > 0 {
					count++
					n /= 10
				}
				return acc + count
			},
			want: 6, // 1 + 2 + 3 digits
		},
		{
			name:      "empty slice returns initial",
			nums:      []int{},
			initial:   42,
			operation: func(acc, curr int) int { return acc + curr },
			want:      42,
		},
		{
			name:    "count even numbers",
			nums:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			initial: 0,
			operation: func(acc, curr int) int {
				if curr%2 == 0 {
					return acc + 1
				}
				return acc
			},
			want: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Reduce(tt.nums, tt.initial, tt.operation)
			if got != tt.want {
				t.Errorf("Reduce() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 4. Compose
func TestCompose(t *testing.T) {
	tests := []struct {
		name string
		f    func(int) int
		g    func(int) int
		x    int
		want int
	}{
		{
			name: "double then add 10",
			f:    func(x int) int { return x + 10 },
			g:    func(x int) int { return x * 2 },
			x:    5,
			want: 20, // (5*2) + 10 = 20
		},
		{
			name: "add 5 then square",
			f:    func(x int) int { return x * x },
			g:    func(x int) int { return x + 5 },
			x:    3,
			want: 64, // (3+5)² = 64
		},
		{
			name: "negate then absolute (custom)",
			f: func(x int) int {
				if x < 0 {
					return -x
				}
				return x
			},
			g:    func(x int) int { return -x },
			x:    7,
			want: 7, // -7 then absolute = 7
		},
		{
			name: "identity functions",
			f:    func(x int) int { return x },
			g:    func(x int) int { return x },
			x:    42,
			want: 42,
		},
		{
			name: "multiple negations",
			f:    func(x int) int { return -x },
			g:    func(x int) int { return -x },
			x:    10,
			want: 10, // -(-10) = 10
		},
		{
			name: "zero handling",
			f:    func(x int) int { return x + 100 },
			g:    func(x int) int { return x * 0 },
			x:    5,
			want: 100, // 5*0 = 0, then 0+100 = 100
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			composed := Compose(tt.f, tt.g)
			got := composed(tt.x)
			if got != tt.want {
				t.Errorf("Compose(f,g)(%d) = %v, want %v", tt.x, got, tt.want)
			}

			// Verify composition works as f(g(x))
			manualResult := tt.f(tt.g(tt.x))
			if got != manualResult {
				t.Errorf("Compose(f,g)(%d) = %v, but f(g(%d)) = %v", tt.x, got, tt.x, manualResult)
			}
		})
	}
}

/*----- Part 5: Pointer Playground & Escape Analysis -----*/

// TestSwapValues tests the SwapValues function
func TestSwapValues(t *testing.T) {
	tests := []struct {
		name  string
		a     int
		b     int
		wantA int
		wantB int
	}{
		{
			name:  "swap positive numbers",
			a:     5,
			b:     10,
			wantA: 10,
			wantB: 5,
		},
		{
			name:  "swap negative numbers",
			a:     -3,
			b:     -7,
			wantA: -7,
			wantB: -3,
		},
		{
			name:  "swap mixed numbers",
			a:     15,
			b:     -20,
			wantA: -20,
			wantB: 15,
		},
		{
			name:  "swap with zero",
			a:     0,
			b:     100,
			wantA: 100,
			wantB: 0,
		},
		{
			name:  "swap same values",
			a:     42,
			b:     42,
			wantA: 42,
			wantB: 42,
		},
		{
			name:  "swap large numbers",
			a:     999,
			b:     -1000,
			wantA: -1000,
			wantB: 999,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotA, gotB := SwapValues(tt.a, tt.b)
			if gotA != tt.wantA || gotB != tt.wantB {
				t.Errorf("SwapValues(%d, %d) = (%d, %d), want (%d, %d)",
					tt.a, tt.b, gotA, gotB, tt.wantA, tt.wantB)
			}
		})
	}
}

// TestSwapPointers tests the SwapPointers function
func TestSwapPointers(t *testing.T) {
	tests := []struct {
		name  string
		a     int
		b     int
		wantA int
		wantB int
	}{
		{
			name:  "swap positive numbers",
			a:     5,
			b:     10,
			wantA: 10,
			wantB: 5,
		},
		{
			name:  "swap negative numbers",
			a:     -3,
			b:     -7,
			wantA: -7,
			wantB: -3,
		},
		{
			name:  "swap mixed numbers",
			a:     15,
			b:     -20,
			wantA: -20,
			wantB: 15,
		},
		{
			name:  "swap with zero",
			a:     0,
			b:     100,
			wantA: 100,
			wantB: 0,
		},
		{
			name:  "swap same values",
			a:     42,
			b:     42,
			wantA: 42,
			wantB: 42,
		},
		{
			name:  "swap large numbers",
			a:     999,
			b:     -1000,
			wantA: -1000,
			wantB: 999,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create original variables
			originalA := tt.a
			originalB := tt.b

			// Take their addresses
			aPtr := &originalA
			bPtr := &originalB

			// Call SwapPointers
			SwapPointers(aPtr, bPtr)

			// Check that the original variables were modified
			if originalA != tt.wantA || originalB != tt.wantB {
				t.Errorf("SwapPointers(%d, %d) resulted in (%d, %d), want (%d, %d)",
					tt.a, tt.b, originalA, originalB, tt.wantA, tt.wantB)
			}

			// Verify that pointers still point to the same memory locations
			if aPtr != &originalA {
				t.Errorf("aPtr changed memory location after SwapPointers")
			}
			if bPtr != &originalB {
				t.Errorf("bPtr changed memory location after SwapPointers")
			}
		})
	}
}
