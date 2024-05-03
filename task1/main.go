package main

import "fmt"

type Human struct {
	Name string
	Age  int
}

func (h *Human) Speak() {
	fmt.Printf("My name is %s and I am %d years old.\n", h.Name, h.Age)
}

type Action struct {
	Human // Встраиваем структуру Human
	Hobby string
}

func main() {
	a := Action{
		Human: Human{Name: "Alice", Age: 30},
		Hobby: "Reading",
	}

	a.Speak() // Выведет: My name is Alice and I am 30 years old.

	fmt.Printf("%s's hobby is %s.\n", a.Name, a.Hobby) // Выведет: Alice's hobby is Reading.
}
