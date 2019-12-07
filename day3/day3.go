package day3

import (
	"strings"
)

func tokenizePath(input string) []string {
	return strings.Split(input, ",")
}

type point struct {
	X, Y int
}

func tracePath(tokens []string, wireNumber int, m map[point][]int) {
	x := 0
	y := 0
	for _, token := range tokens {
		// loop over each token using the above x and y value to add a new point to the hash map with the wire number
		token = token + "a"
		m[point{x, y}] = append(m[point{x, y}], wireNumber)
	}
}

// Run day 3 of advent
func Run(input <-chan string) {
	m := make(map[point][]int)

	wireNumber := 0
	for wirePath := range input {
		tokens := tokenizePath(wirePath)
		tracePath(tokens, wireNumber, m)
		wireNumber++
	}
}
