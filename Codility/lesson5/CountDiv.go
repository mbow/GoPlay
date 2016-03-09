package main

import "fmt"

//A and B are integers within the range [0..2,000,000,000];
//K is an integer within the range [1..2,000,000,000];
//A ≤ B.

func main() {
	//A := 6
	//B := 11
	//K := 2

	//A := 0
	//B := 0
	//K := 11

	//A := 0
	//B := 1
	//K := 11

	// SLOW
	A := 0
	B := 2000000000
	K := 9

	//A := 100
	//B := 1000
	//K := 10

	fmt.Println(Solution(A, B, K))
	fmt.Println(SolutionB(A, B, K))
}

// correct 100% performance 0% ... needs to be o(1)
func Solution(A int, B int, K int) int {

	var xs []int
	for ; A <= B; A++ {
		if A%K == 0 {
			xs = append(xs, A)
		}
	}

	return len(xs)
}

// 100% on both
func SolutionB(A int, B int, K int) int {

	if A%K == 0 {
		return (B-A)/K + 1
	} else {
		return (B - (A - A%K)) / K
	}
}
