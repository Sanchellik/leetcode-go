package main

import "fmt"

func missingNumber(nums []int) int {
	n := len(nums)

	expectedSum := n * (n + 1) / 2

	calculatedSum := 0

	for i := 0; i < n; i++ {
		calculatedSum += nums[i]
	}

	return expectedSum - calculatedSum
}

func main() {
	fmt.Println(missingNumber([]int{0, 1}))
}
