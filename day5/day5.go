package day5

import (
	"fmt"
	"strconv"

	"github.com/millera023/adventofcode/util"
)

// Run day 5 of advent
func Run(input <-chan string) {
	intcodes := util.ChannelToSlice(input)

	// Part 1
	fmt.Println("Day5 -- Part1: ", executeProgram(intcodes))
}

func splitOpCode(opCode string) (string, []string) {
	reversedAndPaddedOpCode := util.Reverse(fmt.Sprintf("0000%v", opCode))

	parsedOpCode := util.Reverse(string(reversedAndPaddedOpCode[:2]))
	paramMode1 := string(reversedAndPaddedOpCode[2])
	paramMode2 := string(reversedAndPaddedOpCode[3])

	return parsedOpCode, []string{paramMode1, paramMode2}
}

func processOpCode(opCodeIndex int, program []string) int {
	currentOpCode, paramModes := splitOpCode(program[opCodeIndex])

	switch currentOpCode {
	case "01":
		paramMode1 := paramModes[0]
		paramMode2 := paramModes[1]
		reg1, reg2 := getRegisterValues(program, program[opCodeIndex+1], program[opCodeIndex+2], paramMode1, paramMode2)
		return reg1 + reg2
	case "02":
		paramMode1 := paramModes[0]
		paramMode2 := paramModes[1]
		reg1, reg2 := getRegisterValues(program, program[opCodeIndex+1], program[opCodeIndex+2], paramMode1, paramMode2)
		return reg1 * reg2
	default:
		panic(fmt.Sprintf("%v %v", "Bad Op Code:", program[opCodeIndex]))
	}

}

func getRegisterValues(program []string, r1 string, r2 string, paramMode1 string, paramMode2 string) (int, int) {
	reg1, _ := strconv.Atoi(r1)
	reg2, _ := strconv.Atoi(r2)

	if paramMode1 == "1" {
		reg1, _ = strconv.Atoi(program[reg1])
	}
	if paramMode2 == "1" {
		reg2, _ = strconv.Atoi(program[reg2])
	}

	return reg1, reg2
}

func executeProgram(program []string) string {
	currentOpCodeIndex := 0
	for {
		if program[currentOpCodeIndex] == "99" {
			break
		}

		currentOpCodeIndex += processOpCode(currentOpCodeIndex, program)
	}

	return program[0]
}
