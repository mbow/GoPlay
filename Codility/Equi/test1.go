package main

import (
	"fmt"
)

func main() {
	A := []int{-1, 3, -4, 5, 1, -6, 2, 1}

	fmt.Printf("%d \n", Solution(A))
	fmt.Printf("%d \n", SolutionB(A))

}
//84%
func Solution(A []int) int {

	for i, _ := range A {
		rSum := 0
		lSum := 0
		fmt.Println("i =", i)
		fmt.Println(A)

		for k, v := range A {
			if k < i {
				lSum += v
			}
		}
		fmt.Println(lSum)

		for m := i + 1; m < len(A); m++ {
			rSum += A[m]
		}
		fmt.Println(rSum)

		if rSum == lSum {
			fmt.Println("hit", i)
			//return i+1
		}
		fmt.Println("")
	}

	return -1
}
//port a 100% from blog
func SolutionB(A []int) int {

	if len(A) == 0 {
		return -1
	}

	total := 0

	for i, _ := range A {
		total += A[i]
	}
	leftSum := 0
	for i, _ := range A {
		rightSum := total - leftSum - A[i]
		if leftSum == rightSum {
			//fmt.Println("hit", i)
			return i
		}
		leftSum += A[i]
	}

	return -1
}
