package main

import "fmt"
import "./Structure"

func main() {

	fmt.Println("main func")

	arr := [1,2,3]

	Structure.ShowList()
	Structure.ShowBubbleSort(&arr)
}
