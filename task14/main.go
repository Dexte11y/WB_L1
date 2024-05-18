// Разработать программу, которая в рантайме способна
// определить тип переменной: int, string, bool, channel из переменной типа interface{}.

package main

import (
	"fmt"
	"reflect"
)

func interfaceArg(i interface{}) {
	fmt.Println("type:", reflect.TypeOf(i))
}

func main() {
	// input := 3
	// input := "qwerty"
	// input := true
	input := make(chan int)
	interfaceArg(input)
}
