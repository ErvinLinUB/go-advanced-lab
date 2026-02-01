package main

import (
	"errors"
	"math"
	"os"
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

/*----- Part 3: Higher-Order Functions -----*/

// 1. Apply - applies operation to each element in slice
func Apply(nums []int, operation func(int) int) []int {
	result := make([]int, len(nums))
	for i, num := range nums {
		result[i] = operation(num)
	}
	return result
}

// 2. Filter - returns elements where predicate is true
func Filter(nums []int, predicate func(int) bool) []int {
	result := []int{}
	for _, num := range nums {
		if predicate(num) {
			result = append(result, num)
		}
	}
	return result
}

// 3. Reduce - reduces slice to single value using operation
func Reduce(nums []int, initial int, operation func(accumulator, current int) int) int {
	result := initial
	for _, num := range nums {
		result = operation(result, num)
	}
	return result
}

// 4. Compose - returns composition f(g(x))
func Compose(f func(int) int, g func(int) int) func(int) int {
	return func(x int) int {
		return f(g(x))
	}
}

/*----- Part 4: Process Explorer -----*/

// ExploreProcess demonstrates process information and memory addresses
func ExploreProcess() {
	// Get current process
	pid := os.Getpid()
	ppid := os.Getppid()

	// Create a slice of integers
	data := []int{1, 2, 3, 4, 5}

	// Get memory addresses
	// &data gets the address of the slice header
	// &data[0] gets the address of the first element in the underlying array
	sliceAddr := &data
	firstElemAddr := &data[0]

	// Print process information
	println("=== Process Information ===")
	println("Current Process ID:", pid)
	println("Parent Process ID:", ppid)
	println("Memory address of slice:", sliceAddr)
	println("Memory address of first element:", firstElemAddr)
	println("Note: Other processes cannot access these memory addresses due to process isolation")

	/*
		Include comments explaining:

		1. What a process ID is
			A Process ID is basically a unique number assigned by the operating system to identify a running process.

		2. Why process isolation is important
			Process isolation is important because it prevents one process from accessing or interfering with the memory of another process.

		3. The difference between the slice header address and element addresses
			The slice header address is where Go stores metadata about the slice (for example: length, capacity, pointer to underlying array), while the element address points to the actual data in memory.
	*/
}

func main() {
	// Call the Part 4 function
	ExploreProcess()
}
