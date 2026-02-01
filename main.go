package main

import (
	"errors"
	"math"
)

/*----- Part 1: Table-Driven Tests & Math Operations -----*/

// 1. Factorial - calculates n!
func Factorial(n int) (int, error) {
	if n < 0 {
		return 0, errors.New("factorial is not defined for negative numbers")
	}

	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result, nil
}

// 2. IsPrime - checks if a number is prime
func IsPrime(n int) (bool, error) {
	if n < 2 {
		return false, errors.New("prime check requires number >= 2")
	}

	// Check if divisible by 2 (special case)
	if n == 2 {
		return true, nil
	}
	if n%2 == 0 {
		return false, nil
	}

	// Check odd divisors up to sqrt(n)
	limit := int(math.Sqrt(float64(n)))
	for i := 3; i <= limit; i += 2 {
		if n%i == 0 {
			return false, nil
		}
	}

	return true, nil
}

// 3. Power - calculates base^exponent
func Power(base, exponent int) (int, error) {
	if exponent < 0 {
		return 0, errors.New("negative exponents not supported")
	}

	result := 1
	for i := 0; i < exponent; i++ {
		result *= base
	}
	return result, nil
}

/*----- Part 2: Function Factory & Closures -----*/

// 1. MakeCounter - returns a closure that increments a counter
func MakeCounter(start int) func() int {
	count := start
	return func() int {
		count++
		return count
	}
}

// 2. MakeMultiplier - returns a closure that multiplies by captured factor
func MakeMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

// 3. MakeAccumulator - returns three closures sharing captured state
func MakeAccumulator(initial int) (func(int), func(int), func() int) {
	accumulator := initial

	add := func(amount int) {
		accumulator += amount
	}

	subtract := func(amount int) {
		accumulator -= amount
	}

	get := func() int {
		return accumulator
	}

	return add, subtract, get
}
