// Разработать программу, которая переворачивает слова в строке.
// Пример: «snow dog sun — sun dog snow».

package main

import (
	"fmt"
	"strings"
)

func reverseWords(sentence string) string {
	words := strings.Fields(sentence)
	reversedWords := make([]string, len(words))

	for i, word := range words {
		reversedWords[len(words)-1-i] = word
	}

	return strings.Join(reversedWords, " ")
}

func main() {
	sentence := "snow dog sun"
	reversedSentence := reverseWords(sentence)
	fmt.Println(reversedSentence)
}
