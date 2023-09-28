// Complete the method/function so that it converts dash/underscore delimited words into camel casing.
// The first word within the output should be capitalized
// only if the original word was capitalized (known as Upper Camel Case, also often referred to as Pascal case).
// The next words should be always capitalized.

// Examples
// "the-stealth-warrior" gets converted to "theStealthWarrior"
// "The_Stealth_Warrior" gets converted to "TheStealthWarrior"
// "The_Stealth-Warrior" gets converted to "TheStealthWarrior"

package main

import (
	"fmt"
	"regexp"
	"unicode/utf8"
)

func main() {
	fmt.Println(ToCamelCase("the-stealth-warrior"))
}

func ToCamelCase(s string) string {
	re := regexp.MustCompile(`[_-]`)
	parts := re.Split(s, -1)
	res := parts[0]

	for i := 1; i < len(parts); i++ {
		if parts[i][0] > 91 {
			res += capitalize(parts[i])
		} else {
			res += parts[i]
		}
	}

	return res
}

func capitalize(text string) string {
	r, size := utf8.DecodeRuneInString(text)

	return string(r-32) + text[size:]
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
