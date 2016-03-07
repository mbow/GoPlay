package main

import (
	"fmt"
	"strconv"
)

var tests = []struct {
	in   int
	want int
}{
	{9, 2},
	{529, 4},
	{20, 1},
	{1041, 5},
	{104165, 2},
	{10418884, 3},
	{15, 0},
	{-1, 0},
	{0, 0},
	{2147483647, 0},
	{2147483648, 0},
}

func main() {
	for _, v := range tests {
		fmt.Println(Solution(v.in), "==", v.want)
		fmt.Println("")
	}
	//return 0
}

func Solution(N int) int {
	gap := 0
	biggap := 0
	setbit := 0

	S := strconv.FormatInt(int64(N), 2)
	fmt.Println(S)
	if N <= 0 || N > 2147483647 {
		return 0
	}

	for i, v := range S {
		//if im ascii 49 (1) im odd and do not dived by 2
		if v%2 != 0 {
			setbit = i
		}
		//fmt.Println(i, setbit)
	}

	if setbit <= 1 {
		return 0
	}

	for i := 0; i < setbit; i++ {
		//if im ascii 49 (1) im odd and do not dived by 2
		if S[i]%2 != 0 {
			for m := i + 1; m < setbit; m++ {
				//if im ascii 48 (0) im even and dived by 2
				if S[m]%2 == 0 {
					gap++
					if biggap < gap {
						biggap = gap
					}
				} else {
					gap = 0
				}
			}
			gap = 0
		}
	}
	return biggap
}
