package main

import (
	"fmt"
)

func isPalindrome(s string) bool {
	if len(s) <= 1 {
		return true
	}
	if s[0] != s[len(s)-1] {
		return false
	}
	return isPalindrome(s[1 : len(s)-1])
}

func max(a, b byte) byte {
	if a > b {
		return a
	}
	return b
}

func highestPalindrome(s string, k int) string {
	if isPalindrome(s) {
		return s
	}
	if k == 0 {
		return "-1"
	}
	if s[0] != s[len(s)-1] {
		s = string(max(s[0], s[len(s)-1])) + highestPalindrome(s[1:len(s)-1], k-1) + string(max(s[0], s[len(s)-1]))
	} else {
		s = string(s[0]) + highestPalindrome(s[1:len(s)-1], k) + string(s[len(s)-1])
	}
	return s
}

func main() {
	s := "3943"
	k := 1
	result := highestPalindrome(s, k)
	fmt.Println(result) // Output: "3993"
}
