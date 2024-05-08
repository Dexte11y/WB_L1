// К каким негативным последствиям может привести данный фрагмент кода, и как это исправить?
// Приведите корректный пример реализации.

package main

import (
	"fmt"
	"strings"
)

// var justString string
// func someFunc() {
//   v := createHugeString(1 << 10)
//   justString = v[:100]
// }

// func main() {
//   someFunc()
// }

func createHugeString(size int) string {
	return strings.Repeat("a", size)
}

func someFunc() {
	var justString string                //постараться избежать глобальной переменной
	v := createHugeString(1 << 10)[:100] //создать сразу подстроку нужного размера
	justString = v
	fmt.Println(justString, len(justString))
}

func main() {
	someFunc()
}
