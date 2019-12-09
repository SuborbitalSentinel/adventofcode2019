package day4

import (
	"fmt"
	"strconv"
)

// Run day4 of advent
func Run() {
	const minPassword = 138241
	const maxPassword = 674034

	part1 := 0
	part2 := 0
	for i := minPassword; i < maxPassword; i++ {
		if alwaysAscendingDigits(i) && atLeastTwoAdjacentDigitsAreSame(i) {
			part1++
		}

		if alwaysAscendingDigits(i) && onlyTwoAdjacentDigitsAreSame(i) {
			part2++

		}
	}

	fmt.Println("Day4 -- Part1: ", part1)
	fmt.Println("Day4 -- Part2: ", part2)
}

func onlyTwoAdjacentDigitsAreSame(input int) bool {
	panic("Finish this method, wrong right now")
	//digits := []rune(strconv.Itoa(input))

	//possiblyMet := false
	//previousDigit := digits[0]
	//for _, digit := range digits[1:] {
	//	if possiblyMet {
	//		if digit != previousDigit {
	//			return true
	//		}
	//		possiblyMet = false
	//	}

	//	if digit == previousDigit {
	//		possiblyMet = true
	//	}

	//	previousDigit = digit
	//}
	//return false
}

func atLeastTwoAdjacentDigitsAreSame(input int) bool {
	digits := []rune(strconv.Itoa(input))
	previousDigit := digits[0]
	for _, digit := range digits[1:] {
		if digit == previousDigit {
			return true
		}
		previousDigit = digit
	}
	return false
}

func alwaysAscendingDigits(input int) bool {
	digits := []rune(strconv.Itoa(input))
	previousDigit := digits[0]
	for _, digit := range digits[1:] {
		if digit < previousDigit {
			return false
		}
		previousDigit = digit
	}
	return true
}
