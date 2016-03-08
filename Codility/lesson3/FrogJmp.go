package main

import "fmt"

func main() {

	//test data
	//X := 10 // 1
	//Y := 85 // 1
	//D := 30 // 1

	//X := 20
	//Y := 20
	//D := 1

	X := 7592
	Y := 872727123
	D := 29

	fmt.Println(Solution(X, Y, D))

}

// 100% on first go.
func Solution(x int, y int, d int) int {

	// If start position is same as end no jumps needed
	if x >= y {
		return 0
	}
	// example expects no negative values
	if x <= 0 || y <= 0 || d <= 0 {
		return 0
	}

	// return if number greater than example spec of 1,000,000,000
	if x > 1e9 || y > 1e9 || d > 1e9 {
		return 0
	}

	// find true length
	y = y - x
	// find jumps needed
	j := y / d
	// if remainder one more jump needed
	if y%d != 0 {
		return j + 1
	} else {
		return j
	}
}
