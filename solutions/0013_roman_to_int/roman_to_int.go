package main

import (
	"fmt"
	"strings"
)

func romanToInt(s string) int {
	stringSplited := strings.Split(s, "")
	listOfNumbers := make([]int, 10)
	amountOfLetters := len(stringSplited)
	for i := 0; i < len(stringSplited); i++ {
		switch stringSplited[i] {
		case "I":
			listOfNumbers[i] = 1
		case "V":
			listOfNumbers[i] = 5
		case "X":
			listOfNumbers[i] = 10
		case "L":
			listOfNumbers[i] = 50
		case "C":
			listOfNumbers[i] = 100
		case "D":
			listOfNumbers[i] = 500
		default:
			listOfNumbers[i] = 1000
		}
	}

	finalNumber := 0

	if amountOfLetters == 1 {
		finalNumber += listOfNumbers[0]
	} else if amountOfLetters == 2 {
		if listOfNumbers[0] >= listOfNumbers[1] {
			finalNumber += listOfNumbers[0] + listOfNumbers[1]
		} else {
			finalNumber += (listOfNumbers[1] - listOfNumbers[0])
		}
	} else {
		for i := 0; i <= amountOfLetters; i++ {
			if listOfNumbers[i] >= listOfNumbers[i+1] {
				finalNumber += listOfNumbers[i]
			} else if listOfNumbers[i] < listOfNumbers[i+1] {
				finalNumber += (listOfNumbers[i+1] - listOfNumbers[i])
				i += 1
			}
		}
	}

	return finalNumber
}

func main() {
	fmt.Println(romanToInt("MCMXCIV"))  // 1994
	fmt.Println(romanToInt("III"))      // 3
	fmt.Println(romanToInt("LVIII"))    // 58
	fmt.Println(romanToInt("IV"))       // 4
	fmt.Println(romanToInt("MCDLXXVI")) // 1476
	fmt.Println(romanToInt("MMCDXXV"))  // 2425
}
