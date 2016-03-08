package main

import (
	"fmt"
	"sort"
)

func main() {

	//A := []int{2, 3, 1, 5}
	//A := []int{1, 2, 3, 4, 6}
	//A := []int{1, 2, 3, 4, 5}
	A := []int{3, 4, 5, 7, 8}
	//A := []int{0, 100000}
	//A := []int{}
	//A := []int{2}
	//A := []int{2, 3, 4, 5}
	//A := []int{1, 2}
	//A := []int{3, 1}
	//A := []int{100}

	fmt.Println(SolutionC(A))
	fmt.Println(SolutionD(A))

}

// 60% then 100% after adding check for missing int on start of slice
func SolutionC(A []int) int {

	//check for empty slice
	if len(A) == 0 {
		return 1
	}

	//check if a slice is already in sorted order else sort it
	if !sort.IntsAreSorted(A) {
		sort.Ints(A)
	}

	//step though sorted array, check for N+1
	for i, v := range A {
		if i < len(A)-1 {
			if A[i+1]-v > 1 {
				//found number before missing number so add 1
				return v + 1
			}
		}
	}

	// Check missing numbers not on end or start of slice
	if A[0] != 1 {
		return 1
	} else {
		return A[len(A)-1] + 1
	}
	//something is wrong
	return -1
}

// Internet way https://en.wikipedia.org/wiki/1_%2B_2_%2B_3_%2B_4_%2B_%E2%8B%AF
// The infinite series whose terms are the natural numbers 1 + 2 + 3 + 4 + · · · is a divergent series.
// The nth partial sum of the series is the triangular number
func SolutionD(A []int) int {

	sum := 0
	//check for empty slice
	if len(A) == 0 {
		return 1
	}

	//check if a slice is already in sorted order else sort it
	if !sort.IntsAreSorted(A) {
		sort.Ints(A)
	}

	// find are slice sum
	for _, v := range A {
		sum += v
	}
	// sum of series n(n+1)/2
	sos := A[len(A)-1] * (A[len(A)-1] + 1) / 2

	// no number is missing missing number is next number in series
	if sos-sum == 0 {
		return A[len(A)-1] + 1
	}

	return sos - sum
}
