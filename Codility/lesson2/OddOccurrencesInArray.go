package main

import "fmt"

func main() {
	//A := []int{9, 3, 9, 3, 9, 7, 9}
	A := []int{9, 9, 9, 9, 3, 3, 3, 3, 6, 6, 6}
	fmt.Println("->", SolutionB(A))
	fmt.Println("->", SolutionC(A))
}
// My way 100% correct but 25% on performance
// Solution is to run two nested loops. The outer loop picks all elements one by one and inner loop counts number
// of occurrences of the element picked by outer loop
func SolutionB(A []int) int {

	for i, v := range A {
		n := 0
		for i, _ := range A {
			if v == A[i] {
				n++
			}
		}
		if n%2 != 0 {
			return A[i]
		}
	}
	return -1
}
// Best way
// Correct way that gets 100% http://www.geeksforgeeks.org/find-the-number-occurring-odd-number-of-times/
// The Best Solution is to do bitwise XOR of all the elements. XOR of all elements gives us odd occurring element.
// Please note that XOR of two elements is 0 if both elements are same and XOR of a number x with 0 is x
func SolutionC(A []int) int {
	res := 0
	for _, v := range A {
		res = res ^ v
	}
	return res
}
