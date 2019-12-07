package day2

import (
	"fmt"
	"strconv"

	"github.com/millera023/adventofcode/util"
)

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

// Run day two of advent
func Run(input <-chan string) {

	intcodes := util.ChannelToArray(input)

	currentOpCodeIndex := 0
	for {
		currentOpCode, _ := strconv.Atoi(intcodes[currentOpCodeIndex])
		if currentOpCode == 99 {
			break
		}

		firstRegister := getRegisterValue(intcodes, currentOpCodeIndex+1)
		secondRegister := getRegisterValue(intcodes, currentOpCodeIndex+2)
		outputRegister, _ := strconv.Atoi(intcodes[currentOpCodeIndex+3])
		//fmt.Println(fmt.Sprintf("%v %v %v %v %v", "Full OpCode index:", currentOpCodeIndex, currentOpCodeIndex+1, currentOpCodeIndex+2, currentOpCodeIndex+3))
		//fmt.Println(fmt.Sprintf("%v %v %v %v %v", "Full OpCode:", currentOpCode, firstRegister, secondRegister, outputRegister))

		//fmt.Println(fmt.Sprintf("%v %v", "Registers before:", intcodes))
		//fmt.Println(fmt.Sprintf("%v %v %v %v %v", "OpCode was:", currentOpCode, "with values:", firstRegister, secondRegister))
		//fmt.Println(fmt.Sprintf("%v %v %v %v", "Output before:", intcodes[outputRegister], "in Register:", outputRegister))

		intcodes[outputRegister] = strconv.Itoa(processOpCode(currentOpCode, firstRegister, secondRegister))

		//fmt.Println(fmt.Sprintf("%v %v %v %v", "Output after:", intcodes[outputRegister], "in Register:", outputRegister))
		//fmt.Println(fmt.Sprintf("%v %v", "Registers after:", intcodes))
		//fmt.Println("")

		currentOpCodeIndex += 4
	}

	fmt.Println("Day2 -- Part1: ", intcodes[0])
}
