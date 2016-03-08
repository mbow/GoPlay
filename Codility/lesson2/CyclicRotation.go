package main

import "fmt"

func main() {
	A := []int{3, 8, 9, 7, 6}
	//A := []int{}
	K := 3
	fmt.Println(Solution(A, K))
}

func Solution(a []int, k int) []int {
	// store int popped off end of array. auto value -1000/1000
	var x int
	// throw out any values of k < 0 || < 100 from question assume section
	// if empty array return...(gotcha on first run 87% to 100% with that added.)
	if k < 0 || k > 100 || len(a) <= 0 {
		return a
	}
	// if number shift is larger than array length, we dived by length and shift by remainder
	n := k % len(a)
	// pop from end of array to beginning (Example k is only ever positive)
	for i := n; i > 0; i-- {
		x, a = a[len(a)-1], a[:len(a)-1]
		a = append([]int{x}, a...)
	}
	return a
}
