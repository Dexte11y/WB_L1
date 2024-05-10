// Разработать программу, которая проверяет, что все символы в строке уникальные
// (true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.

// Например:
// abcd — true
// abCdefAaf — false
// aabcd — false

package main

import (
	"fmt"
	"strings"
)

func unicString(str string) bool {
	str = strings.ToLower(str)
	seen := make(map[rune]bool)
	for _, char := range str {
		if seen[char] {
			return false
		}
		seen[char] = true
	}

	return true
}

func main() {
	words := []string{"abcd", "abCdefAaf", "aabcd"}
	for _, word := range words {
		flag := unicString(word)
		if flag {
			fmt.Println("true")
		} else {
			fmt.Println("false")
		}

	}
}
