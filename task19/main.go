// Разработать программу, которая переворачивает подаваемую
// на ход строку (например: «главрыба — абырвалг»). Символы могут быть unicode.

package main

import (
	"fmt"
)

func main() {
	str := "главрыба"
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		fmt.Println("i: ", runes[i], "j: ", runes[j])
		runes[i], runes[j] = runes[j], runes[i]
	}
	fmt.Println(string(runes))
}
