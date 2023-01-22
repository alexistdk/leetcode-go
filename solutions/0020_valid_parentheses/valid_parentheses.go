package main

import (
	"fmt"
	"strings"
)

func isValid(s string) bool {

	if len(strings.Split(s, ""))%2 != 0 {
		return false
	}

	mappingChars := map[rune]rune{
		'}': '{',
		']': '[',
		')': '(',
	}

	stack := make([]rune, 0, 1)

	// 40 = '(', 41 = ')'
	// 123 = '{', 125 = '}'
	// 91 = '[', 93 = ']'

	for _, ch := range s {
		switch ch {
		case 40, 123, 91:
			stack = append(stack, ch)
		case 41, 125, 93:
			if len(stack) == 0 || stack[len(stack)-1] != mappingChars[ch] {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}

	return len(stack) == 0
}

func main() {
	fmt.Println(isValid("()"))
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("(]"))
	fmt.Println(isValid("([)]"))
	fmt.Println(isValid("{[]}"))
}

// (?:\(\)\{\}\[\])
