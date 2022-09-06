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

	testNumber := 100_000

	start := time.Now()

	numPrimes := countPrimes(testNumber)

	elapsed := time.Since(start)

	fmt.Println("Completed in", elapsed)
	fmt.Println("There are", numPrimes, "primes up to", testNumber)
}
