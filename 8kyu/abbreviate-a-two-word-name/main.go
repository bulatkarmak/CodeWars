// Write a function to convert a name into initials. This kata strictly takes two words with one space in between them.
// The output should be two capital letters with a dot separating them.
// It should look like this:
// Sam Harris => S.H
// patrick feeney => P.F

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(AbbrevName("hello world"))
}

func AbbrevName(name string) string {
	res := ""

	splittedString := strings.Split(name, " ")

	for _, v := range splittedString {
		res += strings.ToUpper(string(v[0]))
		res += "."
	}

	return res[:len(res)-1]

}

// -------- REFACTORED VERSION --------

// func AbbrevName(name string) string {
// 	splittedString := strings.Fields(name) // использование Fields вместо Split для разделения строки по пробелам

// 	var abbrevs []string
// 	for _, word := range splittedString {
// 		abbrevs = append(abbrevs, strings.ToUpper(string(word[0])))
// 	}

// 	return strings.Join(abbrevs, ".")
// }
