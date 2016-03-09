package main

import "fmt"

func main() {
	A := []int{0, 1, 0, 1, 1}
	// think single 0 or 1
	// think double 0 0 , 0 1, 1 0 ,11
	fmt.Println(SolutionC(A))
}

func SolutionC(A []int) int {

	pairs := 0
	//count the numbers of zero discovered while traversing 'A'
	//for each successive '1' in the list, number of pairs will
	//be incremented by the number of zeros discovered before that '1'
	zero_count := 0
	// traverse through the list 'A'
	for _, v := range A {
		if v == 0 {
			//counting the number of zeros discovered
			zero_count += 1
		} else if v == 1 {
			//if '1' is discovered, then number of pairs is incremented
			//by the number of '0's discovered before that '1'
			pairs += zero_count

			//if pairs is greater than 1 billion, return -1
			if pairs > 1000000000 {
				return -1
			}
		}
		//return number of pairs
	}
	return pairs

}
