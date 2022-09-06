package main

import (
	"fmt"
)

const PI = 3.14

func main() {
	var student2 = "chaocaunhe"

	// var a, b, c, d int = 1, 2, 3, 4
	fmt.Println(student2)
	//fmt.Print(a, b, c, d)
	var a, b = 6, "Hello"
	c, d := 7, "World!"

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	fmt.Println(PI)
	var arr44 = [3]int{1, 2, 3}
	fmt.Println("mang arr1", arr44)
	var arr12 = [...]int{1, 2, 3}
	fmt.Println("mang arr12", arr12)
	fmt.Println("\nslice.............")
	myslice2 := []string{"1", "2", "3", "4"}
	fmt.Println(len(myslice2))
	fmt.Println(cap(myslice2))
	fmt.Println(myslice2)
	fmt.Println("\ntạo slice từ capacity")

	arr1 := [6]int{10, 11, 12, 13, 14, 15}
	fmt.Println(arr1)
	myslice := arr1[2:4]
	fmt.Printf("myslice = %v\n", myslice)
	fmt.Printf("length = %d\n", len(myslice))
	fmt.Printf("capacity = %d\n", cap(myslice))

	fmt.Println("\ntạo slice từ make function ")

	myslice1 := make([]int, 5, 10)
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))

	myslice3 := []int{1, 2, 3}
	fmt.Printf("myslice1 = %v\n", myslice3)
	fmt.Printf("length = %d\n", len(myslice3))
	fmt.Printf("capacity = %d\n", cap(myslice3))

	myslice3 = append(myslice3, 20, 21)
	fmt.Printf("myslice1 = %v\n", myslice3)
	fmt.Printf("length = %d\n", len(myslice3))
	fmt.Printf("capacity = %d\n", cap(myslice3))

}
