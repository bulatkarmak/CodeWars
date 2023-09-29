// Consider an array/list of sheep where some sheep may be missing from their place.
// We need a function that counts the number of sheep present in the array (true means present).

// For example,

// [true,  true,  true,  false,
//   true,  true,  true,  true ,
//   true,  false, true,  false,
//   true,  false, false, true ,
//   true,  true,  true,  true ,
//   false, false, true,  true]
// The correct answer would be 17.

// Hint: Don't forget to check for bad values like null/undefined

package main

import "fmt"

func main() {
	tmp := []bool{true, true, true, false,
		true, true, true, true,
		true, false, true, false,
		true, false, false, true,
		true, true, true, true,
		false, false, true, true}

	fmt.Println(CountSheeps(tmp))
}

func CountSheeps(numbers []bool) int {
	count := 0

	for _, v := range numbers {
		if v {
			count += 1
		}
	}

	return count // your code here
}
