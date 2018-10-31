package Structure

import (
	"./list"
	"./sort"
)

type Item struct {
	dataType uintptr
	data uint64
}

func ShowList() {
	list.ShowList()
}

func ShowBubbleSort(arr []Item) {
	sort.BubbleSort(arr)
}
