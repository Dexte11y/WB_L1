// Имеется последовательность строк -
// (cat, cat, dog, cat, tree) создать для нее собственное множество.

package main

import "fmt"

func main() {
	data := []string{"cat", "cat", "dog", "cat", "tree"}
	setData := make(map[string]bool)

	for _, v := range data {
		setData[v] = true
	}
	for k := range setData {
		fmt.Printf("%s ", k)
	}
}
