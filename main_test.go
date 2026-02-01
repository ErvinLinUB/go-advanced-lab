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
