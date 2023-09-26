package main

import "fmt"

func main() {
	fmt.Println(NumberToString(124))
}

func NumberToString(n int) string {
	res := ""
	len := 0
	tmp := 0

	if n == 0 {
		return "0"
	}

	if n < 0 {
		res = "-"
		n = -n
	}

	for temp_n := n; temp_n > 0; temp_n /= 10 {
		len += 1
	}

	for len > 0 {
		tmp = pow(10, len-1)
		res += fmt.Sprint(n / tmp)
		n %= tmp
		len -= 1
	}

	return res
}

func pow(num, pow int) int {
	t_num := num

	if pow == 0 {
		return 1
	}

	pow -= 1
	for pow > 0 {
		num *= t_num
		pow -= 1
	}

	return num
}
