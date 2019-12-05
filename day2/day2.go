package day2

import (
	"fmt"
)

// Run day two of advent
func Run(input <-chan string) {
	part1 := 0
	part2 := 0
	for line := range input {
		fmt.Println(line)
	}

	fmt.Println("Day1 -- Part1: ", part1)
	fmt.Println("Day1 -- Part2: ", part2)
}
