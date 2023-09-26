package main

import "fmt"

func main() {
	fmt.Println(DigitalRoot(493193))
}

func DigitalRoot(n int) int {
	if n < 10 {
		return n
	}

	sum := 0

	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return DigitalRoot(sum)
}
