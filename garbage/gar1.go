package main

import "fmt"

func main() {
	fmt.Printf("garbage1\n")

	var a [4]int
	a[0] = 1
	i := a[0]

	fmt.Println(i)
}
