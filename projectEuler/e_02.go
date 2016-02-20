package main

import "fmt"

func main() {
	cap := 4000000

	a, sum := 1, 0

	for i := 2; i < cap; {
		if i%2 == 0 {
			//print(i, "-") //turn on to print even fibs
			sum += i
		}
		a, i = i, a+i
	}
	fmt.Println(sum) //4613732
}
