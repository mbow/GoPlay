package main

import (
	"fmt"
	"math"
)

func main() {
	//test data
	A := []int{3, 1, 2, 4, 3}
	//A := []int{10, 20, 1, 10, 8}
	//A := []int{-10, 20, -1, 10, -8}
	//A := []int{-1000, 1000}
	fmt.Println(SolutionB(A))
}

// My 100% correct answer
func SolutionB(A []int) int {

	// test says A non-empty zero-indexed array A
	// but test anyway...
	if len(A) <= 0 {
		return 0
	}
	// area to store sum of slice/array
	sum := 0
	// max [] integer range is −1,000..1,000 then +1
	p := 2001
	//Store left side of tape
	lSum := 0
	// calculate sum of slice
	for _, v := range A {
		sum += v
	}
	// step though comparing left side to right side
	for i, v := range A {
		// gotcha don't measure against last item
		if i < len(A)-1 {
			diff := 0
			// add to left side of tape
			lSum += v
			// calculate new right side total of tape
			rSum := sum - lSum
			// work out difference
			diff = lSum - rSum
			// as data can have minus number or x--y convert to whole
			//find absolute(Abs) used import but could use (((n)>=0)?(n):(-(n)))
			diff = int(math.Abs(float64(diff)))
			//fmt.Println("lsum:", lSum, "-", "rSum:", rSum, "=", diff)
			//set the smallest difference is needed
			if p > diff {
				p = diff
			}
		}
	}
	return int(p)
}
