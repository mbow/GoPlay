package main

import (
	"fmt"
	//"reflect"
	"sort"
)

type tests struct {
	in   []int
	want bool
}

var testData = []tests{
	{[]int{1}, true},        //0
	{[]int{1, 1}, true},     //1
	{[]int{1, 2}, true},     //2
	{[]int{2, 1}, true},     //3
	{[]int{2, 2}, true},     //4
	{[]int{1, 1, 1}, true},  //5
	{[]int{1, 1, 2}, true},  //6
	{[]int{1, 2, 1}, true},  //7
	{[]int{1, 2, 2}, true},  //8
	{[]int{1, 2, 3}, true},  //9
	{[]int{1, 3, 1}, true},  //10
	{[]int{1, 3, 2}, true},  //11
	{[]int{1, 3, 3}, true},  //12
	{[]int{2, 1, 1}, true},  //13
	{[]int{2, 1, 2}, true},  //14
	{[]int{2, 1, 3}, true},  //15
	{[]int{2, 2, 1}, true},  //16
	{[]int{2, 2, 3}, true},  //17
	{[]int{2, 3, 1}, false}, //18
	{[]int{2, 3, 2}, true},  //19
	{[]int{2, 3, 3}, true},  //20
	{[]int{3, 1, 1}, true},  //21
	{[]int{3, 1, 2}, false}, //22
	{[]int{3, 1, 3}, true},  //23
	{[]int{3, 2, 1}, true},  //24

	{[]int{2, 3, 2, 4}, true},  //25
	{[]int{1, 2, 4, 3}, true},  //26
	{[]int{1, 4, 3, 2}, true},  //27
	{[]int{1, 3, 4, 2}, false}, //28

	{[]int{1, 5, 3, 3, 7}, true},  //29
	{[]int{1, 3, 5, 3, 4}, false}, //30

	{[]int{1, 5, 3, 3, 3, 3, 3, 3, 3, 3, 7}, true},                                                           //31
	{[]int{1, 5, 3, 3, 3, 7, 3, 3, 3, 3, 7}, false},                                                          //32
	{[]int{6, 3, 3, 3, 3, 4, 4, 4, 5, 5, 5, 5, 6, 6, 7}, false},                                              //33
	{[]int{1, 3, 3, 3, 4, 3, 3, 3, 3, 7}, true},                                                              //34
	{[]int{1, 2, 3, 3, 3, 3, 3, 3, 3, 7, 5}, true},                                                           //35
	{[]int{7, 2, 3, 3, 3, 3, 3, 3, 3, 3, 4, 4, 4, 5}, false},                                                 //36
	{[]int{2, 3, 3, 7, 3, 3, 3, 3, 3, 3, 3, 4, 4, 4, 5}, false},                                              //37
	{[]int{1, 3, 3, 3, 4, 3, 3, 3, 3, 7}, true},                                                              //38
	{[]int{1, 5, 3, 4, 7}, false},                                                                            //39
	{[]int{3, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 1}, true},                                //40
	{[]int{1, 1, 1, 3, 3, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 3, 3, 1, 3, 3, 3}, false}, //41
	{[]int{1, 1, 1, 3, 1, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 1, 2, 2, 2, 2, 2, 2, 2, 3, 3, 1, 3, 3, 3}, false}, //42

	{[]int{3, 5, 7, 6, 9, 11}, true}, //43
	{[]int{1, 4, 5, 6, 7, 2}, false}, //44
	{[]int{1, 11, 5, 6, 7, 2}, true}, //45

	{[]int{3, 1, 1, 2, 2, 3, 3, 1}, true},                                                   //46 f
	{[]int{1, 1, 3, 1, 1, 2, 2, 3, 3, 1}, true},                                             //47 f
	{[]int{1, 1, 3, 1, 1, 2, 2, 3, 3, 1}, true},                                             //48 f
	{[]int{3, 6, 3, 3, 3, 4, 4, 4, 5, 5, 5, 3, 6, 6, 7}, true},                              //49 f
	{[]int{1, 1, 1, 30, 1, 1, 2, 2, 2, 2, 3, 3, 3, 3, 3, 4, 4, 4, 12, 24, 1, 50, 60}, true}, //50 f

}

func main() {

	input := testData
	failed := 0
	totaltests := 0

	for _, v := range input {
		if canBeSortedByOneSwap(v.in) == v.want {
			//fmt.Println("PASS")
		} else {
			failed++
		}
		totaltests++
	}
	fmt.Println("PASS:", totaltests-failed, "|", "Fail:", failed)
}

func canBeSortedByOneSwap(list []int) bool {

	if sort.IntsAreSorted(list) {
		//Nothing to do, already sorted
		return true
	}

	//Make a sorted list to compare
	sortedList := make([]int, len(list))
	copy(sortedList, list)
	sort.Ints(sortedList)

	//Counter for differences in list
	diffCnt := 0

	//compare each item against corect list and count missmatches
	//if more than 2 differences, needs more than one swap so return
	for i, v := range list {
		if v != sortedList[i] {
			diffCnt++
			if diffCnt > 2 {
				return false
			}
		}
	}

	// ret
	if diffCnt == 2 {
		return true
	}

	//catch case something went wrong
	//fmt.Println("ERR") // return ERROR
	return false
}

/* Old bubble sort with one  swap
func swapNcompare(list []int) bool {

	if sort.IntsAreSorted(list) {
		//Nothing to do, already sorted
		return true
	}

	//duplicate the sorted list to compare later
	sortedList := make([]int, len(list))
	copy(sortedList, list)
	sort.Ints(sortedList)
	//fmt.Println(list, "[IN]")
	oneSwapSort(list)
	//fmt.Println(list, "[OUT]")
	// compare if the same only one swap needed
	return compare(list, sortedList)

}

func compare(a []int, b []int) bool {
	return reflect.DeepEqual(a, b)
}

func oneSwapSort(a []int) {
	// pass in the slice and the current index, desending
	for i := len(a) - 1; i >= 0; i-- {
		pos := i
		pos2 := 0
		change := false
		ran := false

		b := make([]int, len(a))
		copy(b, a)

		//{[]int{1, 3, 3, 3, 4, 3, 3, 3, 3, 7}, true},
		//fmt.Println(i, ":", a)
		for pos > 0 && b[pos-1] == b[pos] {
			//fmt.Println("pos--")
			pos--
		}

		for pos > 0 && b[pos-1] > b[pos] {
			//fmt.Println("swap0:", b, pos)
			b[pos-1], b[pos] = b[pos], b[pos-1]
			change = true
			pos--
			//fmt.Println("swap1:", b, pos)
			pos2 = pos

			for pos2 > 0 && b[pos2-1] == b[pos2] && b[pos2] != b[i] && ran != true {
				//fmt.Println("swap2:", b, "@", pos2)
				//b[pos2-1], b[pos2] = b[pos2], b[pos2-1]
				pos2--
				if pos2 > 0 && b[pos2-1] != b[pos2] {
					ran = true
					if pos2 > 0 && b[pos2-1] > b[pos2] {
						//fmt.Println("swap3", pos2)
						pos2--
					}
				}

			}

		}

		if change {
			//fmt.Println("swap:", i, a[i], "with", pos, a[pos])
			//a[i], a[pos] = a[pos], a[i]
			//fmt.Println("swap:", i, a[i], "with", pos2, a[pos2])
			a[i], a[pos2] = a[pos2], a[i]
			return
		}
	}

}
*/
