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

// -------- REFACTORED VERSION --------

// func ToCamelCase(s string) string {
// 	// Разделяем строку по символам '_' или '-'
// 	re := regexp.MustCompile(`[_-]+`)
// 	parts := re.Split(s, -1)

// 	// Если строка пуста или состоит только из разделителей, возвращаем пустую строку
// 	if len(parts) == 0 {
// 		return ""
// 	}

// 	// Оставляем первое слово без изменений
// 	res := parts[0]

// 	// Преобразуем оставшиеся части в CamelCase
// 	for _, part := range parts[1:] {
// 		res += strings.Title(part) // `Title` делает первую букву слова заглавной
// 	}

// 	return res
// }
