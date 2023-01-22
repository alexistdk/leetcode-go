package main

import (
	"strconv"
	"strings"
)

func isPalindrome(x int) bool {
	intToStr := strconv.Itoa(x)
	stringSplited := strings.Split(intToStr, "")
	var isPalindrome bool
	for i := 0; i < len(stringSplited); i++ {
		if stringSplited[i] == stringSplited[len(stringSplited)-(i+1)] {
			isPalindrome = true
		} else {
			return false
		}
	}
	return isPalindrome
}
