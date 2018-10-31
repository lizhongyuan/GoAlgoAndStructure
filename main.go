package main

import "fmt"

/*
type Item struct {
	dataType uintptr
	data uint64
}

item1 := &Item{&"uint64", 1}
*/

type test1 struct {
	a uint64
	b uint64
}

var testA *test1 = &test1{1, 2}

func main() {

	fmt.Println("main func")

	Structure.ShowList()
	Structure.ShowBubbleSort(&arr)
}
