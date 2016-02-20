/* https://projecteuler.net/problem=31
In England the currency is made up of pound, £, and pence, p, and there are eight coins in general circulation:

1p, 2p, 5p, 10p, 20p, 50p, £1 (100p) and £2 (200p).
It is possible to make £2 in the following way:

1×£1 + 1×50p + 2×20p + 1×5p + 1×2p + 3×1p
How many different ways can £2 be made using any number of coins?
*/
package main

import "fmt"

const target int = 200                           // Value in pence
var coins = []int{1, 2, 5, 10, 20, 50, 100, 200} // UK coin denominations

//from lesson 13 of functions
func recursion(aTarget int, aCoins []int, size int) int {

	if aTarget < 0 {
		return 0
	}
	if aTarget == 0 {
		return 1
	}
	if size == 1 {
		return 1
	}
	return recursion(aTarget, aCoins, size-1) + recursion(aTarget-aCoins[size-1], aCoins, size)

}

//Following is how i found most entry's on internet worked, as its alot quicker than recursion
// so i ported http://blog.dreamshire.com/project-euler-31-solution/ from python

func dynamic() {
	ways := make([]int, target+1)
	ways[0] = 1 //This needs be done outside loop unlike python example

	for _, v := range coins {
		for x := v; x <= (target); x++ {
			ways[x] += ways[x-v]
		}
	}
	fmt.Println(ways[target])
}

func main() {
	fmt.Println("Recursion Way:", recursion(target, coins, len(coins)))
	fmt.Print("Dynamic Way  : ")
	dynamic()

}
