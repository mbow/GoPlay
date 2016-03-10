package main

import "fmt"

//import "sort"

func main() {

	//A := []int{1, 5, 3, 3, 7}
	//A := []int{1, 2, 4, 3, 7}
	//A := []int{1, 3, 5}
	//A := []int{1}
	//A := []int{1, 0}
	//A := []int{0, 1}
	A := []int{1, 5, 3, 4, 7}

	fmt.Println(Solution3(A))
}

func Solution3(A []int) bool {

	fmt.Println("Input", A)

	swapCount := 0

	//Slice with one values is already sorted
	if len(A) == 1 {
		fmt.Println("Size of one is in order")
		fmt.Println(A)
		return true
	}
	// Slice has 2 elements check is need to be swapped
	if len(A) == 2 {
		if A[0] > A[1] {
			fmt.Println("Swap pair")
			A[0], A[1] = A[1], A[0]
		}
		fmt.Println(A)
		return true
	}

	//check not already in order
	ordered := true
	for i, _ := range A {
		if i < len(A)-1 {
			//fmt.Println(i)
			//fmt.Println(A[i], "<", A[i+1])
			if !(A[i] < A[i+1]) {
				fmt.Println("Set order bool")
				ordered = false
				break //exit traverse of slice
			}
		}
	}
	if ordered {
		fmt.Println("Already Ordered")
		return true
	}

	// Travers the given array from right to left
	for i := len(A) - 1; i >= 0; i-- {

		//fmt.Println(A[i])
		//Don't compare first element in slice not a pair
		if i != 0 {
			// find numbers not in order
			if A[i] < A[i-1] {
				fmt.Println("Found")
				// Find the other element to be swapped & if the number is duplicate find correct place
				j := i - 1
				for j >= 0 && A[i] < A[j] {
					// Check duplicates
					if i+1 != len(A) {
						//to shift tracks duplicate in a row
						toShift := 1
						fmt.Println(A[i], A[i+1])
						if A[i] == A[i+1] {
							toShift++
						}
						fmt.Println("toShift=", toShift)
						// Swap the pair
						for toShift != 0 {
							toShift--
							//fmt.Println(i, j)
							//fmt.Println(A[i], ":", A[j])
							//fmt.Println(A[i], ":", A[j+toShift])
							A[i+toShift], A[j] = A[j], A[i+toShift] //a[i], a[j] = a[j], a[i]
							//fmt.Println(A)
						}
						j--
						fmt.Println("CAN SWAP", A)

					}
				}
				//update swap count to check valid entry's only have one swap
				swapCount++
				if swapCount > 1 {
					fmt.Println("Too many swaps needed")
					return false
				}

			}

		}

	}

	fmt.Println("default exit true")
	fmt.Println(A)
	return true
}
