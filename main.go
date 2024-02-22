package main

import "fmt"

func SortSlice(slice []int) {
	n := len(slice)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if slice[j] > slice[j+1] {
				// swap
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
}

func IncrementOdd(slice []int) {
	for i := 1; i < len(slice); i += 2 {
		slice[i]++
	}
}

func PrintSlice(slice []int) {
	fmt.Println(slice)
}

func ReverseSlice(slice []int) {
	n := len(slice)
	for i := 0; i < n/2; i++ {
		slice[i], slice[n-i-1] = slice[n-i-1], slice[i]
	}
}

func appendFunc(dst func([]int), src ...func([]int)) func([]int) {
	return func(slice []int) {
		dst(slice)
		for _, f := range src {
			f(slice)
		}
	}
}

func main() {

	slice := []int{4, 2, 7, 1, 9, 5}

	// Sorting the slice
	SortSlice(slice)
	PrintSlice(slice)

	// Incrementing odd positions
	IncrementOdd(slice)
	PrintSlice(slice)

	ReverseSlice(slice)
	PrintSlice(slice)

	combinedFunc := appendFunc(SortSlice, IncrementOdd, ReverseSlice)
	combinedFunc(slice)
	PrintSlice(slice)
}
