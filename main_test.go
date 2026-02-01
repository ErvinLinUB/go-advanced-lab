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
