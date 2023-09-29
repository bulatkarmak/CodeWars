// Your task is to write function factorial.

// https://en.wikipedia.org/wiki/Factorial

package main

import "fmt"

func main() {
	fmt.Println(factorial(5))
}

func factorial(num int) int {
	if num == 0 || num == 1 {
		return 1
	}

	return num * factorial(num-1)
}

// -------- REFACTORED VERSION --------

// func factorial(num int) int {
// 	result := 1

// 	for i := 1; i <= num; i++ {
// 		result *= i
// 	}

// 	return result
// }
