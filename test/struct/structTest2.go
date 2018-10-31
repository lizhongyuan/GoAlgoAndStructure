package main


import "fmt"

type Item struct {
	dataType	string
	intData		uint64
}


func main() {

	item1 := &Item{"int", 1}
	item2 := &Item{"int", 2}

	fmt.Printf("Item.dataType:%s, Item.intData:%d\n", item1.dataType, item1.intData)
	fmt.Printf("Item.dataType:%s, Item.intData:%d\n", item2.dataType, item2.intData)

	// arr := make([]Item, 10)
	arr := []Item{}

	arr = append(arr, *item1)
	arr = append(arr, *item2)

	for i := 0; i < len(arr); i++ {
		fmt.Printf("Item.dataType:%s, Item.intData:%d\n", arr[i].dataType, arr[i].intData)
	}
}

