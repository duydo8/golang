package main

import (
	"fmt"
)

type Animal interface {
	speak()
}
type Dog struct {
}

func (d Dog) speak() {
	fmt.Println("gau gau")
}
func main() {
	var animal Animal
	animal = Dog{}
	animal.speak()
}
