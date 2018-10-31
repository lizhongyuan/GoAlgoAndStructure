package main


func Filter(s []int, fn func(int) bool) []int {

	var curSlice []int

	for _, v := range s {
		if fn(v) {
			curSlice = append(curSlice, v)
		}
	}

	return curSlice
}


func main() {

}
