package main

import "fmt"
import "time"

func isPrime(number int) bool {

	element := 1
	count := 0

	for element <= number {

		if number%element == 0 {
			count += 1
		}

		if count > 2 {
			return false
		}

		element += 1
	}

	return true
}

func countPrimes(number int) int {

	count := 0

	for index := 2; index < number; index++ {
		if isPrime(index) {
			count += 1
		}
	}
	return count
}

func main() {

	testNumber := 100_000
	start := time.Now()

	// for index := 1; index <= testNumber; index++ {
	// 	fmt.Println(index, isPrime(index))
	// }

	fmt.Println(countPrimes(testNumber))

	elapsed := time.Since(start)
	fmt.Println(elapsed)
}
