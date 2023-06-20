package main

import (
	"fmt"
)

func isBalancedBracket(s string) string {
	stack := make([]rune, 0)
	return isBalancedBracketRecursive([]rune(s), stack)
}

func isBalancedBracketRecursive(chars []rune, stack []rune) string {
	if len(chars) == 0 {
		if len(stack) == 0 {
			return "YES"
		}
		return "NO"
	}

	char := chars[0]
	if isOpeningBracket(char) {
		stack = append(stack, char)
	} else {
		if len(stack) == 0 || !isMatchingBracket(stack[len(stack)-1], char) {
			return "NO"
		}
		stack = stack[:len(stack)-1]
	}

	return isBalancedBracketRecursive(chars[1:], stack)
}

func isOpeningBracket(char rune) bool {
	return char == '(' || char == '{' || char == '['
}

func isMatchingBracket(open, close rune) bool {
	switch open {
	case '(':
		return close == ')'
	case '{':
		return close == '}'
	case '[':
		return close == ']'
	default:
		return false
	}
}

func main() {
	input := "{[()]}"
	output := isBalancedBracket(input)
	fmt.Println(output) // Output: YES

	input = "{[(])}"
	output = isBalancedBracket(input)
	fmt.Println(output) // Output: NO

	input = "{(([])[])[]}"
	output = isBalancedBracket(input)
	fmt.Println(output) // Output: YES
}
