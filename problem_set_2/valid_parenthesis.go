package main

import (
	"fmt"
)

func isValid(s string) bool {
	var b string
	b = ""
	var length = len(s)
	for i := 0; i < length; i++ {
		if string(s[i]) == "{" || s[i] == '(' || s[i] == '[' {
			b = b + string(s[i])
		} else {
			var last_index = len(b) - 1
			if (len(b) == 0) || (s[i] == '}' && b[last_index] != '{') || (s[i] == ')' && b[last_index] != '(') || (s[i] == ']' && b[last_index] != '[') {
				return false
			}
			b = b[:last_index]
		}

	}
	return len(b) == 0
}
func main() {
	inputString := "()[]{}"
	result := isValid(inputString)
	fmt.Println(result)
}
