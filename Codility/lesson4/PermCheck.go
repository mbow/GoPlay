package main

import "fmt"
import "sort"

func main() {

	//A := []int{4, 1, 3, 2}
	//A := []int{4, 1, 3}
	A := []int{9, 5, 7, 3, 2, 7, 3, 1, 10, 8} // antiSum
	//A := []int{1, 3, 6, 4, 1, 2}
	//A := []int{2, 3, 1, 5}
	//A := []int{1, 2, 3, 4, 6}
	//A := []int{1, 2, 3, 4, 5}
	//A := []int{3, 4, 5, 7, 8}
	//A := []int{0, 100000}
	//A := []int{}
	//A := []int{2}
	//A := []int{1}
	//A := []int{-2, -3, -1, 1, 0, 2, 4}
	//A := []int{-2, -3, -1, 5, 0, 2, 4}
	//A := []int{4, 5, 6, 2, 7}
	//A := []int{1, 2}
	//A := []int{3, 1}
	//A := []int{-100}
	//A := []int{-2147483648, 2147483647}

	fmt.Println(SolutionD(A))
	fmt.Println(SolutionE(A))

}

//80% first go failed both antiSum 1&2 [9, 5, 7, 3, 2, 7, 3, 1, 10, 8]
// for 100% see solutionE
func SolutionD(A []int) int {

	sum := 0
	//check for empty slice
	if len(A) == 0 {
		return 0
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
		return 1
	}
	return 0
}

// for 100% see solutionE
func SolutionE(A []int) int {

	counter := make([]bool, len(A))

	for _, v := range A {

		if v <= 0 || v > len(A) {
			return 0
		} else {
			if counter[v-1] != false {
				return 0
			} else {
				counter[v-1] = true
			}
		}
	}
	return 1

}
