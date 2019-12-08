package day3

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Run day 3 of advent
func Run(input <-chan string) {
	m := make(map[point][]wire)

	wireNumber := 0
	for wirePath := range input {
		tokens := tokenizePath(wirePath)
		tracePath(tokens, wireNumber, m)
		wireNumber++
	}

	// point 0,0 will always have the wires connected
	delete(m, point{0, 0})

	fmt.Println("Day3 -- Part1: ", part1(m))
	fmt.Println("Day3 -- Part2: ", part2(m))
}

func part1(m map[point][]wire) int {
	minSeenDistance := math.MaxInt32
	for key, value := range m {
		if len(value) == 2 { // Assuming that there were be at most 2 values per point in space
			if value[0].Number != value[1].Number {
				distance := abs(key.X) + abs(key.Y)
				if distance < minSeenDistance {
					minSeenDistance = distance
				}
			}
		}
	}

	return minSeenDistance
}

func part2(m map[point][]wire) int {
	minSeenSteps := math.MaxInt32
	for _, value := range m {
		if len(value) == 2 { // Assuming that there were be at most 2 values per point in space
			if value[0].Number != value[1].Number {
				distance := value[0].Steps + value[1].Steps
				if distance < minSeenSteps {
					minSeenSteps = distance
				}
			}
		}
	}

	return minSeenSteps
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

type point struct {
	X, Y int
}

type wire struct {
	Number, Steps int
}

func tokenizePath(input string) []string {
	return strings.Split(input, ",")
}

func parseToken(input string) (string, int) {
	direction := input[:1]
	distance, _ := strconv.Atoi(input[1:])
	return direction, distance
}

func tracePath(tokens []string, wireNumber int, m map[point][]wire) {
	m[point{0, 0}] = append(m[point{0, 0}], wire{wireNumber, 0})

	x, y, stepCount := 0, 0, 0
	for _, token := range tokens {
		direction, distance := parseToken(token)
		switch direction {
		case "R":
			for i := 1; i <= distance; i++ {
				x++
				stepCount++
				m[point{x, y}] = append(m[point{x, y}], wire{wireNumber, stepCount})
			}
		case "L":
			for i := 1; i <= distance; i++ {
				x--
				stepCount++
				m[point{x, y}] = append(m[point{x, y}], wire{wireNumber, stepCount})
			}
		case "U":
			for i := 1; i <= distance; i++ {
				y++
				stepCount++
				m[point{x, y}] = append(m[point{x, y}], wire{wireNumber, stepCount})
			}
		case "D":
			for i := 1; i <= distance; i++ {
				y--
				stepCount++
				m[point{x, y}] = append(m[point{x, y}], wire{wireNumber, stepCount})
			}
		}
	}
}
