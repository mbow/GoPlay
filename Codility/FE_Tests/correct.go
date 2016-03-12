package main

import (
	"fmt"
	"reflect"
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

	{[]int{3, 1, 1, 2, 2, 3, 3, 1}, true},                                                   //43 f
	{[]int{1, 1, 3, 1, 1, 2, 2, 3, 3, 1}, true},                                             //44 f
	{[]int{1, 1, 1, 30, 1, 1, 2, 2, 2, 2, 3, 3, 3, 3, 3, 4, 4, 4, 12, 24, 1, 50, 60}, true}, //45 f
	{[]int{1, 1, 3, 1, 1, 2, 2, 3, 3, 1}, true},                                             //46 f
	{[]int{3, 6, 3, 3, 3, 4, 4, 4, 5, 5, 5, 3, 6, 6, 7}, true},                              //47 f
}

func main() {

	input := testData
	failed := 0
	totaltests := 0

	for i, v := range input {
		//fmt.Println("")
		if swapNcompare(v.in) == v.want {
			fmt.Println("TEST", i, "PASS")
		} else {

			fmt.Println("TEST", i, "!!FAIL!!")
			failed++
		}
		totaltests++
	}
	fmt.Println("PASS:", totaltests-failed, "|", "Fail:", failed)
}

func swapNcompare(list []int) bool {

	if sort.IntsAreSorted(list) {
		//Nothing to do, already sorted
		return true
	}

	//duplicate the sorted list to compare later
	sortedList := make([]int, len(list))
	copy(sortedList, list)
	sort.Ints(sortedList)

	oneSwapSort(list)

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
		change := false

		b := make([]int, len(a))
		copy(b, a)

		for pos > 0 && b[pos-1] == b[pos] {
			pos--
		}

		for pos > 0 && b[pos-1] > b[pos] {
			b[pos-1], b[pos] = b[pos], b[pos-1]
			change = true
			pos--

		}

		if change {
			//fmt.Println("swap:", upperBound, a[upperBound], "with", pos, a[pos])
			a[i], a[pos] = a[pos], a[i]
			return
		}
	}

}
