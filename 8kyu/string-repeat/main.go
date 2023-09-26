package main

import "fmt"

func main() {
	fmt.Println(RepeatStr(3, "Hi"))
}

func RepeatStr(repetitions int, value string) string {
	res := ""

	for repetitions > 0 {
		res += value
		repetitions -= 1
	}

	return res
}
