package day5

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/millera023/adventofcode/util"
)

// Run day 5 of advent
func Run(input <-chan string) {
	intcodes := util.ChannelToSlice(input)

	// Part 1
	fmt.Println("Day5 -- Part1: ")
	executeProgram(intcodes)
}

func splitOpCode(opCode string) (string, []string) {
	reversedAndPaddedOpCode := util.Reverse(fmt.Sprintf("0000%v", opCode))

	parsedOpCode := util.Reverse(string(reversedAndPaddedOpCode[:2]))
	paramMode1 := string(reversedAndPaddedOpCode[2])
	paramMode2 := string(reversedAndPaddedOpCode[3])

	return parsedOpCode, []string{paramMode1, paramMode2}
}

func processOpCode(opCodeIndex int, program []string) int {
	const alu, io = 4, 2
	currentOpCode, paramModes := splitOpCode(program[opCodeIndex])

	switch currentOpCode {

	case "01":
		paramMode1 := paramModes[0]
		paramMode2 := paramModes[1]
		reg1, reg2 := getRegisterValues(program, program[opCodeIndex+1], program[opCodeIndex+2], paramMode1, paramMode2)

		program[outputIndex(program, opCodeIndex+3)] = strconv.Itoa(reg1 + reg2)
		return alu

	case "02":
		paramMode1 := paramModes[0]
		paramMode2 := paramModes[1]
		reg1, reg2 := getRegisterValues(program, program[opCodeIndex+1], program[opCodeIndex+2], paramMode1, paramMode2)

		program[outputIndex(program, opCodeIndex+3)] = strconv.Itoa(reg1 * reg2)
		return alu

	case "03":
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Input Requested: ")
		input, _ := reader.ReadString('\n')
		input = strings.Replace(input, "\n", "", -1)

		program[outputIndex(program, opCodeIndex+1)] = input
		return io

	case "04":
		paramMode1 := paramModes[0]
		if paramMode1 == "1" {
			fmt.Println("Output: ", program[opCodeIndex+1])
		} else {
			fmt.Println("Output: ", program[outputIndex(program, opCodeIndex+1)])
		}

		return io

	default:
		panic(fmt.Sprintf("%v %v", "Bad Op Code:", program[opCodeIndex]))
	}
}

func outputIndex(program []string, programPointer int) int {
	outputRegisterIndex, _ := strconv.Atoi(program[programPointer])

	return outputRegisterIndex
}

func getRegisterValues(program []string, r1 string, r2 string, paramMode1 string, paramMode2 string) (int, int) {
	reg1, _ := strconv.Atoi(r1)
	reg2, _ := strconv.Atoi(r2)

	if paramMode1 == "0" {
		reg1, _ = strconv.Atoi(program[reg1])
	}
	if paramMode2 == "0" {
		reg2, _ = strconv.Atoi(program[reg2])
	}

	return reg1, reg2
}

func executeProgram(program []string) {
	currentOpCodeIndex := 0
	for {
		if program[currentOpCodeIndex] == "99" {
			fmt.Println("Program Halt")
			break
		}

		currentOpCodeIndex += processOpCode(currentOpCodeIndex, program)
	}
}
