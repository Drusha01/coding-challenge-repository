package main

import "fmt"

func lengthOfLIS(nums []int) int {
	var length = len(nums)
	if length == 0 {
		return 0
	} else {
		var list = make([]int, length)
		for i := 0; i < length; i++ {
			list[i] = 1
		}
		for i := 0; i < length; i++ {
			for j := 0; j < i; j++ {
				if nums[j] < nums[i] && list[i] < list[j]+1 {
					list[i] = list[j] + 1
				}
			}
		}
		var max = 1
		for i := 0; i < length; i++ {
			if list[i] > max {
				max = list[i]
			}

		}
		return max
	}
}
func main() {
	// Example usage
	inputNumbers := []int{10, 9, 2, 5, 3, 7, 101, 18}
	result := lengthOfLIS(inputNumbers)
	fmt.Println(result) // Output: 4
}
