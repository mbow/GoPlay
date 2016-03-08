package main

import "fmt"

func main() {
	//N := 5
	//A := []int{3, 4, 4, 6, 1, 4, 4}

	//N := 1
	//A := []int{3, 4, 4, 6, 1, 4, 4, 100000}

	//N := 100000
	//A := []int{3, 4, 4, 6, 1, 4, 4, 100000}

	//N := 1
	//A := []int{2, 1, 1, 2, 1}

	//N := 5
	//A := []int{5, 5, 5, 5, 5, 1, 6}
	fmt.Println(SolutionB(N, A))
}

func SolutionB(N int, A []int) []int {

	//create slice to store the counters
	result := make([]int, N)

	// we hold max counter value and it value on last op
	max := 0
	maxAtLastCnt := 0

	for _, v := range A {
		op := v
		//if the operations is max counter.
		if op == N+1 {
			maxAtLastCnt = max
		} else {
			//op is between 1 to N, but the index for the slice
			//is between 0 and (N-1). Decrease op to adjust it.
			op--
			if result[op] < maxAtLastCnt {
				result[op] = maxAtLastCnt + 1
			} else {
				result[op]++
			}
			//update the max value if necessary.
			if max < result[op] {
				max = result[op]
			}
		}
	}

	//apply the 'max counter' operation
	//to the slot(s) where it should be applied.
	for j := 0; j < N; j++ {
		if result[j] < maxAtLastCnt {
			result[j] = maxAtLastCnt
		}
	}
	return result
}
