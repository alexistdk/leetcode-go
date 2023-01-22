package main

import (
	"bytes"
	"fmt"
	"strings"
)

func longestCommonPrefix(strs []string) string {
	lenOfStrings := make([]int, len(strs))

	// slice with the lenghts of all of the strings
	for i, str := range strs {
		lenOfStrings[i] = len(str)
	}

	// get the lenght of the shortest string
	minLen := lenOfStrings[0]
	for _, lenght := range lenOfStrings {
		if lenght < minLen {
			minLen = lenght
		}
	}

	lettersInCommon := make([]string, minLen)

	for i := 0; i < minLen; i++ {
		charAux := bytes.NewBufferString(string(strs[0][i])).String()
		var isLetterInAllStrs bool
		for j := 0; j < len(strs); j++ {
			if bytes.NewBufferString(string(strs[j][i])).String() == charAux {
				isLetterInAllStrs = true
			} else {
				isLetterInAllStrs = false
				break
			}
		}
		if isLetterInAllStrs {
			lettersInCommon[i] = charAux
		} else {
			break
		}
	}

	return strings.Join(lettersInCommon[:], "")
}

func main() {
	strs := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(strs))
}
