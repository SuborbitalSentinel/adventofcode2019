package day2

import (
	"fmt"
	"strconv"

	"github.com/millera023/adventofcode/util"
)

// Run day two of advent
func Run(input <-chan string) {
	intcodes := util.ChannelToSlice(input)

	// Part 1
	part1 := make([]string, len(intcodes))
	copy(part1, intcodes)
	fmt.Println("Day2 -- Part1: ", executeProgram(part1, 12, 2))

	// Part 2
	fmt.Println("Day2 -- Part2: ", part2(intcodes))
}

func part2(input []string) int {
	for i := 0; i <= 99; i++ {
		for j := 0; j <= 99; j++ {
			part2 := make([]string, len(input))
			copy(part2, input)
			if executeProgram(part2, i, j) == "19690720" {
				return 100*i + j
			}
		}
	}

	return -1
}

func processOpCode(code int, reg1 int, reg2 int) int {
	switch code {
	case 1:
		return reg1 + reg2
	case 2:
		return reg1 * reg2
	default:
		panic(fmt.Sprintf("%v %v", "Bad Op Code:", code))
	}

}

func getRegisterValue(input []string, idx int) int {
	valueIndex, _ := strconv.Atoi(input[idx])
	value, _ := strconv.Atoi(input[valueIndex])

	return value
}

func executeProgram(input []string, start1 int, start2 int) string {
	input[1] = strconv.Itoa(start1)
	input[2] = strconv.Itoa(start2)

	currentOpCodeIndex := 0
	for {
		currentOpCode, _ := strconv.Atoi(input[currentOpCodeIndex])
		if currentOpCode == 99 {
			break
		}

		firstRegister := getRegisterValue(input, currentOpCodeIndex+1)
		secondRegister := getRegisterValue(input, currentOpCodeIndex+2)
		outputRegister, _ := strconv.Atoi(input[currentOpCodeIndex+3])

		input[outputRegister] = strconv.Itoa(processOpCode(currentOpCode, firstRegister, secondRegister))

		currentOpCodeIndex += 4
	}

	return input[0]
}
