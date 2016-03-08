package main

import (
	"fmt"
)

const (
	MaxInt64  = 1<<63 - 1
	MinInt64  = -1 << 63
	MaxUint64 = 1<<64 - 1
)

func main() {

	A := []int{1, 3, 6, 4, 1, 2}
	//A := []int{2, 3, 1, 5}
	//A := []int{1, 2, 3, 4, 6}
	//A := []int{1, 2, 3, 4, 5}
	//A := []int{3, 4, 5, 7, 8}
	//A := []int{0, 100000}
	//A := []int{}
	//A := []int{2}
	//A := []int{-2, -3, -1, 1, 0, 2, 4}
	//A := []int{-2, -3, -1, 5, 0, 2, 4}
	//A := []int{4, 5, 6, 2, 7}
	//A := []int{1, 2}
	//A := []int{3, 1}
	//A := []int{-100}
	//A := []int{-2147483648, 2147483647}

	fmt.Println(SolutionC(A))

}

// 77% then 100% after adding check for case when input array is not missing any element
func SolutionC(A []int) int {

	//check for empty slice
	if len(A) == 0 {
		return 1
	}
	// cbn Current biggest number hold minimal positive integer that's been found
	// we need track this as bitmap size includes minus integers so we cant use A[i]+1
	cbn := 1
	//slice of bool's length of biggest size of number given - aka look up table
	bitmap := make([]bool, len(A)+1)
	bitmap[0] = true
	//populate bitmap and also find highest positive int in input list.
	for _, v := range A {
		if v > 0 && v <= len(A) {
			//mark valid entry in bitmap table
			bitmap[v] = true
		}
		if v > cbn {
			//update Current biggest number
			cbn = v
		}
	}
	//find the first positive number in bitmap that is false and return it
	for i := 1; i < len(bitmap); i++ {
		if !bitmap[i] {
			return i
		}
	}
	//this is to handle the case when input array is not missing any element.
	return (cbn + 1)
}
