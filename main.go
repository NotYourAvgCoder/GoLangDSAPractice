package main

import (
	"fmt"

	"leetcode.com/Arrays/easy/problem2614_primeInDiagonal"
)

func main() {
	nums := [][]int{{1, 2, 3}, {5, 6, 7}, {9, 10, 11}}
	fmt.Printf("Result from prime in diagonal problem : %v\n", problem2614_primeInDiagonal.DiagonalPrime(nums))
}
