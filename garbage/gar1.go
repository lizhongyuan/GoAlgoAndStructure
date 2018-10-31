package main

import "fmt"

func main() {
	fmt.Printf("garbage1\n")

	var a [4]int
	a[0] = 1
	i := a[0]

	fmt.Println(i)

	// aSlice := a[:]

	// --------------

	// letters := []string{"a", "b", "c"}

	var s []byte
	s = make([]byte, 5, 5)

	fmt.Println(len(s))
	fmt.Println(cap(s))

	// b := []byte{'g', 'o', 'l', 'a', 'n', 'g'}

	c := []uint64{1, 2, 3, 4, 5}

	c[1] = 0

	fmt.Println(c[1:4])
	fmt.Println(c)

	t := make([]uint64, len(c), (cap(c)+1)*2)
	copy(t, c)

	fmt.Println(t)

	d := []uint64{1, 2, 3, 4}

	dAppd := appendByte(d, 5, 6)

	fmt.Println(dAppd)

	e := make([]int, 1)
	e = append(e, 2, 3, 4)

	fmt.Println(e)
}


func appendByte(slice []uint64, data ...uint64) []uint64 {

	m := len(slice)
	n := m + len(data)

	if n > cap(slice) {
		newSlice := make([]uint64, (n + 1) * 2)
		copy(newSlice, slice)
		slice = newSlice
	}

	slice = slice[0:n]
	copy(slice[m:n], data)

	return slice
}


