package main

import (
	"fmt"
	"reflect"
	"sort"
)

type tests struct {
	num  int
	in   []int
	want bool
}

var testData = []tests{
	{1, []int{0}, true},
	{2, []int{0, 1}, true},
	{3, []int{0, 0}, true},
	{4, []int{1, 0}, true},
	{5, []int{1, 1}, true},
	{6, []int{1, 2, 3}, true},
	{7, []int{3, 2, 1}, true}, //<------To do still failed
	{8, []int{1, 2, 1}, true},
	{9, []int{2, 1, 1}, true},
	{10, []int{1, 3, 2, 4}, true},
	{11, []int{1, 2, 4, 3}, true},
	{12, []int{1, 4, 3, 2}, false},
	{13, []int{1, 3, 4, 2}, false},
	{14, []int{1, 5, 3, 3, 7}, true},
	{15, []int{1, 5, 3, 3, 3, 3, 3, 3, 3, 3, 7}, true},
	{16, []int{1, 5, 3, 3, 3, 7, 3, 3, 3, 3, 7}, false},
	{17, []int{6, 3, 3, 3, 3, 4, 4, 4, 5, 5, 5, 5, 6, 6, 7}, true},
	{18, []int{1, 3, 3, 3, 4, 3, 3, 3, 3, 7}, true},
	{19, []int{1, 2, 3, 3, 3, 3, 3, 3, 3, 7, 5}, true},
	{20, []int{7, 2, 3, 3, 3, 3, 3, 3, 3, 3, 4, 4, 4, 5}, true},
	{21, []int{2, 3, 3, 7, 3, 3, 3, 3, 3, 3, 3, 4, 4, 4, 5}, true},
	{22, []int{1, 3, 3, 3, 4, 3, 3, 3, 3, 7}, true},
	{23, []int{1, 5, 3, 4, 7}, true},
}

func main() {

	//total := 0
	input := testData
	for _, v := range input {
		//fmt.Println("TEST", v.num)
		if swapSort(v.in) == v.want {
			fmt.Println("TEST", v.num, "TRUE:PASSED")
		} else {
			fmt.Println("TEST", v.num, "FAILLED")
		}
	}
}

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
					if list[i] == list[k+1] {
						break
					}
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
