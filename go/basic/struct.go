package main

import "fmt"

type StudetnInfo struct {
	id    int
	class string
	age   int
}
type Student struct {
	id   int
	name string
	info StudetnInfo
}

func main() {
	st1 := Student{1, "duydo", StudetnInfo{1, "lop 1", 10}}
	fmt.Println(st1)

	var anonymous = struct {
		email string
		age   int
	}{
		"duydd1@gmail.com", 999,
	}
	fmt.Println(anonymous)
	// con trỏ struct
	pointer := &Student{999, "duy", StudetnInfo{1, "lop 1", 10}}
	fmt.Println(&pointer)
	fmt.Println(pointer.id)
	fmt.Println(pointer.name)
	//
	type NoName struct {
		string
		int
	}
	n := NoName{"duy", 1}
	fmt.Println(n)

}
