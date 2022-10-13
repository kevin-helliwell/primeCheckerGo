package main

import (
	"fmt"
	"time"
)

func isEven(number int) bool {
	return number%2 == 0
}

func isOdd(number int) bool {
	return number%2 != 0

}

func isPrime(number int) bool {

	element := 3
	count := 0

	if isEven(number) {
		return false
	}

	if isOdd(number) {
		for element <= number {
			if number%element == 0 {
				count += 1
			}
			if count > 1 {
				return false
			}
			element += 2
		}
	}
	return true

}

func countPrimes(number int) int {

	count := 0

	for index := 1; index < number; index++ {
		if isPrime(index) {
			count += 1
		}
	}
	return count
}

func main() {

	testNumbers := []int{10_000, 25_000, 50_000, 75_000, 100_000, 125_000, 150_000, 175_000, 200_000}

	for index := 0; index < len(testNumbers); index++ {

		testNumber := testNumbers[index]
		start := time.Now()

		numPrimes := countPrimes(testNumber)

		elapsed := time.Since(start)

		fmt.Println("Completed in", elapsed)
		fmt.Println("There are", numPrimes, "primes up to", testNumber)
	}
}
