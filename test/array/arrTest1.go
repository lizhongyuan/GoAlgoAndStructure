package main

import "fmt"

func getArrayAvr(array []float64) float64 {

	sum := 0.0
	var avg float64

	switch len(array) {
	case 0:
		avg = 0
	default:
		for _, v := range array {
			sum += v
		}

		avg = sum / float64(len(array))
	}

	// fmt.Println(avg)
	return avg
}

func main() {

	// slice
	array := []float64{1,2,3,4,5,6}

	// var avg = getArrayAvr(array)
	avg := getArrayAvr(array)

	fmt.Printf("avg: %.2f", avg)
}


