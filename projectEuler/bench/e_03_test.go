package main

import "fmt"
import "math/big"
import "testing"

func BenchmarkPrime(b *testing.B) {
	number := 600851475143

	checkprime := func(n int) bool {
		i := big.NewInt(int64(n))
		return i.ProbablyPrime(n)
	}

	primelist := make([]int, 0)
	for i := 2; i*i <= b.N; i++ {
		if number%i == 0 && checkprime(i) {
			primelist = append(primelist, i)
		}
	}
	fmt.Println(primelist[:]) //
}
