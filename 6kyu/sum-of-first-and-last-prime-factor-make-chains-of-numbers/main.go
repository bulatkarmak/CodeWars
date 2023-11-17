// https://www.codewars.com/kata/5629b94e34e04f8fb200008e/train/go

package main

import "fmt"

func main() {
	fmt.Println(SflpfData(83, 163))
}

func SflpfData(k, nMax int) []int {
	res := []int{}
	for i := k; i <= nMax; i++ {
		if Sflpf(i) == k {
			res = append(res, i)
		}
	}

	if len(res) == 0 {
		return nil
	}

	// Код с nil добавил, т.к. зачем-то CodeWars попросил такой вывод, если функция не выдает никаких чисел

	return res
}

func Sflpf(num int) int {
	primeFactors := FindPrimeFactors(num)
	return primeFactors[0] + primeFactors[len(primeFactors)-1]
}

func FindPrimeFactors(num int) []int {
	res := make([]int, 0)

	for factor := 2; num > 1; factor++ {
		for num%factor == 0 {
			res = append(res, factor)
			num = num / factor
		}
	}
	return res
}
