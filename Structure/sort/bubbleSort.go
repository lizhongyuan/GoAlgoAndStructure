package sort

import "fmt"

type Item struct {
	dataType uintptr
	data uint64
}

func BubbleSort(arr []Item) {
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d", arr[i].data)
	}
}
