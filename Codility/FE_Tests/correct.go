package main

import "fmt"

func main() {

	//A := []int{1, 5, 3, 3, 7} //true
	//A := []int{1, 5, 3, 3, 3, 3, 3, 3, 3, 3, 7} //true
	//A := []int{1, 5, 3, 3, 3, 7, 3, 3, 3, 3, 7} //false
	//A := []int{1, 3, 3, 3, 4, 3, 3, 3, 3, 7} //true
	//A := []int{1, 2, 3, 3, 3, 3, 3, 3, 3, 7, 5} //true
	A := []int{7, 2, 3, 3, 3, 3, 3, 3, 3, 3, 4, 4, 4, 5} // true // fails counter not dealing with more than one flush
	//A := []int{1, 2, 4, 3, 7} //true
	//A := []int{1, 3, 5} //true
	//A := []int{1} //true
	//A := []int{1, 0} //true
	//A := []int{0, 1} /true
	//A := []int{1, 5, 3, 4, 7} //true

	//fmt.Println(Solution4(A))
	fmt.Println(A)
	fmt.Println("RESULT=", swapSort(A))
}

//fmt.Println(list[next], "==", list[next+toShift], "||", list[next]+1, "==", list[next+toShift])
//if list[next] == list[next+toShift] || list[next]+1 == list[next+toShift] {
//
func swapSort(list []int) bool {
	swapCount := 0
	toShift := 0
	for itemCount := len(list) - 1; ; itemCount-- {
		hasChanged := false
		for current := 0; current < itemCount; current++ {
			next := current + 1
			if list[current] > list[next] {
				temp := list[current]
				list = append(list[:current], list[current+1:]...)
				//list = list[current], list[:len(list)-1]
				//insert
				fmt.Println(list, "Atemp=", temp)
				for shift := current; shift < len(list)-1; shift++ {
					fmt.Println(list[next], "==", list[next+toShift], "||", list[next]+1, "==", list[next+toShift])
					if list[next] == list[next+toShift] || list[next]+1 == list[next+toShift] {
						toShift++
					}
				}
				fmt.Println("toshift", toShift, "next", next)
				if current == 0 && {
					list = insert(list, toShift+next+1, temp)
				}else{
					list = insert(list, toShift+next, temp)
				}
				swapCount++
				if swapCount >= 2 {
					fmt.Println("Too many swaps needed")
					fmt.Println(list, "SC:", swapCount)
					return false
				}
			}
			hasChanged = true
		}
		if !hasChanged {
			break
		}

	}
	fmt.Println(list, "SC:", swapCount)
	return true

}

func insert(s []int, at int, val int) []int {
	// Make sure there is enough room
	s = append(s, 0)
	fmt.Println(s, "insertA")
	// Move all elements of s up one slot
	copy(s[at+1:], s[at:])
	fmt.Println(s, "insertB")
	// Insert the new element at the now free position
	s[at] = val
	fmt.Println(s, "insertC")
	return s
}
