package problem2614_primeInDiagonal

import "fmt"

func isPrime(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func DiagonalPrime(nums [][]int) int {
	numberOfElementsInDiagonal := len(nums)
	i := 0
	j := 0
	k := 0
	l := len(nums[0]) - 1
	res := 0

	for n := 1; n <= numberOfElementsInDiagonal; n++ {
		fmt.Printf("Element on one diagonal : %v and another %v\n", nums[i][j], nums[k][l])
		if isPrime(nums[i][j]) && nums[i][j] > res {
			res = nums[i][j]
		}
		if isPrime(nums[k][l]) && nums[k][l] > res {
			res = nums[k][l]
		}
		i += 1
		j += 1
		k += 1
		l -= 1
	}
	return res
}
