package main

import (
	"fmt"
)

func main() {
	//test data
	A := []int{1, 3, 1, 4, 2, 3, 5, 4}
	//A := []int{1, 3, 1, 4, 2, 3, 5, 4}
	//A := []int{2, 3, 1, 5}
	//A := []int{1, 2, 3, 4, 6}
	//A := []int{1, 2, 3, 4, 5}
	//A := []int{3, 4, 5, 7, 8}
	//A := []int{0, 100000}
	//A := []int{}
	//A := []int{2}
	//A := []int{2, 3, 4, 5}
	//A := []int{1, 2}
	//A := []int{3, 1}
	//A := []int{100}

	X := 5
	//X := 0
	//X := 1
	//X := 100000

	fmt.Println(Solution(X, A))
}

// My 100% correct answer
func Solution(X int, A []int) int {

	// test says A non-empty zero-indexed array & positive X
	// but test anyway...
	if len(A) <= 0 || X < 1 {
		return 0
	}
	// Make slice of booleans (number frog hops + 1 = steps), we will mark off as true when we match a leaf to a step
	// we are already on step 0 so mark as true
	leafPath := make([]bool, X+1)
	leafPath[0] = true
	// Set steps remaining till frog reaches river bank
	todo := X

	//step over slice
	for i, v := range A {
		//Where did the leaf fall
		aLeaf := v
		//leaf fall within river / needed steps
		if aLeaf < len(leafPath) {
			//has a leaf already dropped here
			if !leafPath[aLeaf] {
				//one less step need
				todo--
				//set the leaf in the path as visible
				leafPath[aLeaf] = true
			}
		}
		//reached river bank
		if todo == 0 {
			return i
		}
	}
	//frog cant jump, not enough leafs
	return -1
}
