package main

import (
	"fmt"
	"reflect"
	"sort"
)

func main() {

	//A := []int{1, 5, 3, 3, 7} //true
	//A := []int{1, 5, 3, 3, 3, 3, 3, 3, 3, 3, 7} //true
	//A := []int{1, 5, 3, 3, 3, 7, 3, 3, 3, 3, 7} //false
	A := []int{6, 3, 3, 3, 3, 4, 4, 4, 5, 5, 5, 5, 6, 6, 7} //true
	//A := []int{1, 3, 3, 3, 4, 3, 3, 3, 3, 7} //true
	//A := []int{1, 2, 3, 3, 3, 3, 3, 3, 3, 7, 5} //true
	//A := []int{7, 2, 3, 3, 3, 3, 3, 3, 3, 3, 4, 4, 4, 5} // true
	//A := []int{2, 3, 7, 3, 3, 3, 3, 3, 3, 3, 4, 4, 4, 5} //true
	//A := []int{1, 2, 4, 3, 7} //true
	//A := []int{1, 3, 5} //true
	//A := []int{1} //true
	//A := []int{0} //true
	//A := []int{1, 0} //true
	//A := []int{0, 1} // true
	//A := []int{1, 5, 3, 4, 7} //true

	//fmt.Println(Solution4(A))
	fmt.Println(A, ":A:")
	fmt.Println("RESULT=", swapSort(A))
}

//fmt.Println(list[next], "==", list[next+toShift], "||", list[next]+1, "==", list[next+toShift])
//if list[next] == list[next+toShift] || list[next]+1 == list[next+toShift] {
//
func swapSort(list []int) bool {

	if sort.IntsAreSorted(list) {
		//Nothing to do, already sorted
		return true
	}

	//duplicate the sorted list to compare later
	sortedList := make([]int, len(list))
	copy(sortedList, list)
	sort.Ints(sortedList)

	// count any number that are duplicate or duplicates in a flush
	dup := 0

	// while you check if current number is greater in value than next item in list
	// if so check if number that are duplicate or duplicates in a flush
	for i := 0; i < len(list)-1; i++ {
		if list[i] > list[i+1] {
			dup++
			for k := i + 1; k < len(list)-1; k++ {
				//fmt.Println(list[k], "==", list[k+1], "or", list[k]+1, "==", list[k+1], ":A")
				if list[k] == list[k+1] || list[k]+1 == list[k+1] {
					dup++
				} else {
					break
				}
			}
			//store the value to swap
			val := list[i]
			//remove value from list
			list = append(list[:i], list[i+1:]...)
			//insert value at new location
			list = insert(list, dup+i, val)
			// compare lists, if not same more than one swap needed return false
			fmt.Println(list, ":B:")
			return compare(list, sortedList)
		}
	}

	//something went wrong as IntsAreSorted should of returned true
	fmt.Println("ERR?:", list)
	return true

}

func insert(s []int, at int, val int) []int {
	// Make sure there is enough room
	s = append(s, 0)
	// Move all elements of s up one slot
	copy(s[at+1:], s[at:])
	// Insert the new element at the now free position
	s[at] = val
	//return the swapped element slice
	return s
}

func compare(a []int, b []int) bool {
	return reflect.DeepEqual(a, b)
}
