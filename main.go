package main

import "fmt"
import "time"

func isPrime(number int) bool {

	element := 1
	count := 0

	if number%2 == 0 {
		return false
	}

	if number%2 == 1 {
		for element <= number {
			if number%element == 0 {
				count += 1
			}
			if count > 2 {
				return false
			}
			element += 2
		}
	}
	return true

	// for element <= number {

	// 	if number%element == 0 {
	// 		count += 1
	// 	}

	// 	if count > 2 {
	// 		return false
	// 	}

	// 	element += 1
	// }

	// return true
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

	testNumber := 200_000
	
	start := time.Now()
	
	numPrimes := countPrimes(testNumber)
	
	elapsed := time.Since(start)
	
	fmt.Println("Completed in", elapsed)
	fmt.Println("There are", numPrimes, "primes up to", testNumber)
}
