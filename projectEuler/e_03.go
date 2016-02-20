package main

import "fmt"
import "math/big"

func main() {
	number := 600851475143

	checkprimes := func(n int) bool {
		i := big.NewInt(int64(n))
		return i.ProbablyPrime(20)
	}

	primes := make([]int, 0)
	for i := 2; i*i <= number; i++ {
		if number%i == 0 && checkprimes(i) {
			primes = append(primes, i)
		}
	}
	fmt.Println(primes[len(primes)-1]) //6857
}
