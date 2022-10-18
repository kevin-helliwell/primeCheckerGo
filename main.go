package main

import (
	"fmt"
	"time"
)

// Checks if the number given is divisible by 2
// If the number is divisible by 2, the number cannot be prime

func isEven(number int) bool {
	return number%2 == 0
}

// Checks if the number given is not divisible by 2
// If the number is not divisible by 2, it is odd and either composite (not prime) or prime

func isOdd(number int) bool {
	return number%2 != 0

}

func isPrime(number int) bool {

	// Used to check for divisible factors of our given number
	// Starts at 3 since we're only checking odd numbers
	element := 3

	// Increments when a divisible factor of our given number is found
	count := 0

	// Guard clause (prime numbers cannot be even)
	if isEven(number) {
		return false
	}

	// If our given number is not even, it must be odd
	if isOdd(number) {
		for element <= number {
			if number%element == 0 {
				count += 1 // Increments when a divisible factor of our number is found
			}
			if count > 1 {
				return false // Short-circuits the loop (prime numbers only have one divisible factor other than 1)
			}
			element += 2 // All odd numbers are two positions away from each other
		}
	}
	return true // If the loop concludes, the given number must be prime

}

// Counts primes based on above function up to a given number
// Iterates only through odd numbers since even numbers cannot be prime

func countPrimes(number int) int {

	count := 0

	for index := 1; index <= number; index = index + 2 {
		if isPrime(index) {
			count += 1
		}
	}
	return count
}

func main() {

	// Iterates through an array of increasingly large numbers and times the execution for testing purposes

	testNumbers := []int{10_000, 25_000, 50_000, 75_000, 100_000, 125_000, 150_000, 175_000, 200_000}

	var totalTime time.Duration = 0

	for index := 0; index < len(testNumbers); index++ {

		testNumber := testNumbers[index]
		start := time.Now()

		numPrimes := countPrimes(testNumber)

		elapsed := time.Since(start)
		totalTime += elapsed

		fmt.Println("There are", numPrimes, "primes up to", testNumber, "- Completed in", elapsed)

	}
	fmt.Println("Total time - Completed in", totalTime)
}
