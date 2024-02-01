package problem2996smallestint

func MissingInteger(nums []int) int {
	// sequentialNums := []int{}
	greatestSequentialNum := nums[0]
	isInSlice := make(map[int]bool)
	currentSum := 0
	for index, num := range nums {
		// fmt.Printf("Processing : %v\n", num)
		if index == 0 {
			currentSum += num
		} else if num != nums[index-1]+1 {
			if greatestSequentialNum < currentSum {
				greatestSequentialNum = currentSum
			}
			currentSum = num
		} else {
			currentSum += num
		}
		isInSlice[num] = true
	}
	if greatestSequentialNum < currentSum {
		greatestSequentialNum = currentSum
	}
	if isInSlice[greatestSequentialNum] {
		return greatestSequentialNum + 1
	}
	return greatestSequentialNum
}
