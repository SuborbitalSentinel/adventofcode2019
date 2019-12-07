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

// Run day two of advent
func Run(input <-chan string) {

	intcodes := util.ChannelToArray(input)

	currentOpCodeIndex := 0
	for {
		currentOpCode, _ := strconv.Atoi(intcodes[currentOpCodeIndex])
		if currentOpCode == 99 {
			break
		}

		firstRegister, _ := strconv.Atoi(intcodes[currentOpCodeIndex+1])
		secondRegister, _ := strconv.Atoi(intcodes[currentOpCodeIndex+2])
		outputRegister, _ := strconv.Atoi(intcodes[currentOpCodeIndex+3])

		intcodes[outputRegister] = strconv.Itoa(processOpCode(currentOpCode, firstRegister, secondRegister))

		currentOpCodeIndex += 4
	}

	fmt.Println("Day2 -- Part1: ", intcodes[0])
}
