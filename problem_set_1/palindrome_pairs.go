package main

import "fmt"

func isPalindrome(s string) bool {
	var length = len(s)
	var half_lengh = int(length / 2)
	for i := 0; i < half_lengh; i++ {
		if s[i] != s[length-1-i] {
			return false
		}
	}
	return true
}
func palindromePairs(words []string) [][]int {
	var length = len(words)
	var result = make([][]int, 0)
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			if i != j {
				if isPalindrome(words[i] + words[j]) {
					var temp = make([]int, 2)
					temp[0] = i
					temp[1] = j
					result = append(result, temp)
				}
			}
		}

	}
	return result
}

func main() {
	// Example usage
	inputWords := []string{"bat", "tab", "cat"}
	result := palindromePairs(inputWords)
	fmt.Println(result) // Output: [[0 1] [1 0]]
}
